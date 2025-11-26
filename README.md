# 🧱 Go Gin Domain Driven Design (DDD) Boilerplate

Boilerplate backend service dengan:

- ⚙️ **Gin** HTTP framework
- 🧠 **Domain-Driven Design + Clean Architecture**
- 🔐 **JWT Authentication**
- 🗄️ **PostgreSQL & MySQL (pluggable via ENV)**
- 🐳 **Docker & docker-compose**
- 📜 **Swagger / OpenAPI documentation**
- ✅ **Unit tests & GitHub Actions CI**

Cocok sebagai starter kit untuk microservice atau monolith kecil yang ingin rapi dari awal.

---

## 📁 Project Structure

```bash
.
├── cmd/
│   └── api/
│       └── main.go           # entrypoint REST API
├── internal/
│   ├── config/               # load ENV & app config
│   ├── domain/               # DDD domain layer (entities & repository contracts)
│   │   └── user/
│   ├── usecase/              # application service / business logic
│   ├── infrastructure/       # db, repository impl, security, etc
│   │   ├── db/
│   │   ├── repository/
│   │   └── security/
│   └── interfaces/
│       └── http/             # Gin handlers, router, middleware
├── pkg/
│   ├── logger/               # simple logger wrapper
│   └── response/             # uniform API response
├── docs/                     # generated Swagger files (swag)
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
└── .github/
    └── workflows/
        └── ci.yml            # GitHub Actions: build & test
```

- Domain (internal/domain)
  Berisi entity & interface repository (pure business). Tidak tahu apa-apa tentang DB, HTTP, dll.

- Usecase (internal/usecase)
  Menyusun business flow (Register, Login, GetProfile, dll). Hanya bergantung pada interface domain & service (JWT, hasher).

- Infrastructure (internal/infrastructure)
  Implementasi nyata: koneksi DB (Postgres/MySQL), repository SQL, JWT, bcrypt, dll.

- Interfaces (internal/interfaces/http)
  Gin handler, routing, middleware, Swagger binding.

## 🚀 Getting Started

1. Prerequisites
   Go 1.22+
   Docker & docker-compose
   (Opsional, untuk Swagger): swag CLI
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. Clone & Module Name
   ```bash
   git clone https://github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit.git
   cd your-app

   # ganti module name di go.mod kalau perlu
   # module github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit
   ```