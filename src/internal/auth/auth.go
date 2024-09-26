package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"momonga_blog/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Clime struct {
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Exp int64  `json:"exp"`
	Alg string `json:"alg" default:"HS256"`
	Typ string `json:"typ" default:"JWT"`
}

var stretchCost = 12
var alg = "HS256"
var accessTokenExp = 12 * time.Hour

// パスワードのハッシュ化
func HashPassword(planePassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(planePassword), stretchCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// パスワード比較
func ComparePassword(hashedPassword, planePassword string) error {
	result := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(planePassword))
	if result != nil {
		return result
	}
	return nil
}

// リフレッシュトークン作成
func CreateRefreshToken(length int) (string, error) {
	token := make([]byte, length)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(token), nil
}

// フレッシュトークンの有効期限作成 CreateRefreshTokenExpire(30 * 24 * time.Hour)
func CreateRefreshTokenExpire(days int) time.Time {
	return time.Now().Add(time.Duration(days) * 24 * time.Hour)
}

// アクセストークン作成
func CreateAccessToken(clime Clime) (string, error) {
	// jwt作成
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": clime.Sub,
		"aud": clime.Aud,
		"exp": clime.Exp,
		"alg": clime.Alg,
		"typ": clime.Typ,
	})
	cnf, err := config.GetConfig()
	if err != nil {
		return "", err
	}
	tokenString, err := jwt.SignedString(cnf.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// クレーム作成
func NewClime(uuid string) Clime {
	return Clime{
		Sub: uuid,
		Aud: "admin",
		Exp: time.Now().Add(accessTokenExp).Unix(),
	}
}

// アクセストークン解析
func parseAccessToken(tokenString string) (*jwt.Token, error) {
	cnf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return cnf.SecretKey, nil
	})
}

// アクセストークンの認証
func AuthAccessToken(tokenString string) (string, error) {
	token, err := parseAccessToken(tokenString)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", fmt.Errorf("signature invalid")
		} else if err == jwt.ErrTokenExpired {
			return "", fmt.Errorf("token is expired")
		} else {
			return "", fmt.Errorf("token is invalid")
		}
	}
	claims := token.Claims.(jwt.MapClaims)

	// user取得に変更(今はuuidを返す)
	return claims["sub"].(string), nil
}

