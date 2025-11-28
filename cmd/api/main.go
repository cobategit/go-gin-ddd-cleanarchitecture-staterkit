package main

import (
	"log"

	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/config"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/domain/d_user"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/infrastructure/db"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/infrastructure/repository/repo_user"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/infrastructure/security"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/interfaces/http"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/interfaces/http/handler"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/usecase/uc_user"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// @title           Practice Go v1
// @version         1.0
// @description     Boilerplate Gin + DDD + JWT + Postgres/MySQL
// @host            localhost:5006
// @BasePath        /
// @schemes         http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()

	var sqlxDB *sqlx.DB
	switch cfg.DBDriver {
	case "postgres":
		sqlxDB = db.NewPostgresDB(cfg)
	case "mysql":
		sqlxDB = db.NewMySQLDB(cfg)
	default:
		log.Fatalf("unsupported DB_DRIVER: %s", cfg.DBDriver)
	}

	jwtService := security.NewJWTService(cfg)
	hasher := security.NewBcryptHasher()

	var repo usecaseUserRepository // adapter ke domain.Repository

	// kita pakai alias type untuk menghindari import siklik
	repo = chooseUserRepo(cfg.DBDriver, sqlxDB)

	userUC := uc_user.NewUserUseCase(repo, jwtService, hasher)
	userHandler := handler.NewUserHandler(userUC)

	router := http.NewRouter(userHandler, jwtService)

	logger.InitLogger("Starting server on :%s", "info", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		logger.InitLogger("Server stopped: %v", "error", err)
	}
}

type usecaseUserRepository interface {
	userRepoInterface
}

// kita definisikan interface minimal yang sama dengan domain.Repository
type userRepoInterface interface {
	GetByID(id int64) (*userRepoDomainUser, error)
	GetByEmail(email string) (*userRepoDomainUser, error)
	Create(user *userRepoDomainUser) error
}

// type alias biar singkat
type userRepoDomainUser = d_user.UserE

// import alias
// (secara praktis, di file nyata cukup import domain & implement Repo langsung,
// di sini hanya untuk contoh; boleh dipermudah).
func chooseUserRepo(driver string, dbConn *sqlx.DB) usecaseUserRepository {
	switch driver {
	case "postgres":
		return repo_user.NewPostgresUserRepository(dbConn)
	case "mysql":
		return repo_user.NewMySQLUserRepository(dbConn)
	default:
		panic("unsupported driver")
	}
}
