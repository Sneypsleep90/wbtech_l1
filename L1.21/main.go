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
