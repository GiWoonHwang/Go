/*
함수는 여러 문장을 묶어서 실행하는 코드 블럭의 단위이다. Go에서 함수는 func 키워드를 사용하여 정의한다. func 뒤에 함수명을 적고 괄호 ( ) 안에 그 함수에 전달하는 파라미터들을 적게 된다.
함수 파라미터는 0개 이상 사용할 수 있는데, 각 파라미터는 파라미터명 뒤에 int, string 등의 파라미터 타입을 적어서 정의한다. 함수의 리턴 타입은 파라미터 괄호 ( ) 뒤에 적게 되는데,
이는 C와 같은 다른 언어에서 리턴 타입을 함수명 앞에 쓰는 것과 대조적이다. 함수는 패키지 안에 정의되며 호출되는 함수가 호출하는 함수의 반드시 앞에 위치해야 할 필요는 없다.
아래 예제는 say라는 함수를 정의한 예이다. say() 함수는 문자열 msg 파라미터를 하나 갖고 있으며, 리턴 값이 없으므로 별도의 리턴타입을 정의하지 않았다.
*/

package main

func main1() {
	msg := "Hello"
	say1(msg)
}

func say1(msg string) {
	println(msg)
}

/*
Go에서 파라미터를 전달하는 방식은 크게 Pass By Value와 Pass By Reference로 나뉜다.

Pass By Value
위의 [1. 함수]의 예제에서는 msg의 값 "Hello" 문자열이 복사되어 함수 say()에 전달된다. 즉, 만약 파라미터 msg의 값이 say() 함수 내에서 변경된다하더라도 호출함수 main()에서의 msg 변수는 변함이 없다.
Pass By Reference
아래의 예제에서처럼 msg 변수앞에 & 부호를 붙이면 msg 변수의 주소를 표시하게 된다. 흔히 포인터라 불리우는 이 용법을 사용하면 함수에 msg 변수의 값을 복사하지 않고 msg 변수의 주소를 전달하게 된다. 피호출 함수 say()에서는 *string 과 같이 파라미터가 포인터임을 표시하고 이때 say 함수의 msg는 문자열이 아니라 문자열을 갖는 메모리 영역의 주소를 갖게 된다. msg 주소에 데이타를 쓰기 위해서는 *msg = "" 과 같이 앞에 *를 붙이는데 이를 흔히 Dereferencing 이라 한다.

*/

func main() {
	msg := "Hello"
	say(&msg)
	println(msg) //변경된 메시지 출력
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}
