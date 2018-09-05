package task

import "github.com/FernandoCagale/serverless-task/pkg/entity"

//UseCase use case interface
type UseCase interface {
	Create(task *entity.Task) (err error)
	Update(id int, task *entity.Task) (err error)
	Delete(id int) (err error)
	FindByID(id int) (task *entity.Task, err error)
	FindAll() (tasks []*entity.Task, err error)
}
