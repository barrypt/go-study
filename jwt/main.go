package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type myClaims struct {
	UserNmae string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	jwt_key := []byte("yingxiaozhu") // 加密key
	// 加密一个token
	claims := myClaims{
		UserNmae: "qpt",
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Nanosecond)), // 一分钟之前开始生效
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),       // 两个小时后失效
			Issuer:    "签发人",                                               // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(jwt_key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(token_string) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IumiluWwj-S4uyIsImV4cCI6MTY0OTk1NDkzOSwiaXNzIjoi562-5Y-R5Lq6IiwibmJmIjoxNjQ5OTQ3Njc5fQ.d8a24gGacP7Af_zy2NdUJvGO1-rHJENZXzV3dA_AESA

	// 解密token
	parseToken, err := jwt.ParseWithClaims(token_string, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt_key, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	bys, _ := json.Marshal(parseToken.Claims.(*myClaims))
	fmt.Println(string(bys)) //

	rasToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	sigStr, _ := rasToken.SigningString()
	fmt.Println(sigStr)

	key, _ := os.ReadFile("./key")
	rasKey, _ := jwt.ParseRSAPrivateKeyFromPEM(key)

	token1, _ := rasToken.SignedString(rasKey)
	fmt.Println("rasToken", token1)
	pubFile, _ := os.ReadFile("./pub")
	pub, _ := jwt.ParseRSAPublicKeyFromPEM(pubFile)
	// 解密token
	parseToken1, err1 := jwt.ParseWithClaims(token1, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err1 != nil {
		fmt.Println(err.Error())
	}
	bys1, _ := json.Marshal(parseToken1.Claims.(*myClaims))
	fmt.Println(string(bys1)) //

}
