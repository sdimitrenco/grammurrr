package usecases

import (
	"github.com/sdimitrenco/grammurrr/internal/domains"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/internal/repositories"
)

type WordUseCase struct {
	repo   repositories.WordRepository
	logger *logging.Logger
}

func NewWordUseCase(repo repositories.WordRepository, logger *logging.Logger) *WordUseCase {
	return &WordUseCase{
		repo:   repo,
		logger: logger,
	}
}

func (ws *WordUseCase) AddWord(word domains.Word) (domains.Word, error) {
	w, e := ws.FoundWordByName(word)

	if e == nil {
		return w, e
	}

	return ws.repo.Create(w)
}

func (ws *WordUseCase) FoundWordByName(word domains.Word) (domains.Word, error) {
	return ws.repo.FindByWord(word)
}
