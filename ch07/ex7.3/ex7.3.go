package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는", avg, "입니다.")
}

func main() {
	PrintAvgScore("홍길동", 80, 74, 95)
	PrintAvgScore("서태지", 88, 92, 53)
	PrintAvgScore("강진우", 78, 73, 78)
}
