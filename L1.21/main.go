/*
Паттерн адаптер можно применить,когда нужно использовать уже существующий класс,но его интерфейс не подходит коду.
пример, есть сервис, который нельзя менять, а мой код ожидает другой интерфейс. Адаптер помогает
«скентовать» старый и новый код, чтобы всё работало вместе.



*/

package main

import "fmt"

type Notifier interface {
	SendNotification(message string)
}

type EmailSender struct{}

func (e *EmailSender) SendEmail(emailText string) {
	fmt.Println("Отправка email:", emailText)
}

type EmailAdapter struct {
	EmailService *EmailSender
}

func (a *EmailAdapter) SendNotification(message string) {
	a.EmailService.SendEmail(message)
}

func Client(n Notifier, msg string) {
	n.SendNotification(msg)
}

func main() {
	oldEmail := &EmailSender{}
	adapter := &EmailAdapter{EmailService: oldEmail}
	Client(adapter, "Встреча в 15:00")
}
