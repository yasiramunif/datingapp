package repositories

import (
	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"datingapp-api/models"
)

// IUserRepository interface for table users
type IUserRepository interface {
	BeginTrans()
	RollbackTrans()
	CommitTrans()
	GetAllUsers() ([]models.Profile, error)
	GetUserById(int64) (models.User, error)
	GetUserByUsername(string) (models.User, error)
	CreateUser(models.User) (int64, error)
	CreateProfile(models.Profile) (int64, error)
}

// UserRepo struct
type UserRepo struct {
	db orm.Ormer
}

// NewUserRepo initial user repository
func NewUserRepo(o orm.Ormer) IUserRepository {
	return &UserRepo{
		db: o,
	}
}

// BeginTrans method for beginning database transactions
func (repo *UserRepo) BeginTrans() {
	repo.db.Begin()
}

// RollbackTrans method for rollbacking database transactions
func (repo *UserRepo) RollbackTrans() {
	repo.db.Rollback()
}

// CommitTrans method for committing database transactions
func (repo *UserRepo) CommitTrans() {
	repo.db.Commit()
}

// GetAllUsers func to get all users
func (repo *UserRepo) GetAllUsers() (users []models.Profile, err error) {
	_, err = repo.db.QueryTable("user_profiles").
		// RelatedSel("User").
		All(&users)
	return users, err
}

// GetUserById func to get user by id
func (repo *UserRepo) GetUserById(id int64) (user models.User, err error) {
	err = repo.db.QueryTable("users").
		Filter("id", id).
		One(&user)
	return user, err
}

// GetUserByUsername func to get user by username
func (repo *UserRepo) GetUserByUsername(username string) (user models.User, err error) {
	err = repo.db.QueryTable("users").
		Filter("username", username).
		One(&user)

	repo.db.LoadRelated(&user, "Profile")
	return user, err
}

func (repo *UserRepo) CreateUser(user models.User) (int64, error) {
	id, err := repo.db.Insert(&user)
	return id, err
}

func (repo *UserRepo) CreateProfile(profile models.Profile) (int64, error) {
	id, err := repo.db.Insert(&profile)
	return id, err
}