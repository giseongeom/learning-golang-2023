package main

import (
	"fmt"
)

func main() {
	/*
					slice := []int{0, 1, 2, 3, 4, 5}
					s := slice[2:5]
					fmt.Println(s)
					s = s[1:]
					fmt.Println(s)

					부분 slice를 이용하고, 처음과 마지막 인덱스를 전부 사용해서  3,4,5 값을 출력
					s1 := slice[3:]
					s2 := s1[0:]
					fmt.Println(s2)

					다음 코드에서의 s2의 capacity 값은?
					fmt.Println(len(slice))
					fmt.Println(cap(slice))

					s2 := slice[1:3:4]
					fmt.Println(cap(s2))

					s3 := slice[1:3]
					fmt.Println(cap(s3))

					다음 코드에서 slice3, slice4  값들은? length와 capacity값은?

					slice3 := make([]int,3,5)
					slice4 := append(slice3, 3,4)

				    fmt.Println(len(slice3), cap(slice3))
				    fmt.Println(slice3)
				    fmt.Println(len(slice4), cap(slice4))
				    fmt.Println(slice4)

					slice3[1] = 1
				    fmt.Println(len(slice3), cap(slice3))
				    fmt.Println(slice3)
				    fmt.Println(len(slice4), cap(slice4))
				    fmt.Println(slice4)

					slice3 = append(slice3, 5)
				    fmt.Println(len(slice3), cap(slice3))
				    fmt.Println(slice3)
				    fmt.Println(len(slice4), cap(slice4))
				    fmt.Println(slice4)

			source := []int{0, 1, 2}
			target := make([]int, len(source), cap(source)*2)
			copy(target, source)
			target[0] = 100
			fmt.Println(target) // 100,1,2

			append와 sub-slice를 이용해서 Slice의 4번째 index를 삭제 후 결과 확인,
			3번째 index에 30 추가

		slice := []int{0, 1, 2, 3, 4, 5}
		s1 := slice[0:3]
		s2 := slice[4:]
		s3 := append(s1, s2...)
		fmt.Println(s3) // 0, 1, 2, 4, 5
	*/

	// 3번째 index에 30 추가
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)

	s1 := slice[0:2]     // 0, 1
	//fmt.Println("s1:", s1)

	s2 := make([]int, 2, 4)

	c1 := copy(s2, s1)
	fmt.Println(c1, "copied")

	//fmt.Println(len(s2), cap(s2))
	s2 = append(s2, 30) // 0, 1, 30
	//fmt.Println("s2:", s2)

    s3 := slice[2:] // 2, 3, 4, 5
	s4 := append(s2, s3...) // 0, 1, 30, 2, 3, 4, 5

	fmt.Println("")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s3[0]:", &s3[0])
	fmt.Println("s4:", s4)
	fmt.Println("slice:", slice)

	fmt.Println("")
	fmt.Println("s1[0]:", &s1[0])
	fmt.Println("slice[0]:", &slice[0])
}
