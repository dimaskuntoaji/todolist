package http


import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"todolist/controllers/activity"
	"todolist/model"
	"gorm.io/gorm"
)

type activityHandler struct {
	activityUsecase activity.ActivityUsecase
}

func NewActivityHandler(activityUsecase activity.ActivityUsecase) *activityHandler {
	return &activityHandler{activityUsecase}
}

func (h *activityHandler) CreateActivity() fiber.Handler {
	type CreateActivityRequest struct {
		Title string `json:"title"`
		Email string `json:"email"`
	}
	return func(c *fiber.Ctx) error {
		request := &CreateActivityRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(model.Response{
					Status:  "Bad Request",
					Message: "Your request could not be completed as it contains errors. Please check and try again",
				})
		}

		activity, err := h.activityUsecase.Create(model.Activity{Title: request.Title, Email: request.Email})
		if err != nil {
			if err.Error() == "null title" {
				return c.Status(fiber.StatusBadRequest).
					JSON(model.Response{
						Status:  "Bad Request",
						Message: "title cannot be null",
					})
			}

			return c.Status(fiber.StatusBadGateway).
				JSON(model.Response{
					Status:  "Bad Gateway",
					Message: "Oops! Something went wrong while trying to reach the server. Please try again later.",
				})
		}

		return c.Status(fiber.StatusCreated).
			JSON(model.Response{
				Status:  "Success",
				Message: "Success",
				Data:    activity,
			})
	}
}

func (h *activityHandler) UpdateActivity() fiber.Handler {
	type UpdateActivityRequest struct {
		Title string `json:"title"`
		Email string `json:"email"`
	}
	return func(c *fiber.Ctx) error {
		activityID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(model.Response{
					Status:  "Bad Request",
					Message: "Invalid type of activity ID",
				})
		}

		request := &UpdateActivityRequest{}
		err = c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(model.Response{
					Status:  "Bad Request",
					Message: "Your request could not be completed as it contains errors. Please check and try again",
				})
		}

		activity, err := h.activityUsecase.Update(model.Activity{Title: request.Title, Email: request.Email}, activityID)
		if err != nil {
			if err.Error() == "null struct" {
				return c.Status(fiber.StatusBadRequest).
					JSON(model.Response{
						Status:  "Bad Request",
						Message: "title cannot be null",
					})
			}

			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).
					JSON(model.Response{
						Status:  "Not Found",
						Message: fmt.Sprintf("Activity with ID %d Not Found", activityID),
					})
			}

			return c.Status(fiber.StatusBadGateway).
				JSON(model.Response{
					Status:  "Bad Gateway",
					Message: "Oops! Something went wrong while trying to reach the server. Please try again later.",
				})
		}

		return c.Status(fiber.StatusOK).
			JSON(model.Response{
				Status:  "Success",
				Message: "Success",
				Data:    activity,
			})
	}
}

func (h *activityHandler) DeleteActivity() fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(model.Response{
					Status:  "Bad Request",
					Message: "Invalid type of activity ID",
				})
		}

		err = h.activityUsecase.Delete(activityID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).
					JSON(model.Response{
						Status:  "Not Found",
						Message: fmt.Sprintf("Activity with ID %d Not Found", activityID),
					})
			}

			return c.Status(fiber.StatusBadGateway).
				JSON(model.Response{
					Status:  "Bad Gateway",
					Message: "Oops! Something went wrong while trying to reach the server. Please try again later.",
				})
		}

		return c.Status(fiber.StatusOK).
			JSON(model.Response{
				Status:  "Success",
				Message: "Success",
				Data:    struct{}{},
			})
	}
}

func (h *activityHandler) GetAllActivities() fiber.Handler {
	return func(c *fiber.Ctx) error {
		activities, err := h.activityUsecase.GetAll()
		if err != nil {
			return c.Status(fiber.StatusBadGateway).
				JSON(model.Response{
					Status:  "Bad Gateway",
					Message: "Oops! Something went wrong while trying to reach the server. Please try again later.",
				})
		}

		return c.Status(fiber.StatusOK).
			JSON(model.Response{
				Status:  "Success",
				Message: "Success",
				Data:    activities,
			})
	}
}

func (h *activityHandler) GetActivityByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(model.Response{
					Status:  "Bad Request",
					Message: "Invalid type of activity ID",
				})
		}

		activity, err := h.activityUsecase.GetByID(activityID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).
					JSON(model.Response{
						Status:  "Not Found",
						Message: fmt.Sprintf("Activity with ID %d Not Found", activityID),
					})
			}

			return c.Status(fiber.StatusBadGateway).
				JSON(model.Response{
					Status:  "Bad Gateway",
					Message: "Oops! Something went wrong while trying to reach the server. Please try again later.",
				})
		}

		return c.Status(fiber.StatusOK).
			JSON(model.Response{
				Status:  "Success",
				Message: "Success",
				Data:    activity,
			})
	}
}

