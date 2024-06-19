package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	// もしmap型の変数queryMapが文字列"page"をキーに持っているのであれば、pにはpage キーに対応する値 queryMap["page"] が、ok には true が格納される
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		// パラメータ page に対応する 1 つ目の値を採用し、それを数値に変換する
		var err error
		page, err = strconv.Atoi(p[0])

		// 数値に変換できない値だった場合は 400 番エラーを返す
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
		// パラメータ page が存在しなかった場合
	} else {
		// パラメータ page=1 がついていたときと同じ処理をしたい
		page = 1
	}

	resString := fmt.Sprintf("Article List (page: %d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		// クエリパラメータが数字でない場合は、400 Bad Requestを返す
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
