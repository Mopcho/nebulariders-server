.PHONY: auth-server game-server run-all

auth-server:
	cd auth-server && go run main.go

game-server:
	cd game-server && go run main.go player.go

run-all:
	@trap 'kill 0' SIGINT; \
	cd auth-server && go run main.go types.go utils.go & \
	cd game-server && go run main.go player.go & \
	wait