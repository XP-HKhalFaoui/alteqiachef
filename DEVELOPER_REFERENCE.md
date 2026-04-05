# 👨‍💻 AlteqiaChef - Developer Quick Reference

Quick reference for common development tasks and commands.

---

## 🎯 Quick Commands

### Start Application

**Windows PowerShell:**
```powershell
.\setup.ps1 -Action setup
```

**Windows CMD:**
```cmd
setup.bat setup
```

**All Platforms:**
```bash
cd setupchef && docker-compose up -d
```

### Stop Application

```bash
cd setupchef && docker-compose down
```

### View Status

**PowerShell:**
```powershell
.\setup.ps1 -Action status
```

**All Platforms:**
```bash
cd setupchef && docker-compose ps
```

### View Logs

**PowerShell:**
```powershell
.\setup.ps1 -Action logs
```

**All Platforms:**
```bash
cd setupchef && docker-compose logs -f
```

---

## 🛠️ Docker Commands

| Command | Purpose |
|---------|---------|
| `docker-compose up -d` | Start containers in background |
| `docker-compose down` | Stop and remove containers |
| `docker-compose ps` | View running containers |
| `docker-compose logs -f` | Stream logs |
| `docker-compose logs <service>` | View specific service logs |
| `docker-compose restart` | Restart all containers |
| `docker-compose pull` | Pull latest images |
| `docker exec -it <container> bash` | Enter container shell |

---

## 🗄️ Database Commands

### Access Database Shell

```bash
# Using docker exec
docker exec -it pos-postgres psql -U postgres -d pos_system
```

### Common SQL Queries

```sql
-- List all tables
\dt

-- List all users
\du

-- View users table
SELECT id, username, email, role, created_at FROM users;

-- Count orders
SELECT COUNT(*) FROM orders;

-- View recent orders
SELECT * FROM orders ORDER BY created_at DESC LIMIT 10;

-- Exit
\q
```

### Backup Database

```bash
docker exec pos-postgres pg_dump -U postgres -d pos_system > backup.sql
```

### Reset Database

```bash
cd setupchef
./db-reset.sh
```

---

## 🔙 Backend (Go) Development

### Local Development Setup

```bash
cd backend

# Install dependencies
go mod download
go mod tidy

# Run locally (requires PostgreSQL running)
go run main.go

# Run tests
go test ./...

# Format code
go fmt ./...

# Lint code
golangci-lint run
```

### Environment Variables

Create `.env` or set environment:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres123
DB_NAME=pos_system
```

### API Endpoints

| Method | Endpoint | Purpose |
|--------|----------|---------|
| `GET` | `/api/v1/health` | Health check |
| `POST` | `/api/v1/auth/login` | User login |
| `POST` | `/api/v1/auth/register` | User registration |
| `GET` | `/api/v1/orders` | List orders |
| `POST` | `/api/v1/orders` | Create order |
| `GET` | `/api/v1/products` | List products |

---

## ⚛️ Frontend (React) Development

### Local Development Setup

```bash
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

# Format code
pnpm format
```

### Environment Variables

Create `.env.local`:

```env
VITE_API_URL=http://localhost:8080
```

### Project Structure

```
frontend/
├── src/
│   ├── api/           # API client
│   ├── components/    # React components
│   ├── hooks/         # Custom React hooks
│   ├── lib/           # Utility functions
│   ├── routes/        # TanStack Router routes
│   ├── services/      # Services (sound, etc)
│   ├── types/         # TypeScript types
│   └── main.tsx       # Entry point
├── public/            # Static assets
└── package.json       # Dependencies
```

---

## 🧪 Testing

### Run Backend Tests

```bash
cd backend
go test ./...
go test -v ./...        # Verbose
go test -cover ./...    # With coverage
```

### Run Frontend Tests

```bash
cd frontend
pnpm test
pnpm test --ui         # With UI
pnpm test --coverage   # With coverage
```

---

## 📦 Building

### Build Backend Image

```bash
cd backend
docker build -f Dockerfile -t alteqiachef-backend:latest .
```

### Build Frontend Image

```bash
cd frontend
docker build -f Dockerfile -t alteqiachef-frontend:latest .
```

### Push to Registry

```bash
# Tag image
docker tag alteqiachef-backend:latest docker23xis/alteqiachef-backend:latest

