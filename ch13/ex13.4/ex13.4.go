package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User  //필드명 생략
	Price int
	Level int
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser{
		User{"화랑", "hwarang", 40, 10},
		250,
		3,
	}
	fmt.Printf("User: %s, ID: %s, Age: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP User: %s, ID: %s, Age: %d, VIP Level: %d, USER Level: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,
		vip.User.Level,
	)
}
