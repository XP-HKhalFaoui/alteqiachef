# 📋 AlteqiaChef - Docker Auto-Installation Quick Reference

## 🚀 Get Started in 3 Steps

### Step 1: Open Terminal
```powershell
# Windows PowerShell
cd alteqiachef

# OR Windows CMD
cd alteqiachef\setupchef
```

### Step 2: Run Setup
```powershell
# PowerShell
.\setup.ps1 -Action setup

# OR CMD
setup.bat setup
```

### Step 3: Access Application
```
Open browser to: http://localhost:3000
```

---

## ✨ What the Script Does Automatically

| Step | Action | Status |
|------|--------|--------|
| 1 | Check Docker installed | ✅ Automatic |
| 2 | Install Docker if needed | ✅ Automatic |
| 3 | Verify Docker Compose | ✅ Automatic |
| 4 | Check system ports | ✅ Automatic |
| 5 | Pull Docker images | ✅ Automatic |
| 6 | Start containers | ✅ Automatic |
| 7 | Show success message | ✅ Automatic |

---

## 🔧 Available Commands

### PowerShell (setup.ps1)
```powershell
.\setup.ps1 -Action setup     # Full setup (with Docker install)
.\setup.ps1 -Action start     # Start containers
.\setup.ps1 -Action stop      # Stop containers
.\setup.ps1 -Action restart   # Restart containers
.\setup.ps1 -Action status    # Show status
.\setup.ps1 -Action logs      # View logs
.\setup.ps1 -Action clean     # Remove everything
.\setup.ps1 -Action help      # Show help
```

### CMD (setup.bat)
```cmd
setup.bat setup    # Full setup (with Docker install)
setup.bat start    # Start containers
setup.bat stop     # Stop containers
setup.bat restart  # Restart containers
setup.bat status   # Show status
setup.bat logs     # View logs
setup.bat clean    # Remove everything
setup.bat help     # Show help
```

---

## 📍 Access Points

After setup completes:

| Service | URL | Port |
|---------|-----|------|
| Frontend | http://localhost:3000 | 3000 |
| Backend | http://localhost:8080 | 8080 |
| Database | localhost | 5432 |

---

## 💻 Installation Methods

The script automatically tries:

1. **Chocolatey** (Fastest)
   - If Chocolatey is installed, uses it
   - Installs Docker via: `choco install docker-desktop -y`

2. **Direct Download** (Reliable)
   - If Chocolatey not available
   - Downloads Docker Desktop installer
   - Runs installer automatically

3. **Manual Installation** (Fallback)
   - If auto-install fails
   - Provides link to download manually
   - User can install and run setup again

---

## 🎯 Common Use Cases

### First Time User (No Docker)
```powershell
.\setup.ps1 -Action setup
# Answer: y (to install Docker)
# Computer will restart
# Run script again after restart
# Application ready!
```

### Docker Already Installed
```powershell
.\setup.ps1 -Action setup
# Docker found automatically
# Skips installation
# Starts POS system
# Application ready!
```

### Check What's Running
```powershell
.\setup.ps1 -Action status
# Shows running containers
# Shows service health
# Shows port usage
```

### Stop and Restart
```powershell
.\setup.ps1 -Action stop
# Services stop

.\setup.ps1 -Action restart
# Services restart
```

---

## ✅ Verification

After setup completes, verify everything works:

```powershell
# Check Docker
docker --version

# Check containers
docker-compose ps

# Check logs
docker-compose logs

# Test frontend
curl http://localhost:3000

# Test backend
curl http://localhost:8080/api/v1/health
```

---

## 🆘 If Something Goes Wrong

### Docker Installation Fails
```powershell
# Try running as Administrator
# Right-click PowerShell → "Run as Administrator"
.\setup.ps1 -Action setup
```

### Port Already in Use
```powershell
# Check which process uses port
netstat -ano | findstr :3000

# Stop Docker completely
.\setup.ps1 -Action stop

# Clean everything
.\setup.ps1 -Action clean

# Start fresh
.\setup.ps1 -Action setup
```