# Push to Docker Hub
docker push docker23xis/alteqiachef-backend:latest
```

---

## 🔍 Debugging

### Check Backend Health

```bash
curl http://localhost:8080/api/v1/health
```

### View Backend Logs

```bash
docker-compose logs pos-backend -f
```

### View Frontend Logs

```bash
docker-compose logs pos-frontend -f
```

### Enter Backend Container

```bash
docker exec -it pos-backend bash
```

### Enter Frontend Container

```bash
docker exec -it pos-frontend bash
```

### Database Connection Test

```bash
docker exec pos-postgres psql -U postgres -d pos_system -c "SELECT 1;"
```

---

## 📝 Git Workflow

### Before Committing

1. **Format code:**
   ```bash
   # Backend
   cd backend && go fmt ./...
   
   # Frontend
   cd frontend && pnpm format
   ```

2. **Run tests:**
   ```bash
   # Backend
   cd backend && go test ./...
   
   # Frontend
   cd frontend && pnpm test
   ```

3. **Run linter:**
   ```bash
   # Backend
   cd backend && golangci-lint run
   
   # Frontend
   cd frontend && pnpm lint
   ```

### Common Git Commands

```bash
# Check status
git status

# Add changes
git add .

# Commit changes
git commit -m "feat: description of changes"

# Push to remote
git push origin main

# Pull latest
git pull origin main

# Create branch
git checkout -b feature/your-feature
```

---

## 🚀 Deployment

### Production Build

```bash
# Build images
docker-compose build

# Or build specific service
docker build -f backend/Dockerfile -t alteqiachef-backend:1.0 ./backend
docker build -f frontend/Dockerfile -t alteqiachef-frontend:1.0 ./frontend

# Push to registry
docker push alteqiachef-backend:1.0
docker push alteqiachef-frontend:1.0
```

### Update Production

```bash
# Pull latest images
docker-compose pull

# Restart services
docker-compose restart

# Or full restart
docker-compose down
docker-compose up -d
```

---

## 📊 Performance Monitoring

### Monitor Container Resources

```bash
docker stats
```

### Check Disk Usage

```bash
docker system df
```

### Monitor Logs in Real-time

```bash
docker-compose logs -f --tail 100
```

---

## 🧹 Cleanup

### Clean Unused Images

```bash
docker image prune -a
```

### Clean Unused Volumes

```bash
docker volume prune
```

### Clean Everything

```bash
docker system prune -a
```

### Full Reset (⚠️ Removes all data)

```bash
cd setupchef
docker-compose down -v
docker-compose up -d
```

---

## 📚 File Locations

| Component | Location |
|-----------|----------|
| Backend Code | `/backend` |
| Frontend Code | `/frontend` |
| Database Schema | `/database/init/01_schema.sql` |
| Seed Data | `/database/init/02_seed_data.sql` |
| Docker Config | `/setupchef/docker-compose.yml` |
| Setup Scripts | `/setupchef/*.sh` |
| Configuration | `.env` (root) |

---

## 🔗 Useful Links

- [Backend API Routes](../backend/internal/api/routes.go)
- [Database Schema](../database/init/01_schema.sql)
- [Frontend Routes](../frontend/src/routes/__root.tsx)
- [Environment Setup](./ENVIRONMENT_SETUP.md)
- [Full Setup Guide](./SETUP.md)

---

## 💡 Tips & Tricks

### Faster Rebuilds

Use BuildKit for faster Docker builds:
```bash
export DOCKER_BUILDKIT=1
docker build .
```

### Persistent Shell History

Mount bash history:
```bash
docker exec -it <container> bash
history  # Now persists
```

### View Real-time Container Activity

```bash
docker-compose logs -f
docker stats
```

### Database Snapshots

```bash
# Before making changes
docker-compose exec postgres pg_dump -U postgres -d pos_system > before.sql

# Make changes...

# If needed, restore
docker-compose exec -T postgres psql -U postgres -d pos_system < before.sql
```

---

## 🆘 Common Issues

### Port Already in Use

```powershell
# Windows PowerShell - Find and kill process
Get-Process -Id (Get-NetTCPConnection -LocalPort 3000).OwningProcess | Stop-Process
```

### Containers Keep Crashing

```bash
# Check logs
docker-compose logs

# Rebuild from scratch
docker-compose down -v
docker-compose build --no-cache
docker-compose up -d
```

### Database Won't Connect

```bash
# Verify database is running
docker-compose ps

# Check database logs
docker-compose logs pos-postgres

# Reset database
./setupchef/db-reset.sh
```

---

**Quick Links:**
- 📖 [Full Setup Guide](./SETUP.md)
- ⚡ [Quick Start](./QUICKSTART.md)
- 🔧 [Environment Setup](./ENVIRONMENT_SETUP.md)

---

**Version:** 1.0
**Last Updated:** November 28, 2025
