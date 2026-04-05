# 🔧 AlteqiaChef - Environment Setup Guide

This guide helps you set up your environment for running AlteqiaChef on Windows, macOS, or Linux.

---

## 📋 Table of Contents

1. [System Requirements](#system-requirements)
2. [Installing Docker](#installing-docker)
3. [Windows Setup](#windows-setup)
4. [macOS Setup](#macos-setup)
5. [Linux Setup](#linux-setup)
6. [Verifying Installation](#verifying-installation)
7. [Troubleshooting](#troubleshooting)

---

## 📦 System Requirements

### Minimum
- **RAM:** 4GB (8GB recommended)
- **Disk Space:** 5GB free
- **CPU:** Dual-core processor
- **Internet:** For pulling Docker images

### Recommended
- **RAM:** 8GB or more
- **Disk Space:** 10GB or more
- **CPU:** Quad-core or better
- **SSD:** For better performance

---

## 🐳 Installing Docker

### What is Docker?

Docker is a containerization platform that bundles your application and all its dependencies into isolated containers. This ensures the app runs the same on any machine.

### Why Docker?

- ✅ **Consistency:** Works the same on Windows, macOS, Linux, and servers
- ✅ **Simplicity:** No need to install Go, Node, PostgreSQL separately
- ✅ **Isolation:** Each service runs independently
- ✅ **Scalability:** Easy to scale services

---

## 🪟 Windows Setup

### Step 1: Install Docker Desktop

1. **Download** Docker Desktop from https://www.docker.com/products/docker-desktop
2. **Run** the installer (`DockerDesktopInstaller.exe`)
3. **Select** installation options:
   - ✅ "Install required Windows components for WSL 2"
   - ✅ "Start Docker Desktop when you log in"
4. **Complete** the installation
5. **Restart** your computer when prompted

### Step 2: Verify Docker Installation

#### PowerShell
```powershell
# Check Docker version
docker --version
docker-compose --version

# Test Docker
docker run hello-world
```

#### Command Prompt (CMD)
```cmd
REM Check Docker version
docker --version
docker-compose --version

REM Test Docker
docker run hello-world
```

### Step 3: Configure Docker Resources

1. **Open** Docker Desktop
2. **Click** the gear icon (Settings)
3. **Go to** "Resources"
4. **Set:**
   - **CPUs:** At least 2 (or half your system cores)
   - **Memory:** At least 4GB (8GB recommended)
   - **Swap:** 1GB minimum
5. **Apply & Restart**

### Step 4: Enable WSL 2 (Optional but Recommended)

For better performance:

1. **Open** PowerShell as Administrator
2. **Run:**
   ```powershell
   wsl --install
   wsl --set-default-version 2
   ```
3. **Restart** your computer

---

## 🍎 macOS Setup

### Step 1: Install Docker Desktop

#### Using Homebrew (Recommended)
```bash
brew install --cask docker
```

#### Manual Download
1. **Visit** https://www.docker.com/products/docker-desktop
2. **Download** for macOS
3. **Open** the .dmg file
4. **Drag** Docker to Applications folder

### Step 2: Start Docker Desktop

1. **Open** Applications folder
2. **Double-click** Docker
3. **Allow** system password if prompted
4. **Wait** for Docker to start (icon appears in menu bar)

### Step 3: Verify Installation

```bash
docker --version
docker-compose --version
docker run hello-world
```

### Step 4: Optimize Resources

1. **Click** Docker icon in menu bar
2. **Select** Preferences
3. **Go to** Resources
4. **Set:**
   - **CPUs:** 2-4
   - **Memory:** 4-8GB
5. **Apply**

---

## 🐧 Linux Setup

### Ubuntu/Debian

```bash
# Update package manager
sudo apt-get update
sudo apt-get upgrade

# Install Docker
sudo apt-get install docker.io docker-compose

# Add your user to docker group (avoid sudo)
sudo usermod -aG docker $USER

# Apply group changes
newgrp docker

# Verify installation
docker --version
docker-compose --version
docker run hello-world
```

### Fedora/RHEL/CentOS

```bash
# Install Docker
sudo yum install docker docker-compose

# Start Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Add your user to docker group
sudo usermod -aG docker $USER

# Verify installation
docker --version
docker-compose --version
```

### Arch Linux

```bash
# Install Docker
sudo pacman -S docker docker-compose

# Start Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Add user to docker group
sudo usermod -aG docker $USER

# Verify installation
docker --version
docker-compose --version
```

---

## ✅ Verifying Installation

### Check Docker Service

```bash
# Check Docker daemon is running
docker ps

# Should output: CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
# (may be empty if no containers running)
```

### Check Docker Compose

```bash
# Check Docker Compose version
docker-compose --version

# Should output: docker-compose version X.XX.X, build XXXXXXX
```

### Test Installation

```bash
# Pull and run a test container
docker run --rm hello-world

# Should output: Hello from Docker!
```

### Check Resources

```bash
# View Docker system information
docker system info

# Look for:
# - Containers: 0
# - Images: varies
# - Server Version: 20.10 or higher
```

---

## 🆘 Troubleshooting

### Docker not starting (Windows)

**Problem:** Docker Desktop won't start

**Solution:**
1. Check if WSL 2 is installed: `wsl -l`
2. Update WSL 2 kernel: https://aka.ms/wsl2kernel
3. Restart Docker Desktop
4. If still failing, reinstall Docker:
   ```powershell
   # Uninstall
   wsl --unregister docker-desktop
   wsl --unregister docker-desktop-data
   # Then reinstall Docker Desktop
   ```

### Docker runs but containers fail

**Problem:** Containers won't start

**Solution:**
1. Check Docker daemon: `docker ps`
2. Check resources: Open Docker Desktop → Settings → Resources
3. Restart Docker Desktop
4. Check logs: `docker logs <container_id>`

### Permission denied (Linux)

**Problem:** "permission denied" when running docker

**Solution:**
```bash
# Add user to docker group
sudo usermod -aG docker $USER

# Apply new group settings
newgrp docker

# Verify
docker ps
```

### Out of disk space

**Problem:** "No space left on device"

**Solution:**
```bash
# Clean up unused Docker resources
docker system prune -a

# Remove old images
docker image prune -a

# Remove unused volumes
docker volume prune

# Check disk usage
docker system df
```

### Ports in use

**Problem:** "Port X is already allocated"

**Solution:**

**Windows PowerShell:**
```powershell
# Find process using port
netstat -ano | findstr :3000

# Kill process (replace PID with the number)
taskkill /PID <PID> /F
```

**macOS/Linux:**
```bash
# Find process using port
lsof -i :3000

# Kill process (replace PID with the number)
kill -9 <PID>
```

### Network issues

**Problem:** Container can't connect to network

**Solution:**
```bash
# Restart Docker
docker restart

# Reset network
docker network prune

# Restart containers
docker-compose down
docker-compose up -d
```

---

## 📊 System Health Check

```bash
# Complete Docker system info
docker system df

# View Docker version
docker --version

# Check container runtime
docker info | grep -i runtime

# List all networks
docker network ls

# List all volumes
docker volume ls

# Check disk usage
docker system df
```

---

## 🚀 Next Steps

After verifying Docker installation:

1. **Navigate** to project directory:
   ```bash
   cd alteqiachef
   ```

2. **Run setup** (choose your method):
   
   **Windows PowerShell:**
   ```powershell
   .\setup.ps1 -Action setup
   ```
   
   **Windows CMD:**
   ```cmd
   setup.bat setup
   ```
   
   **macOS/Linux:**
   ```bash
   cd setupchef
   docker-compose up -d
   ```

3. **Verify** services:
   ```bash
   docker-compose ps
   ```

4. **Access** application:
   - Frontend: http://localhost:3000
   - Backend: http://localhost:8080

---

## 📚 Additional Resources

- [Docker Documentation](https://docs.docker.com/)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [WSL 2 Setup (Windows)](https://docs.microsoft.com/en-us/windows/wsl/install)

---

## 🆘 Still Having Issues?

1. **Check logs:** `docker logs <container_name>`
2. **Read SETUP.md** for application-specific issues
3. **Review QUICKSTART.md** for quick reference
4. **Check Docker status:** `docker ps` and `docker-compose ps`
5. **Restart everything:** `docker-compose down -v` and `docker-compose up -d`

---

**Version:** 1.0
**Last Updated:** November 28, 2025
