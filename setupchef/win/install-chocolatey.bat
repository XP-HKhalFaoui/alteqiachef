@echo off
REM AlteqiaChef - Chocolatey Installation Script for Windows
REM Run: install-chocolatey.bat
REM This script installs Chocolatey package manager

setlocal enabledelayedexpansion

REM Color codes for output
set "RED=[91m"
set "GREEN=[92m"
set "YELLOW=[93m"
set "BLUE=[94m"
set "RESET=[0m"

cls
echo.
echo %BLUE%╔════════════════════════════════════════════════════════╗%RESET%
echo %BLUE%║      AlteqiaChef - Chocolatey Installation Script      ║%RESET%
echo %BLUE%╚════════════════════════════════════════════════════════╝%RESET%
echo.

REM Check if running as administrator
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo %RED%[X] ERROR: This script must be run as Administrator%RESET%
    echo.
    echo %YELLOW%To run as Administrator:%RESET%
    echo 1. Press Win+X
    echo 2. Select "Windows Terminal (Admin)" or "Command Prompt (Admin)"
    echo 3. Navigate to the script directory
    echo 4. Run: install-chocolatey.bat
    echo.
    echo %YELLOW%Press any key to exit...%RESET%
    pause >nul
    exit /b 1
)

echo %GREEN%[OK] Running as Administrator%RESET%
echo.

REM Check if Chocolatey is already installed
echo %YELLOW%[*] Checking if Chocolatey is already installed...%RESET%
choco --version >nul 2>&1
if %errorlevel% equ 0 (
    echo %GREEN%[OK] Chocolatey is already installed!%RESET%
    echo.
    for /f "tokens=*" %%i in ('choco --version') do set "CHOCO_VERSION=%%i"
    echo %GREEN%Version: !CHOCO_VERSION!%RESET%
    echo.
    echo %YELLOW%Press any key to continue...%RESET%
    pause >nul
    exit /b 0
)

echo %RED%[X] Chocolatey not found%RESET%
echo.

REM Check PowerShell ExecutionPolicy
echo %YELLOW%[*] Checking PowerShell ExecutionPolicy...%RESET%
for /f %%i in ('powershell Get-ExecutionPolicy') do set "EXEC_POLICY=%%i"
echo %YELLOW%Current Policy: !EXEC_POLICY!%RESET%

if "!EXEC_POLICY!"=="Restricted" (
    echo %YELLOW%[*] Setting ExecutionPolicy to AllSigned...%RESET%
    powershell -Command "Set-ExecutionPolicy AllSigned" >nul 2>&1
    echo %GREEN%[OK] ExecutionPolicy updated%RESET%
) else (
    echo %GREEN%[OK] ExecutionPolicy is compatible%RESET%
)
echo.

REM Install Chocolatey
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %YELLOW%[*] Installing Chocolatey Package Manager...%RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.

@powershell -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"

if %errorlevel% neq 0 (
    echo.
    echo %RED%════════════════════════════════════════════════════════%RESET%
    echo %RED%[X] Installation Failed!%RESET%
    echo %RED%════════════════════════════════════════════════════════%RESET%
    echo.
    echo %YELLOW%Troubleshooting steps:%RESET%
    echo.
    echo 1. %YELLOW%Check for existing installation:%RESET%
    echo    dir C:\ProgramData\chocolatey
    echo.
    echo 2. %YELLOW%If the folder exists, backup and remove it:%RESET%
    echo    move C:\ProgramData\chocolatey C:\ProgramData\chocolatey.bak
    echo.
    echo 3. %YELLOW%Try again:%RESET%
    echo    install-chocolatey.bat
    echo.
    echo 4. %YELLOW%For more help, visit:%RESET%
    echo    https://chocolatey.org/install
    echo.
    echo %YELLOW%Press any key to exit...%RESET%
    pause >nul
    exit /b 1
)

echo.
echo %GREEN%════════════════════════════════════════════════════════%RESET%
echo %GREEN%[OK] Chocolatey Installation Successful!%RESET%
echo %GREEN%════════════════════════════════════════════════════════%RESET%
echo.

REM Verify installation
echo %YELLOW%[*] Verifying installation...%RESET%
choco --version >nul 2>&1
if %errorlevel% equ 0 (
    for /f "tokens=*" %%i in ('choco --version') do set "CHOCO_VERSION=%%i"
    echo %GREEN%[OK] Version: !CHOCO_VERSION!%RESET%
) else (
    echo %YELLOW%[*] Please restart your terminal and run: choco --version%RESET%
)
echo.

REM Install Docker Desktop automatically
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %YELLOW%[*] Installing Docker Desktop automatically...%RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.

choco install docker-desktop -y --no-progress
if %errorlevel% neq 0 (
    echo.
    echo %RED%[X] Docker installation failed!%RESET%
    echo.
    echo %YELLOW%Try running manually:%RESET%
    echo    choco install docker-desktop -y
    echo.
    echo %YELLOW%Press any key to exit...%RESET%
    pause >nul
    endlocal
    exit /b 1
)

echo.
echo %GREEN%[OK] Docker Desktop installation started!%RESET%
echo.

REM Auto-restart the computer
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo %YELLOW%[*] Restarting your computer in 30 seconds...%RESET%
echo %BLUE%════════════════════════════════════════════════════════%RESET%
echo.
echo %YELLOW%To cancel the restart, run:%RESET%
echo    shutdown /a
echo.
echo %YELLOW%Restarting in:%RESET%

for /L %%i in (30,-1,1) do (
    echo %%i seconds...
    timeout /t 1 /nobreak >nul
)

echo.
echo %YELLOW%[*] Restarting now...%RESET%
shutdown /r /t 0 /c "Docker installation complete. Restarting to apply changes..."

endlocal
exit /b 0
