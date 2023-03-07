package main

import "fmt"

//import "math"

/*
for 문을 이용해서 0부터 9까지 숫자 출력
for 문을 이용해서 51부터 100까지 숫자 합산 출력
for 문에서 조건문만을 사용해서 3^10 보다 작은 3의 배수 합산 출력
for range문과  이차원 int 배열 ([2][5] int) 이용해서 1부터 9까지 출력
for 이중 중첩문을 사용해서 구구단 출력
책 233 Page
책 251 Page
*/

func main() {
	/*
	   //for 문을 이용해서 0부터 9까지 숫자 출력

	   		for i := 0 ; i < 10 ; i++ {
	   			fmt.Println(i)
	   		}
	*/

	/*
	   //for 문을 이용해서 51부터 100까지 숫자 합산 출력
	   	sum := 0
	   	for i := 51; i <= 100; i++ {
	   		sum += i
	   	}
	   	fmt.Println(sum)
	*/

	/*
		//for 문에서 조건문만을 사용해서 3^10 보다 작은 3의 배수 합산 출력
		sum := 0
		i := 0
		//limit := int(math.Pow(3, 10)) // 59049
		//fmt.Println(limit)
		limit := 59049

		for i < limit {
			sum = sum + i
			i = i + 3
			//fmt.Println(i)
		}

		fmt.Println(sum)
		// 581101209
	*/

	/*
		//for range문과  이차원 int 배열 ([2][5] int) 이용해서 1부터 9까지 출력

		var a1 = [2][5]int{
			{0, 1, 2, 3, 4},
			{5, 6, 7, 8, 9},
		}

		for _, num_a1 := range a1 {
			for _, num_a1_child := range num_a1 {
				fmt.Println(num_a1_child)
			}
		}
	*/

	//for 이중 중첩문을 사용해서 구구단 출력
	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			fmt.Printf("%d * %d = %d\n", x, y, x * y)
		}
		fmt.Println("********************")
	}

}
