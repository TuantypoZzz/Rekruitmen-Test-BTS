package controller

import (
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ChecklistController struct {
	service.ChecklistService
}

func NewChecklistController(ChecklistService *service.ChecklistService) *ChecklistController {
	return &ChecklistController{ChecklistService: *ChecklistService}
}

func (controller ChecklistController) Route(api fiber.Router) {
	checklists := api.Group("/checklist")
	checklists.Post("/", controller.Create)
	// checklists.Put("/:id", controller.Update)
	checklists.Delete("/:id", controller.Delete)
	checklists.Get("/:id", controller.FindById)
	checklists.Get("/", controller.FindAll)
}

func (controller ChecklistController) Create(c *fiber.Ctx) error {
	var request model.CreateChecklist
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	user := c.Locals("userData").(entity.User)

	response := controller.ChecklistService.Create(c.Context(), request, user.Username)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// func (controller ChecklistController) Update(c *fiber.Ctx) error {
// 	var request model.ChecklistCreateOrUpdateModel
// 	id, err := strconv.Atoi(c.Params("id"))
// 	err = c.BodyParser(&request)
// 	exception.PanicLogging(err)

// 	response := controller.ChecklistService.Update(c.Context(), request, id)
// 	return c.JSON(model.GeneralResponse{
// 		Code:    200,
// 		Message: "Success",
// 		Data:    response,
// 	})
// }

func (controller ChecklistController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	exception.PanicLogging(err)
	user := c.Locals("userData").(entity.User)

	controller.ChecklistService.Delete(c.Context(), id, user.Username)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

func (controller ChecklistController) FindById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	exception.PanicLogging(err)

	result := controller.ChecklistService.FindById(c.Context(), id)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller ChecklistController) FindAll(c *fiber.Ctx) error {
	user := c.Locals("userData").(entity.User)
	result := controller.ChecklistService.FindAll(c.Context(), user.Username)
	return c.JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
