package usecases

import (

	"github.com/sdimitrenco/grammurrr/internal/domains"
	"github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"
	"github.com/sdimitrenco/grammurrr/internal/repositories"
)


type WordService struct {
	repo repositories.WordRepository
	logger *logging.Logger
}

func NewWordService(repo repositories.WordRepository, logger *logging.Logger) *WordService {
	return &WordService{
		repo: repo,
		logger: logger,
	}
} 

func (ws *WordService) AddWord(word, land string)  (domains.Word, error) {
	w, e := ws.FoundWordByName(word, land)

	if(e == nil) {
		return w, e
	}

	w = domains.Word{
		Word: word,
		Lang: land,
	}

	return ws.repo.Create(w)
}

func (ws *WordService) FoundWordByName(word, lang string) (domains.Word, error) {
	return ws.repo.FindByWord(word, lang)
}