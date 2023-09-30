package observer

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type Observer interface {
	SendMessage(message string)
}

type ConcreteSubject struct {
	observers []Observer
	message   string
}

func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	var indexToRemove int
	for i, o := range s.observers {
		if o == observer {
			indexToRemove = i
			break
		}
	}
	s.observers = append(s.observers[:indexToRemove], s.observers[indexToRemove+1:]...)
}

func (s *ConcreteSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.SendMessage(s.message)
	}
}

func (s *ConcreteSubject) UpdateMessage(message string) {
	s.message = message
	s.NotifyObservers()
}
