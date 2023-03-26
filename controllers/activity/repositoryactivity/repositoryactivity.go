package repositoryactivity

import (
	"todolist/controllers/activity"
	"todolist/model"
	"gorm.io/gorm"
)

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) activity.ActivityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) Create(activity model.Activity) (model.Activity, error) {
	tx := r.db.Create(&activity)
	return activity, tx.Error
}

func (r *activityRepository) Update(activity model.Activity, activityID int) (model.Activity, error) {
	tx := r.db.Where("activity_id = ?", activityID).Updates(&activity)
	return activity, tx.Error
}

func (r *activityRepository) Delete(activityID int) error {
	return r.db.Where("activity_id = ?", activityID).Delete(&model.Activity{}).Error
}


func (r *activityRepository) FindAll() ([]model.Activity, error) {
	activities := []model.Activity{}
	tx := r.db.Find(&activities)
	return activities, tx.Error
}

func (r *activityRepository) FindByID(activityID int) (model.Activity, error) {
	activity := model.Activity{}
	tx := r.db.Where("activity_id = ?", activityID).First(&activity)
	return activity, tx.Error
}

