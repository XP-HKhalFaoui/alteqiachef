# 🍽️ AlteqiaChef - Complete Setup & Documentation Index

Welcome to AlteqiaChef! This document serves as your central hub for all setup and configuration information.

---

## 📚 Documentation Overview

### 🚀 Getting Started (Choose Your Path)

#### **I just want to run it (5 minutes)**
→ Read: **[QUICKSTART.md](./QUICKSTART.md)**
- Fastest way to get up and running
- Essential commands only
- Quick troubleshooting

#### **I want detailed setup instructions**
→ Read: **[SETUP.md](./SETUP.md)**
- Complete setup guide
- All services explained
- Comprehensive troubleshooting
- Database management
- User management

#### **I need to set up my environment first**
→ Read: **[ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md)**
- Docker installation for Windows, macOS, Linux
- System requirements
- Verification steps
- Environment troubleshooting

#### **I'm a developer**
→ Read: **[DEVELOPER_REFERENCE.md](./DEVELOPER_REFERENCE.md)**
- Backend (Go) development setup
- Frontend (React) development setup
- Testing and building
- Git workflow
- Debugging tips

---

## 🎯 Quick Navigation

### By Use Case

| I want to... | Read... | Time |
|--------------|---------|------|
| Run the app quickly | [QUICKSTART.md](./QUICKSTART.md) | 5 min |
| Understand full setup | [SETUP.md](./SETUP.md) | 15 min |
| Set up Docker | [ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md) | 20 min |
| Develop locally | [DEVELOPER_REFERENCE.md](./DEVELOPER_REFERENCE.md) | 30 min |
| Create an admin | [SETUP.md #User Management](./SETUP.md#-user-management) | 5 min |
| Reset database | [SETUP.md #Database Management](./SETUP.md#-database-management) | 2 min |
| Fix issues | [SETUP.md #Troubleshooting](./SETUP.md#-troubleshooting) | varies |

### By Experience Level

#### 🟢 Beginner
1. Install Docker → [ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md)
2. Run application → [QUICKSTART.md](./QUICKSTART.md)
3. Create admin user → [SETUP.md](./SETUP.md)
4. Access app at http://localhost:3000

#### 🟡 Intermediate
1. Follow Beginner path
2. Review [SETUP.md](./SETUP.md) for detailed architecture
3. Learn user management: [SETUP.md #User Management](./SETUP.md#-user-management)
4. Practice database operations: [SETUP.md #Database Management](./SETUP.md#-database-management)

#### 🔴 Advanced
1. Set up development environment → [DEVELOPER_REFERENCE.md](./DEVELOPER_REFERENCE.md)
2. Clone repository and work locally
3. Modify code and test
4. Build custom images
5. Deploy to production

---

## 🚀 Setup Scripts

### Windows PowerShell

```powershell
# Setup (pull images, start containers)
.\setup.ps1 -Action setup

# View status
.\setup.ps1 -Action status

# View logs
.\setup.ps1 -Action logs

# Stop services
.\setup.ps1 -Action stop

# Restart services
.\setup.ps1 -Action restart
```

### Windows Command Prompt

```cmd
# Setup
setup.bat setup

# View status
setup.bat status

# View logs
setup.bat logs

# Stop services
setup.bat stop

# Restart services
setup.bat restart
```

### macOS / Linux

```bash
cd setupchef

# Start services
docker-compose up -d

# View status
docker-compose ps

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

---

## 📍 Application URLs

| Service | URL |
|---------|-----|
| 🌐 **Frontend** | http://localhost:3000 |
| 🔌 **Backend API** | http://localhost:8080 |
| 🗄️ **Database** | localhost:5432 |

---

## 🗂️ Project Structure

```
alteqiachef/
├── 📖 Documentation
│   ├── SETUP.md                    # Full setup guide
│   ├── QUICKSTART.md               # 5-minute quick start
│   ├── ENVIRONMENT_SETUP.md        # Docker & environment setup
│   ├── DEVELOPER_REFERENCE.md      # Developer quick reference
│   ├── README.md                   # Project overview
│   └── docs/                       # Additional documentation
│
├── 🐳 Docker & Deployment
│   ├── setupchef/                  # Setup scripts and docker-compose
│   │   ├── docker-compose.yml
│   │   ├── install-pos.sh
│   │   ├── create-admin.sh
│   │   ├── db-reset.sh
│   │   └── ...
│   ├── docker-compose.yml          # Production setup
│   ├── docker-compose.dev.yml      # Development setup
│   ├── setup.ps1                   # Windows PowerShell script
│   └── setup.bat                   # Windows CMD script
│
├── 🔙 Backend (Go + Gin)
│   ├── backend/
│   │   ├── main.go
│   │   ├── go.mod
│   │   ├── internal/
│   │   │   ├── api/                # HTTP routes
│   │   │   ├── handlers/           # Request handlers
│   │   │   ├── models/             # Data models
│   │   │   ├── repository/         # Database queries
│   │   │   └── middleware/         # Auth & validation
│   │   └── Dockerfile
│
├── ⚛️ Frontend (React + TypeScript)
│   ├── frontend/
│   │   ├── src/
│   │   │   ├── api/                # API client
│   │   │   ├── components/         # React components
│   │   │   ├── routes/             # TanStack Router
│   │   │   ├── hooks/              # Custom hooks
│   │   │   ├── types/              # TypeScript types
│   │   │   └── main.tsx
│   │   ├── package.json
│   │   └── Dockerfile
│
├── 🗄️ Database
│   └── database/
│       └── init/
│           ├── 01_schema.sql       # Database schema
│           └── 02_seed_data.sql    # Sample data
│
└── 📚 Additional Files
    ├── Makefile                    # Make commands
    ├── LICENSE                     # MIT License
    └── start.sh                    # Shell startup script
```

---

## ⚡ Common Commands

### Start Application

```bash
# Windows PowerShell
.\setup.ps1 -Action setup

# All platforms
cd setupchef && docker-compose up -d
```

### Stop Application

```bash
cd setupchef && docker-compose down
```

### View Logs

```bash
cd setupchef && docker-compose logs -f
```

### Access Database

```bash
docker exec -it pos-postgres psql -U postgres -d pos_system
```

### Create Admin User

```bash
cd setupchef && ./create-admin.sh
```

### Reset Database

```bash
cd setupchef && ./db-reset.sh
```

---

## 🔧 System Requirements

| Component | Minimum | Recommended |
|-----------|---------|-------------|
| **RAM** | 4GB | 8GB+ |
| **Disk Space** | 5GB | 10GB+ |
| **CPU** | Dual-core | Quad-core+ |
| **OS** | Windows 10/11, macOS, Linux | Modern versions |
| **Docker** | 20.10+ | Latest |

---

## 📋 Setup Checklist

- [ ] **Environment Setup**
  - [ ] Docker installed
  - [ ] Docker Compose installed
  - [ ] Ports 3000, 8080, 5432 available
  - [ ] 4GB+ RAM available

- [ ] **Initial Setup**
  - [ ] Run setup script (PowerShell/CMD/bash)
  - [ ] Verify all containers running: `docker-compose ps`
  - [ ] Access frontend: http://localhost:3000

- [ ] **User Setup**
  - [ ] Create admin user: `./create-admin.sh`
  - [ ] Login with admin credentials
  - [ ] Create demo users if needed

- [ ] **Verification**
  - [ ] Frontend loads and is responsive
  - [ ] Backend API is accessible
  - [ ] Database contains initial data
  - [ ] Can navigate between different roles

---

## 🆘 Troubleshooting Guide

### Issue: Containers won't start
**Solution:** Check [SETUP.md #Troubleshooting](./SETUP.md#-troubleshooting)

### Issue: Port already in use
**Solution:** Check [QUICKSTART.md #Troubleshooting](./QUICKSTART.md#quick-troubleshooting)

### Issue: Database connection error
**Solution:** Check [SETUP.md #Database Management](./SETUP.md#-database-management)

### Issue: Frontend can't connect to backend
**Solution:** Check [SETUP.md #Troubleshooting](./SETUP.md#-troubleshooting)

### Issue: Docker not installed
**Solution:** Follow [ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md)

---

## 📚 Technology Stack

### Backend
- **Language:** Go 1.21+
- **Framework:** Gin
- **Database:** PostgreSQL 15
- **API:** RESTful with JWT authentication

### Frontend
- **Framework:** React 18.3+
- **Language:** TypeScript 5.6+
- **UI Library:** shadcn/ui
- **Styling:** TailwindCSS
- **Routing:** TanStack Router

### Infrastructure
- **Containerization:** Docker
- **Orchestration:** Docker Compose
- **Database:** PostgreSQL 15 Alpine

---

## 🔗 Useful Resources

### Official Documentation
- [Docker Docs](https://docs.docker.com/)
- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://react.dev)
- [PostgreSQL Docs](https://www.postgresql.org/docs/)

### This Project
- [Full Setup Guide](./SETUP.md)
- [Quick Start](./QUICKSTART.md)
- [Environment Setup](./ENVIRONMENT_SETUP.md)
- [Developer Reference](./DEVELOPER_REFERENCE.md)
- [Project README](./README.md)

### Related Files
- [Docker Compose Config](./setupchef/docker-compose.yml)
- [Database Schema](./database/init/01_schema.sql)
- [API Routes](./backend/internal/api/routes.go)

---

## 💬 Support & Questions

If you encounter issues:

1. **Check the relevant guide** based on your use case
2. **Review troubleshooting sections** in the documentation
3. **Check logs:** `docker-compose logs`
4. **Verify Docker installation:** `docker --version`
5. **Test connectivity:** `curl http://localhost:8080/api/v1/health`

---

## 📄 Document Versions

| Document | Version | Last Updated |
|----------|---------|--------------|
| SETUP.md | 1.0 | Nov 28, 2025 |
| QUICKSTART.md | 1.0 | Nov 28, 2025 |
| ENVIRONMENT_SETUP.md | 1.0 | Nov 28, 2025 |
| DEVELOPER_REFERENCE.md | 1.0 | Nov 28, 2025 |
| setup.ps1 | 1.0 | Nov 28, 2025 |
| setup.bat | 1.0 | Nov 28, 2025 |

---

## 📝 Getting Help

### Documentation Index
This file you're reading now serves as a comprehensive index to all setup documentation.

### By Platform
- **Windows:** See [ENVIRONMENT_SETUP.md #Windows Setup](./ENVIRONMENT_SETUP.md#-windows-setup)
- **macOS:** See [ENVIRONMENT_SETUP.md #macOS Setup](./ENVIRONMENT_SETUP.md#-macos-setup)
- **Linux:** See [ENVIRONMENT_SETUP.md #Linux Setup](./ENVIRONMENT_SETUP.md#-linux-setup)

### By Role
- **End Users:** Follow [QUICKSTART.md](./QUICKSTART.md)
- **DevOps Engineers:** Follow [SETUP.md](./SETUP.md)
- **Developers:** Follow [DEVELOPER_REFERENCE.md](./DEVELOPER_REFERENCE.md)

---

## ✅ Next Steps

1. **Choose your setup method:**
   - New user? → Read [QUICKSTART.md](./QUICKSTART.md)
   - Need Docker? → Read [ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md)
   - Want details? → Read [SETUP.md](./SETUP.md)

2. **Install Docker** (if not already installed)

3. **Run the setup script:**
   ```powershell
   # PowerShell
   .\setup.ps1 -Action setup
   ```

4. **Access the application:**
   - Open http://localhost:3000

5. **Create admin user:**
   ```bash
   cd setupchef && ./create-admin.sh
   ```

---

## 📞 Quick Links

| Need | Link |
|------|------|
| Quick start | [QUICKSTART.md](./QUICKSTART.md) |
| Full setup | [SETUP.md](./SETUP.md) |
| Docker setup | [ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md) |
| Developer info | [DEVELOPER_REFERENCE.md](./DEVELOPER_REFERENCE.md) |
| Project info | [README.md](./README.md) |

---

## 🎉 Welcome to AlteqiaChef!

You now have everything you need to set up and run the complete POS system. Start with the guide that matches your needs and you'll be up and running in minutes!

**Happy building! 🚀**

---

**Created:** November 28, 2025
**Version:** 1.0
**Status:** Ready to Use
