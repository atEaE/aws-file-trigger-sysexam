/* [ユーザー]
 *
 */
CREATE TABLE users (
    id SERIAL PRIMARY KEY,           --: ユーザーID
    name VARCHAR(100) NOT NULL,      --: ユーザー名
    email VARCHAR(100) NOT NULL,     --: E-MAILアドレス
    password VARCHAR(100) NOT NULL,  --: パスワード
    token VARCHAR(100)               --: アクセストークン
);

/* [フライト]
 *
 */
CREATE TABLE flights (
    id SERIAL PRIMARY KEY,           --: ID
    flight_date DATE NOT NULL,    --: 便運行日
    carrier_cd CHAR(2) NOT NULL,     --: 便キャリアコード
    flight_no VARCHAR(4) NOT NULL,   --: 便番号
    origin_apo_cd CHAR(3) NOT NULL,  --: 発地空港
    dest_app_cd CHAR(3) NOT NULL,    --: 着地空港
    cancelled BOOLEAN,               --: キャンセルフラグ
    diverted BOOLEAN                 --: ダイバートフラグ
);