# BookStore Backend - Deployment Guide

## 🐳 Docker Deployment

### Prerequisites
- Docker installed
- Docker Compose installed (optional)

### Build Docker Image

```bash
cd BE
docker build -t bookstore-backend .
```

### Run Container

```bash
docker run -d \
  --name bookstore-api \
  -p 3000:3000 \
  -e PORT=3000 \
  -e ALLOWED_ORIGINS=http://localhost:5173 \
  -v $(pwd)/data:/root/data \
  bookstore-backend
```

### Check Container Status

```bash
docker ps
docker logs bookstore-api
```

### Test API

```bash
curl http://localhost:3000/health
curl http://localhost:3000/api/books
```

### Stop Container

```bash
docker stop bookstore-api
docker rm bookstore-api
```

## 🚀 Docker Compose (Full Stack)

### Start All Services

```bash
# From project root
docker-compose up -d
```

### Check Status

```bash
docker-compose ps
docker-compose logs -f backend
```

### Stop All Services

```bash
docker-compose down
```

### Rebuild After Changes

```bash
docker-compose up -d --build
```

## 📦 Production Build

### Build Optimized Binary

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build \
  -ldflags="-w -s" \
  -o bookstore-server \
  cmd/server/main.go
```

### Run with Systemd

Create `/etc/systemd/system/bookstore.service`:

```ini
[Unit]
Description=BookStore API Server
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/bookstore
ExecStart=/opt/bookstore/bookstore-server
Restart=on-failure
RestartSec=5s

Environment="PORT=3000"
Environment="DB_PATH=/opt/bookstore/data/bookstore.db"
Environment="ALLOWED_ORIGINS=https://yourdomain.com"

[Install]
WantedBy=multi-user.target
```

Enable and start:

```bash
sudo systemctl enable bookstore
sudo systemctl start bookstore
sudo systemctl status bookstore
```

## ☁️ Cloud Deployment

### Heroku

```bash
# Install Heroku CLI
heroku login
heroku create bookstore-api

# Add buildpack for Go
heroku buildpacks:set heroku/go

# Deploy
git push heroku main

# Check logs
heroku logs --tail
```

### Railway

```bash
# Install Railway CLI
npm i -g @railway/cli
railway login

# Initialize project
railway init
railway up

# Set environment variables
railway variables set PORT=3000
railway variables set ALLOWED_ORIGINS=*
```

### Fly.io

Create `fly.toml`:

```toml
app = "bookstore-api"

[build]
  dockerfile = "Dockerfile"

[[services]]
  internal_port = 3000
  protocol = "tcp"

  [[services.ports]]
    handlers = ["http"]
    port = 80

  [[services.ports]]
    handlers = ["tls", "http"]
    port = 443
```

Deploy:

```bash
fly launch
fly deploy
fly open
```

## 🔒 Production Checklist

- [ ] Set proper CORS origins
- [ ] Use environment variables for sensitive data
- [ ] Enable HTTPS
- [ ] Set up database backups
- [ ] Configure logging
- [ ] Set up monitoring
- [ ] Add rate limiting
- [ ] Implement authentication
- [ ] Use reverse proxy (nginx/caddy)
- [ ] Set up CI/CD pipeline

## 📊 Monitoring

### Health Check Endpoint

```bash
curl http://localhost:3000/health
```

Response:
```json
{
  "status": "healthy",
  "message": "Server is running",
  "time": "2026-04-15T17:23:52+07:00",
  "database": "connected",
  "book_count": 12
}
```

### Metrics to Monitor

- Response time
- Error rate
- Database connection status
- Memory usage
- CPU usage
- Request count

## 🔧 Environment Variables

### Required
- `PORT` - Server port (default: 3000)
- `DB_PATH` - Database file path
- `ALLOWED_ORIGINS` - CORS allowed origins

### Optional
- `LOG_LEVEL` - Logging level
- `MAX_CONNECTIONS` - Max database connections

## 📝 Backup Strategy

### Database Backup

```bash
# Backup
cp data/bookstore.db backups/bookstore-$(date +%Y%m%d).db

# Restore
cp backups/bookstore-20260415.db data/bookstore.db
```

### Automated Backup Script

```bash
#!/bin/bash
BACKUP_DIR="/backups"
DB_PATH="./data/bookstore.db"
DATE=$(date +%Y%m%d_%H%M%S)

mkdir -p $BACKUP_DIR
cp $DB_PATH $BACKUP_DIR/bookstore-$DATE.db

# Keep only last 7 days
find $BACKUP_DIR -name "bookstore-*.db" -mtime +7 -delete
```

## 🚨 Troubleshooting

### Container Won't Start

```bash
docker logs bookstore-api
docker inspect bookstore-api
```

### Database Issues

```bash
# Check database file
ls -lh data/bookstore.db

# Verify database
sqlite3 data/bookstore.db "PRAGMA integrity_check;"
```

### Port Already in Use

```bash
# Find process using port
lsof -i :3000

# Kill process
kill -9 <PID>
```

## 📚 Additional Resources

- [Docker Documentation](https://docs.docker.com/)
- [Go Deployment Best Practices](https://golang.org/doc/)
- [SQLite Production Considerations](https://www.sqlite.org/whentouse.html)
