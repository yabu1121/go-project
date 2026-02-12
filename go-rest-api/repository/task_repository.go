package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db}
}

// GetAllTasksは全てのタスクを取得する
func (tr *TaskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// GetTaskByIdは指定されたIDのタスクを取得する
func (tr *TaskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ? AND tasks.id = ?", userId, taskId).First(task).Error; err != nil {
		return err
	}
	return nil
}

// CreateTaskはタスクを作成する
func (tr *TaskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTaskはタスクを更新する
func (tr *TaskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("user_id = ? AND id = ?", userId, taskId).Updates(map[string]interface{}{"title": task.Title})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *TaskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 0 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
