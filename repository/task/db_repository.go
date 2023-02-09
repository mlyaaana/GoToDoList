package task

import (
	"gorm.io/gorm"
	"todolist/database/models"
	"todolist/domain"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) Create(task *domain.Task) error {
	return r.db.Create(&models.Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
		CreatedAt:   task.CreatedAt,
		Deadline:    task.Deadline,
	}).Error
}

func (r *DBRepository) Get(id string) (*domain.Task, error) {
	task := &models.Task{}
	err := r.db.Where(&models.Task{Id: id}).First(task).Error
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
		CreatedAt:   task.CreatedAt,
		Deadline:    task.Deadline,
	}, nil
}

func (r *DBRepository) List() ([]*domain.Task, error) {
	var tasksModels []models.Task
	tasksDomain := make([]*domain.Task, 0)
	err := r.db.Find(tasksModels).Error
	if err != nil {
		return nil, err
	}

	for _, t := range tasksModels {
		tasksDomain = append(tasksDomain, &domain.Task{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Done:        t.Done,
			CreatedAt:   t.CreatedAt,
			Deadline:    t.Deadline,
		})
	}

	return tasksDomain, nil
}

func (r *DBRepository) Delete(id string) {
	r.db.Delete(&models.Task{Id: id})
}

func (r *DBRepository) Done(id string) error {
	task := &models.Task{}
	err := r.db.Where(&models.Task{Id: id}).First(task).Error
	if err != nil {
		return err
	}
	task.Done = !task.Done
	return r.db.Save(task).Error
}

func (r *DBRepository) ListCompletedTasks() ([]*domain.Task, error) {
	var tasksModels []models.Task
	var tasksDomain []*domain.Task
	err := r.db.Where(&models.Task{Done: true}).Find(&tasksModels).Error
	if err != nil {
		return nil, err
	}

	for _, t := range tasksModels {
		tasksDomain = append(tasksDomain, &domain.Task{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
			Done:        t.Done,
			CreatedAt:   t.CreatedAt,
			Deadline:    t.Deadline,
		})
	}

	return tasksDomain, nil
}
