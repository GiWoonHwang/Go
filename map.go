/*
5Map은 키(Key)에 대응하는 값(Value)을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조이다. Go 언어는 Map 타입을 내장하고 있는데, "map[Key타입]Value타입" 과 같이 선언할 수 있다. 예를 들어 정수를 키로하고 문자열을 값으로 하는 맵 변수 idMap을 선언하기 위해서는 다음과 같이 할 수 있다.

1
var idMap map[int]string
이때 선언된 변수 idMap은 (map은 reference 타입이므로) nil 값을 갖으며, 이를 Nil Map이라 부른다. Nil map에는 어떤 데이타를 쓸 수 없는데, map을 초기화하기 위해 make()함수를 사용할 수 있다 (map 리터럴을 사용할 수도 있는 이는 아래 참조).

1
idMap = make(map[int]string)
make() 함수의 첫번째 파라미터로 map 키워드와 [키타입]값타입 을 지정하는데, 이때의 make()함수는 해시테이블 자료구조를 메모리에 생성하고 그 메모리를 가리키는 map value를 리턴한다 (map value는 내부적으로 runtime.hmap 구조체를 가리키는 포인터이다). 따라서 idMap 변수는 이 해시테이블을 가리키는 map을 가리키게 된다.

map은 make() 함수를 써서 초기화할 수도 있지만, 리터럴(literal)을 사용해 초기화할 수도 있다. 리터럴 초기화는 "map[Key타입]Value타입 { key:value }" 와 같이 Map 타입 뒤 { } 괄호 안에 "키: 값" 들을 열거하면 된다.
*/

package main

import "fmt"

func main() {
	var m map[int]string // [키타입]밸류타입
	//이때 선언된 변수 idMap은 (map은 reference 타입이므로) nil 값을 갖으며, 이를 Nil Map이라 부른다. Nil map에는 어떤 데이타를 쓸 수 없는데, map을 초기화하기 위해 make()함수를 사용할 수 있다.
	m = make([int]string)
	// 추가 혹은 갱신
	m[901] = "apple"
	m[134] = "grape"
	m[777] = "tomato"

	str := m[134]
	println(str)
	noData := m[999]
	println(noData) // 값이 없으면 nil 혹은 zero 리턴

	delete(m, 777)
	//fmt.Print()
}

func test() { // 키 체크
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}

	val, exists := tickers["MSFT"]
	if !exists {
		println("No MSFT ticker")
	}
}

// for 루프를 통한 map 열거

func test1() {
	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
