package main

import (
	"tmbot/bot"
	"tmbot/client"
	"tmbot/database"
	"tmbot/services"
)

func main() {

	/*

		Бот
			Получать сообщения (Клиент)
			Обрабатывать сообщения (Процессор)
			Отправлять сообщения (Клиент)

		Клиент
			Отправлять сообщения
			Получать сообщения

		Процессор
			Распознать команду
			Обработать команду (Сервис)

		Сервис
			Отправить ошибку
			Отправить ответ на команду
			Вести диалог
			Взаимодействие с базой данных (Хранилище)

		Хранилище
			Подключиться к БД
			Получить/Сохранить/Отредактировать/Удалить элемент
	*/

	postgresDB := database.NewDatabase()
	service := services.NewServices(postgresDB)
	telegram := client.New()
	b := bot.NewBot(telegram, service)
	_ = b
}
