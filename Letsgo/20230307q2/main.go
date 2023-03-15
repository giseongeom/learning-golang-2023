package main

import "fmt"

// if 조건문을 이용해서 temperature 변수 값이 28보다 크면 "에어컨을 켠다" , 3 도 이하이면 "히터를 켠다" , 그 외는 "전기를 끈다"를 출력
func main() {
	temperature := 31
	fmt.Println("현재온도: ", temperature, "도")

	switch temperature != 0 {
	case temperature > 28:
			fmt.Println("에어컨을 켠다.")
		if temperature == 30 {
			fmt.Println("에어컨 온도를 1도 내려주세요.")
		} else if temperature == 31 {
			fmt.Println("에어컨 온도를 2도 내려주세요.")
		}
	case temperature == 28:
		fmt.Println("덥죠?")
	case temperature < 0:
		fmt.Println("추워요")
	default:
		fmt.Println("적정한 온도죠")
	}

	/*
		if temperature > 28 {
			fmt.Println("에어컨을 켠다")
			if temperature > 30 {
				fmt.Println("에어컨을 온도를 25도로 조절한다. Save money")
			}
		}

		if temperature <= 28 && temperature > 3 {
			fmt.Println("전기를 끈다")
		}

		if temperature <= 3 {
			fmt.Println("히터를 켠다")
		}
	*/
}
