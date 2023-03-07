/*
함수명을 갖지 않는 함수를 익명함수(Anonymous Function)이라 부른다. 일반적으로 익명함수는 그 함수 전체를 변수에 할당하거나 다른 함수의 파라미터에 직접 정의되어 사용되곤 한다. 아래 예제는 main() 함수 안에서 익명함수가 선언되어 sum이라는 변수에 할당되고 있음을 보여준다. sum에 할당된 익명함수를 보면 func 바로 뒤에 함수명이 생략되었음을 알 수 있다. 이점을 제외하고 함수의 정의 내용을 동일한다. 일단 익명함수가 변수에 할당된 이후에는 변수명이 함수명과 같이 취급되며 "변수명(파라미터들)" 형식으로 함수를 호출할 수 있다.
*/

package main

func main() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s

	}

	result := sum(1, 2, 3, 4, 5)
	println(result)
}

/*
Go 프로그래밍 언어에서 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급되며, 따라서 다른 함수의 파라미터로 전달하거나 다른 함수의 리턴값으로도 사용될 수 있다. 즉, 함수의 입력 파라미터나 리턴 파라미터로서 함수 자체가 사용될 수 있다.
함수를 다른 함수의 파라미터로 전달하기 위해서는 익명함수를 변수에 할당한 후 이 변수를 전달하는 방법과 직접 다른 함수 호출 파라미터에 함수를 적는 방법이 있다.
*/

func main2() {
	//변수 add 에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}

	// add 함수 전달
	r1 := calc(add, 10, 20)
	println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의함
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

/*
type 문은 구조체(struct), 인터페이스 등 Custom Type(혹은 User Defined Type)을 정의하기 위해 사용된다. type 문은 또한 함수 원형을 정의하는데 사용될 수 있다.
즉, 위 예제에서 func(x int, y int) int 함수 원형이 코드 상에 계속 반복됨을 볼 수 있는데, 이 경우 type 문을 정의함으로써 해당 함수의 원형을 간단히 표현할 수 있다.
*/

type calculator func(int, int) int

// calculator 원형 사용
func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

// 이렇게 함수의 원형을 정의하고 함수를 타 메서드에 전달하고 리턴받는 기능을 타 언어에서 흔히 델리게이트(Delegate)라 부른다. Go는 이러한 Delegate 기능을 제공하고 있다.
