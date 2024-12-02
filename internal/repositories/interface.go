package repositories

import "github.com/FelipePn10/microservice-go/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
