package main

// import (
// 	"fmt"       //표준 출력 패키지
// 	"io/ioutil" //파일과 데이터 스트림 패키지
// 	"net/http"  //http 클라이언트 및 서버 패키지
// )

// func main() {
// 	//GET 호출
// 	//http.Get함수를 사용하여 주어진 url에 http get 요청을 보낸다.
// 	//서버에서 응답을 받아 resp변수에 저장하고, 발생한 에러는 err변수에 저장
// 	resp, err := http.Get("https://naver.com")
// 	//http요청에 오류 발생시, err변수가 nil이 아니라면 panic함수를 호출하여 프로그램 중단
// 	if err != nil {
// 		panic(err)
// 	}

// 	//HTTP 응답의 본문을 닫는 작업을 지연시킴 / 함수가 봔환하기 전에 실행
// 	defer resp.Body.Close()

// 	//결과 출력
// 	//http응답의 본문에서 모든 데이터를 읽어옴.
// 	//읽어온 데이터를 data변수에 저장, 에러는 err에 저장
// 	data, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//데이터를 문자열로 변환하고 출력.
// 	//결과적으로 네이버 웹사이트의 html내용이 출력
// 	fmt.Printf("%s\n", string(data))
// }
