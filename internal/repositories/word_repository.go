package repositories

type WordRepository interface {
	FindAll() ([]Word, error)
	FindById(id int) (Word, error)
	FindByWord(word string) (Word, error)
	FindTranslate(word string) (Word, error)
	Create(word Word) (Word, error)
	Update(word Word) (Word, error)
	Delete(id int) error
}

type Word string