# 🐧 AlteqiaChef - Complete Linux & Fedora Setup Guide

Comprehensive guide for setting up AlteqiaChef on all Linux distributions, with special focus on Fedora.

---

## ✨ What's New

Your AlteqiaChef application now has **complete Linux support** with:

✅ **Automatic Setup Script** - Works on all major Linux distributions
✅ **Docker Auto-Installation** - Detects distro and installs Docker appropriately
✅ **Fedora Quick Start** - 5-minute setup for Fedora users
✅ **Comprehensive Documentation** - Step-by-step guides for all scenarios

---

## 🚀 Quick Start (3 Steps)

### For Fedora Users

```bash
# Step 1: Navigate to project
cd alteqiachef

# Step 2: Run setup (make script executable first)
chmod +x setup.sh
./setup.sh setup

# Step 3: Open browser
# http://localhost:3000
```

**That's it!** Script handles everything automatically.

### For Other Linux Users

The same setup works for:
- Ubuntu 20.04+
- Debian 10+
- RHEL 8+
- CentOS 8+

Just run the same 3 commands!

---

## 📋 Supported Linux Distributions

| Distro | Version | Package Manager | Support |
|--------|---------|-----------------|---------|
| **Fedora** | 35+ | DNF | ✅ Full |
| **RHEL** | 8+ | DNF | ✅ Full |
| **CentOS** | 8+ | DNF | ✅ Full |
| **Ubuntu** | 20.04+ | APT | ✅ Full |
| **Debian** | 10+ | APT | ✅ Full |

---

## 🔍 How the Setup Script Works

### The setup.sh Script

```bash
./setup.sh setup
```

This automatically:

1. ✅ **Detects your Linux distribution**
   - Checks /etc/os-release
   - Identifies package manager (DNF, APT, etc)

2. ✅ **Checks Docker installation**
   - Looks for docker command
   - If not found, offers to install

3. ✅ **Installs Docker** (if needed)
   - Fedora/RHEL/CentOS: Uses DNF
   - Ubuntu/Debian: Uses APT
   - Configures repositories
   - Installs docker-ce and docker-compose

4. ✅ **Starts Docker daemon**
   - Enables auto-start with systemd
   - Starts Docker service

5. ✅ **Configures permissions**
   - Adds current user to docker group
   - Allows docker commands without sudo

6. ✅ **Checks system**
   - Verifies ports are available
   - Checks disk space

7. ✅ **Pulls Docker images**
   - Downloads from Docker registry
   - Shows progress

8. ✅ **Starts containers**
   - PostgreSQL database
   - Go backend API
   - React frontend

9. ✅ **Verifies services**
   - Checks all containers running
   - Tests frontend (3000)
   - Tests backend (8080)
   - Tests database (5432)

10. ✅ **Shows success message**
    - Application ready
    - Access URLs provided

---

## 📖 Available Documentation

### FEDORA_QUICKSTART.md
**Perfect for: Fedora users who want to get started fast**
- 5-minute setup
- Basic commands
- Quick troubleshooting
- No prerequisites knowledge needed

### LINUX_SETUP.md
**Perfect for: All Linux users wanting comprehensive guide**
- Supported distributions explained
- Automated setup with script
- Manual step-by-step setup
- Common tasks (logs, database, restart)
- Troubleshooting guide
- Security considerations
- System administration
- Systemd integration

### LINUX_IMPLEMENTATION.md
**Perfect for: Understanding the implementation**
- Architecture and design
- Installation flow diagram
- Features overview
- System requirements
- Verification steps

---

## 🎮 Available Commands

```bash
# Full setup with Docker auto-installation
./setup.sh setup

# Start containers
./setup.sh start

# Stop containers
./setup.sh stop

# Restart all containers
./setup.sh restart

# Show service status
./setup.sh status

# Stream logs from all services
./setup.sh logs

# Remove containers and volumes (WARNING: loses data)
./setup.sh clean

# Show help message
./setup.sh help
```

---

## 📍 Application URLs

After setup completes, access at:

- **Frontend:** http://localhost:3000
- **Backend API:** http://localhost:8080
- **Database:** localhost:5432 (PostgreSQL)

---

## 🔧 Manual Setup (Step by Step)

If you prefer to set up manually:

### Step 1: Install Docker

**Fedora/RHEL/CentOS:**
```bash
sudo dnf install -y docker-ce docker-compose-plugin
```

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get install -y docker.io docker-compose
```

### Step 2: Start Docker

```bash
# Start daemon
sudo systemctl start docker

