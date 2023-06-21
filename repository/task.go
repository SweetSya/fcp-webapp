package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	if fd := t.db.Model(
		model.Task{}).Where(
		"id = ?", task.ID).Updates(
		task); fd.Error != nil {
		return fd.Error
	}
	return nil // TODO: replace this
}

func (t *taskRepository) Delete(id int) error {
	if fd := t.db.Delete(model.Task{}, id); fd.Error != nil {
		return fd.Error
	}
	return nil // TODO: replace this
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	var tasks []model.Task

	if fd := t.db.Find(&tasks); fd.Error != nil {
		return []model.Task{}, fd.Error
	}

	return tasks, nil // TODO: replace this
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	var tasksCat []model.TaskCategory
	if fd := t.db.Table("tasks").Select("tasks.id as ID, tasks.title as Title, categories.name as Category").Joins("left join categories on categories.id = tasks.category_id").Where("category_id = ?", id).Limit(1).Scan(&tasksCat); fd.Error != nil {
		return []model.TaskCategory{}, fd.Error
	}
	return tasksCat, nil // TODO: replace this
}
