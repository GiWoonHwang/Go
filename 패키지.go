/*
1. go
Go는 패키지(Package)를 통해 코드의 모듈화, 코드의 재사용 기능을 제공한다. Go는 패키지를 사용해서 작은 단위의 컴포넌트를 작성하고, 이러한 작은 패키지들을 활용해서 프로그램을 작성할 것을 권장한다.

Go는 실제 프로그램 개발에 필요한 많은 패키지들을 표준 라이브러리로 제공한다. 이러한 표준 라이브러리 패키지들은 GOROOT/pkg 안에 존재한다. GOROOT 환경변수는 Go 설치 디렉토리를 가리키는데, 보통 Go 설치시 자동으로 추가된다. 즉, 윈도우즈에서 Go를 설치했을 경우 디폴트로 C:\go 에 설치되며, GOROOT는 C:\go를 가리킨다.

Go에 사용하는 표준패키지는 https://golang.org/pkg 에 자세히 설명되어 있다. Go 프로그래밍에서 표준패키지를 자주 불러 사용하므로 이 링크를 자주 참조하게 될 것이다.


2. main
일반적으로 패키지는 라이브러리로서 사용되지만, "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다. 패키지명이 main 인 경우,

컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다. 그리고 이 main 패키지 안의 main() 함수가 프로그램의 시작점 즉 Entry Point가 된다. 패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

3. import
다른 패키지를 프로그램에서 사용하기 위해서는 import 를 사용하여 패키지를 포함시킨다. 예를 들어, Go의 표준 라이브러리인 fmt 패키지를 사용하기 위하여, import "fmt" 와 같이 해당 패키지를 포함시킬 것을 선언해 준다. Import 후에는 아래 예제처럼 fmt 패키지의 Println() 함수를 호출하여 사용할 수 있다.
패키지를 import 할 때, Go 컴파일러는 GOROOT 혹은 GOPATH 환경변수를 검색하는데, 표준 패키지는 GOROOT/pkg 에서 그리고 사용자 패키지나 3rd Party 패키지의 경우 GOPATH/pkg 에서 패키지를 찾게 된다.

GOROOT 환경변수는 Go 설치시 자동으로 시스템에 설정되지만, GOPATH는 사용자가 지정해 주어야 한다. GOPATH 환경변수는 3rd Party 패키지를 갖는 라이브러리 디렉토리나 사용자 패키지가 있는 작업 디렉토리를 지정하게 되는데, 복수 개일 경우 세미콜론(윈도우즈의 경우)을 사용하여 연결한다.

4. 패키지 scope
패키지 내에는 함수, 구조체, 인터페이스, 메서드 등이 존재하는데, 이들의 이름(Identifier)이 첫문자를 대문자로 시작하면 이는 public 으로 사용할 수 있다. 즉, 패키지 외부에서 이들을 호출하거나 사용할 수 있게 된다. 반면, 이름이 소문자로 시작하면 이는 non-public 으로 패키지 내부에서만 사용될 수 있다.

5. 패키지 init 함수와 alias
개발자가 패키지를 작성할 때, 패키지 실행시 처음으로 호출되는 init() 함수를 작성할 수 있다. 즉, init 함수는 패키지가 로드되면서 실행되는 함수로 별도의 호출 없이 자동으로 호출된다.

package testlib

var pop map[string]string

func init() {   // 패키지 로드시 map 초기화
    pop = make(map[string]string)
}
경우에 따라 패키지를 import 하면서 단지 그 패키지 안의 init() 함수만을 호출하고자 하는 케이스가 있다. 이런 경우는 패키지 import 시 _ 라는 alias 를 지정한다. 아래는 other/xlib 패키지를 호출하면서 _ alias를 지정한 예이다.


package main
import _ "other/xlib"
만약 패키지 이름이 동일하지만, 서로 다른 버젼 혹은 서로 다른 위치에서 로딩하고자 할 때는, 패키지 alias를 사용해서 구분할 수 있다.

import (
    mongo "other/mongo/db"
    mysql "other/mysql/db"
)
func main() {
    mondb := mongo.Get()
    mydb := mysql.Get()

6. 사용자 정의 패키지 생성

*/