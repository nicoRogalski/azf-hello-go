package service

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

type Message struct {
	Id   string
	Code string
	Text string
}

func (h HelloService) GetHelloMessage() *Message {
	return &Message{
		Id:   "1234",
		Code: "hello",
		Text: "Welcome to azure function with Go",
	}
}
