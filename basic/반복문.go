/*
go 프로그래밍 언어에서 반복문은 for 루프를 이용한다. go의 반복문은 for 하나 밖에 없다. 초기값, 조건식 ,증감식 등은 경우에 따라 생략할 수 있다. 다만 초기값; 조건식; 증감을 둘러싸는 괄호를 생략하는데, 괄호를 쓰면 에러가 난다.
*/
package main

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

// 초기값과 증감식을 생략하고 조건식만을 사용할 수 있는데, 다른 언어의 while 루프와 같이 쓰이도록 한다.

func main2() {
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)
}

// for 루프로 무한루프를 만드려면 초기값, 조건식, 증감 모두를 생략하면 된다.
func main3() {
	for {
		println("infinite loop")
	}
}

// for range 문은 컬렉션으로 부터 한 요소 씩 가져와 차례로 for 블럭의 문장들을 실행한다. 이는 다른 언어의 foreach와 비슷한 용법이다. for 인덱스,요소값:= range 컬렉션과 같이 루프를 구성한다

func main4() {
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		println(index, name)
	}
}

// for문에서 즉시 빠져나올 필요가 있을 때 break, for 루프 중간에서 나머지 문장들을 실행하지 않고 for 루프 시작부분으로 가려면 continue 그리고 기타 임의으 문장으로 이동하기 위해 goto 문
func main5() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")
}
