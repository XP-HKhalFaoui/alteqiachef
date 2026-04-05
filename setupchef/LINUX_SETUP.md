# 🐧 AlteqiaChef - Linux (Fedora/RHEL/CentOS) Setup Guide

Complete setup guide for AlteqiaChef on Linux systems, with a focus on Fedora, RHEL, and CentOS.

---

## ⚡ Quick Start (3 Steps)

### Step 1: Clone and Navigate
```bash
cd alteqiachef
```

### Step 2: Run Setup Script
```bash
chmod +x setup.sh
./setup.sh setup
```

### Step 3: Access Application
```
Open: http://localhost:3000
```

---

## 📦 Supported Linux Distributions

| Distribution | Package Manager | Status |
|--------------|-----------------|--------|
| **Fedora 35+** | DNF | ✅ Fully Supported |
| **RHEL 8+** | DNF | ✅ Fully Supported |
| **CentOS 8+** | DNF | ✅ Fully Supported |
| **Ubuntu 20.04+** | APT | ✅ Fully Supported |
| **Debian 10+** | APT | ✅ Fully Supported |

---

## 🔧 Prerequisites

### System Requirements
- **OS:** Fedora 35+, RHEL 8+, CentOS 8+, or Ubuntu 20.04+
- **RAM:** 4GB minimum (8GB recommended)
- **Disk Space:** 5GB free
- **CPU:** Dual-core processor
- **Internet:** For pulling Docker images

### Required Packages
- **docker** - Container runtime
- **docker-compose** - Container orchestration
- **curl** - For health checks
- **netstat** - For port checking

---

## 🚀 Automated Setup (Recommended)

### The Easiest Way

```bash
cd alteqiachef
chmod +x setup.sh
./setup.sh setup
```

The script will automatically:

1. ✅ Detect your Linux distribution
2. ✅ Install Docker if not found
3. ✅ Configure Docker daemon
4. ✅ Add your user to docker group
5. ✅ Pull Docker images
6. ✅ Start all containers
7. ✅ Verify services are running

---

## 🛠️ Manual Setup (Step by Step)

### Step 1: Install Docker

#### Fedora/RHEL/CentOS

```bash
# Update system
sudo dnf update -y

# Add Docker repository
sudo dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo

# Install Docker
sudo dnf install -y docker-ce docker-ce-cli containerd.io

# Install Docker Compose (if not included)
sudo dnf install -y docker-compose-plugin
```

#### Ubuntu/Debian

```bash
# Update system
sudo apt-get update
sudo apt-get upgrade -y

# Install dependencies
sudo apt-get install -y apt-transport-https ca-certificates curl gnupg lsb-release

# Add Docker GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add Docker repository
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# Install Docker Compose
sudo apt-get install -y docker-compose
```

### Step 2: Start Docker Service

```bash
# Start Docker daemon
sudo systemctl start docker

# Enable auto-start
sudo systemctl enable docker

# Verify Docker is running
sudo systemctl status docker
```

### Step 3: Configure Docker for Current User

```bash
# Add current user to docker group
sudo usermod -aG docker $USER

# Apply group changes (choose one):

# Option 1: Log out and log back in
# Then log back in to activate group changes

# Option 2: Apply changes immediately
newgrp docker

# Verify Docker works without sudo
docker ps
```

### Step 4: Clone AlteqiaChef Repository

```bash
# Clone the repository (if not already done)
git clone https://github.com/XP-HKhalFaoui/alteqiachef.git
cd alteqiachef
```

### Step 5: Start AlteqiaChef

```bash
# Navigate to setup directory
cd setupchef

# Pull latest images
docker-compose pull

# Start containers
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f

# Stop viewing logs: Ctrl+C
```

### Step 6: Verify Services

```bash
# Check all containers are running
docker-compose ps

# Test frontend
curl http://localhost:3000

# Test backend
curl http://localhost:8080/api/v1/health

# Test database
docker exec pos-postgres psql -U postgres -d pos_system -c "SELECT 1;"
```

---

## 📋 Setup Script Commands

### Available Actions

```bash
./setup.sh setup       # Complete setup with Docker installation
./setup.sh start       # Start containers
./setup.sh stop        # Stop containers
./setup.sh restart     # Restart all containers
./setup.sh status      # Show service status and health
./setup.sh logs        # Stream logs from all services
./setup.sh clean       # Remove containers and volumes
./setup.sh help        # Show help message
```

### Examples

```bash
# Full setup
./setup.sh setup

# Check what's running
./setup.sh status

# View logs in real-time
./setup.sh logs

# Stop all services
./setup.sh stop

# Restart services
./setup.sh restart
```

