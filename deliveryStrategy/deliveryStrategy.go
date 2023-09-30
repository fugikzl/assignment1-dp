package deliveryStrategy

import "fmt"

//общий интерфейс для планов доставки
type DeliveryStrategy interface {
	Deliever(posylka string) string
}

//Премиумный план доставки
type PremiumDeliveryTrategy struct{}

func (s *PremiumDeliveryTrategy) Deliever(posylka string) string {
	return fmt.Sprintf("Package %s deliveried with premium plan very fast!\n", posylka)
}

//Дефолтный план доставки
type DefaultDeliveryStrategy struct{}

func (s *DefaultDeliveryStrategy) Deliever(posylka string) string {
	return fmt.Sprintf("Package %s delivered with default plan in 3-7 days.\n", posylka)
}
