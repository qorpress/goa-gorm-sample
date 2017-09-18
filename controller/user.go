package controller

import (
	"github.com/ReSTARTR/goa-gorm-sample/app"
	"github.com/ReSTARTR/goa-gorm-sample/db"
	"github.com/ReSTARTR/goa-gorm-sample/models"
	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here
	udb := models.NewUserDB(db.DB)
	u := models.UserFromUserPayload(ctx.Payload)
	if err := udb.Add(ctx, u); err != nil {
		return err
	}

	// UserController_Create: end_implement
	res := &app.ApplicationVndUser{}
	return ctx.OK(res)
}

// Index runs the index action.
func (c *UserController) Index(ctx *app.IndexUserContext) error {
	// UserController_Index: start_implement

	// Put your logic here
	udb := models.NewUserDB(db.DB)
	users := udb.ListApplicationVndUser(ctx)

	// UserController_Index: end_implement
	res := app.ApplicationVndUserCollection{}
	res = users
	return ctx.OK(res)
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here
	udb := models.NewUserDB(db.DB)
	var user *app.ApplicationVndUser
	var err error
	user, err = udb.OneApplicationVndUser(ctx, ctx.UserID)
	if err != nil {
		return err
	}

	// UserController_Show: end_implement
	res := &app.ApplicationVndUser{}
	res = user
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here
	udb := models.NewUserDB(db.DB)
	if err := udb.UpdateFromUserPayload(ctx, ctx.Payload, ctx.UserID); err != nil {
		return nil
	}

	// UserController_Update: end_implement
	res := &app.ApplicationVndUser{}
	return ctx.OK(res)
}
