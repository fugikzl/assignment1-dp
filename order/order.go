package order

import (
	"ass1/deliveryStrategy"
	"ass1/observer"
	"ass1/user"
)

type Order struct {
	deliveryStrategy deliveryStrategy.DeliveryStrategy //реализации стратегии для разных видов доставки
	posylka          string
	user             user.User
	subject          observer.ConcreteSubject
}

//конструктор заказа
func NewOrder(deliveryStrategy deliveryStrategy.DeliveryStrategy, posylka string, user user.User) *Order {
	return &Order{
		deliveryStrategy: deliveryStrategy,
		posylka:          posylka,
		user:             user,
		subject:          *observer.NewConcreteSubject(),
	}
}

func (o *Order) RegisterNotificationReceiver(receiver observer.Observer) {
	o.subject.RegisterObserver(receiver)
}

func (o *Order) RemoveNotificationReceiver(receiver observer.Observer) {
	o.subject.RemoveObserver(receiver)
}

//метод исполнения заказа
func (o *Order) DelieverPosylka() {
	deliveryresult := o.deliveryStrategy.Deliever(o.posylka)
	o.subject.UpdateMessage(deliveryresult)
}
