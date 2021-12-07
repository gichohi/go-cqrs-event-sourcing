package events

type ProductCreatedEventHandler interface {
	Handle()
}

func (event *ProductCreatedEvent) Handle() {

}