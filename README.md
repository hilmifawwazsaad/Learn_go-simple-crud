# Go Simple CRUD API

API sederhana untuk belajar CRUD (Create, Read, Update, Delete) menggunakan Go.

## Live Demo

**https://go-simple-crud.zeabur.app/**

## Struktur Project

```
go-simple-crud/
├── main.go              # Entry point & routing
├── models/
│   └── category.go      # Struct definitions
├── data/
│   └── category.go      # Seed data
└── handlers/
    └── category.go      # Handler functions
```

## Endpoints

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/` | Welcome message |
| GET | `/health` | Health check |
| GET | `/categories` | Ambil semua kategori |
| POST | `/categories` | Tambah kategori baru |
| GET | `/categories/{id}` | Ambil detail kategori |
| PUT | `/categories/{id}` | Update kategori |
| DELETE | `/categories/{id}` | Hapus kategori |

## Contoh Request

### Get All Categories
```bash
curl https://go-simple-crud.zeabur.app/categories
```

### Get Category by ID
```bash
curl https://go-simple-crud.zeabur.app/categories/1
```

### Create Category
```bash
curl -X POST https://go-simple-crud.zeabur.app/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Gaming", "description": "Peralatan gaming"}'
```

### Update Category
```bash
curl -X PUT https://go-simple-crud.zeabur.app/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Elektronik Update", "description": "Deskripsi baru"}'
```

### Delete Category
```bash
curl -X DELETE https://go-simple-crud.zeabur.app/categories/1
```

## Run Locally

```bash
# Clone repository
git clone https://github.com/username/go-simple-crud.git
cd go-simple-crud

# Run
go run main.go

# Server berjalan di http://localhost:8081
```

## Tech Stack

- **Go** - Programming language
- **net/http** - Standard library HTTP server
- **Zeabur** - Deployment platform
