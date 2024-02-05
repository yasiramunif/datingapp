package repositories

import (
	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"datingapp-api/models"
)

// IUserRepository interface for table users
type IFeatureRepository interface {
	BeginTrans()
	RollbackTrans()
	CommitTrans()
	GetAllFeatures() ([]models.Feature, error)
	GetFeatureById(int64) (models.Feature, error)
}

// FeatureRepo struct
type FeatureRepo struct {
	db orm.Ormer
}

// NewFeatureRepo initial user repository
func NewFeatureRepo(o orm.Ormer) IFeatureRepository {
	return &FeatureRepo{
		db: o,
	}
}

// BeginTrans method for beginning database transactions
func (repo *FeatureRepo) BeginTrans() {
	repo.db.Begin()
}

// RollbackTrans method for rollbacking database transactions
func (repo *FeatureRepo) RollbackTrans() {
	repo.db.Rollback()
}

// CommitTrans method for committing database transactions
func (repo *FeatureRepo) CommitTrans() {
	repo.db.Commit()
}

// GetAllFeatures func to get all features
func (repo *FeatureRepo) GetAllFeatures() (features []models.Feature, err error) {
	_, err = repo.db.QueryTable("premium_features").
		All(&features)
	return features, err
}

// GetFeatuerById func to get feature by id
func (repo *FeatureRepo) GetFeatureById(id int64) (feature models.Feature, err error) {
	err = repo.db.QueryTable("premium_features").
		Filter("id", id).
		One(&feature)
	return feature, err
}
