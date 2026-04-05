# 🚀 AlteqiaChef - Setup Guide

> **Complete setup instructions for the AlteqiaChef Point of Sale (POS) System**

A modern, enterprise-grade POS system built with **Go** backend and **React** frontend, fully containerized with Docker.

---

## 📋 Table of Contents

1. [System Requirements](#system-requirements)
2. [Quick Start](#quick-start)
3. [Detailed Setup](#detailed-setup)
4. [Docker Setup](#docker-setup)
5. [Development Setup](#development-setup)
6. [Database Management](#database-management)
7. [User Management](#user-management)
8. [Troubleshooting](#troubleshooting)
9. [Environment Variables](#environment-variables)

---

## 📦 System Requirements

### Minimum Requirements
- **Operating System:** Windows, macOS, or Linux
- **Docker:** v20.10+ ([Download](https://www.docker.com/get-started))
- **Docker Compose:** v1.29+
- **RAM:** 4GB minimum (8GB recommended)
- **Disk Space:** 5GB for Docker images and database

### Development Requirements (Optional)
- **Go:** v1.21+ ([Download](https://golang.org/dl/))
- **Node.js:** v18+ ([Download](https://nodejs.org/))
- **PostgreSQL Client:** `psql` command-line tool

---

## ⚡ Quick Start

### 1. **Clone & Navigate to Project**
```bash
cd alteqiachef
cd setupchef
```

### 2. **Start the Application**
```bash
# Using the install script
./install-pos.sh

# OR using Docker Compose directly
docker-compose pull
docker-compose up -d
docker-compose ps
```

### 3. **Access the Application**
- **Frontend:** http://localhost:3000
- **Backend API:** http://localhost:8080
- **PostgreSQL:** localhost:5432

### 4. **Create Admin User**
```bash
cd ..
./setupchef/create-admin.sh
```

---

## 🔧 Detailed Setup

### Option 1: Using Install Script (Recommended)

The `install-pos.sh` script automates the entire setup process:

```bash
#!/bin/bash
# Pull images
docker-compose pull

# Start services
docker-compose up -d

# Check status
docker-compose ps
```

**Run it:**
```bash
cd setupchef
chmod +x install-pos.sh
./install-pos.sh
```

### Option 2: Manual Docker Setup

```bash
# Navigate to setupchef directory
cd setupchef

# Pull the latest images
docker-compose pull

# Start all containers in the background
docker-compose up -d

# Verify all services are running
docker-compose ps

# View logs
docker-compose logs -f
```

### Option 3: Using Makefile Commands

```bash
# From project root
make up          # Start containers
make down        # Stop containers
make restart     # Restart containers
make logs        # View all logs
make status      # Show service status
```

---

## 🐳 Docker Setup

### Docker Compose Configuration

The application uses three main services:

#### **1. PostgreSQL Database**
```yaml
postgres:
  image: postgres:15-alpine
  container_name: pos-postgres
  ports:
    - "5432:5432"
  environment:
    POSTGRES_DB: pos_system
    POSTGRES_USER: postgres
    POSTGRES_PASSWORD: postgres123
```

#### **2. Go Backend API**
```yaml
backend:
  image: docker23xis/alteqiachef-backend:latest
  container_name: pos-backend
  ports:
    - "8080:8080"
  environment:
    - DB_HOST=postgres
    - DB_PORT=5432
    - DB_USER=postgres
    - DB_PASSWORD=postgres123
    - DB_NAME=pos_system
```

#### **3. React Frontend**
```yaml
frontend:
  image: docker23xis/alteqiachef-frontend:latest
  container_name: pos-frontend
  ports:
    - "3000:3000"
  environment:
    - VITE_API_URL=http://localhost:8080
```

### Docker Compose Commands

```bash
# Start containers
docker-compose up -d

# Stop containers
docker-compose down

# View running containers
docker-compose ps

# View container logs
docker-compose logs -f

# Rebuild images (if needed)
docker-compose build --no-cache

# Stop and remove all (WARNING: removes volumes)
docker-compose down -v
```

---

## 💻 Development Setup

### For Backend Development (Go)

```bash
# Navigate to backend directory
cd backend

# Install dependencies
go mod download
go mod tidy

# Run the application (requires PostgreSQL running)
go run main.go

# Run tests
go test ./...

# Build binary
go build -o main

# Hot reload with Air (optional)
go install github.com/cosmtrek/air@latest
air
```

### For Frontend Development (React)

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
pnpm install

# Start development server
pnpm dev

# Build for production
pnpm build

# Run tests
pnpm test

# Lint code
pnpm lint
```

### Environment Variables for Development

**Backend (.env in `/backend`)**
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres123
DB_NAME=pos_system
```

**Frontend (.env in `/frontend`)**
```env
VITE_API_URL=http://localhost:8080
```

---

## 🗄️ Database Management

### Reset Database

```bash
# Using the script
./setupchef/db-reset.sh

# This will:
# 1. Drop all existing tables and data
# 2. Recreate database schema
# 3. Load fresh seed data
```

### Access PostgreSQL Shell

```bash
# Using make command
make db-shell

# OR using Docker directly
docker exec -it pos-postgres psql -U postgres -d pos_system

# Common commands:
\dt              # List all tables
\du              # List all users
SELECT * FROM users;  # View users
\q               # Exit
```

### Backup Database

```bash
# Using script
./setupchef/backup.sh

# Manual backup
docker exec pos-postgres pg_dump -U postgres -d pos_system > backup.sql
```

### Restore Database

```bash
# Using script
./setupchef/restore.sh

# Manual restore
docker exec -i pos-postgres psql -U postgres -d pos_system < backup.sql
```

---

## 👥 User Management

### Create Super Admin User

```bash
./setupchef/create-admin.sh
```

The script will prompt you for:
- Username
- Email
- First Name
- Last Name
- Password

### Create Demo Users

```bash
./setupchef/create-demo-users.sh
```

This creates sample users for testing:
- **Admin Account**
- **Server Account**
- **Kitchen Account**
- **Counter Account**

### Remove All User Data

```bash
# WARNING: This removes all data!
./setupchef/remove-data.sh
```

---

## 🔍 Troubleshooting

### Container Won't Start

**Problem:** Containers fail to start or keep restarting

**Solution:**
```bash
# Check container logs
docker-compose logs pos-backend
docker-compose logs pos-frontend
docker-compose logs pos-postgres

# Remove containers and start fresh
docker-compose down -v
docker-compose pull
docker-compose up -d
```

### Database Connection Error

**Problem:** Backend can't connect to database

**Solution:**
```bash
# Verify database container is running
docker-compose ps

# Check database connectivity
docker exec pos-backend psql -h postgres -U postgres -d pos_system -c "SELECT 1;"

# Reset database
./setupchef/db-reset.sh
```

### Port Already in Use

**Problem:** Ports 3000, 8080, or 5432 are already in use

**Solution:**
```bash
# Find process using port (macOS/Linux)
lsof -i :3000
lsof -i :8080
lsof -i :5432

# Kill process
kill -9 <PID>

# Or change ports in docker-compose.yml
# Then restart containers
docker-compose down
docker-compose up -d
```

### Frontend Can't Connect to Backend

**Problem:** Frontend shows connection errors

**Solution:**
```bash
# Check VITE_API_URL environment variable
docker-compose logs pos-frontend

# Verify backend is running
curl http://localhost:8080/api/v1/health

# Update VITE_API_URL in docker-compose.yml if needed
```

### Out of Disk Space

**Problem:** Docker complains about disk space

**Solution:**
```bash
# Clean up Docker resources
docker system prune -a

# Remove all volumes (WARNING: deletes data)
docker volume prune

# Check disk usage
docker system df
```

---

## 🔐 Environment Variables

### Backend Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Backend API port |
| `DB_HOST` | `postgres` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `postgres123` | Database password |
| `DB_NAME` | `pos_system` | Database name |

### Frontend Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `VITE_API_URL` | `http://localhost:8080` | Backend API base URL |

### Database Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `POSTGRES_DB` | `pos_system` | Database name |
| `POSTGRES_USER` | `postgres` | Database user |
| `POSTGRES_PASSWORD` | `postgres123` | Database password |

---

## 📊 Service Architecture

```
┌─────────────────────────────────────┐
│      AlteqiaChef POS System         │
├─────────────────────────────────────┤
│                                     │
│  ┌──────────────┐  ┌─────────────┐ │
│  │   Frontend   │  │   Backend   │ │
│  │   (React)    │  │   (Go Gin)  │ │
│  │ :3000        │  │   :8080     │ │
│  └──────────────┘  └─────────────┘ │
│         │               │           │
│         └───────┬───────┘           │
│                 │                   │
│         ┌───────▼────────┐          │
│         │   PostgreSQL   │          │
│         │   :5432        │          │
│         │   pos_system   │          │
│         └────────────────┘          │
│                                     │
└─────────────────────────────────────┘
```

---

## ✅ Verification Checklist

After setup, verify everything is working:

- [ ] All three containers are running: `docker-compose ps`
- [ ] Frontend accessible: http://localhost:3000
- [ ] Backend accessible: http://localhost:8080/api/v1/health
- [ ] Database responding: `make db-shell` or similar command
- [ ] Admin user created and can login
- [ ] Can access different roles (Admin, Server, Kitchen, Counter)

---

## 📝 Common Tasks

### View Application Logs
```bash
make logs              # All logs
make logs-backend      # Backend only
make logs-frontend     # Frontend only
make logs-db           # Database only
```

### Restart Services
```bash
make restart           # Restart all services
docker-compose restart pos-backend  # Restart specific service
```

### Database Shell Access
```bash
make db-shell
# Or directly
docker exec -it pos-postgres psql -U postgres -d pos_system
```

### Remove All Data (Destructive)
```bash
make clean             # Remove volumes and images
./setupchef/remove-data.sh  # Remove data only
```

---

## 🆘 Getting Help

If you encounter issues:

1. Check the logs: `docker-compose logs`
2. Review this troubleshooting section
3. Verify all ports are available
4. Ensure Docker and Docker Compose are properly installed
5. Check that you have sufficient disk space and RAM

---

## 📚 Additional Resources

- [Docker Documentation](https://docs.docker.com/)
- [Go Documentation](https://golang.org/doc/)
- [React Documentation](https://react.dev)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

---

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Last Updated:** November 28, 2025
**Version:** 1.0
