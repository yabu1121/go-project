//usecaseはinterfaceに依存する

package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	// interfaceはメソッドの一覧 
	// () (a,b) ってなっているのは引数と返り値の型
	// a, bを返す
	// errorはエラーが起きた時に返す
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// NewUserUsecaseはuserUsecaseのコンストラクタ
// 引数にIUserRepositoryを受け取り、IUserUsecaseを返す
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}


func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	// bcrypt.GenerateFromPasswordはパスワードをハッシュ化する
	// 10 は　暗号化のコスト(数字が大きいほど暗号化に時間がかかる)
	// 返り値としてハッシュ化されたパスワードとエラーが返ってくる
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10); 
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		Email: user.Email,
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID :newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}


func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// jwtとはJSON Web Tokenの略で、何らかの情報を安全にやり取りするための仕組み
	// jwt.NewWithClaimsはjwtトークンを作成する
	// jwt.SigningMethodHS256はjwtトークンの署名方法
	// jwt.MapClaimsはjwtトークンのクレーム
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// user_idはstoredIdのIDを代入
		"user_id": storedUser.ID,
		// jwtトークンの期限をtime.Now()で今の時刻、Add()で追加、add(time.Hour * 12)で12時間後、Unix()でUNIXタイムスタンプに変換
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	// token.SignedString()はjwtトークンに署名する
	// []byte(os.Getenv("SECRET"))はjwtトークンの署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

