package main

import "fmt"

/*
- if 조건문을 이용해서 light 변수 값이 green이면 "길을 건넌다" ,  그 외의 값이면 "대기한다"를 출력
- if 조건문을 이용해서 temperature 변수 값이 28보다 크면 "에어컨을 켠다" , 3 도 이하이면 "히터를 켠다" , 그 외는 "전기를 끈다"를 출력
- 책 197 Page
*/

//- if 조건문을 이용해서 light 변수 값이 green이면 "길을 건넌다" ,  그 외의 값이면 "대기한다"를 출력
func main() {
	light := "black"
	if light == "green" {
		fmt.Println("길을 건넌다")
	} else {
		fmt.Println("대기한다")
	}
}