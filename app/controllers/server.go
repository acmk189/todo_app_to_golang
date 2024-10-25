package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/acmk189/golang_udemy_todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	// htmlファイルの define を明示的に指定
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// URLとハンドラ(レスポンス処理)の紐付け -> URLを登録
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	// nil はデフォルトのマルチプレクサ -> page not found を表示
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
