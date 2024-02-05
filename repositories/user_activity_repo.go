package repositories

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"datingapp-api/models"
	"fmt"
)

// IUserActivityRepository interface for table useractivitys
type IUserActivityRepository interface {
	BeginTrans()
	RollbackTrans()
	CommitTrans()
	ViewUser(int64) (models.Profile, error)
	CountViewUserInADay(int64) (int, error)
	CreateUserActivity(models.UserActivity) (int64, error)
}

// UserActivityRepo struct
type UserActivityRepo struct {
	db orm.Ormer
}

// NewUserActivityRepo initial useractivity repository
func NewUserActivityRepo(o orm.Ormer) IUserActivityRepository {
	return &UserActivityRepo{
		db: o,
	}
}

// BeginTrans method for beginning database transactions
func (repo *UserActivityRepo) BeginTrans() {
	repo.db.Begin()
}

// RollbackTrans method for rollbacking database transactions
func (repo *UserActivityRepo) RollbackTrans() {
	repo.db.Rollback()
}

// CommitTrans method for committing database transactions
func (repo *UserActivityRepo) CommitTrans() {
	repo.db.Commit()
}

// ViewUser func to view all use
func (repo *UserActivityRepo) ViewUser(userId int64) (user models.Profile, err error) {
	baseQuery := `
		select * from user_profiles up
		left join user_activities ua on up.user_id = ua.other_user_id
		where up.user_id <> %d and up.user_id not in (
			select other_user_id from user_activities ua
			where ua.user_id = %d and date(created_at) = date(now()) 
		)
		limit 1
	`
	query := fmt.Sprintf(baseQuery, userId, userId)
	logs.Info(query)
	err = repo.db.Raw(query).QueryRow(&user)
	return user, err
}

// ViewUser func to view all use
func (repo *UserActivityRepo) CountViewUserInADay(userId int64) (countView int, err error) {
	baseQuery := `
		select count(id) from user_activities ua
		where ua.user_id = %d and date(created_at) = date(now())
	`
	query := fmt.Sprintf(baseQuery, userId)
	logs.Info(query)
	err = repo.db.Raw(query).QueryRow(&countView)
	return countView, err
}

func (repo *UserActivityRepo) CreateUserActivity(useractivity models.UserActivity) (int64, error) {
	id, err := repo.db.Insert(&useractivity)
	return id, err
}
