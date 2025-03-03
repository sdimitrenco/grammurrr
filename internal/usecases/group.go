package usecases

import (
	"github.com/sdimitrenco/grammurrr/internal/domains"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/internal/repositories"
)

type GroupUseCase struct {
	repo   repositories.GroupRepository
	logger *logging.Logger
}

func NewGroupUseCase(repo repositories.GroupRepository, logger *logging.Logger) *GroupUseCase {
	return &GroupUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (ws *GroupUseCase) AddGroup(group domains.WordGroup) (domains.WordGroup, error) {
	g, e := ws.FoundGroupByName(group)

	if e == nil {
		return g, e
	}


	return ws.repo.Create(g)
}

func (ws *GroupUseCase) FoundGroupByName(group domains.WordGroup) (domains.WordGroup, error) {
	return ws.repo.FindByGroupName(group)
}
