package activity

import "todolist/model"

type ActivityRepository interface {
	FindAll() ([]model.Activity, error)
	FindByID(activityID int) (model.Activity, error)
	Create(activity model.Activity) (model.Activity, error)
	Update(activity model.Activity, activityID int) (model.Activity, error)
	Delete(activityID int) error
}

type ActivityUsecase interface {
	GetAll() ([]model.Activity, error)
	GetByID(activityID int) (model.Activity, error)
	Create(activity model.Activity) (model.Activity, error)
	Update(activity model.Activity, activityID int) (model.Activity, error)
	Delete(activityID int) error
}