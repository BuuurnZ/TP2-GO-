package main

import (
	"fmt"
	"strings"
)

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Email envoyé à %s: %s\n", e.Email, message)
	return nil
}

func (e EmailNotifier) GetType() string {
	return "Email"
}

type SMSNotifier struct {
	Phone string
}

func (s SMSNotifier) Send(message string) error {
	if !strings.HasPrefix(s.Phone, "06") {
		return fmt.Errorf("numéro de téléphone invalide: %s (doit commencer par 06)", s.Phone)
	}
	fmt.Printf("SMS envoyé au %s: %s\n", s.Phone, message)
	return nil
}

func (s SMSNotifier) GetType() string {
	return "SMS"
}

type PushNotifier struct {
	DeviceID string
}

func (p PushNotifier) Send(message string) error {
	fmt.Printf("Push envoyé sur %s: %s\n", p.DeviceID, message)
	return nil
}

func (p PushNotifier) GetType() string {
	return "Push"
}
