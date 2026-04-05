#!/usr/bin/env pwsh
<#
.SYNOPSIS
    AlteqiaChef Setup Script for Windows (PowerShell)
    
.DESCRIPTION
    Automated setup script for the AlteqiaChef POS System on Windows.
    Pulls Docker images, starts containers, and initializes the database.
    
.PARAMETER Action
    Action to perform: 'setup', 'start', 'stop', 'restart', 'status', 'logs', 'clean'
    
.EXAMPLE
    .\setup.ps1 -Action setup
    .\setup.ps1 -Action status
    .\setup.ps1 -Action logs
    
.NOTES
    Requires: Docker Desktop with Docker Compose
    Author: AlteqiaChef Team
    Version: 1.0
#>

param(
    [Parameter(Mandatory=$false)]
    [ValidateSet('setup', 'start', 'stop', 'restart', 'status', 'logs', 'clean', 'help')]
    [string]$Action = 'help'
)

# Color definitions
$Colors = @{
    Reset   = "`e[0m"
    Bold    = "`e[1m"
    Blue    = "`e[34m"
    Green   = "`e[32m"
    Yellow  = "`e[33m"
    Red     = "`e[31m"
}

function Write-ColorOutput($message, $color = 'Reset') {
    Write-Host "$($Colors[$color])$message$($Colors.Reset)"
}

function Write-Header($title) {
    Write-Host ""
    Write-ColorOutput "╔════════════════════════════════════════════════════════╗" "Blue"
    Write-ColorOutput "║ $title.PadRight(55, ' ') ║" "Blue"
    Write-ColorOutput "╚════════════════════════════════════════════════════════╝" "Blue"
    Write-Host ""
}

function Check-Docker {
    Write-ColorOutput "🔍 Checking Docker installation..." "Yellow"
    
    try {
        $dockerVersion = docker --version
        Write-ColorOutput "✅ Docker found: $dockerVersion" "Green"
        return $true
    }
    catch {
        Write-ColorOutput "❌ Docker is not installed or not in PATH" "Red"
        return $false
    }
}

function Install-Docker {
    Write-Header "Docker Installation Required"
    
    Write-ColorOutput "Docker is required to run AlteqiaChef. Installing..." "Yellow"
    Write-Host ""
    
    # Check if running as administrator
    $isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")
    
    if (-not $isAdmin) {
        Write-ColorOutput "⚠️  This script needs to run as Administrator to install Docker" "Yellow"
        Write-ColorOutput "   Attempting to restart with elevated privileges..." "Yellow"
        Write-Host ""
        
        # Relaunch script as administrator
        $scriptPath = $MyInvocation.MyCommand.Path
        $params = "-Action `"setup`""
        Start-Process powershell -Verb RunAs -ArgumentList "-File `"$scriptPath`" $params"
        exit 0
    }
    
    # Check if Chocolatey is installed
    Write-ColorOutput "📦 Checking for Chocolatey package manager..." "Yellow"
    try {
        $chocoVersion = choco --version 2>$null
        if ($chocoVersion) {
            Write-ColorOutput "✅ Chocolatey found: $chocoVersion" "Green"
            Write-Host ""
            
            Write-ColorOutput "📥 Installing Docker Desktop via Chocolatey..." "Yellow"
            choco install docker-desktop -y --no-progress
            
            if ($LASTEXITCODE -eq 0) {
                Write-ColorOutput "✅ Docker Desktop installed successfully!" "Green"
                Write-ColorOutput "⏳ Please restart your computer and run setup again" "Yellow"
                exit 0
            }
        }
    }
    catch {
        Write-ColorOutput "ℹ️  Chocolatey not found, will use alternative method" "Yellow"
    }
    
    Write-Host ""
    Write-ColorOutput "📥 Alternative: Installing Docker Desktop..." "Yellow"
    Write-ColorOutput "   Downloading Docker Desktop installer..." "Yellow"
    
    $dockerInstallerUrl = "https://desktop.docker.com/win/main/amd64/Docker%20Desktop%20Installer.exe"
    $installerPath = "$env:TEMP\DockerInstaller.exe"
    
    try {
        # Download Docker installer
        $ProgressPreference = 'SilentlyContinue'
        Invoke-WebRequest -Uri $dockerInstallerUrl -OutFile $installerPath -UseBasicParsing
        
        Write-ColorOutput "✅ Docker Desktop installer downloaded" "Green"
        Write-ColorOutput "📦 Starting installation..." "Yellow"
        
        # Run installer
        & $installerPath install --quiet
        
        # Wait for installation
        Start-Sleep -Seconds 30
        
        Write-ColorOutput "✅ Docker Desktop installed!" "Green"
        Write-ColorOutput "⏳ Please restart your computer to complete installation" "Yellow"
        Write-ColorOutput "   Then run: .\setup.ps1 -Action setup" "Yellow"
        
        # Cleanup
        Remove-Item $installerPath -Force -ErrorAction SilentlyContinue
        exit 0
    }
    catch {
        Write-ColorOutput "❌ Failed to download Docker installer" "Red"
        Write-ColorOutput "   Please download manually from: https://www.docker.com/products/docker-desktop" "Yellow"
        Write-ColorOutput "   Then run this script again" "Yellow"
        exit 1
    }
}

