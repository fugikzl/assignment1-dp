package main

import (
	"ass1/deliveryStrategy"
	"ass1/order"
	"ass1/singletons"
	"ass1/user"
)

/**
	Background: ну это типо инет магаз, где есть несколько видов доставки и глобальный админ дэшборд(синглтон типа хз). Ну и крч виды доставки это стратегия пон, а обсерверы - это юзер и дэшборд которых надо регать в заказе(типа композиция пон).
**/
func main() {

	AdminDashboard := singletons.NewDashboard()

	JohnDoe := user.NewUser("John Doe")
	Anonim := user.NewUser("Anonim")

	order1 := order.NewOrder(&deliveryStrategy.DefaultDeliveryStrategy{}, "vibrator", *JohnDoe)
	order1.RegisterNotificationReceiver(JohnDoe)
	order1.RegisterNotificationReceiver(AdminDashboard)
	order1.DelieverPosylka()

	order2 := order.NewOrder(&deliveryStrategy.PremiumDeliveryTrategy{}, "iphone 15 max", *JohnDoe)
	order2.RegisterNotificationReceiver(JohnDoe)
	order2.RegisterNotificationReceiver(AdminDashboard)
	order2.DelieverPosylka()

	order3 := order.NewOrder(&deliveryStrategy.DefaultDeliveryStrategy{}, "poni", *Anonim)
	order3.RegisterNotificationReceiver(Anonim)
	order3.RegisterNotificationReceiver(AdminDashboard)
	order3.DelieverPosylka()

	AdminDashboard.SeeHistory()
}
