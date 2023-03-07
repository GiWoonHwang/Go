/*
go 프로그래밍 언어는 다음과 같은 기본적인 데이터 타입들을 가지고 있다.

1. 불리언
	bool
2. 문자열
	string: string은 한번 생성되면 수정될 수 없는 immutable 타입이다.
3. 정수형 타입
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
4. float 및 복소수 타입
	float32 float64 complex64 complex128
5. 기타 타입
	byte: uint8과 동일하며 바이트 코드에 사용
	rune: int32과 동일하며 유니코드 코드포인트에 사용한다


2. 문자열
	문자열은 '' 와 "" 를 사용하여 표현할 수 있다.
	''로 둘러 쌓인 문자열은 raw string literal이라 부르는데, 이 안에 있는 문자열은 별도로 해석되지 않고, raw string 그대로의 값을 갖는다.
	예를 들어 문자열 안에 \n이 있는 경우 이는 newline으로 해석되지 않는다. 또한, backquoute는 복수 라인의 문자열을 표현할 때 자주 사용된다.

	""로 둘러 싸인 문자열은 interpreted string literal이라 부르는데, 복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 escape 문자열들은 특별한 의미로 해석된다. 예를 들어, 문자열 안에 \n이 있을 경우, 이는 newline으로 해석된다.
	이중인용부호를 이용해 문자열을 여러 아니에 걸쳐 쓰기 위해서는 + 연산자를 이용해 결합하여 사용한다.


*/

package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인.
	rawLiteral := `아리랑\n
아리랑\n
  아라리요`

	// Interpreted String Literal
	interLiteral := "아리랑아리랑\n아리리요"
	// 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
	// interLiteral := "아리랑아리랑\n" +
	//                 "아리리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
}

/*
데이터 타입 변환
하나의 데이터 타입에서 다른 데티어 타입으로 변환하기 위해서는 T(v)와 같이 표현하고 이를 type conversion이라 부르는데, 여기서 t는 변환하고자 하는 타입을 표시하고, v는 변환될 값을 지정한 것이다. 예를 들어 정수 100을 float로 변경하기 위하여
float32(100)처럼 표현하고, 문자열을 바이트 배열로 변경하기 위해여 []byte("abc")로 표현할 수 있다.
*/

func main2() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
