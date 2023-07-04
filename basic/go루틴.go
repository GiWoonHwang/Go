/*
쓰레드
쓰레드(Thread)는 프로그램내에서 실행 흐름을 의미합니다. 프로그램은 일반적으로는 하나의 실행 흐름(쓰레드)을 가지지만, 경우에 따라 하나 이상의 쓰레드를 갖는 경우도 있습니다. 이를 멀티 쓰레드라고 합니다.

CPU는 단순한 계산기입니다. 따라서 주어진 값을 계산만 할 뿐, 이 값이 어디서 왔고, 어디로 가는지는 신경쓰지 않습니다. 멀티 쓰레드인 경우, OS에서 쓰레드를 관리하고, 쓰레드의 개수가 CPU보다 많은 경우, 쓰레드를 교체해가면서 CPU를 사용하도록 합니다. 이를 컨텍스트 스위칭(Context Switching)이라고 합니다.

컨텍스트 스위칭은 하나의 CPU가 여러 쓰레드를 다룰 때, 쓰레드를 전환시키며 CPU를 사용하도록 하는 것을 의미합니다. 이렇게 컨텍스트 스위칭이 발생하면 전환 비용이 발생하므로 성능이 저하되는 문제가 발생할 수 있습니다.

반대로, CPU의 개수가 쓰레드의 개수와 동일하다면, 컨텍스트 스위칭이 발생하지 않으므로 성능에 아무 문제가 발생하지 않습니다.

1. go 루틴
Go루틴(goroutine)은 Go 런타임이 관리하는 Lightweight 논리적 (혹은 가상적) 쓰레드(주1)이다. Go에서 "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행한다. goroutine은 비동기적으로(asynchronously) 함수루틴을 실행하므로, 여러 코드를 동시에(Concurrently) 실행하는데 사용된다.

(주1)goroutine은 OS 쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위하여 만든 것으로, 기본적으로 Go 런타임이 자체 관리한다. Go 런타임 상에서 관리되는 작업단위인 여러 goroutine들은 종종 하나의 OS 쓰레드 1개로도 실행되곤 한다. 즉, Go루틴들은 OS 쓰레드와 1 대 1로 대응되지 않고, Multiplexing으로 훨씬 적은 OS 쓰레드를 사용한다. 메모리 측면에서도 OS 쓰레드가 1 메가바이트의 스택을 갖는 반면, goroutine은 이보다 훨씬 작은 몇 킬로바이트의 스택을 갖는다(필요시 동적으로 증가). Go 런타임은 Go루틴을 관리하면서 Go 채널을 통해 Go루틴 간의 통신을 쉽게 할 수 있도록 하였다.

아래 예제에서 main 함수를 보면, 먼저 say()라는 함수를 동기적으로 호출하고, 다음으로 동일한 say() 함수를 비동기적으로 3번 호출하고 있다. 첫번째 동기적 호출은 say() 함수가 완전히 끝났을 때 다음 문장으로 이동하고, 다음 3개의 go say() 비동기 호출은 별도의 Go루틴들에서 동작하면서, 메인루틴은 계속 다음 문장(여기서는 time.Sleep)을 실행한다. 여기서 goroutine들은 그 실행순서가 일정하지 않으므로 프로그램 실행시 마다 다른 출력 결과를 나타낼 수 있다.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)
}

/*
익명함수 go루틴
*/
func main2() {
	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() //Go루틴 모두 끝날 때까지 대기
}

/*
여기서 sync.WaitGroup을 사용하고 있는데, 이는 기본적으로 여러 Go루틴들이 끝날 때까지 기다리는 역활을 한다. WaitGroup을 사용하기 위해서는 먼저 Add() 메소드에 몇 개의 Go루틴을 기다릴 것인지 지정하고, 각 Go루틴에서 Done() 메서드를 호출한다 (여기서는 defer 를 사용하였다).
그리고 메인루틴에서는 Wait() 메서드를 호출하여, Go루틴들이 모두 끝나기를 기다린다.
*/

/*
다중 cpu 처리
Go는 디폴트로 1개의 CPU를 사용한다. 즉, 여러 개의 Go 루틴을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리한다 (Concurrent 처리). 만약 머신이 복수개의 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬처리 (Parallel 처리)하게 할 수 있는데, 병렬처리를 위해서는 아래와 같이 runtime.GOMAXPROCS(CPU수) 함수를 호출하여야 한다
(여기서 CPU 수는 Logical CPU 수를 가리킨다).
*/

func main3() {
	// 4개의 CPU 사용
	runtime.GOMAXPROCS(4)

	//...
}
