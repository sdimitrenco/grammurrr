package repositories

import "github.com/sdimitrenco/grammurrr/internal/domains"

type GroupRepository interface {
	FindAll() ([]domains.WordGroup, error)
	FindById(group domains.WordGroup) (domains.WordGroup, error)
	FindByGroupName(group domains.WordGroup) (domains.WordGroup, error)
	Create(group domains.WordGroup) (domains.WordGroup, error)
	Update(group domains.WordGroup) (domains.WordGroup, error)
	Delete(group domains.WordGroup) error
}

