package controllers

import "github.com/wecanooo/gosari/core/context"

type IUserController interface {
	Index(*context.AppContext) error
	Show(*context.AppContext) error
	Create(*context.AppContext) error
}

type UserController struct{}

func NewUserController() IUserController {
	return &UserController{}
}

func (u *UserController) Index(c *context.AppContext) error {
	return c.SuccessJSON("Index")
}

func (u *UserController) Show(c *context.AppContext) error {
	return c.SuccessJSON("Show")
}

func (u *UserController) Create(c *context.AppContext) error {
	return c.SuccessJSON("create")
}
