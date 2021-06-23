package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Balance       = 1000
	EarnPoint     = 500
	LosePoint     = 100
	VictoryPoint  = 5000
	GameoverPoint = 0
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	balance := Balance
	for {
		fmt.Print("1~5사이의 값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요:")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5사이의 값만 입력하세요:")
		} else {
			r := rand.Intn(5) + 1
			if n == r {
				balance += EarnPoint
				fmt.Println("축하합니다. 맞추셨습니다. 남은 돈:", balance)
				if balance >= VictoryPoint {
					fmt.Println("You Win")
					break
				}
			} else {
				balance -= LosePoint
				fmt.Println("꽝 아쉽지만 다음 기회를.. 남은 돈:", balance)
				if balance <= GameoverPoint {
					fmt.Println("Game Over")
					break
				}
			}
		}
	}
}
