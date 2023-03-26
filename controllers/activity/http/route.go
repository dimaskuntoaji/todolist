package http

import "github.com/gofiber/fiber/v2"

func MapRoutes(r fiber.Router, h activityHandler) {
	r.Get("/activity-groups", h.GetAllActivities())
	r.Get("/activity-groups/:id", h.GetActivityByID())
	r.Post("/activity-groups", h.CreateActivity())
	r.Patch("/activity-groups/:id", h.UpdateActivity())
	r.Delete("/activity-groups/:id", h.DeleteActivity())
}
