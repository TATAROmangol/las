package entities

type Order struct{
	Id string
	Item string 
	Quantity int
}

func NewOrder(id, item string, quantity int) Order{
	return Order{id, item, quantity}
}