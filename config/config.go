package config

import "os"

func JWTKey() string {
	secret := os.Getenv("SECRET_JWT")
	return secret
}
