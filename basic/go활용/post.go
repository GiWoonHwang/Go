/*
Go의 표준 패키지인 http 패키지는 웹 관련 클라이언트 및 서버 기능을 제공한다. 그 중 http.Post() 메서드는 웹서버로 간단히 데이타를 POST 하는데 사용된다.

아래 예제는 테스트 웹사이트인 httpbin.org에 임의의 텍스트를 POST를 사용해서 보내는 코드이다. Post()메서드의 첫번째 파라미터는 Post를 받는 URL을 적고, 두번째는 Request Body의 MIME 타입을, 그리고 마지막에는 전송할 데이타를 (io.Reader로) 보낸다.

httpbin.org/post 는 전송한 데이타를 그대로 리턴하는데, Response의 data를 체크하면 전송 데이타가 그대로 리턴됨을 볼 수 있다.
*/

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 간단한 http.Post 예제
	reqBody := bytes.NewBufferString("Post plain text")
	resp, err := http.Post("http://httpbin.org/post", "text/plain", reqBody) // 주소, body 타입, 전송할 데이터
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

/*
http.PostForm()은 Form 데이타를 보내는데 유용한 메서드이다. 즉, 일반 웹페이지에서 Submit 버튼을 누르면, 입력 컨트롤들의 데이타를 Form 데이타로 서버에 전송하는데, 이 PostForm() 함수를 사용하면 동일한 효과를 내며 데이타를 쉽게 웹서버로 보낼 수 있다.
아래 예제는 Name과 Age 데이타를 Form 형식으로 웹서버에 전송하는 예이다. httpbin.org/post 는 전송한 데이타를 그대로 리턴하므로 Response에서는 form 필드에 Name과 Age 데이타가 동일함을 볼 수 있다.
*/
func main1() {
	// 간단한 http.PostForm 예제
	resp, err := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

/*
JSON 데이타를 Post하는 것은 위의 Plain Text를 Post하는 것과 비슷하다. 다만, 데이타가 JSON 포맷이므로 http.Post()의 두번째 파라미터에 application/json를 적고, 세번째 파라미터에 JSON으로 인코딩된 데이타를 전달하면 된다.
아래 예제에서 JSON 데이타는 encoding/json 표준 패키지의 Marshal() 함수를 써서 임의의 구조체 데이타를 JSON으로 변경하는 방법을 썼다.
*/
type Person struct {
	Name string
	Age  int
}

func main2() {
	person := Person{"Alex", 10}
	pbytes, _ := json.Marshal(person) // json 객체로 변환
	buff := bytes.NewBuffer(pbytes)   // 근데 이걸 왜 또 해주는걸까 ?
	resp, err := http.Post("http://httpbin.org/post", "application/json", buff)

	//...(elided)...
}
