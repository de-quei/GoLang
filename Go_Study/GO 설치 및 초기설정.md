# 1. GO 설치 및 초기 설정 (2023.08.06)
> windows x64 기준 / VSCode 사용

1. Go 설치
   * 자신에게 맞는 버전의 Go를 설치해줍니다! 👉🏻 [GO Download](https://go.dev/dl/)
   * 아무것도 건들지 말고 **next**만 눌러주세요!
   * 잘 따라하셨다면 GO는 **C:\Program Files\Go**에 저장되어있습니다.

</br>

2. 설치 확인
   * cmd에 명령어를 사용하여 설치를 확인해줄겁니다!
     * 버전 확인
      ```
      go version
      ```
      * 설치 확인
      ```
      go
      ```
      * 환경변수 확인
      ```
      go env
      ```

      </br>

3. 프로젝트 폴더 생성
   * C:\ 밑에 Go_study라는 폴더를 생성하고, 그 안에 **bin, src, pkg** 폴더를 생성해줍니다.

</br>
      
4. 환경변수 설정
   * 저희가 설정해야할 환경변수는 총 2가지 입니다!
   * **제어판 > 환경변수 설정 검색 > 환경변수 설정 > 시스템 변수**에 추가해줍니다!
     * GOROOT → 설치 시 자동 설정됨
     * GOPATH → C:\프로젝트 폴더명\ ( Ex.C:\Go_study\ )
     * GOBIN → $GOPATN\bin
  
> go env를 실행하여 환경변수가 올바르게 설정되었는지 확인해줍니다.

</br>

5. VScode설정
   * vscode를 열고 만든 Go_study 프로젝트 폴더를 열어줍니다.
   * 확장팩에서 Go를 검색하고 install해줍니다.
   * ctrl + shift + p를 누르고 검색창에 **go install/update tools**를 검색합니다.
   * 나오는 항목들을 모두 체크후 install해주면 console에 "You are ready to Go :)"가 나오는 것을 확인하실 수 있습니다!

> 이제부터 많은 사람들이 난관에 봉착합니다...

</br>

6. 테스트 코드 실행
   * 우리는 테스트 코드를 실행할겁니다! ( 에러가 많이 나는게 정상입니다...! )
   * src폴더 안에 test.go파일을 생성하고 이 코드를 붙여주세요!
   ```go
   package main

   import "fmt"

   func main() {
	   fmt.Println("hello go!")
   }
   ```
   * **이 코드만 붙여도 오른쪽 하단에 에러가 뜰텐데 그거 다 install해줍니다!**

</br>

7. git mod init
   * 저같은 경우는 "gopls was not able to find modules in your workspace" 이런 에러가 났는데.. 두 가지 방법을 해봤습니다.
     * 방법 1 / setting.json 파일 수정
         * 파일>기본설정>설정에 들어가서 **gopls**를 검색해주세요.
         * 그럼 밑에 setting.json파일이 뜨는데 클릭해주세요.
         * setting.json파일에 밑 코드를 붙여넣고 저장하고 vscode를 다시 껐다 켜주세요.
         ```json
         "gopls": {
            "experimentalWorkspaceModule": true,
         }
         ``` 
         * 근데 저는 이걸 해도.. "Invalid settings: gopls setting "experimentalWorkspaceModule" is deprecated" 이게 뜨면서 안되더라구요..
  
     * 방법 2/git mod init
       * cmd에 이 코드를 붙여넣어 실행해주세요.
       ```
       go mod init github.com/lecture
       ```
       * 저는 이걸 실행하니까 vscode상에 go.mod파일이 추가되어 더이상 에러가 발생하지 않았습니다!

8. 실행
   * 더 이상 무서울건 없습니다. 테스트코드를 실행해보면 콘솔에 hello go!가 출력됩니다!

</br>

개발자로 일하는 사람도 아니고.. 개발 배운지 1년 반 밖에 안된 마이스터고 고등학생이라서 자료찾기도 다른 언어에 비해 너무 어려웠고 해결 안될거같아서 몇시간을 쩔쩔맨거같은데 그래도 해결해서 기분이 좋습니다! 나중에 필요한게 있으면 또 추가하겠습니다!