function Verify-Docker {
    Write-ColorOutput "🔍 Verifying Docker and Docker Compose..." "Yellow"
    
    try {
        $composeVersion = docker-compose --version
        Write-ColorOutput "✅ Docker Compose found: $composeVersion" "Green"
        return $true
    }
    catch {
        Write-ColorOutput "❌ Docker Compose is not installed" "Red"
        return $false
    }
}

function Check-Ports {
    Write-ColorOutput "🔍 Checking required ports..." "Yellow"
    
    $ports = @(3000, 8080, 5432)
    $portsInUse = @()
    
    foreach ($port in $ports) {
        $connection = Test-NetConnection -ComputerName localhost -Port $port -WarningAction SilentlyContinue
        if ($connection.TcpTestSucceeded) {
            $portsInUse += $port
        }
    }
    
    if ($portsInUse.Count -gt 0) {
        Write-ColorOutput "⚠️  Port(s) already in use: $($portsInUse -join ', ')" "Yellow"
        Write-ColorOutput "   These may be from previous containers or other applications" "Yellow"
        Write-ColorOutput "   Run 'docker-compose down' to stop running containers" "Yellow"
    }
    else {
        Write-ColorOutput "✅ All required ports are available" "Green"
    }
}

function Setup-Application {
    Write-Header "AlteqiaChef Setup"
    
    # Check and install Docker if needed
    if (-not (Check-Docker)) {
        Write-Host ""
        $installDocker = Read-Host "Docker not found. Would you like to install it? (y/n)"
        
        if ($installDocker -eq 'y' -or $installDocker -eq 'yes') {
            Install-Docker
        }
        else {
            Write-ColorOutput "❌ Docker is required. Exiting." "Red"
            exit 1
        }
    }
    
    # Verify Docker Compose
    if (-not (Verify-Docker)) {
        Write-ColorOutput "❌ Docker Compose is required but not found" "Red"
        exit 1
    }
    
    Check-Ports
    
    Write-ColorOutput "`n📥 Pulling latest Docker images..." "Yellow"
    docker-compose pull
    
    if ($LASTEXITCODE -ne 0) {
        Write-ColorOutput "❌ Failed to pull images" "Red"
        exit 1
    }
    
    Write-ColorOutput "`n🚀 Starting containers..." "Yellow"
    docker-compose up -d
    
    if ($LASTEXITCODE -ne 0) {
        Write-ColorOutput "❌ Failed to start containers" "Red"
        exit 1
    }
    
    Write-ColorOutput "`n⏳ Waiting for services to start..." "Yellow"
    Start-Sleep -Seconds 5
    
    Write-ColorOutput "`n✅ Checking service status..." "Yellow"
    docker-compose ps
    
    Write-Header "Setup Complete! 🎉"
    
    Write-ColorOutput "📍 Access your application:" "Green"
    Write-ColorOutput "   Frontend: http://localhost:3000" "Green"
    Write-ColorOutput "   Backend:  http://localhost:8080" "Green"
    Write-ColorOutput "   Database: localhost:5432" "Green"
    
    Write-ColorOutput "`n📝 Next steps:" "Green"
    Write-ColorOutput "   1. Open http://localhost:3000 in your browser" "Green"
    Write-ColorOutput "   2. Create an admin user (see SETUP.md)" "Green"
    Write-ColorOutput "   3. Start managing your restaurant!" "Green"
    
    Write-Host ""
}

function Start-Application {
    Write-Header "Starting AlteqiaChef"
    
    docker-compose up -d
    Write-Host ""
    docker-compose ps
}

