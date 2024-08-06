module github.com/Mopcho/nebulariders-server/game-server

go 1.21.6

require github.com/gorilla/websocket v1.5.3

require github.com/joho/godotenv v1.5.1

require github.com/Mopcho/nebulariders-server/common/mopHttp v0.0.0

require github.com/golang-jwt/jwt/v5 v5.2.1

replace github.com/Mopcho/nebulariders-server/common/mopHttp => ../common/mopHttp
