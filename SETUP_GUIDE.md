# BookStore Backend - Setup Guide

## 📋 Prerequisites

- Go 1.21 hoặc cao hơn
- GCC compiler (cho CGO/SQLite)
- curl và jq (cho testing - optional)

### Kiểm tra prerequisites

```bash
go version
gcc --version
```

## 🚀 Installation Steps

### 1. Clone và navigate
```bash
cd /home/kiennt1/TLEcommerce/BE
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Build application
```bash
CGO_ENABLED=1 go build -o bookstore-server cmd/server/main.go
```

### 4. Run server
```bash
./bookstore-server
```

Hoặc sử dụng script:
```bash
./run.sh
```

## ✅ Verification

Server khởi động thành công khi thấy output:
```
2026/04/15 17:05:28 Server starting on http://localhost:3000

 ┌───────────────────────────────────────────────────┐ 
 │                   BookStore API                   │ 
 │                  Fiber v2.52.12                   │ 
 │               http://127.0.0.1:3000               │ 
 └───────────────────────────────────────────────────┘
```

### Test API endpoints
```bash
./test_api.sh
```

Hoặc manual test:
```bash
curl http://localhost:3000/api/books
curl http://localhost:3000/api/books/1
curl "http://localhost:3000/api/books/search?q=harry"
```

## 📁 Project Files

```
BE/
├── cmd/server/main.go              # Entry point
├── internal/
│   ├── database/sqlite.go          # Database setup
│   ├── models/
│   │   ├── book.go                 # Book model
│   │   └── order.go                # Order models
│   ├── repository/
│   │   ├── book_repository.go      # Book data access
│   │   └── order_repository.go     # Order data access
│   └── handlers/
│       ├── book_handler.go         # Book HTTP handlers
│       └── order_handler.go        # Order HTTP handlers
├── migrations/001_init.sql         # Database schema
├── data/bookstore.db               # SQLite database (auto-created)
├── go.mod                          # Go dependencies
├── go.sum                          # Dependency checksums
├── .env.example                    # Environment template
├── .gitignore                      # Git ignore rules
├── run.sh                          # Run script
├── test_api.sh                     # API test script
├── README.md                       # Main documentation
├── PROJECT_OVERVIEW.md             # Project overview
└── SETUP_GUIDE.md                  # This file
```

## ⚙️ Configuration

### Environment Variables

Tạo file `.env` (optional):
```bash
cp .env.example .env
```

Edit `.env`:
```
PORT=3000
DB_PATH=./data/bookstore.db
ALLOWED_ORIGINS=http://localhost:5173
```

### Default Values
- Port: 3000
- Database: ./data/bookstore.db
- CORS Origins: http://localhost:5173

## 🗄️ Database

### Auto-initialization
Database được tự động:
1. Tạo file tại `./data/bookstore.db`
2. Run migrations từ `migrations/001_init.sql`
3. Seed 12 books vào database

### Manual database check
```bash
sqlite3 data/bookstore.db "SELECT COUNT(*) FROM books;"
sqlite3 data/bookstore.db "SELECT * FROM books LIMIT 3;"
```

## 🔧 Troubleshooting

### Error: "CGO_ENABLED=0"
**Problem**: SQLite requires CGO
**Solution**: Always set `CGO_ENABLED=1`:
```bash
CGO_ENABLED=1 go build -o bookstore-server cmd/server/main.go
```

### Error: "gcc not found"
**Problem**: GCC compiler not installed
**Solution**: Install GCC:
```bash
# Ubuntu/Debian
sudo apt-get install build-essential

# Fedora/RHEL
sudo dnf install gcc

# macOS
xcode-select --install
```

### Port already in use
**Problem**: Port 3000 đã được sử dụng
**Solution**: 
1. Kill process: `lsof -ti:3000 | xargs kill`
2. Hoặc đổi port: `PORT=3001 ./bookstore-server`

### Database locked
**Problem**: SQLite database bị lock
**Solution**: 
1. Stop tất cả instances của server
2. Delete `data/bookstore.db`
3. Restart server (sẽ tự tạo lại)

## 🧪 Testing

### Run all tests
```bash
./test_api.sh
```

### Manual testing

**Get all books:**
```bash
curl http://localhost:3000/api/books | jq
```

**Get book by ID:**
```bash
curl http://localhost:3000/api/books/1 | jq
```

**Search books:**
```bash
curl "http://localhost:3000/api/books/search?q=harry" | jq
```

**Create order:**
```bash
curl -X POST http://localhost:3000/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer": {
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "0123456789",
      "address": "123 Street",
      "note": "Test"
    },
    "items": [{"id": 1, "quantity": 2, "price": 12.99}],
    "total": 25.98
  }' | jq
```

## 🔄 Development Workflow

### 1. Make code changes
Edit files trong `internal/` directory

### 2. Rebuild
```bash
CGO_ENABLED=1 go build -o bookstore-server cmd/server/main.go
```

### 3. Restart server
```bash
./bookstore-server
```

### Hot reload (optional)
Install air:
```bash
go install github.com/cosmtrek/air@latest
```

Create `.air.toml` và run:
```bash
air
```

## 📚 Next Steps

1. ✅ Server đang chạy
2. ✅ API endpoints hoạt động
3. ✅ Database có seed data
4. 🔄 Connect với Frontend
5. 🔄 Test full integration

## 🆘 Support

Nếu gặp vấn đề:
1. Check server logs
2. Verify prerequisites installed
3. Check port availability
4. Review error messages
5. Check database file permissions

## 🎉 Success!

Nếu bạn thấy output sau khi chạy `./test_api.sh`:
```
=========================================
All tests completed!
=========================================
```

Backend đã sẵn sàng để integrate với Frontend! 🚀
