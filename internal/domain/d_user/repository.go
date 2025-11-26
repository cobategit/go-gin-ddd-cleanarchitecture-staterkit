package d_user

// Repository mendefinisikan kontrak untuk menyimpan user.
// Implementasinya ada di layer infrastructure.
type Repository interface {
	GetByID(id int64) (*UserE, error)
	GetByEmail(email string) (*UserE, error)
	Create(User *UserE) error
}
