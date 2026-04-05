# 🐳 AlteqiaChef - Docker Installation Guide

This guide helps you install Docker before running AlteqiaChef.

---

## 📋 Quick Installation

### Windows PowerShell (Automatic)
```powershell
cd alteqiachef
.\setup.ps1 -Action setup
# Script will detect Docker and install it automatically if needed
```

### Windows CMD (Automatic)
```cmd
cd alteqiachef
setup.bat setup
# Script will detect Docker and install it automatically if needed
```

---

## 🔄 Installation Methods

### Method 1: Automatic Installation (Recommended)

The setup scripts now **automatically detect and install Docker** if it's not found.

**PowerShell:**
```powershell
.\setup.ps1 -Action setup
# Choose 'y' when prompted to install Docker
```

**CMD:**
```cmd
setup.bat setup
# Installation will start automatically if Docker is not found
```

The script will:
1. ✅ Check if Docker is installed
2. ✅ If not found, offer to install it
3. ✅ Download Docker Desktop installer
4. ✅ Run the installer automatically
5. ✅ Prompt for computer restart
6. ✅ Continue with POS setup after restart

---

### Method 2: Manual Installation

#### Windows

1. **Download Docker Desktop:**
   - Visit: https://www.docker.com/products/docker-desktop
   - Click "Download for Windows"
   - Wait for download to complete

2. **Run the installer:**
   - Double-click `DockerDesktopInstaller.exe`
   - Follow the installation wizard
   - Accept default options
   - Click "Install"

3. **Restart your computer:**
   - Restart is required to complete installation
   - Save any open work first

4. **Verify installation:**
   ```powershell
   docker --version
   docker-compose --version
   ```

5. **Run AlteqiaChef setup:**
   ```powershell
   .\setup.ps1 -Action setup
   ```

#### macOS

1. **Using Homebrew (Easiest):**
   ```bash
   brew install --cask docker
   ```

2. **Or download manually:**
   - Visit: https://www.docker.com/products/docker-desktop
   - Download for macOS (Intel or Apple Silicon)
   - Open the .dmg file
   - Drag Docker to Applications folder
   - Launch Docker from Applications

3. **Verify installation:**
   ```bash
   docker --version
   ```

#### Linux

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get install docker.io docker-compose
sudo usermod -aG docker $USER
newgrp docker
```

**Fedora/CentOS:**
```bash
sudo yum install docker docker-compose
sudo systemctl start docker
sudo systemctl enable docker
sudo usermod -aG docker $USER
```

---

## ⚙️ Automatic Installation Details

### PowerShell Script Installation Process

The PowerShell script (`setup.ps1`) includes:

1. **Docker Detection:**
   - Checks if Docker is already installed
   - Returns installation status

2. **Docker Installation:**
   - Requests Administrator elevation if needed
   - Checks for Chocolatey package manager
   - Uses Chocolatey if available (faster)
   - Falls back to direct download if needed
   - Downloads Docker Desktop installer
   - Runs installer silently
   - Prompts for computer restart

3. **Installation Verification:**
   - After restart, script verifies Docker installation
   - Continues with POS setup if Docker is ready

### CMD Script Installation Process

The CMD script (`setup.bat`) includes:

1. **Docker Detection:**
   - Checks if Docker is installed
   - Checks for Docker Compose

2. **Docker Installation via Chocolatey:**
   - Detects Chocolatey installation
   - If missing, installs Chocolatey first
   - Installs Docker Desktop via Chocolatey
   - Provides manual installation link if needed

3. **Post-Installation:**
   - Prompts for computer restart
   - Shows instructions to run setup again

---

## ✅ System Requirements for Docker

Before installing Docker, ensure your system meets:

| Requirement | Details |
|-------------|---------|
| **Windows Version** | Windows 10 (build 19041+) or Windows 11 |
| **Processor** | 64-bit processor with virtualization support |
| **RAM** | Minimum 4GB (8GB recommended) |
| **Disk Space** | 5GB free space |
| **Virtualization** | Hyper-V or WSL 2 enabled |

### Check Windows Version

**PowerShell:**
```powershell
[System.Environment]::OSVersion.VersionString
```

**CMD:**
```cmd
wmic os get version
```

Should show: Windows 10 (19041+) or Windows 11

### Enable Hyper-V (Windows 10)

1. Open **Control Panel**
2. Select **Programs**
3. Click **Turn Windows features on or off**
4. Check **Hyper-V**
5. Click **OK**
6. Restart your computer

### Enable WSL 2 (Recommended)

**PowerShell as Administrator:**
```powershell
wsl --install
wsl --set-default-version 2
```

---

## 🔍 Troubleshooting Installation

### Installation Fails

**Problem:** Docker installation fails

**Solution:**
1. Run script as Administrator
2. Check Windows version: Must be 19041 or higher
3. Enable Hyper-V or WSL 2
4. Restart computer
5. Try again

### "Docker: not found" After Installation

**Problem:** Docker installed but not in PATH

**Solution:**
1. Restart your computer (full restart required)
2. Open new PowerShell/CMD window
3. Try again: `docker --version`

### Installation Hangs

**Problem:** Installation appears to be stuck

**Solution:**
1. Wait at least 5 minutes (first installation takes time)
2. Check if Docker Desktop is running in background
3. If still stuck, restart and try manual installation

### Insufficient Disk Space

**Problem:** Installation fails due to disk space

**Solution:**
1. Free up at least 5GB of disk space
2. Delete temporary files
3. Try installation again

### Administrator Privileges

**Problem:** Script says it needs Administrator access

**Solution:**
1. Right-click PowerShell and select "Run as administrator"
2. Run script again

---

## 📊 Installation Verification

### Check Installation Success

**PowerShell:**
```powershell
# Check Docker
docker --version
# Should output: Docker version XX.X.X, build XXXXX

