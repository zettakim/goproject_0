package main

import (
	"goproject/ch20-1/fedex"
	"goproject/ch20-1/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("조르바", fedexSender)

	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린왕자", koreaPostSender)
	SendBook("조르바", koreaPostSender)
}
