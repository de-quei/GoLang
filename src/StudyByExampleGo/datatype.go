package main

func main() {

	//1. 문자열
	// Raw String Literal / 복수라인
	//rawLiteral := `아리랑\n아리랑\n아라리요`

	// Interpreted String Literal
	//interLiteral := "아리랑아리랑\n아라리요"
	// +를 사용하여 두 라인에 걸쳐 사용할 수도 있음.
	//interLiteral := "아리랑아리랑\n" + "아라리요"

	//fmt.Println(rawLiteral)
	//fmt.Println()
	//fmt.Println(interLiteral)

	//2. 타입 변환 (Type Conversion)
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)
	//출력결과 : +1.000000e+002 100
	//uint로 변환된 정수 int타입의 100 / uint로 변환된 값을 float32로 변환

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
	//출력결과 : 0xc000051f18 ABC
	//byte slice로 변환된 문자열 ABC / byte슬라이스로 변환된 값을 다시 문자열로 변환
}

/* 출력결과 */
//아리랑\n아리랑\n아라리요
//
//아리랑아리랑
//아라리요
