package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// GET 호출
	resp, err := http.Get("http://csharp.news") // url에 get
	if err != nil {                             // 에러처리
		panic(err)
	}

	defer resp.Body.Close() // finally 함수로 닫아줌

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body) // 신기하게 닫은 후 데이터를 읽어옴, data 변수에 resp.body값을 대입한다.
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data)) // 콘솔 출력
}

/*
http.Get()은 한 문장으로 HTTP GET을 호출하는 장점이 있지만, Request시 헤더를 추가한다던가, Request 스트림을 추가하는 것과 같은 세밀한 컨트롤을 할 수 없는 단점이 있다. Request시 보다 많은 컨트롤이 필요하다면, Request 객체를 (NewRequest 생성자를 통해) 직접 생성해서 http.Client 객체를 통해 호출하면 된다.

아래 예제는 http.NewRequest() 생성자를 통해 Request 객체를 생성하고, 여기에 임의의 헤더를 추가하고, http.Client 객체를 통해 호출하는 코드이다. HTTP 호출결과는 Response 객체인데, 이 객체의 Body 필드를 통해 실제 Request 결과를 가져올 수 있다.
*/

func main2() {
	// Request 객체 생성
	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil) // get이니까 null ?
	if err != nil {                                                        // 에러처리
		panic(err)
	}

	//필요시 헤더 추가 가능
	req.Header.Add("User-Agent", "Crawler")

	// Client객체에서 Request 실행
	client := &http.Client{} // 포인터로 줌 ??
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 항상 닫은 후에 데이터를 읽는다.

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
}
