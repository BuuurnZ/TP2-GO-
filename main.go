package main

import (
	"fmt"
	"time"
)

type Notifier interface {
	Send(message string) error
	GetType() string
}

type Storer interface {
	Save(notification Notification)
	GetAll() []Notification
}

type Notification struct {
	Message   string
	Type      string
	Timestamp time.Time
}

func main() {
	fmt.Println("=== Système de Notifications ===")
	
	storer := NewMemoryStorer()
	
	notifiers := []Notifier{
		EmailNotifier{Email: "user@example.com"},
		SMSNotifier{Phone: "0612345678"},
		SMSNotifier{Phone: "0123456789"},
		PushNotifier{DeviceID: "device123"},
		EmailNotifier{Email: "admin@example.com"},
	}
	
	message := "Test de notification"
	
	for _, notifier := range notifiers {
		err := notifier.Send(message)
		if err != nil {
			fmt.Printf("Erreur lors de l'envoi %s: %v\n", notifier.GetType(), err)
		} else {
			notification := Notification{
				Message:   message,
				Type:      notifier.GetType(),
				Timestamp: time.Now(),
			}
			storer.Save(notification)
		}
	}
	
	fmt.Println("\n=== Historique des notifications ===")
	allNotifications := storer.GetAll()
	for i, notif := range allNotifications {
		fmt.Printf("%d. [%s] %s - %s\n", i+1, notif.Type, notif.Message, notif.Timestamp.Format("2006-01-02 15:04:05"))
	}
	
	fmt.Printf("\nTotal: %d notifications archivées\n", storer.GetCount())
}