# Check Docker Compose
docker-compose --version
# Should output: docker-compose version X.XX.X, build XXXXX

# Test Docker
docker run hello-world
# Should run successfully and show welcome message
```

**CMD:**
```cmd
REM Check Docker
docker --version

REM Check Docker Compose
docker-compose --version

REM Test Docker
docker run hello-world
```

---

## 🚀 Next Steps After Installation

After Docker is installed:

1. **Restart your computer** (if prompted)

2. **Verify installation:**
   ```powershell
   docker --version
   ```

3. **Run AlteqiaChef setup:**
   ```powershell
   cd alteqiachef
   .\setup.ps1 -Action setup
   ```

4. **Access application:**
   - Open: http://localhost:3000

---

## 🔐 Docker Security Considerations

### Windows Security

- Docker requires Hyper-V, which has security implications
- Only install on trusted computers
- Keep Docker updated for security patches

### User Permissions

- Docker commands require elevated privileges
- Consider the security implications for your network

### Image Security

- Pull images only from trusted sources
- Verify image hashes when possible
- Keep images updated

---

## 📚 Additional Resources

- [Docker Official Website](https://www.docker.com/)
- [Docker Installation Guide](https://docs.docker.com/install/)
- [Docker Desktop for Windows](https://docs.docker.com/desktop/install/windows-install/)
- [Docker Desktop for macOS](https://docs.docker.com/desktop/install/mac-install/)
- [Docker on Linux](https://docs.docker.com/engine/install/linux/)
- [WSL 2 Setup](https://docs.microsoft.com/en-us/windows/wsl/install)

---

## 💡 Pro Tips

### Faster Installation

1. Use Chocolatey for package management
2. Install WSL 2 for better performance
3. Use SSD for faster Docker operations

### Better Performance

1. Allocate sufficient resources (Settings → Resources)
2. Enable WSL 2 backend
3. Clean up old images: `docker system prune -a`

### Troubleshooting Tools

```powershell
# Show Docker system info
docker system info

# View Docker disk usage
docker system df

# Clean up unused resources
docker system prune -a

# View all images
docker image ls

# View all containers
docker container ls -a
```

---

## 🆘 Still Having Issues?

1. **Check system requirements:**
   - Windows version must be 19041+
   - 4GB RAM minimum
   - Hyper-V or WSL 2 enabled

2. **Restart your computer:**
   - Many issues are resolved by restart
   - Ensure full restart (not sleep mode)

3. **Try manual installation:**
   - Download from https://www.docker.com/products/docker-desktop
   - Follow official installation steps

4. **Check logs:**
   - Docker Desktop → Troubleshoot → Check logs
   - Review error messages carefully

5. **Update your system:**
   - Windows Update to latest version
   - Ensure security patches are applied

---

## 📝 Installation Checklist

- [ ] Windows version is 19041 or higher
- [ ] Hyper-V or WSL 2 is enabled
- [ ] 4GB+ RAM is available
- [ ] 5GB+ disk space is free
- [ ] Administrator access available
- [ ] Internet connection is stable
- [ ] Run setup script
- [ ] Docker installation completes
- [ ] Computer restarted (if prompted)
- [ ] Docker verified with: `docker --version`
- [ ] Ready to run AlteqiaChef setup

---

## 🎉 Installation Complete!

Once Docker is installed and verified:

```powershell
# Run AlteqiaChef setup
cd alteqiachef
.\setup.ps1 -Action setup

# Access the application
# Frontend: http://localhost:3000
# Backend:  http://localhost:8080
```

---

**Version:** 1.0
**Last Updated:** November 28, 2025
**Compatible with:** Windows 10/11, macOS, Linux