---

## 📍 Application URLs

After setup completes:

| Service | URL |
|---------|-----|
| **Frontend** | http://localhost:3000 |
| **Backend API** | http://localhost:8080 |
| **Database** | localhost:5432 |

---

## 🔍 Verification

### Check Installation

```bash
# Check Docker version
docker --version

# Check Docker Compose
docker-compose --version

# Test Docker
docker run hello-world

# Check containers
docker-compose ps

# View logs
docker-compose logs pos-backend
```

### Verify Services

```bash
# Check frontend
curl http://localhost:3000

# Check backend
curl http://localhost:8080/api/v1/health

# Check database
docker exec pos-postgres psql -U postgres -d pos_system -c "SELECT version();"
```

---

## 👤 Create Admin User

```bash
# Navigate to setupchef
cd setupchef

# Run admin creation script
./create-admin.sh

# Follow the prompts:
# - Username
# - Email
# - First Name
# - Last Name
# - Password
```

---

## 🔄 Common Tasks

### View Logs

```bash
# View all logs
docker-compose logs

# View specific service
docker-compose logs pos-backend
docker-compose logs pos-frontend
docker-compose logs pos-postgres

# Stream logs (real-time)
docker-compose logs -f

# Last 100 lines
docker-compose logs --tail 100

# Stop viewing: Ctrl+C
```

### Access Database

```bash
# Connect to PostgreSQL
docker exec -it pos-postgres psql -U postgres -d pos_system

# Common commands:
\dt              # List tables
\du              # List users
SELECT * FROM users;  # View users
\q               # Exit
```

### Restart Services

```bash
# Restart all services
docker-compose restart

# Restart specific service
docker-compose restart pos-backend
docker-compose restart pos-frontend
docker-compose restart pos-postgres

# Full restart (stop + start)
docker-compose down
docker-compose up -d
```

### Backup Database

```bash
# Create backup
docker exec pos-postgres pg_dump -U postgres -d pos_system > backup.sql

# Check backup
ls -lh backup.sql
```

### Restore Database

```bash
# Restore from backup
docker exec -i pos-postgres psql -U postgres -d pos_system < backup.sql
```

### Reset Database

```bash
cd setupchef
./db-reset.sh
```

---

## 🆘 Troubleshooting

### Docker not found

**Problem:** Command not found: docker

**Solution:**
```bash
# Verify installation
which docker

# If not found, install Docker (see Prerequisites)
# Then check PATH
export PATH=$PATH:/usr/bin
docker --version
```

### Permission denied

**Problem:** Got permission denied while trying to connect to Docker daemon

**Solution:**
```bash
# Add user to docker group
sudo usermod -aG docker $USER

# Log out and back in, or run:
newgrp docker

# Verify
docker ps
```

### Docker daemon not running

**Problem:** Cannot connect to Docker daemon

**Solution:**
```bash
# Start Docker daemon
sudo systemctl start docker

# Enable auto-start
sudo systemctl enable docker

# Check status
sudo systemctl status docker
```

### Port already in use

**Problem:** Port 3000/8080/5432 already in use

**Solution:**
```bash
# Find what's using the port
sudo netstat -tuln | grep :3000

# Stop existing containers
docker-compose down

# Or kill the process
sudo fuser -k 3000/tcp

# Start fresh
docker-compose up -d
```

### Containers won't start

**Problem:** Containers keep exiting

**Solution:**
```bash
# Check logs
docker-compose logs

# Rebuild containers
docker-compose down
docker-compose build --no-cache
docker-compose up -d

# Check individual service logs
docker-compose logs pos-backend
```

### Out of disk space

**Problem:** No space left on device

**Solution:**
```bash
# Check disk usage
df -h

# Clean Docker resources
docker system prune -a

# Remove old images
docker image prune -a

# Remove unused volumes
docker volume prune

# Check what was freed
docker system df
```

### Database connection error

**Problem:** Backend can't connect to database

**Solution:**
```bash
# Verify database container is running
docker-compose ps

# Check database logs
docker-compose logs pos-postgres

# Test database connection
docker exec pos-postgres psql -U postgres -d pos_system -c "SELECT 1;"

# Reset database if needed
cd setupchef && ./db-reset.sh
```

---

## 🔐 Security Considerations

### User Permissions

- ✅ Docker commands work without `sudo` after adding user to group
- ✅ Only trusted users should have docker group access
- ✅ Docker group has root-equivalent privileges

### Network Security

