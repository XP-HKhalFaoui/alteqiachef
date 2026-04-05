# 🚀 AlteqiaChef - Quick Start Guide

## ⚡ 30-Second Setup

### Windows PowerShell
```powershell
cd alteqiachef
.\setup.ps1 -Action setup
```

### Windows Command Prompt / macOS / Linux
```bash
cd alteqiachef/setupchef
docker-compose up -d
docker-compose ps
```

**Then open:** http://localhost:3000

---

## 📍 Application URLs

| Service | URL |
|---------|-----|
| 🌐 Frontend | http://localhost:3000 |
| 🔌 Backend API | http://localhost:8080 |
| 🗄️ Database | localhost:5432 |

---

## 🎮 Essential Commands

### PowerShell (Windows)
```powershell
# Setup
.\setup.ps1 -Action setup

# View status
.\setup.ps1 -Action status

# View logs
.\setup.ps1 -Action logs

# Restart
.\setup.ps1 -Action restart

# Stop
.\setup.ps1 -Action stop
```

### Docker Compose (All Platforms)
```bash
cd setupchef

# Start
docker-compose up -d

# Stop
docker-compose down

# View status
docker-compose ps

# View logs
docker-compose logs -f

# Restart
docker-compose restart
```

### Make Commands (if available)
```bash
make up          # Start
make down        # Stop
make restart     # Restart
make logs        # View logs
make status      # Status
```

---

## 👤 Create Admin User

```bash
cd setupchef
./create-admin.sh
```

Or use the interactive setup script.

---

## 🆘 Quick Troubleshooting

### Containers won't start?
```bash
docker-compose down
docker-compose up -d
```

### Port already in use?
```bash
# Windows: Check and kill process using port
netstat -ano | findstr :3000

# Then restart
docker-compose restart
```

### Database connection error?
```bash
docker-compose logs pos-postgres
docker-compose restart pos-postgres
```

### Reset everything?
```bash
docker-compose down -v
docker-compose up -d
```

---

## 📊 Check Service Health

```bash
# View all running containers
docker-compose ps

# Check specific service
docker-compose logs pos-backend
docker-compose logs pos-frontend
docker-compose logs pos-postgres
```

---

## 🗝️ Default Credentials

| Service | Username | Password |
|---------|----------|----------|
| PostgreSQL | postgres | postgres123 |
| Database | pos_system | - |

---

## 📚 More Information

- **Full Setup Guide:** See `SETUP.md`
- **Development Guide:** See `README.md`
- **Docker Docs:** https://docs.docker.com/

---

**Need help?** Check `SETUP.md` for detailed troubleshooting and information.
