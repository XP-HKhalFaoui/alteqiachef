# 🐧 AlteqiaChef - Fedora Quick Start

Fastest way to get AlteqiaChef running on Fedora Linux.

---

## ⚡ 5-Minute Setup

### Prerequisites Check
```bash
# Have these ready:
# - Fedora 35+ installed
# - 4GB+ RAM
# - 5GB+ disk space
# - Internet connection
```

### Step 1: Download Project
```bash
# Clone repository (or navigate if already cloned)
git clone https://github.com/XP-HKhalFaoui/alteqiachef.git
cd alteqiachef
```

### Step 2: Run Setup Script
```bash
# Make script executable
chmod +x setup.sh

# Run automated setup
./setup.sh setup
```

**The script will:**
- ✅ Detect Fedora system
- ✅ Install Docker (if needed)
- ✅ Install Docker Compose (if needed)
- ✅ Start Docker daemon
- ✅ Configure user permissions
- ✅ Pull Docker images
- ✅ Start all containers

### Step 3: Access Application
```
Open browser to: http://localhost:3000
```

Done! 🎉

---

## 🚀 What Happens After Setup

```
1. Fedora system detected ✓
2. Docker installed (if needed) ✓
3. Containers starting...
4. Wait 30 seconds ✓
5. Application ready!
```

---

## 📍 Application URLs

```
Frontend: http://localhost:3000
Backend:  http://localhost:8080
Database: localhost:5432
```

---

## 👤 Create Admin User

```bash
# After setup completes
cd setupchef
./create-admin.sh

# Or use Docker directly
docker exec -it pos-backend ./scripts/create-admin.sh
```

---

## 🎮 Common Commands

```bash
# Check what's running
./setup.sh status

# View logs
./setup.sh logs

# Stop everything
./setup.sh stop

# Start again
./setup.sh start

# Restart
./setup.sh restart
```

---

## 🆘 Troubleshooting

### Docker not found
```bash
# Install manually
sudo dnf install -y docker-ce docker-compose-plugin
sudo systemctl start docker
sudo systemctl enable docker

# Then run setup again
./setup.sh setup
```

### Permission denied
```bash
# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker

# Then try again
docker ps
```

### Port already in use
```bash
# Stop containers
./setup.sh stop

# Start fresh
./setup.sh setup
```

### Containers won't start
```bash
# Check logs
docker-compose logs

# Restart Docker
sudo systemctl restart docker

# Try again
./setup.sh start
```

---

## ✅ Verify Everything Works

```bash
# Check containers running
docker ps
# Should show 3 containers: postgres, backend, frontend

# Test frontend
curl http://localhost:3000
# Should return HTML (200 status)

# Test backend
curl http://localhost:8080/api/v1/health
# Should return OK (200 status)

# Test database
docker exec pos-postgres psql -U postgres -d pos_system -c "SELECT 1;"
# Should return: 1
```

---

## 📚 Full Documentation

- **Full Linux Guide:** [LINUX_SETUP.md](./LINUX_SETUP.md)
- **Quick Reference:** [QUICK_REFERENCE.md](./QUICK_REFERENCE.md)
- **Complete Setup:** [SETUP.md](./SETUP.md)
- **Developer Guide:** [DEVELOPER_REFERENCE.md](./DEVELOPER_REFERENCE.md)

---

## 💡 Pro Tips

### Run in Background
```bash
# Start services
./setup.sh start

# Close terminal - services keep running
```

### View Live Logs
```bash
# Watch logs as things happen
./setup.sh logs

# Press Ctrl+C to stop
```

### Backup Database
```bash
# Backup
docker exec pos-postgres pg_dump -U postgres -d pos_system > backup.sql

# Restore if needed
docker exec -i pos-postgres psql -U postgres -d pos_system < backup.sql
```

### Database Access
```bash
# Connect to database
docker exec -it pos-postgres psql -U postgres -d pos_system

# List tables: \dt
# Exit: \q
```

---

## 🎯 Next Steps

1. **Setup complete!** ✅
2. **Open browser:** http://localhost:3000 ✅
3. **Create admin user:** `cd setupchef && ./create-admin.sh` ✅
4. **Login and use:** Start managing! ✅

---

## 📞 Quick Reference

```bash
./setup.sh setup       # Full setup
./setup.sh status      # Show status
./setup.sh logs        # View logs
./setup.sh stop        # Stop services
./setup.sh start       # Start services
./setup.sh restart     # Restart
./setup.sh clean       # Remove everything
```

---

**Ready to go?**
```bash
chmod +x setup.sh
./setup.sh setup
```

Then open: **http://localhost:3000**

---

**Version:** 1.0  
**For:** Fedora 35+  
**Last Updated:** November 28, 2025