### Docker Not Recognized
```powershell
# Restart your computer (important!)
# Then try again

# If still not working:
docker --version
# Should show Docker version

# If not, install manually:
# https://www.docker.com/products/docker-desktop
```

### View Detailed Logs
```powershell
# See all logs
.\setup.ps1 -Action logs

# See specific service logs
docker-compose logs pos-backend
docker-compose logs pos-frontend
docker-compose logs pos-postgres
```

---

## 📚 Documentation

- **Quick Reference** ← You are here
- [DOCKER_AUTO_INSTALL.md](./DOCKER_AUTO_INSTALL.md) - Feature overview
- [DOCKER_INSTALLATION.md](./DOCKER_INSTALLATION.md) - Detailed Docker setup
- [SETUP.md](./SETUP.md) - Complete setup guide
- [QUICKSTART.md](./QUICKSTART.md) - 5-minute quick start

---

## 🎯 Next Steps After Setup

### 1. Create Admin User
```bash
cd setupchef
./create-admin.sh
# Or on Windows: create-admin.sh (in Git Bash)
```

### 2. Create Demo Users
```bash
cd setupchef
./create-demo-users.sh
```

### 3. Login to Application
- Open: http://localhost:3000
- Enter admin credentials
- Start using the POS system!

---

## ⚙️ System Requirements

- **Windows:** 10 (build 19041+) or 11
- **RAM:** 4GB minimum (8GB recommended)
- **Disk:** 5GB free space
- **CPU:** Dual-core minimum
- **Internet:** For pulling Docker images

---

## 🔒 Security Notes

✅ **Safe to Run**
- Scripts are documented and transparent
- No hidden operations
- Requests user confirmation
- Uses official Docker sources

⚠️ **Run on Trusted Machines**
- Docker requires system-level access
- Only install on secure networks
- Keep credentials secure

---

## 💡 Pro Tips

### Faster Setup
- Have stable internet connection
- Use SSD for better performance
- Close unnecessary applications
- Allocate 8GB+ RAM to Docker

### Better Performance
- Use WSL 2 backend (Windows)
- Enable virtualization in BIOS
- Update Docker to latest version

### Troubleshooting Tools
```powershell
# Show system info
docker system info

# Show disk usage
docker system df

# Clean up unused resources
docker system prune -a

# View image details
docker image ls
```

---

## 📞 Quick Commands Reference

```powershell
# Setup & Install
.\setup.ps1 -Action setup          # Install Docker + POS

# Manage Services
.\setup.ps1 -Action start          # Start services
.\setup.ps1 -Action stop           # Stop services
.\setup.ps1 -Action restart        # Restart services
.\setup.ps1 -Action status         # Check status

# View Information
.\setup.ps1 -Action logs           # View all logs
.\setup.ps1 -Action help           # Show help

# Cleanup
.\setup.ps1 -Action clean          # Remove everything

# Direct Docker Commands
docker ps                          # Show containers
docker logs <container>            # View logs
docker exec -it <container> bash   # Enter container
docker-compose ps                  # Show all services
docker-compose logs -f             # Stream logs
```

---

## 🎉 That's It!

Your AlteqiaChef POS system is ready to go!

**To start:**
```powershell
.\setup.ps1 -Action setup
```

**Then open:**
```
http://localhost:3000
```

---

## 📋 Checklist

- [ ] Ran `.\setup.ps1 -Action setup`
- [ ] Script detected Docker (or installed it)
- [ ] Containers are running (`docker-compose ps`)
- [ ] Frontend accessible (http://localhost:3000)
- [ ] Backend accessible (http://localhost:8080)
- [ ] Database responding (port 5432)
- [ ] Created admin user
- [ ] Can login to application

---

**Version:** 1.0
**Last Updated:** November 28, 2025
**Status:** Ready to Use ✅
