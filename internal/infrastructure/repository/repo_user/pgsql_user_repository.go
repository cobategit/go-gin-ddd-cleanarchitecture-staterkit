package repo_user

import (
	domain "github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/domain/d_user"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int64) (*domain.UserE, error) {
	var user domain.UserE
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*domain.UserE, error) {
	var user domain.UserE
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Create(user *domain.UserE) error {
	query := `
INSERT INTO users (email, password, name)
VALUES ($1, $2, $3)
RETURNING id, email, created_at, updated_at;
`
	return r.db.QueryRowx(
		query,
		user.Email,
		user.Password,
		user.Name,
	).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
}
