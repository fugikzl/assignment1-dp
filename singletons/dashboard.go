package singletons

import "fmt"

type Dashboard struct {
	receivedMessages []string
}

func NewDashboard() *Dashboard {
	return &Dashboard{
		receivedMessages: []string{},
	}
}

func (d *Dashboard) SendMessage(message string) {
	fmt.Printf("Dashboard received message: %s\n", message)
	d.receivedMessages = append(d.receivedMessages, message)
}

func (d *Dashboard) SeeHistory() {
	fmt.Println("Dashboard history:")
	fmt.Println(d.receivedMessages)
}
