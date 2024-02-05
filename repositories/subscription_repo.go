package repositories

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"datingapp-api/models"
)

// IUserRepository interface for table users
type ISubscriptionRepository interface {
	BeginTrans()
	RollbackTrans()
	CommitTrans()
	CreateSubscription(models.Subscription) (int64, error)
	GetSubscriptionById(int64) (models.Subscription, error)
	GetSubscriptionByUserId(int64) (models.Subscription, error)
}

// SubscriptionRepo struct
type SubscriptionRepo struct {
	db orm.Ormer
}

// NewSubscriptionRepo initial user repository
func NewSubscriptionRepo(o orm.Ormer) ISubscriptionRepository {
	return &SubscriptionRepo{
		db: o,
	}
}

// BeginTrans method for beginning database transactions
func (repo *SubscriptionRepo) BeginTrans() {
	repo.db.Begin()
}

// RollbackTrans method for rollbacking database transactions
func (repo *SubscriptionRepo) RollbackTrans() {
	repo.db.Rollback()
}

// CommitTrans method for committing database transactions
func (repo *SubscriptionRepo) CommitTrans() {
	repo.db.Commit()
}

// CreateSubscription func to create user subscription
func (repo *SubscriptionRepo) CreateSubscription(subscription models.Subscription) (int64, error) {
	id, err := repo.db.Insert(&subscription)
	return id, err
}

// GetSubscriptionById func to get subscription by id
func (repo *SubscriptionRepo) GetSubscriptionById(id int64) (subscription models.Subscription, err error) {
	err = repo.db.QueryTable("user_subscriptions").
		Filter("id", id).
		One(&subscription)
	return subscription, err
}

// GetSubscriptionByUserId func to get subscription by user id
func (repo *SubscriptionRepo) GetSubscriptionByUserId(userId int64) (subscription models.Subscription, err error) {
	logs.Info("user id %d", userId)
	err = repo.db.QueryTable("user_subscriptions").
		Filter("user_id", userId).
		One(&subscription)
	return subscription, err
}