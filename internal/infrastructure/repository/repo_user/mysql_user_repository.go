package repo_user

import (
	domain "github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/domain/d_user"
	"github.com/jmoiron/sqlx"
)

type MySQLUserRepository struct {
	db *sqlx.DB
}

func NewMySQLUserRepository(db *sqlx.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) GetByID(id int64) (*domain.UserE, error) {
	var user domain.UserE
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) GetByEmail(email string) (*domain.UserE, error) {
	var user domain.UserE
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) Create(user *domain.UserE) error {
	query := `
INSERT INTO users (email, password, name, created_at, updated_at)
VALUES (?, ?, ?, NOW(), NOW())
`
	res, err := r.db.Exec(query, user.Email, user.Password, user.Name)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err == nil {
		user.ID = id
	}
	return nil
}
