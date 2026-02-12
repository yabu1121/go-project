package controller

import (
	"crypto/rand"
	"encoding/hex"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// Goではコンストラクタは関数で定義する、どのような関数がコンストラクタとし扱われるのかは決まっていないが、
// 一般的には以下のような特徴を持つ関数がコンストラクタとして扱われる
// 1. 型がUserControllerを返している
// 2. 引数にIUserUsecaseを受け取っている
// 3. userControllerのポインタを返している
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// SignUpはユーザーがサインアップする
func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

// Loginはユーザーがログインする
// cookieにtokenを保存する
// cookieとはブラウザに保存される小さなテキストファイル
// これに保存することで一時的な認証情報を保存することができる
func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// cookie構造体を初期化
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// Logoutはユーザーがログアウトする
// cookieを削除する
func (uc *userController) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	// Echo の CSRF ミドルウェアのコンテキストに依存せず、
	// 自前でトークンを生成して CSRF 用クッキーとレスポンスを返す。

	// ランダムな 32 バイトのトークンを生成（64 文字の hex 文字列になる）
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate csrf token"})
	}
	token := hex.EncodeToString(b)

	// Echo の CSRF ミドルウェアがデフォルトで期待するクッキー名と同じ "_csrf" を使用
	cookie := &http.Cookie{
		Name:     "_csrf",
		Value:    token,
		Path:     "/",
		Domain:   "", // API ドメインに縛らない（localhost 環境向け）
		HttpOnly: true,
		// CSRF ミドルウェアのデフォルトと合わせて SameSite は指定しない（Default）
		// ローカル HTTP で動かすため Secure は false のまま
	}
	http.SetCookie(c.Response(), cookie)

	return c.JSON(http.StatusOK, map[string]string{"csrf_token": token})
}