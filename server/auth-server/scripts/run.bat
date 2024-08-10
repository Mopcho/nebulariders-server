@echo off
setlocal

echo Starting the auth server...
go run main.go types.go utils.go validations.go

endlocal