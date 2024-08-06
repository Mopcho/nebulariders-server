.PHONY: auth-server game-server run-all

auth-server:
	@trap 'kill 0' SIGINT; \
	cd auth-server && go run main.go types.go utils.go validations.go
	wait

game-server:
	@trap 'kill 0' SIGINT; \
	cd game-server && go run main.go player.go types.go game.go

run-all:
	@trap 'kill 0' SIGINT; \
	cd auth-server && go run main.go types.go utils.go validations.go & \
	cd game-server && go run main.go player.go types.go game.go & \
	wait