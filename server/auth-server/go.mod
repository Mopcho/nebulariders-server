module github.com/Mopcho/nebulariders-server/auth

go 1.21.6

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
)

require github.com/joho/godotenv v1.5.1

require golang.org/x/crypto v0.25.0

require github.com/Mopcho/nebulariders-server/common/mopHttp v0.0.0

replace github.com/Mopcho/nebulariders-server/common/mopHttp => ../common/mopHttp