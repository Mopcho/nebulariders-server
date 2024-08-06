@echo off
setlocal

if "%1" == "--build" (
    call "%~dp0build.bat"
    if %ERRORLEVEL% NEQ 0 (
        echo Build failed, not starting the server.
        exit /b %ERRORLEVEL%
    )
)

echo Starting the game server...
bin\game-server.exe

endlocal