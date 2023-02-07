/*
a 는 5 값을 갖는 변수, b 는 18 값을 갖는 변수, c 는 3.5 값을 갖는 변수,  d는 6.5 값을 갖는 변수,  e 는 127값을 갖는 int8, g는 -128값을 갖는 int8
fmt.Println 함수 이용
산술 연산자와 a와 b 변수를 이용해서  결과 값이 23, 3,  90, -13 이 나오도록
산술 연산자와 c와 d 변수를 이용해서  결과 값이 10.0, 3.0,  1.857 값이 나오도록
산술 연산자와 a와 c 변수를 이용해서  결과 값이 8, 8.5, 17.5, 0.7 값이 나오도록
a에 대해 값 1증가 , b  값에 -1 감소 시켜서 module연산 후 결과 공유
a & 3  값은?
b | 3 값은 ?
a 값에서 << 1 값, >> 3 값은 ?
b 값에서 << 2  값,  >> 2 값은 ?
e < e +1
g  > g - 1
`! (2 < 5 || 10 <5)
a = b = 5 를 하면 에러가 발생. 이걸 b 변수를 이용해서 두 라인으로 작성
c와 d의 값을 할당(대입) 연산자를 통해서 바꿔치기 (한줄로)
158 Page 연습 문제
*/

package main

import "fmt"

func main() {
	a := 5
	b := 18
	c := 3.5
	d := 6.5
	var e int8 = 127
	var g int8 = -128

	// 산술 연산자와 a와 b 변수를 이용해서  결과 값이 23, 3,  90, -13 이 나오도록
	fmt.Printf("a + b = %d\n", a + b)
	fmt.Printf("b / a = %d\n", b / a)
	fmt.Printf("b * a = %d\n", b * a)
	fmt.Printf("a - b = %d\n", a - b)

	//산술 연산자와 c와 d 변수를 이용해서  결과 값이 10.0, 3.0,  1.857 값이 나오도록
	fmt.Printf("c + d = %02.1f\n", c + d)
	fmt.Printf("d - c = %02.1f\n", d - c)
	fmt.Printf("d / c = %01.3f\n", d / c)

	//산술 연산자와 a와 c 변수를 이용해서  결과 값이 8, 8.5, 17.5, 0.7 값이 나오도록
	fmt.Printf("a + c = int %d\n", a + int(c))
	fmt.Printf("a + c = float64 %1.1f\n", float64(a) + c)
	fmt.Printf("a * c = float64 %2.1f\n", float64(a) * c)
	fmt.Printf("c / a = float64 %2.1f\n", c / float64(a))

	//a에 대해 값 1증가 , b  값에 -1 감소 시켜서 module연산 후 결과 공유
	fmt.Printf("before\ta: %d, b: %d\n", a, b)
	a = a + 1
	b = b - 1
	fmt.Printf("after\ta: %d, b: %d\n", a, b)
	fmt.Println("a % b = int", a % b)

	//a & 3  값은?
	fmt.Printf("a & 3 = %08b & %08b ==> %d %08b\n", a, 3, a & 3,  a & 3)
	fmt.Printf("b | 3 = %08b | %08b ==> %d %08b\n", b, 3, b | 3,  b | 3)

	//a 값에서 << 1 값, >> 3 값은 ?
	fmt.Printf("a: %d (%08b)\n", a, a)
	a <<= 1
	fmt.Printf("a << 1 = %d %08b\n", a, a)
	a >>= 3
	fmt.Printf("a >> 3 = %d %08b\n", a, a)

	//b 값에서 << 2  값,  >> 2 값은 ?
	fmt.Printf("b: %d (%08b)\n", b, b)
	b <<= 2
	fmt.Printf("b << 2 = %d %08b\n", b, b)
	b >>= 2
	fmt.Printf("b >> 2 = %d %08b\n", b, b)

	//e < e +1
	fmt.Printf("e: %d (%08b)\n", e, e)
	fmt.Printf("e + 1: %d (%08b)\n", e + 1, e + 1)
	fmt.Printf("e < e + 1: %v\n", e < (e + 1))

	//g  > g - 1
	fmt.Printf("g: %d (%08b)\n", g, g)
	fmt.Printf("g - 1: %d (%08b)\n", g - 1, g - 1)
	fmt.Printf("g > g - 1: %v\n", g > (g - 1))

	//!(2 < 5 || 10 <5)
	fmt.Printf("!(2 < 5 || 10 <5: %v\n", !(2 < 5 || 10 <5))

	//a = b = 5 를 하면 에러가 발생. 이걸 b 변수를 이용해서 두 라인으로 작성
	fmt.Printf("before\ta: %d, b: %d\n", a, b)
	a = 5
	b = a
	fmt.Printf("after\ta: %d, b: %d\n", a, b)

	//c와 d의 값을 할당(대입) 연산자를 통해서 바꿔치기 (한줄로)
	fmt.Printf("before\tc: %f, d: %f\n", c, d)
	c, d = d, c
	fmt.Printf("before\tc: %f, d: %f\n", c, d)
}