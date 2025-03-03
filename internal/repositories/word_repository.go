package repositories

import "github.com/sdimitrenco/grammurrr/internal/domains"

type WordRepository interface {
	FindAll() ([]domains.Word, error)
	FindById(word domains.Word) (domains.Word, error)
	FindByWord(word domains.Word) (domains.Word, error)
	FindTranslate(word domains.Word) (domains.Word, error)
	Create(word domains.Word) (domains.Word, error)
	Update(word domains.Word) (domains.Word, error)
	Delete(id int) error
}

