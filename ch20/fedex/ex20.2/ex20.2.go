package main

import "goproject/ch20/fedex"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{}
	SendBook("어린왕자", sender)
	SendBook("조르바", sender)
}