function Stop-Application {
    Write-Header "Stopping AlteqiaChef"
    
    docker-compose down
    Write-ColorOutput "✅ All containers stopped" "Green"
}

function Restart-Application {
    Write-Header "Restarting AlteqiaChef"
    
    docker-compose restart
    Write-Host ""
    docker-compose ps
}

function Show-Status {
    Write-Header "Service Status"
    
    docker-compose ps
    
    Write-ColorOutput "`n📊 Service Health:" "Blue"
    Write-Host ""
    
    # Check Frontend
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:3000" -TimeoutSec 2 -SkipHttpStatusCodeCheck
        Write-ColorOutput "✅ Frontend (3000): Running" "Green"
    }
    catch {
        Write-ColorOutput "❌ Frontend (3000): Not responding" "Red"
    }
    
    # Check Backend
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:8080/api/v1/health" -TimeoutSec 2 -SkipHttpStatusCodeCheck
        Write-ColorOutput "✅ Backend (8080): Running" "Green"
    }
    catch {
        Write-ColorOutput "❌ Backend (8080): Not responding" "Red"
    }
    
    # Check Database
    try {
        $response = docker exec pos-postgres pg_isready -U postgres 2>$null
        if ($LASTEXITCODE -eq 0) {
            Write-ColorOutput "✅ Database (5432): Running" "Green"
        }
        else {
            Write-ColorOutput "⚠️  Database (5432): Container exists but check connection" "Yellow"
        }
    }
    catch {
        Write-ColorOutput "❌ Database (5432): Not responding" "Red"
    }
    
    Write-Host ""
}

function Show-Logs {
    Write-Header "Application Logs"
    Write-ColorOutput "Press Ctrl+C to stop viewing logs" "Yellow"
    Write-Host ""
    docker-compose logs -f
}

function Clean-Application {
    Write-Header "Cleanup"
    
    Write-ColorOutput "⚠️  This will:" "Yellow"
    Write-ColorOutput "   • Stop all containers" "Yellow"
    Write-ColorOutput "   • Remove containers and networks" "Yellow"
    Write-ColorOutput "   • Remove volumes (DATABASE WILL BE LOST!)" "Red"
    Write-Host ""
    
    $confirm = Read-Host "Are you sure? Type 'yes' to confirm"
    
    if ($confirm -eq 'yes') {
        Write-ColorOutput "`n🗑️  Removing containers and volumes..." "Yellow"
        docker-compose down -v
        Write-ColorOutput "✅ Cleanup complete" "Green"
    }
    else {
        Write-ColorOutput "❌ Cleanup cancelled" "Red"
    }
}

function Show-Help {
    Write-Header "AlteqiaChef Setup Script Help"
    
    Write-ColorOutput "USAGE:" "Bold"
    Write-Host "  .\setup.ps1 -Action <action>`n"
    
    Write-ColorOutput "ACTIONS:" "Bold"
    Write-Host "  setup      : Complete setup (pull images, start containers, initialize db)"
    Write-Host "  start      : Start containers"
    Write-Host "  stop       : Stop containers"
    Write-Host "  restart    : Restart all containers"
    Write-Host "  status     : Show service status"
    Write-Host "  logs       : Stream logs from all services"
    Write-Host "  clean      : Remove containers and volumes (WARNING: loses data)"
    Write-Host "  help       : Show this help message"
    Write-Host ""
    
    Write-ColorOutput "EXAMPLES:" "Bold"
    Write-Host "  .\setup.ps1 -Action setup"
    Write-Host "  .\setup.ps1 -Action status"
    Write-Host "  .\setup.ps1 -Action logs"
    Write-Host ""
    
    Write-ColorOutput "LOCATIONS:" "Bold"
    Write-Host "  Frontend: http://localhost:3000"
    Write-Host "  Backend:  http://localhost:8080"
    Write-Host "  Database: localhost:5432"
    Write-Host ""
}

# Main execution
switch ($Action) {
    'setup' {
        Setup-Application
    }
    'start' {
        Start-Application
    }
    'stop' {
        Stop-Application
    }
    'restart' {
        Restart-Application
    }
    'status' {
        Show-Status
    }
    'logs' {
        Show-Logs
    }
    'clean' {
        Clean-Application
    }
    'help' {
        Show-Help
    }
    default {
        Show-Help
    }
}
