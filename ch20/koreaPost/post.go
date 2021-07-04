package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) send(parcel string) {
	fmt.Printf("KorePost sends %v parcel\n", parcel)
}
