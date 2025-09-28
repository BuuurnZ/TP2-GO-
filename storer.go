package main

type MemoryStorer struct {
	notifications []Notification
}

func NewMemoryStorer() *MemoryStorer {
	return &MemoryStorer{
		notifications: make([]Notification, 0),
	}
}

func (m *MemoryStorer) Save(notification Notification) {
	m.notifications = append(m.notifications, notification)
}

func (m *MemoryStorer) GetAll() []Notification {
	return m.notifications
}

func (m *MemoryStorer) GetCount() int {
	return len(m.notifications)
}
