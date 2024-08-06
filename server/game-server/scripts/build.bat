@echo off
setlocal

echo Building the game server...
go build -o bin\game-server.exe ./cmd/game-server

if %ERRORLEVEL% NEQ 0 (
    echo Failed to build the game server.
    exit /b %ERRORLEVEL%
)

echo Build completed successfully.

endlocal