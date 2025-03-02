package repositories

import "github.com/sdimitrenco/grammurrr/internal/domains"

type WordRepository interface {
	FindAll() ([]domains.Word, error)
	FindById(id int) (domains.Word, error)
	FindByWord(word, lang string) (domains.Word, error)
	FindTranslate(word string) (domains.Word, error)
	Create(word domains.Word) (domains.Word, error)
	Update(word domains.Word) (domains.Word, error)
	Delete(id int) error
}

