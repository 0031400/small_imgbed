package auth

import (
	"encoding/base64"
	"small_imgbed/config"
)

func Auth(authHeader string) bool {
	return authHeader == "Basic "+base64.StdEncoding.EncodeToString([]byte(config.C.Auth.Username+":"+config.C.Auth.Password))
}
