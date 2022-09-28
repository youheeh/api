package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// set DB as a global variable for MySQL connection that can be used on another package
// 別のパッケージで使用できる MySQL 接続のグローバル変数として DB を設定します
var DB *sql.DB

func InitDatabase() (*sql.DB, error) {
	var err error

	// open MySQL connection [MySQL接続を開く]
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mission")
	// returns an error value if an error occurs [エラーが発生した場合は、エラー値を返します]
	if err != nil {
		return nil, err
	}

	return DB, nil
}
