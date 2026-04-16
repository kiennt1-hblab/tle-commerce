# BookStore E-commerce Backend

REST API backend cho BookStore E-commerce được xây dựng với Golang, Fiber framework, và SQLite database.

## 🛠️ Tech Stack

- **Go 1.21+** - Programming Language
- **Fiber v2** - Web Framework
- **SQLite3** - Database
- **sqlx** - SQL Extensions for Go
- **validator** - Struct Validation

## 📦 Installation

```bash
cd BE

go mod download

go run cmd/server/main.go
```

## 🚀 Running

### Option 1: Using run script (recommended)
```bash
./run.sh
```

### Option 2: Manual build and run
```bash
CGO_ENABLED=1 go build -o bookstore-server cmd/server/main.go
./bookstore-server
```

### Option 3: Direct run with go run
```bash
CGO_ENABLED=1 go run cmd/server/main.go
```

Server sẽ chạy tại `http://localhost:3000`

## 📁 Project Structure

```
BE/
├── cmd/
│   └── server/
│       └── main.go           # Entry point
├── internal/
│   ├── handlers/             # HTTP handlers
│   │   ├── book_handler.go
│   │   └── order_handler.go
│   ├── models/               # Data models
│   │   ├── book.go
│   │   └── order.go
│   ├── repository/           # Database layer
│   │   ├── book_repository.go
│   │   └── order_repository.go
│   └── database/             # Database setup
│       └── sqlite.go
├── migrations/               # SQL migration files
│   └── 001_init.sql
├── data/                     # SQLite database file
│   └── bookstore.db
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## 🔌 API Endpoints

### Health Check

- `GET /health` - Server health status

### Books

- `GET /api/books` - Lấy danh sách tất cả sách
- `GET /api/books/:id` - Lấy chi tiết một cuốn sách
- `GET /api/books/search?q=query` - Tìm kiếm sách theo title/author

### Orders

- `POST /api/orders` - Tạo đơn hàng mới

#### Request Body Example:
```json
{
  "customer": {
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "0123456789",
    "address": "123 Street, City",
    "note": "Optional note"
  },
  "items": [
    {
      "id": 1,
      "quantity": 2,
      "price": 12.99
    }
  ],
  "total": 25.98
}
```

## ⚙️ Configuration

Tạo file `.env` hoặc set environment variables:

```bash
PORT=3000
DB_PATH=./data/bookstore.db
ALLOWED_ORIGINS=http://localhost:5173
```

## 📊 Database

Database sử dụng SQLite3 với 3 tables:
- `books` - Lưu thông tin sách
- `orders` - Lưu thông tin đơn hàng
- `order_items` - Lưu chi tiết items trong đơn hàng

Database được tự động migrate và seed với 12 cuốn sách khi khởi động lần đầu.

## 🔧 Development

```bash
go run cmd/server/main.go
```

## 🏗️ Build

### Using Make
```bash
make build
make run
```

### Manual Build
```bash
go build -o bookstore-server cmd/server/main.go
./bookstore-server
```

### Using Docker
```bash
docker build -t bookstore-backend .
docker run -p 3000:3000 bookstore-backend
```

## 📝 Notes

- Database file sẽ được tạo tự động tại `./data/bookstore.db`
- Migrations chạy tự động khi server start
- Seed data chỉ chạy nếu database chưa có books
- CORS được config để cho phép frontend Vue.js gọi API
