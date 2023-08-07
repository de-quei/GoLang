package main

// 1. 변수 (var)
// 선언 --> var 변수명 데이터타입
var a int
// 선언 & 할당 --> var 변수명 데이터타입 = 값
var f float32 = 11.0

// 값 할당 
// 만약 선언된 변수가 Go 프로그램 내에서 사용되지 않는다면, 에러 발생!
a = 10
f = 12.0

// 동일한 타입의 변수가 복수 개 존재할 경우
// 선언 --> var 변수명1, 변수명2, 변수명 3... 데이터타입
var i, j, k int

// 선언 & 할당 --> var 변수명1, 변수명2, 변수명3... 데이터타입 = 값1, 값2, 값3...
var i, j, k int = 1, 2, 3

// 2. 상수 (const)
//선언 및 할당 --> const 변수명 데이터타입 = 값
const c int = 10
const s string = "Hi"
//타입 추론 가능
const a = 10
const g = "Go"

//여러개의 상수를 묶어서 지정
//const ( 변수명 = 값 ...)
const (
	Visa = "Visa"
	Master = "Master"
	Amex = "American Express"
)