# Enable auto-start
sudo systemctl enable docker

# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker
```

### Step 3: Clone Repository

```bash
git clone https://github.com/XP-HKhalFaoui/alteqiachef.git
cd alteqiachef
```

### Step 4: Start Application

```bash
cd setupchef
docker-compose up -d
docker-compose ps
```

### Step 5: Verify Services

```bash
# Check frontend
curl http://localhost:3000

# Check backend
curl http://localhost:8080/api/v1/health

# Check database
docker exec pos-postgres psql -U postgres -d pos_system -c "SELECT 1;"
```

---

## 👤 Create Admin User

```bash
# Navigate to setupchef
cd setupchef

# Run admin creation script
./create-admin.sh

# Follow prompts for:
# - Username
# - Email
# - First Name
# - Last Name
# - Password
```

---

## 🛠️ Common Tasks

### View Logs
```bash
# All logs
./setup.sh logs

# Specific service
docker-compose logs pos-backend
docker-compose logs pos-frontend
docker-compose logs pos-postgres

# Last 100 lines
docker-compose logs --tail 100 -f
```

### Database Access
```bash
# Connect to PostgreSQL
docker exec -it pos-postgres psql -U postgres -d pos_system

# Common commands in psql:
\dt              # List tables
\du              # List users
SELECT * FROM users;  # View users
\q               # Exit
```

### Backup Database
```bash
# Create backup
docker exec pos-postgres pg_dump -U postgres -d pos_system > backup.sql

# Restore from backup
docker exec -i pos-postgres psql -U postgres -d pos_system < backup.sql
```

### Reset Database
```bash
cd setupchef
./db-reset.sh
```

### Restart Services
```bash
./setup.sh restart

# Or specific service
docker-compose restart pos-backend
```

---

## 🆘 Troubleshooting

### Docker not found

**Solution:**
```bash
# Install Docker for your distro
# Fedora:
sudo dnf install -y docker-ce docker-compose-plugin

# Ubuntu/Debian:
sudo apt-get install -y docker.io docker-compose

# Then start Docker
sudo systemctl start docker
```

### Permission denied

**Solution:**
```bash
# Add user to docker group
sudo usermod -aG docker $USER

# Apply changes immediately
newgrp docker

# Or log out and log back in
```

### Port already in use

**Solution:**
```bash
# Find what's using the port
sudo lsof -i :3000

# Stop existing containers
docker-compose down

# Start fresh
./setup.sh setup
```

### Containers won't start

**Solution:**
```bash
# Check logs
docker-compose logs

# Restart Docker daemon
sudo systemctl restart docker

# Try again
./setup.sh start
```

---

## ✅ Verification Checklist

After setup, verify:

- [ ] Docker installed: `docker --version`
- [ ] Compose installed: `docker-compose --version`
- [ ] Containers running: `docker-compose ps` (3 containers)
- [ ] Frontend accessible: http://localhost:3000
- [ ] Backend responding: http://localhost:8080/api/v1/health
- [ ] Database accessible: localhost:5432
- [ ] Admin user created
- [ ] Can login to application

---

## 🔐 Security Best Practices

✅ **User Permissions:**
- Docker commands work without sudo (after group setup)
- Only trusted users should be in docker group
- Docker group has root-equivalent privileges

✅ **Network Security:**
- Services only accessible on localhost by default
- Use firewall rules if exposing to network

✅ **Data Security:**
- Database data persisted in Docker volumes
- Regular backups recommended
- Use strong admin passwords

✅ **Container Security:**
- Use official images from Docker Hub
- Keep containers updated
- Review logs for suspicious activity

---

## 📊 System Requirements

- **RAM:** 4GB minimum (8GB recommended)
- **Disk Space:** 5GB free (10GB recommended)
- **CPU:** Dual-core processor
- **Internet:** For pulling Docker images
- **Kernel:** Linux 3.10+
- **systemd:** For service management

---

## 🔄 Service Management

### Check Service Status
```bash
./setup.sh status
# Shows all containers and health
```

### Monitor in Real-Time
```bash
./setup.sh logs
# Press Ctrl+C to stop
```

### Restart Services
```bash
./setup.sh restart
# Or restart specific service
docker-compose restart pos-backend
```

### Stop All Services
```bash
./setup.sh stop
# Services stopped but data preserved
```

### Start Again
```bash
./setup.sh start
# Services restart with same data
```

---

## 📈 Performance Tips

### Optimize Docker Resources
```bash
# Check resource usage
docker stats

