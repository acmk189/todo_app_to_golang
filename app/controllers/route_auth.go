package controllers

import (
	"log"
	"net/http"

	"github.com/acmk189/golang_udemy_todo_app/app/models"
)

// signupのハンドラ作成
func signup(w http.ResponseWriter, r *http.Request) {
	// HTTPリクエストメソッドによって処理を分ける
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		//　入力フォームから取得したデータを解析(暗記)
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		user := models.User{
			// PostFormValue でinput_typeをもとに取得
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		// topページにリダイレクト
		http.Redirect(w, r, "/", 302)
	}
}

// loginのハンドラ作成
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// authenticateのハンドラ作成
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		// エラーの場合は認証失敗 -> ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		// パスワード認証成功 -> セッション作成
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		// cookie情報作成(パターンとして覚える)
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		// ログイン成功後リダイレクト(ログイン後ページないので一旦top)
		http.Redirect(w, r, "/", 302)
	} else {
		// パスワード認証失敗
		http.Redirect(w, r, "/login", 302)
	}
}

func logout(w http.ResponseWriter, h *http.Request) {
	cookie, err := h.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, h, "/login", 302)
}
