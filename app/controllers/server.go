package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/acmk189/golang_udemy_todo_app/app/models"
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

// 受け取ったCookieがDBに存在するか判定
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// URLとハンドラ(レスポンス処理)の紐付け -> URLを登録
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	// nil はデフォルトのマルチプレクサ -> page not found を表示
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
