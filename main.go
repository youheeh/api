package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuiexample/database"
	"github.com/yuiexample/domain/blog"
)

func main() {
	// init MySQL database [MySQL databaseの初期化]
	db, err := database.InitDatabase()
	// If an error occurs, log it and stop the server [エラーが発生した場合は、ログに記録してサーバーを停止します]
	if err != nil {
		log.Fatal(err)
	}
	// close database connections when application stops [アプリケーション停止時にデータベース接続を閉じる]
	defer db.Close()

	// init mux router [mux routerの初期化]
	r := mux.NewRouter()

	// get list of blog [blogのリストを取得する]
	r.HandleFunc("/v1/blogs", blog.List).Methods(http.MethodGet)
	// create a blog [blogを作成する]
	r.HandleFunc("/v1/blogs", blog.Create).Methods(http.MethodPost)
	// get a blog by given {id} [指定された {id} でblogを取得する]
	r.HandleFunc("/v1/blogs/{id}", blog.GetByID).Methods(http.MethodGet)
	// update a blog by given {id} [指定された {id} でblogを更新します]
	r.HandleFunc("/v1/blogs/{id}", blog.UpdateByID).Methods(http.MethodPut)
	// delete a blog by given {id} [指定された {id} でblogを削除する]
	r.HandleFunc("/v1/blogs/{id}", blog.DeleteByID).Methods(http.MethodDelete)
	//[GET] /v1/shops
	r.HandleFunc("/v1/shops", blog.ShopsList).Methods(http.MethodGet)

	// start the server with port 8080 [ポート 8080 でサーバーを起動します]
	err = http.ListenAndServe(":8080", r)
	// If an error occurs, log it and stop the server [エラーが発生した場合は、ログに記録してサーバーを停止します]
	if err != nil {
		log.Fatal(err)
	}
}
