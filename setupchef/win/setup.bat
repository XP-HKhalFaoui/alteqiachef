@echo off
REM AlteqiaChef Setup Script for Windows (CMD)
REM Run: setup.bat <command>
REM Commands: setup, start, stop, restart, status, logs, clean

setlocal enabledelayedexpansion

REM Color codes for output
set "RED=[91m"
set "GREEN=[92m"
set "YELLOW=[93m"
set "BLUE=[94m"
set "RESET=[0m"

REM Default action
set "ACTION=%1"
if "%ACTION%"=="" set "ACTION=help"

cls
goto %ACTION%

:help
echo.
echo %BLUE%╔════════════════════════════════════════════════════════╗%RESET%
echo %BLUE%║         AlteqiaChef Setup Script for Windows           ║%RESET%
echo %BLUE%╚════════════════════════════════════════════════════════╝%RESET%
echo.
echo %YELLOW%USAGE:%RESET%
echo.
echo   setup.bat ^<command^>
echo.
echo %YELLOW%COMMANDS:%RESET%
echo.
echo   setup      : Complete setup (pull images, start containers)
echo   start      : Start containers
echo   stop       : Stop containers
echo   restart    : Restart all containers
echo   status     : Show service status
echo   logs       : Stream logs from services
echo   clean      : Remove containers and volumes
echo   help       : Show this message
echo.
echo %YELLOW%EXAMPLES:%RESET%
echo.
echo   setup.bat setup
echo   setup.bat status
echo   setup.bat logs
echo.
echo %YELLOW%URLS:%RESET%
echo.
echo   Frontend: http://localhost:3000
echo   Backend:  http://localhost:8080
echo   Database: localhost:5432
echo.
goto end

:setup
cls
echo.
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %BLUE%           AlteqiaChef Setup - Starting              %RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.

echo %YELLOW%[*] Checking Docker installation...%RESET%
docker --version >nul 2>&1
if %errorlevel% neq 0 (
    echo %RED%[X] Docker is not installed or not in PATH%RESET%
    echo.
    call :install-docker-chocolatey
    goto end
)
echo %GREEN%[OK] Docker found%RESET%

docker-compose --version >nul 2>&1
if %errorlevel% neq 0 (
    echo %RED%[X] Docker Compose is not installed%RESET%
    goto end
)
echo %GREEN%[OK] Docker Compose found%RESET%

echo.
echo %YELLOW%[*] Pulling latest Docker images...%RESET%
cd setupchef
docker-compose pull
if %errorlevel% neq 0 (
    echo %RED%[X] Failed to pull images%RESET%
    cd ..
    goto end
)

echo.
echo %YELLOW%[*] Starting containers...%RESET%
docker-compose up -d
if %errorlevel% neq 0 (
    echo %RED%[X] Failed to start containers%RESET%
    cd ..
    goto end
)

echo.
echo %YELLOW%[*] Waiting for services to start...%RESET%
timeout /t 5 /nobreak

echo.
echo %YELLOW%[*] Checking service status...%RESET%
docker-compose ps

cd ..

echo.
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %GREEN%Setup Complete! ^^!%RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.
echo %GREEN%[OK] Access your application:%RESET%
echo     Frontend: http://localhost:3000
echo     Backend:  http://localhost:8080
echo     Database: localhost:5432
echo.
echo %GREEN%[OK] Next steps:%RESET%
echo     1. Open http://localhost:3000 in your browser
echo     2. Create an admin user
echo     3. Start managing your restaurant!
echo.
goto end

:start
cls
echo.
echo %BLUE%Starting AlteqiaChef...%RESET%
echo.
cd setupchef
docker-compose up -d
docker-compose ps
cd ..
goto end

:stop
cls
echo.
echo %BLUE%Stopping AlteqiaChef...%RESET%
echo.
cd setupchef
docker-compose down
echo %GREEN%[OK] All containers stopped%RESET%
cd ..
goto end

:restart
cls
echo.
echo %BLUE%Restarting AlteqiaChef...%RESET%
echo.
cd setupchef
docker-compose restart
docker-compose ps
cd ..
goto end

:status
cls
echo.
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %BLUE%                   Service Status                       %RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.
cd setupchef
docker-compose ps
echo.
echo %BLUE%Service Health:%RESET%
echo.

REM Check Frontend
for /f %%i in ('powershell -Command "try { (Invoke-WebRequest -Uri 'http://localhost:3000' -TimeoutSec 2 -SkipHttpStatusCodeCheck).StatusCode; exit 0 } catch { exit 1 }" 2^>nul') do set "FRONTEND=%%i"
if "%FRONTEND%"=="200" (
    echo %GREEN%[OK] Frontend ^(3000^): Running%RESET%
) else (
    echo %RED%[X] Frontend ^(3000^): Not responding%RESET%
)

REM Check Backend
for /f %%i in ('powershell -Command "try { (Invoke-WebRequest -Uri 'http://localhost:8080/api/v1/health' -TimeoutSec 2 -SkipHttpStatusCodeCheck).StatusCode; exit 0 } catch { exit 1 }" 2^>nul') do set "BACKEND=%%i"
if "%BACKEND%"=="200" (
    echo %GREEN%[OK] Backend ^(8080^): Running%RESET%
) else (
    echo %RED%[X] Backend ^(8080^): Not responding%RESET%
)

echo %GREEN%[OK] Database ^(5432^): Check container logs%RESET%
echo.
cd ..
goto end

:logs
cls
echo.
echo %BLUE%Application Logs%RESET%
echo %YELLOW%Press Ctrl+C to stop viewing logs%RESET%
echo.
cd setupchef
docker-compose logs -f
cd ..
goto end

:clean
cls
echo.
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %RED%                   WARNING: Cleanup                       %RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.
echo This will:
echo   - Stop all containers
echo   - Remove containers and networks
echo   - Remove volumes ^(DATABASE WILL BE LOST!^)
echo.
set /p "CONFIRM=Are you sure? Type 'yes' to confirm: "
if /i "%CONFIRM%"=="yes" (
    echo.
    echo %YELLOW%[*] Removing containers and volumes...%RESET%
    cd setupchef
    docker-compose down -v
    cd ..
    echo %GREEN%[OK] Cleanup complete%RESET%
) else (
    echo %RED%[X] Cleanup cancelled%RESET%
)
echo.
goto end

:end
endlocal

:install-docker-chocolatey
echo.
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %BLUE%         Chocolatey and Docker Installation Required     %RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.
echo %RED%[X] Docker is not installed or not in PATH%RESET%
echo.
echo %YELLOW%Required Software:%RESET%
echo.
echo 1. %YELLOW%Install Chocolatey (Package Manager)%RESET%
echo    Run the installation script:
echo    install-chocolatey.bat
echo.
echo 2. %YELLOW%Install Docker Desktop using Chocolatey%RESET%
echo    After installing Chocolatey, run:
echo    choco install docker-desktop -y
echo.
echo 3. %YELLOW%Restart your computer%RESET%
echo.
echo 4. %YELLOW%Run this script again%RESET%
echo    setup.bat setup
echo.
echo %YELLOW%Or install manually from:%RESET%
echo    https://chocolatey.org/install
echo.
echo %YELLOW%Press any key to exit...%RESET%
pause >nul
exit /b 0

:end
endlocal
