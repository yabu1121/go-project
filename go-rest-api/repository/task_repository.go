package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTask(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, TaskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, TaskId uint) error
	DeleteTask(task *model.Task, userId uint, TaskId uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db}
}


// GetAllTaskは全てのタスクを取得する
func (tr *TaskRepository)GetAllTask(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?",userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}


// GetTaskByIdは指定されたIDのタスクを取得する
func (tr *TaskRepository)GetTaskById(task *model.Task, userId uint, TaskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ? AND task_id = ?", userId, TaskId).First(task).Error; err != nil {
		return err
	}
	return nil
}


// CreateTaskはタスクを作成する
func (tr *TaskRepository)CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil	
}


// UpdateTaskはタスクを更新する
func (tr *TaskRepository)UpdateTask(task *model.Task, userId uint, TaskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("user_id = ? AND task_id = ?", userId, TaskId).Updates(map[string]interface{}{"title": task.Title})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
func (tr *TaskRepository)DeleteTask(task *model.Task, userId uint, TaskId uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", TaskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}