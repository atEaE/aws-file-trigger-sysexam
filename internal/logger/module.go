package logger

import "log"

// OutputLog : [ID]SETP : MSGの形式でログを出力をします。
func OutputLog(id, step string, msg ...string) {
	log.Printf("[%s]%s : %s", id, step, msg)
}
