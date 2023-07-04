/*
1. JSON
JSON (JavaScript Object Notation)은 데이타를 교환하는 한 포맷으로서 그 단순함과 유연함 때문에 널리 사용되고 있다. 특히 웹 브라우져와 웹서버 사이에 데이타를 교환하는데 많이 사용되고 있다. JSON 포맷은 기본적으로 Key-Value Pair의 컬렉션이다.

Go 에서 JSON을 사용하기 위해서는 표준패키지 encoding/json 을 사용하면 된다. 표준패키지는 https://golang.org/pkg/ 에 모두 자세히 설명되어 있는데, 이 중 JSON 관련 패키지는 https://golang.org/pkg/encoding/json/ 에서 참조할 수 있다.

2. JSON 인코딩
Go 데이타를 JSON 포맷으로 변환(인코딩)하기 위해서는 encoding/json 패키지의 Marshal() 함수를 사용한다. 흔히 Go 구조체 혹은 map 데이타를 JSON으로 인코딩하게 되는데, 해당 Go 데이타 값을 json.Marshal()의 파라미터로 전달하면, JSON으로 인코딩된 바이트배열과 에러객체를 리턴한다. 만약 JSON으로 인코딩된 바이트배열을 다시 문자열로 변경할 필요가 있다면, string(바이트배열)과 같이 변경할 수 있다.

한가지 유의할 점은 JSON의 Key는 문자열이어야 한다. Go 구조체의 경우 자동으로 필드명을 문자열로 사용하게 되지만, map인 경우는 map[string]T 처럼 Key가 string인 map만 지원한다 (Value T는 어느 타입이든 상관 없음).
*/

package main

import (
	"encoding/json"
	"fmt"
)

// Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {

	// Go 데이타
	mem := Member{"Alex", 10, true}

	// JSON 인코딩
	jsonBytes, err := json.Marshal(mem) // [123 34 78 97 109 101 34 58 34 65 108 101 120 34 44 34 65 103 101 34 58 49 48 44 34 65 99 116 105 118 101 34 58 116 114 117 101 125]
	if err != nil {
		panic(err)
	}

	// JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)
}

/*
JSON으로 인코딩된 데이타를 다시 디코딩하기 위해서는 encoding/json 패키지의 Unmarshal() 함수를 사용한다.
Unmarshal() 함수의 첫번째 파라미터에는 JSON 데이타를, 두번째 파라미터에는 출력할 구조체(혹은 map)를 포인터로 지정한다. 리턴값은 에러객체이고, 에러가 없을 경우, 두번째 파라미터에 원래 데이타가 복원된다.
*/
func main2() {
	// 테스트용 JSON 데이타
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	// JSON 디코딩
	var mem Member
	err := json.Unmarshal(jsonBytes, &mem) // json 데이터, 출력할 구조체를 포인터로 지정
	if err != nil {
		panic(err)
	}

	// mem 구조체 필드 엑세스
	fmt.Println(mem.Name, mem.Age, mem.Active)
}