# View container details
docker inspect pos-backend
```

### Disk Cleanup
```bash
# Remove unused images
docker image prune -a

# Remove unused volumes
docker volume prune

# Remove unused containers
docker container prune
```

### Network Performance
```bash
# Check network settings
docker network ls
docker network inspect alteqiachef_pos-network
```

---

## 🔄 Update & Maintenance

### Update Docker
```bash
# Fedora/RHEL/CentOS
sudo dnf upgrade docker-ce

# Ubuntu/Debian
sudo apt-get upgrade docker-ce
```

### Update Images
```bash
# Pull latest versions
docker-compose pull

# Restart with new images
docker-compose up -d
```

### View Docker Logs
```bash
# System logs
sudo journalctl -u docker

# Docker daemon info
docker system info
```

---

## 💡 Pro Tips

### Run as Background Service
```bash
# Start services
./setup.sh start

# Close terminal - services continue running
# Services auto-restart on reboot (if enabled)
```

### Create Systemd Service File
```bash
# Create /etc/systemd/system/alteqiachef.service
[Unit]
Description=AlteqiaChef POS System
Requires=docker.service
After=docker.service

[Service]
Type=simple
User=$USER
WorkingDirectory=/path/to/alteqiachef/setupchef
ExecStart=/usr/bin/docker-compose up
ExecStop=/usr/bin/docker-compose down

[Install]
WantedBy=multi-user.target
```

Then:
```bash
sudo systemctl daemon-reload
sudo systemctl enable alteqiachef
sudo systemctl start alteqiachef
```

### Multiple Instances
```bash
# Create separate directory for each instance
mkdir alteqiachef-prod
mkdir alteqiachef-staging

# Use different ports in each docker-compose.yml
# Modify port mappings (e.g., 3001, 3002)
```

---

## 📚 Additional Resources

### Official Documentation
- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Docs](https://docs.docker.com/compose/)
- [Linux Installation Guide](https://docs.docker.com/engine/install/linux/)

### Distribution-Specific
- [Fedora Docker Guide](https://docs.fedoraproject.org/en-US/)
- [Ubuntu Docker Guide](https://docs.docker.com/engine/install/ubuntu/)
- [Debian Docker Guide](https://docs.docker.com/engine/install/debian/)

---

## 📞 Quick Reference

```bash
# Setup
chmod +x setup.sh
./setup.sh setup

# Status
./setup.sh status

# Logs
./setup.sh logs

# Stop
./setup.sh stop

# Start
./setup.sh start

# Restart
./setup.sh restart

# Help
./setup.sh help

# Database
docker exec -it pos-postgres psql -U postgres -d pos_system

# Backup
docker exec pos-postgres pg_dump -U postgres -d pos_system > backup.sql
```

---

## 🎯 Next Steps

1. **Run Setup:**
   ```bash
   chmod +x setup.sh
   ./setup.sh setup
   ```

2. **Access Application:**
   ```
   http://localhost:3000
   ```

3. **Create Admin User:**
   ```bash
   cd setupchef && ./create-admin.sh
   ```

4. **Start Using:**
   - Login with admin credentials
   - Create demo users
   - Manage your restaurant!

---

## ✅ Feature Checklist

- ✅ Fedora 35+ support
- ✅ RHEL 8+ support
- ✅ CentOS 8+ support
- ✅ Ubuntu 20.04+ support
- ✅ Debian 10+ support
- ✅ Automatic Docker detection
- ✅ Automatic Docker installation
- ✅ User permission configuration
- ✅ Service health checks
- ✅ Color-coded output
- ✅ Error handling
- ✅ Clear documentation
- ✅ Quick start guide
- ✅ Troubleshooting guide

---

## 🎉 Ready to Go!

Everything is set up and ready for Linux/Fedora deployment!

**Get started now:**
```bash
chmod +x setup.sh
./setup.sh setup
```

Then open: **http://localhost:3000**

---

## 📝 File Manifest

```
✅ setup.sh                      New
✅ FEDORA_QUICKSTART.md          New
✅ LINUX_SETUP.md                New
✅ LINUX_IMPLEMENTATION.md       New
✅ setup.ps1                     Existing (Windows)
✅ setup.bat                     Existing (Windows)
```

---

**Implementation Status:** ✅ COMPLETE
**Version:** 1.0
**Date:** November 28, 2025
**Supported:** Fedora 35+, RHEL 8+, CentOS 8+, Ubuntu 20.04+, Debian 10+
**Ready for Production:** YES 🎉
