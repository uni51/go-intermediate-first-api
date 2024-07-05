package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/services"
)

func main() {
	// 環境変数から値を取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")

	// 接続文字列を作成
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// 1. サーバー全体で使用する sql.DB 型を一つ生成する
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	// 2. sql.DB型をもとに、サーバー全体で使用するサービス構造体MyAppServiceを一つ生成する
	ser := services.NewMyAppService(db)
	// 3. MyAppService型をもとに、サーバー全体で使用するコントローラ構造体MyAppControllerを一つ生成する
	con := controllers.NewMyAppController(ser)

	// 4. コントローラ型 MyAppController のハンドラメソッドとパスとの関連付けを行う
	r := mux.NewRouter()

	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
