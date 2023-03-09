/*
Go의 표준 패키지인 net/http 패키지는 웹 관련 서버 (및 클라이언트) 기능을 제공한다. Go에서 HTTP 서버를 만들기 위해 중요한 http 패키지 메서드로 ListenAndServe(), Handle(), HandleFunc() 등을 들 수 있다.
ListenAndServe() 메서드는 지정된 포트에 웹 서버를 열고 클라이언트 Request를 받아들여 새 Go 루틴에 작업을 할당하는 일을 한다. Handle()과 HandleFunc() 메서드는 요청된 Request Path에 어떤 Request 핸들러를 사용할 지를 지정하는 라우팅 역활을 한다.

아래 예제는 간단한 HTTP 서버를 구현한 예로써, 브라우져에서 http://localhost:5000/hello 라고 치면, HandleFunc()의 /hello Path에 대한 익명함수를 실행하여 Hello World를 출력하게 된다.
익명함수 안의 http.ResponseWriter 파라미터는 HTTP Response에 무언가를 쓸 수 있게 하며, http.Request 파라미터는 입력된 Request 요청을 검토할 수 있게 한다. 마지막의 ListenAndServe() 메서드는 여기서 2개의 파라미터를 갖고 있는데,
첫번째는 포트 5000 에서 Request를 Listen 할 것을 지정하고, 두번째는 어떤 ServeMux를 사용할 지를 지정하는데 nil인 경우 DefaultServeMux를 사용한다. (ServeMux는 기본적으로 HTTP Request Router (혹은 Multiplexor) 인데, 일반적으로 내장된 DefaultServeMux을 사용하지만, 개발자가 별도로 ServeMux를 만들어 Routing 부분을 세밀하게 제어할 수 있다).
DefaultServeMux를 사용하는 경우, Handle() 혹은 HandleFunc()을 사용하여 라우팅 패턴을 추가하게 된다.
*/

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
