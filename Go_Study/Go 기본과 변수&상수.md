# 2. Go 기본과 변수와 상수
> 공부 교재 👉🏻 [예제로 배우는 Go 프로그래밍](http://golang.site/)
> </br>
> OS : Windows 10 / 개발환경 : VSCode

1. 간단한 Hello! 프로그램
   * 만약 Go 세팅이 안되셨다면 이 글을 꼭 확인하고 와주세요! 👉🏻 [Go 설치 및 초기 설정](https://github.com/de-quei/GoLang/blob/master/Go_Study/GO%20%EC%84%A4%EC%B9%98%20%EB%B0%8F%20%EC%B4%88%EA%B8%B0%EC%84%A4%EC%A0%95.md)
   * 저는 src폴더 안에 **StudyByExampleGo**라는 폴더를 따로 만든 후, 이 폴더안에 .go파일을 만들고 있습니다!
   * 제 기준으로 src > StudyByExampleGo > hello.go라는 파일을 추가해주세요!
   * 아래 코드를 붙여넣고 저장해주세요.
   ```go
   package main //함수들이 소속된 패키지명을 지정

   //메인함수 정의
   //Go는 main패키지 내의 Entry Point인 main()함수를 찾아 프로그램을 실행합니다.
   func main(){ 
    println("Hello Go!")
   }
   ```
   * 실행을 해줄건데 2가지 방법이 있습니다.
        1. VSCode에서 Ctrl + F5 실행
            * 디버그 콘솔창에 Hello Go!가 출력됩니다.
        2. 명령어를 사용한 터미널 출력
            * cd를 사용하여 C:\GO\src\StudyByExampleGo 경로까지 가주세요.
            * 밑 명령어를 붙이고 Enter를 눌러주세요.
            ```
            go run hello.go
            ```
            * 터미널에 Hello Go!가 출력됩니다.