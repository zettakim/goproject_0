package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}
	fmt.Printf("User: %s, ID: %s, Age: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP User: %s, ID: %s, Age: %d, VIP Level: %d, VIP가격: %d\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
