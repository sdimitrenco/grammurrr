package controllers


func  (b *BotControllerImpl) Start () string {
	return "Привет! Я бот для заучивания слов.\n" +
				"/addgroup [название] — создать группу\n" +
				"/addword [группа] [иностранное слово] - [перевод] — добавить слово\n" +
				"/train [группа] — начать тренировку"
	
}

