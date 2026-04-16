# BookStore E-commerce Backend - Project Overview

## 🎯 Mục tiêu
Backend REST API cho BookStore E-commerce application, xây dựng với Golang, Fiber framework, và SQLite database.

## ✅ Features đã implement

### Core API Endpoints
- [x] GET /api/books - Lấy danh sách tất cả sách
- [x] GET /api/books/:id - Lấy chi tiết một cuốn sách
- [x] GET /api/books/search?q=query - Tìm kiếm sách theo title/author
- [x] POST /api/orders - Tạo đơn hàng mới

### Technical Features
- [x] SQLite database với auto-migration
- [x] Seed data (12 books)
- [x] CORS middleware cho frontend integration
- [x] Request validation với go-playground/validator
- [x] Error handling
- [x] Logger middleware
- [x] Graceful shutdown
- [x] Transaction support cho orders

## 🏗️ Architecture

### Clean Architecture Layers

```
cmd/server/main.go          # Application entry point
    ↓
handlers/                   # HTTP request handlers
    ↓
repository/                 # Database operations
    ↓
database/                   # SQLite connection & migrations
```

### Components

1. **Models** (`internal/models/`)
   - Book: Định nghĩa structure cho sách
   - Order & OrderItem: Định nghĩa structure cho đơn hàng
   - CreateOrderRequest: DTO cho order creation với validation

2. **Repository Layer** (`internal/repository/`)
   - BookRepository: GetAll, GetByID, Search
   - OrderRepository: Create (với transaction)

3. **Handler Layer** (`internal/handlers/`)
   - BookHandler: Xử lý HTTP requests cho books
   - OrderHandler: Xử lý HTTP requests cho orders

4. **Database Layer** (`internal/database/`)
   - InitDB: Khởi tạo SQLite connection
   - RunMigrations: Chạy SQL migrations
   - SeedBooks: Seed initial data

## 📊 Database Schema

### Tables
- **books**: id, title, author, price, description, cover, category, rating, stock, created_at, updated_at
- **orders**: id, customer_name, customer_email, customer_phone, customer_address, customer_note, total, created_at
- **order_items**: id, order_id, book_id, quantity, price

### Indexes
- idx_books_category
- idx_books_title
- idx_order_items_order_id

## 🔧 Tech Stack

- **Go 1.21+** - Programming Language
- **Fiber v2** - Web Framework (Express-inspired, fast)
- **SQLite3** - Embedded Database
- **sqlx** - SQL Extensions for Go
- **validator/v10** - Struct Validation

## 📦 Dependencies

```
github.com/gofiber/fiber/v2
github.com/jmoiron/sqlx
github.com/mattn/go-sqlite3
github.com/go-playground/validator/v10
```

## 🚀 Running the Server

### Quick Start
```bash
./run.sh
```

### Manual Build
```bash
CGO_ENABLED=1 go build -o bookstore-server cmd/server/main.go
./bookstore-server
```

Server chạy tại: `http://localhost:3000`

## ⚙️ Configuration

Environment variables (optional):
- `PORT` - Server port (default: 3000)
- `DB_PATH` - Database file path (default: ./data/bookstore.db)
- `ALLOWED_ORIGINS` - CORS origins (default: http://localhost:5173)

## 📝 API Examples

### Get all books
```bash
curl http://localhost:3000/api/books
```

### Get book by ID
```bash
curl http://localhost:3000/api/books/1
```

### Search books
```bash
curl "http://localhost:3000/api/books/search?q=harry"
```

### Create order
```bash
curl -X POST http://localhost:3000/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer": {
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "0123456789",
      "address": "123 Street, City",
      "note": "Optional note"
    },
    "items": [
      {"id": 1, "quantity": 2, "price": 12.99}
    ],
    "total": 25.98
  }'
```

## 🌱 Seed Data

Database được seed với 12 books covering các categories:
- Fiction (To Kill a Mockingbird, The Great Gatsby)
- Science Fiction (1984, Dune)
- Romance (Pride and Prejudice)
- Fantasy (The Hobbit, Harry Potter)
- Mystery (The Da Vinci Code)
- Philosophy (The Alchemist)
- Non-Fiction (Sapiens, Thinking Fast and Slow)
- Self-Help (Atomic Habits)

## 🔒 Data Validation

Order creation validates:
- Required fields: customer name, email, phone, address
- Email format validation
- Minimum quantity: 1
- Minimum price: 0
- At least 1 item in order

## 🐛 Known Limitations

1. Không có authentication/authorization
2. Không có user management
3. Không có order history endpoint
4. Không có admin CRUD cho books
5. Stock không được update khi order
6. Không có pagination

## 🚀 Future Enhancements

### Potential Features
- [ ] User authentication & authorization (JWT)
- [ ] Admin endpoints (CRUD books)
- [ ] Order history per user
- [ ] Stock management
- [ ] Pagination & filtering
- [ ] Payment integration
- [ ] Email notifications
- [ ] Rate limiting
- [ ] API documentation (Swagger)
- [ ] Unit tests
- [ ] Docker support

## 📚 Integration với Frontend

Frontend Vue.js cần config:
```env
VITE_API_URL=http://localhost:3000/api
```

Backend đã config CORS để accept requests từ `http://localhost:5173` (Vite dev server).

## 🎉 Success Metrics

✅ Server khởi động thành công
✅ Database được tạo và migrate
✅ 12 books được seed
✅ Tất cả 4 API endpoints hoạt động
✅ CORS configured đúng
✅ Validation hoạt động
✅ Transaction support cho orders
✅ Graceful shutdown implemented
