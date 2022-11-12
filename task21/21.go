// Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// мы находимся в США и нужно зарядить телефон, но зарядка у нас самая обычная, европейская

//------------------------------- Европейский штепсель -------------------------------
// europePlug - у нас есть обычный европейский штепсель
type europePlug struct {
}

// insertEuropePlugIntoSocket - вставляем штепсель в розетку
func (plug europePlug) insertEuropePlugIntoSocket() {
	fmt.Printf("Вставляем европейский штепсель в розетку.\n")
}

//----------------------------------------------------------------------------

//------------------------------- Американский штепсель -----------------------------
// usPlug - в США все пользуются такими:
type usPlug struct {
}

// insertUsPlugIntoSocket (который без проблем вставляется в розетку)
func (plug usPlug) insertUsPlugIntoSocket() {
	fmt.Printf("Вставляем американский штепсель в розетку.\n")
}

//----------------------------------------------------------------------------

// insertUsPlug - мы можем вставить только американский штепсель,
// но очень хочется воспользоваться своим проводом...
type insertUsPlug interface {
	insertUsPlugIntoSocket()
}

//----------------------------- Адаптер ------------------------------------
// euToUsaAdapter (воспользуемся переходником - адаптером, для того, чтобы зарядить телефон)
type euToUsaAdapter struct {
	Adapter *europePlug
}

// insertUsPlugIntoSocket - адаптер вида европейский-американский спокойно вставляется в розетку
func (adapter euToUsaAdapter) insertUsPlugIntoSocket() {
	fmt.Printf("Вставляем европейский штепсель с помощью адаптера.\n")
}

//  ------------------------------ Клиент ---------------------------------------

// Client - это мы
type Client struct {
}

// ChargePhone - пользоваться розеткой и заряжать телефон мы умеем,
// но можем в данной ситуации использовать только американским штепселем
func (client Client) ChargePhone(plug insertUsPlug) {
	plug.insertUsPlugIntoSocket()
}

//----------------------------------------------------------------------------

func main() {
	client := Client{}
	usPlugExemplar := usPlug{}
	client.ChargePhone(usPlugExemplar)

	euPlugExemplar := europePlug{}
	adapter := euToUsaAdapter{Adapter: &euPlugExemplar}
	client.ChargePhone(adapter)
}
