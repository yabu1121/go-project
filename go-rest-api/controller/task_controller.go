package controller

import (
	"go-rest-api/usecase"

	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks (c echo.Context) error
	GetTaskById (c echo.Context) error
	CreateTask (c echo.Context) error
	UpdateTask (c echo.Context) error
	DeleteTask (c echo.Context) error
}

type TaskController struct {
	tu ITaskUsecase
}

func NewTaskController (tu usecase.ITaskUsecase) ITaskController {
	return &TaskController{tu}
}



func (tc *TaskController)GetAllTasks (c echo.Context) error
func (tc *TaskController)GetTaskById (c echo.Context) error
func (tc *TaskController)CreateTask (c echo.Context) error
func (tc *TaskController)UpdateTask (c echo.Context) error
func (tc *TaskController)DeleteTask (c echo.Context) error