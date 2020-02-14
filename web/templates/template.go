package template

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/echo"
)

// TmplEngine : テンプレートエンジン構造体(テンプレート設定を元に「テンプレート」と付随する「関数」を管理します)
type TmplEngine struct {
	config      TmplConfig
	tmplMap     map[string]*template.Template
	tmplMutex   sync.RWMutex
	fileHandler FileHandler
}

// TmplConfig : テンプレート設定構造体(テンプレートの読み込み必要な情報、区切り文字等を定義します)
type TmplConfig struct {
	Root         string
	Extension    string
	Master       string
	Partials     []string
	Functions    template.FuncMap
	DisableCache bool
	Delims       Delims
}

// Delims : デリミタ構造体(テンプレートの目印となるデリミタを設定します)
type Delims struct {
	Left  string
	Right string
}

// FileHandler : テンプレート設定に則ってテンプレートファイルのStreamを読み込み、文字列を返却します。
type FileHandler func(config TmplConfig, tplFile string) (content string, err error)

// New : テンプレートエンジンのインスタンスを生成します。
func New(conf TmplConfig) *TmplEngine {
	return &TmplEngine{
		config:      conf,
		tmplMap:     make(map[string]*template.Template),
		tmplMutex:   sync.RWMutex{},
		fileHandler: DefaultFileHandler(),
	}
}

// Render : テンプレートに則って動的にページを描画します。
func (t *TmplEngine) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.executeRender(w, name, data)
}

func (t *TmplEngine) executeRender(out io.Writer, name string, data interface{}) error {
	useMaster := true
	if filepath.Ext(name) == t.config.Extension {
		useMaster = false
		name = strings.TrimSuffix(name, t.config.Extension)
	}
	return t.executeTemplate(out, name, data, useMaster)
}

func (t *TmplEngine) executeTemplate(out io.Writer, name string, data interface{}, useMaster bool) error {
	var tpl *template.Template
	var err error

	allFunctions := make(template.FuncMap, 0)
	allFunctions["include"] = func(layout string) (template.HTML, error) {
		buf := new(bytes.Buffer)
		err := t.executeTemplate(buf, layout, data, false)
		return template.HTML(buf.String()), err
	}

	// Get the plugin collection
	for k, v := range t.config.Functions {
		allFunctions[k] = v
	}

	t.tmplMutex.RLock()
	tpl, loadResult := t.tmplMap[name]
	t.tmplMutex.RUnlock()

	exeName := name
	if useMaster && t.config.Master != "" {
		exeName = t.config.Master
	}

	if !loadResult || t.config.DisableCache {
		tplList := make([]string, 0)
		if useMaster {
			//render()
			if t.config.Master != "" {
				tplList = append(tplList, t.config.Master)
			}
		}
		tplList = append(tplList, name)
		tplList = append(tplList, t.config.Partials...)

		tpl = template.New(name).Funcs(allFunctions).Delims(t.config.Delims.Left, t.config.Delims.Right)
		for _, v := range tplList {
			var data string
			data, err = t.fileHandler(t.config, v)
			if err != nil {
				return err
			}
			var tmpl *template.Template
			if v == name {
				tmpl = tpl
			} else {
				tmpl = tpl.New(v)
			}

			_, err = tmpl.Parse(data)
			if err != nil {
				return fmt.Errorf("TemplateEngine render parser name:%v, error: %v", v, err)
			}
		}
		t.tmplMutex.Lock()
		t.tmplMap[name] = tpl
		t.tmplMutex.Unlock()
	}

	err = tpl.Funcs(allFunctions).ExecuteTemplate(out, exeName, data)
	if err != nil {
		return fmt.Errorf("TemplateEngine execute template error: %v", err)
	}

	return nil
}

// SetFileHandler : FileHandler関数を設定します。
func (t *TmplEngine) SetFileHandler(handle FileHandler) {
	if handle == nil {
		panic("FileHandler can't set nil!")
	}
	t.fileHandler = handle
}

// DefaultFileHandler : FileHandlerが設定されない場合のデフォルトのハンドラ関数です。
func DefaultFileHandler() FileHandler {
	return func(config TmplConfig, tplFile string) (content string, err error) {
		// Get the absolute path of the root template
		path, err := filepath.Abs(config.Root + string(os.PathSeparator) + tplFile + config.Extension)
		if err != nil {
			return "", fmt.Errorf("TemplateEngine path: %v error %v", path, err)
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("TemplateEngine render read name :%v, path: %v, error :%v", tplFile, path, err)
		}

		return string(data), nil
	}
}