package main

import "fmt"

func main() {

	//1. 문자열
	// Raw String Literal / 복수라인
	rawLiteral := `아리랑\n아리랑\n아라리요`

	// Interpreted String Literal
	interLiteral := "아리랑아리랑\n아라리요"
	// +를 사용하여 두 라인에 걸쳐 사용할 수도 있음.
	//interLiteral := "아리랑아리랑\n" + "아라리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

}

/* 출력결과 */
//아리랑\n아리랑\n아라리요
//
//아리랑아리랑
//아라리요
