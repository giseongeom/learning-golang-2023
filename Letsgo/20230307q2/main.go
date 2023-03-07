package main

import "fmt"

// if 조건문을 이용해서 temperature 변수 값이 28보다 크면 "에어컨을 켠다" , 3 도 이하이면 "히터를 켠다" , 그 외는 "전기를 끈다"를 출력
func main() {
	temperature := 3

	if temperature > 28 {
		fmt.Println("에어컨을 켠다")
	}

	if temperature <= 28 && temperature > 3 {
		fmt.Println("전기를 끈다")
	}

	if temperature <= 3 {
		fmt.Println("히터를 켠다")
	}
}
