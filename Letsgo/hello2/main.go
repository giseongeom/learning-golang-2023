package main

import "fmt"

func main() {
	// var 를 이용해서 int , float, bool, string  타입 선언 및 해당 값 표기
	var a1 int = 1
	var a2 float64 = 1.2
	var a3 bool = true
	var a4 string = "merong"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(" ")

	//:= 를 이용해서 int , float, bool, string  타입 선언 및 해당 값 표기
	b1 := 1
	b2 := 1.2
	b3 := true
	b4 := "merong"
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(" ")

	//aa 는 int64 , bb 는 float64  타입으로 선언 및 값 세팅 후 cc 변수에  bb값의 integer 로 형 변환 ,  dd 변수에 aa값과 0.5를 곱한  float형 변환 해서 표기
	var aa int64 = 2
	var bb float64 = 3.1415
	var cc int64 = int64(bb)
	var dd float64 = float64(aa) * 0.5
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(" ")

	/*
		// https://stackoverflow.com/questions/37582550/twos-complement-and-fmt-printf
		var my_i int8 = 5
		var my_ni int8 = -5
		fmt.Printf("%d\n", my_i)
		fmt.Printf("%b\n", my_i)
		fmt.Printf("%b\n", uint8(my_ni))
		fmt.Println(" ")

	*/
	//15의 보수는?
	var c1 int16 = 15
	var c2 int16 = -15
	fmt.Println(c1)
	fmt.Printf("%016b\n", c1)
	fmt.Printf("%016b\n", c2)
	fmt.Printf("%b\n", uint16(c2))
	fmt.Println(" ")

	//23의 2의 보수는?
	var d1 int8 = 23
	var d2 int8 = -23
	fmt.Println(d1)
	fmt.Printf("%08b\n", d1)
	fmt.Printf("%08b\n", d2)
	fmt.Printf("%b\n", uint8(d2))
	fmt.Println(" ")

	//23의 보수는?
	fmt.Println("23의 보수는?")
	var e1 int8 = 23
	fmt.Println(e1)
	fmt.Printf("%b\n", uint8(^e1)) // 11101000
	fmt.Println(" ")

	/*
			"안녕하세요?"  문자열과 "Go 어렵지 않아요!!" 문자열을 합쳐서 표기
		    탭 기호  이용해서 라인 하나로 표현하기
		    라인  기호 이용해서 두 라인으로 표현 하기
		    + 기호 이용 해서 라인 하나로 표현하기
		     라인 기호 이용하지 않고 두 라인으로 표현하기
		    "안녕하세요?"  문자열에 new line 기호 표현 하기
	*/

	f1 := "안녕하세요?"
	f2 := "Go 어렵지 않아요!!"
	fmt.Printf("%s\t %s\n", f1, f2)
	fmt.Printf("%s\n%s\n", f1, f2)
	fmt.Printf("%s\n", f1+" "+" "+f2)
	fmt.Println(" ")

	// https://freshman.tech/snippets/go/multiline-strings/
	fmt.Println("f3")
	f3 := `안녕하세요?"
Go 어렵지 않아요!!
Love`
	fmt.Printf("%s\n", f3)
	fmt.Println(" ")

	fmt.Printf("%s\n", f1+"\\n")
	fmt.Println(" ")

	const (
		C1 uint = iota + 1 // 1
		C2
		C3
	)
	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)

}
