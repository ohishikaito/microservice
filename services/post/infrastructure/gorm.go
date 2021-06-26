package infrastructure

import (
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConnect() *gorm.DB {
	var database string
	database = "root:finder0501@tcp(user_db_1)/microservices_user_development?charset=utf8&parseTime=true&Asia%2FTokyo"

	db, err := gorm.Open("mysql", database)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ログファイルを開く。CREATE=作成 WRONLYで読み書き APPENDで後ろに追加
	logfile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	// SetOutPutで出力先を指定 MultiWriterで2つの出力先を指定できる
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	// ログを吐いた時間を出す
	log.SetFlags(log.Ldate | log.Ltime)
	// LogMode true でログを吐き出すように
	db.LogMode(true)
	// 標準出力で出力するように。SetLogger先は一つしか選べないっぽい
	db.SetLogger(log.New(os.Stdout, "", 0))
	return db
}
