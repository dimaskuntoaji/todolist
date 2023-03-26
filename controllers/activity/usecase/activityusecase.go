package usecase

import (
	"fmt"

	"todolist/controllers/activity"
	"todolist/model"
)

type activityUsecase struct {
	activityRepository activity.ActivityRepository
}

func NewActivityUsecase(activityRepository activity.ActivityRepository) activity.ActivityUsecase {
	return &activityUsecase{activityRepository}
}

func (u *activityUsecase) GetAll() ([]model.Activity, error) {
	return u.activityRepository.FindAll()
}

func (u *activityUsecase) GetByID(activityID int) (model.Activity, error) {
	return u.activityRepository.FindByID(activityID)
}

func (u *activityUsecase) Create(activity model.Activity) (model.Activity, error) {
	if activity.Title == "" {
		return model.Activity{}, fmt.Errorf("null title")
	}

	return u.activityRepository.Create(activity)
}

func (u *activityUsecase) Update(activity model.Activity, activityID int) (model.Activity, error) {
	if (activity == model.Activity{}) {
		return model.Activity{}, fmt.Errorf("null struct")
	}

	foundActivity, err := u.activityRepository.FindByID(activityID)
	if err != nil {
		return model.Activity{}, err
	}

	if activity.Title != "" {
		foundActivity.Title = activity.Title
	}

	if activity.Email != "" {
		foundActivity.Email = activity.Email
	}

	return u.activityRepository.Update(foundActivity, activityID)
}

func (u *activityUsecase) Delete(activityID int) error {
	if _, err := u.activityRepository.FindByID(activityID); err != nil {
		return err
	}

	return u.activityRepository.Delete(activityID)
}
