package main

import "fmt"

func main() {

	var c int = 10
	p := &c

	// struct 값 표기 (초기값 선언에 맞춰서..)
	fmt.Printf("struct1: %v\n", p)
	// struct 값 표기 (변수명과 함께..)
	fmt.Printf("struct2: %+v\n", p)
	// struct 값 표기 (struct type과 함께...)
	fmt.Printf("struct3: %#v\n", p)
	// data type 표기
	fmt.Printf("type: %T\n", p)
	//
	fmt.Printf("bool: %t\n", true)
	// 10진수 값 표기
	fmt.Printf("int: %d\n", 123)
	// Bianry 값 표기
	fmt.Printf("bin: %b\n", 14)
	// char 표기
	fmt.Printf("char: %c\n", 33)
	// 16진수 (Hex) 값 표기
	fmt.Printf("hex: %x\n", 456)
	// 문자열에  대해서도 16진수 (Hex) 값으로  표기 가능
	fmt.Printf("str3: %x\n", "hex this")
	// 소수 값 표기
	fmt.Printf("float1: %f\n", 78.9)
	// 지수값 표기
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)
	// 문자열 표기
	fmt.Printf("str1: %s\n", "\"string\"")
	// 문자열 표기 (특수문자 그대로..)
	fmt.Printf("str2: %q\n", "\"string\"")
	// 포인터 주소 값
	fmt.Printf("pointer: %p\n", &p)
	// 숫자만큼 자리 수 확보 해서 표기 (오른쪽 기준)
	fmt.Printf("width1: |%06d|%06d|\n", 12, 345)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	// 소수점 기준으로 정수와 소수 부분 자리수 확보해서 표기
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	// 소수점 기준으로 정수와 소수 부분 자리수 확보해서 표기 (왼쪽 기준)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// 문자열도 자리수 확보 (오른쪽 기준)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	// 문자열도 자리수 확보 (왼쪽 )
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
}