- ✅ Services only accessible on localhost by default
- ✅ Use firewall rules if exposing to network
- ✅ Change default database password in production

### Data Security

- ✅ Database data persisted in Docker volume
- ✅ Regular backups recommended
- ✅ Use strong passwords for admin accounts

---

## 📊 System Administration

### Monitor Services

```bash
# View running containers
docker ps

# View all containers (including stopped)
docker ps -a

# View resource usage
docker stats

# View detailed container info
docker inspect pos-backend
```

### Manage Volumes

```bash
# List all volumes
docker volume ls

# Inspect a volume
docker volume inspect alteqiachef_postgres_data

# Remove unused volumes
docker volume prune
```

### Manage Networks

```bash
# List networks
docker network ls

# Inspect network
docker network inspect alteqiachef_pos-network
```

### Docker Cleanup

```bash
# Remove unused containers
docker container prune

# Remove unused images
docker image prune -a

# Remove unused volumes
docker volume prune

# Remove all unused resources
docker system prune -a
```

---

## 🔄 Update Services

### Update Images

```bash
# Pull latest images
docker-compose pull

# Restart services with new images
docker-compose up -d

# Check running version
docker-compose ps
```

### Update Docker

**Fedora/RHEL/CentOS:**
```bash
sudo dnf upgrade docker-ce
sudo systemctl restart docker
```

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get upgrade docker-ce
sudo systemctl restart docker
```

---

## 📝 Service Commands Reference

```bash
# View service status
./setup.sh status

# View real-time logs
./setup.sh logs

# Start services
./setup.sh start

# Stop services
./setup.sh stop

# Restart services
./setup.sh restart

# Complete cleanup
./setup.sh clean

# Help
./setup.sh help
```

---

## 💾 Systemd Integration (Optional)

### Create Systemd Service

Create `/etc/systemd/system/alteqiachef.service`:

```ini
[Unit]
Description=AlteqiaChef POS System
Requires=docker.service
After=docker.service

[Service]
Type=simple
User=YOUR_USERNAME
WorkingDirectory=/path/to/alteqiachef/setupchef
ExecStart=/usr/bin/docker-compose up
ExecStop=/usr/bin/docker-compose down
Restart=always

[Install]
WantedBy=multi-user.target
```

### Use Systemd Service

```bash
# Enable service
sudo systemctl enable alteqiachef

# Start service
sudo systemctl start alteqiachef

# Check status
sudo systemctl status alteqiachef

# View logs
sudo journalctl -u alteqiachef -f

# Stop service
sudo systemctl stop alteqiachef
```

---

## 📚 Additional Resources

### Official Documentation
- [Docker Official Docs](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Fedora Docker Guide](https://docs.fedoraproject.org/en-US/fedora-silverblue/getting-started/)

### Linux Guides
- [Linux Docker Installation](https://docs.docker.com/engine/install/)
- [Manage Docker as Non-Root User](https://docs.docker.com/engine/install/linux-postinstall/)
- [Docker Security](https://docs.docker.com/engine/security/)

---

## ✅ Verification Checklist

After setup, verify:

- [ ] Docker installed: `docker --version`
- [ ] Docker Compose installed: `docker-compose --version`
- [ ] Docker daemon running: `sudo systemctl status docker`
- [ ] User in docker group: `groups $USER`
- [ ] Containers running: `docker-compose ps` (all 3)
- [ ] Frontend accessible: http://localhost:3000
- [ ] Backend responding: http://localhost:8080/api/v1/health
- [ ] Database accessible: Port 5432 responds
- [ ] Admin user created
- [ ] Can login to application

---

## 🎉 Next Steps

1. **Access the application:**
   ```
   http://localhost:3000
   ```

2. **Create admin user:**
   ```bash
   cd setupchef && ./create-admin.sh
   ```

3. **Create demo users (optional):**
   ```bash
   cd setupchef && ./create-demo-users.sh
   ```

4. **Start using:**
   - Login with admin credentials
   - Manage your restaurant!

---

## 📞 Quick Commands

```bash
# Setup
./setup.sh setup

# Status
./setup.sh status

# Logs
./setup.sh logs

# Stop
./setup.sh stop

# Restart
./setup.sh restart

# Database shell
docker exec -it pos-postgres psql -U postgres -d pos_system

# View logs
docker-compose logs -f

# Clean
./setup.sh clean
```

---

**Version:** 1.0
**Last Updated:** November 28, 2025
**Compatible with:** Fedora 35+, RHEL 8+, CentOS 8+, Ubuntu 20.04+, Debian 10+
