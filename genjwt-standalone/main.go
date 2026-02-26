package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// appClaims 与 serp-server-seo 中 common/middleware/ssoAuth.go 的 AppClaims 保持一致
type appClaims struct {
	Uid int64 `json:"uid,omitempty"`
	jwt.RegisteredClaims
}

func main() {
	configYml := flag.String("c", "", "配置文件路径（如 serp-server-seo-master 下的 resources/config.dev.yaml）")
	userID := flag.Int64("u", 37, "用户 ID（sys_user.id）")
	flag.Parse()

	if *configYml == "" {
		fmt.Println("用法: genjwt-standalone -c <配置文件路径> [-u 用户ID]")
		fmt.Println("示例: genjwt-standalone -c D:\\seo\\serp-server-seo-master\\resources\\config.dev.yaml -u 37")
		flag.Usage()
		return
	}

	v := viper.New()
	v.SetConfigFile(*configYml)
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("读取配置失败: %v\n", err)
		return
	}

	signKey := v.GetString("jwt.sign-key")
	if signKey == "" {
		fmt.Println("错误: 配置中未找到 jwt.sign-key")
		return
	}
	expiresMin := v.GetInt("jwt.expires")
	if expiresMin <= 0 {
		expiresMin = 10080
	}
	issuer := v.GetString("jwt.issuer")
	subject := v.GetString("jwt.subject")

	now := time.Now()
	expiresAt := now.Add(time.Duration(expiresMin) * time.Minute)
	claims := appClaims{
		Uid: *userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(signKey))
	if err != nil {
		fmt.Printf("生成 Token 失败: %v\n", err)
		return
	}

	fmt.Println("将下方 Token 填入 Apifox Authorization，值为: Bearer <Token>")
	fmt.Println()
	fmt.Println(tokenStr)
}
