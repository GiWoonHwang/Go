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
	// msg := "Hello"
	// say(&msg)
	// println(msg) //변경된 메시지 출력
	main2()

}

func say(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}

/*
가변인자 함수
함수에 고정된 수의 파라미터들을 전달하지 않고 다양한 숫자의 파라미터를 전달하고자 할 때 가변 파라미터를 나타내는 ... (3개의 마침표)을 사용한다. 즉 문자열 가변 파라미터를 나타내기 위해서 ...string 과 같이 표현한다.
가변 파라미터를 갖는 함수를 호출할 때는 0개, 1개, 2개, 혹은 n개의 동일타입 파라미터를 전달할 수 있다. 이렇게 가변 파라미터를 받아들이는 함수를 Variadic Function (가변인자함수)라고 부른다.
아래 예제는 say 함수에 4개의 문자열을 전달할 수도 있고, 1개의 문자열을 전달할 수도 있음을 예시하고 있다
*/

func main2() {
	say2("This", "is", "a", "book")
	say2("Hi")
}

func say2(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

/*
Go 프로그래밍 언어에서 함수는 리턴값이 없을 수도, 리턴값이 하나 일 수도, 또는 리턴값이 복수 개일 수도 있다. C 언어에서 void 혹은 하나의 값만을 리턴하는 것과 대조적으로 Go 언어는 복수개의 값을 리턴할 수 있다.
Go 언어는 또한 Named Return Parameter 라는 기능을 제공하는데, 이는 리턴되는 값들을 (함수에 정의된) 리턴 파라미터들에 할당할 수 있는 기능이다.

함수에서 리턴값이 있는 경우는 func 문의 (파라미터 괄호 다음) 마지막에 리턴값의 타입을 정의해 준다. 그리고 값을 리턴하기 위해 함수내에서 return 키워드를 사용한다. 예를 들어, 아래 예제는 sum() 함수의 리턴 타입이 int 임을 표시하고 있으며, sum 함수 마지막에 return s 와 같이 정수 s의 값을 리턴하고 있다.
*/

func main3() {
	total := sum(1, 7, 3, 5, 9)
	println(total)
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

// 복수 개의 값을 리턴하기 위해서는 해당 리턴 타입들을 괄호() 안에 적어준다.
func main4() {
	count, total := sum2(1, 7, 3, 5, 9)
	println(count, total)
}

func sum2(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

/*
Go에서 Named Return Parameter들에 리턴값들을 할당하여 리턴할 수 있는데, 이는 리턴되는 값들이 여러 개일 때, 코드 가독성을 높이는 장점이 있다. 예를 들어,
위의 sum() 함수를 Named Return Parameter를 이용하여 고쳐쓰면 다음과 같다. 아래 예제에서 func 라인의 마지막 리턴타입 정의 부분을 보면, (int, int) 가 아니라 (count int, total int) 처럼 정의되어 있음을 볼 수 있다.
즉, 리턴 파라미터명과 그 타입을 함께 정의한 것이다. 그리고 함수 내에서는 이 count, total에 결과값을 직접 할당하고 있음을 볼 수 있다. 또한 마지막에 return 문이 있는 것을 볼 수 있는데,
실제 return 문에는 아무 값들을 리턴하지 않지만, 그래도 리턴되는 값이 있을 경우에는 빈 return 문을 반드시 써 주어야 한다 (이를 생략하면 에러 발생).
*/

func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
