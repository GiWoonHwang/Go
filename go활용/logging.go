/*
Go에서 로그를 사용하기 위하여 Go의 표준 패키지 "log"를 이용할 수 있다. 가장 단순한 예제로 아래 예와 같이 log 패키지의 Print*() 같은 출력 함수들을 직접 사용하면, 디폴트 출력 Stdout에 날짜/시간(디폴트 표준플래그)와 함께 해당 로그 내용을 화면에 출력된다.
만약 굳이 날짜/시간을 없애고자 한다면, log.SetFlags(0) 을 지정한다 (플래그는 아래 참조).
*/

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0) // 날짜/시간을 없애준다.
	log.Println("Logging")
}

/*
log 패키지는 여러가지 로거들을 지원하기 위해 "Logger" 라는 타입(인터페이스)를 제공하고 있다. 사실 위의 예제는 개발자가 별도의 Logger를 생성하지 않았기 때문에 표준 Logger를 자동으로 사용한 것이다.

새로운 로거(Logger)를 만들기 위해 log.New() 함수를 사용하게 된다. log.New()는 3개의 파라미터를 받아들이는데, 첫번째는 io.Writer 인터페이스를 지원하는 타입으로 표준콘솔출력(os.Stdout), 표준에러(os.Stderr), 파일포인터 혹은 io.Writer를 지원하는 모든 타겟이 사용될 수 있다. 두번째 파라미터는 로그출력의 가장 처음에 적는 Prefix로서 프로그램명,카테고리 등을 기재할 수 있다. 세번째 파라미터는 로그플래그로 표준플래그(log.LstdFlags), 날짜플래그(log.Ldate), 시간플래그(log.Ltime), 파일위치플래그(log.Lshortfile, log.Llongfile) 등을 | (OR 연산자)로 묶어 지정할 수 있다.

아래 예제는 표준콘솔출력(Stdout)으로 로그를 보내는 myLogger를 만들어 로깅을 하는 코드이다. log.New()에 정의된 대로 myLogger는 "INFO:" 라는 Prefix와 날짜/시간(표준플래그)을 먼저 왼쪽에 고정해서 출력하고, Print 문의 내용을 다음에 출력한다

*/

var myLogger *log.Logger

func main2() {
	myLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags) // io.write 인터페이스를 지원하는 타입, 로그출력에 가장 처음 찍는 prefix, 표준플래그,날짜플래그,시간플래그 등 연산자를 통해 묶어서 지정

	//....
	run()

	myLogger.Println("End of Program")
}

func run() {
	myLogger.Print("Test")
}

/*
보통 많은 경우 로그파일에 로그를 출력하기 원하는데, 파일에 로깅하기 위해서는 먼저 로그파일을 오픈하고 파일포인터를 log.New()의 첫번째 파라미터에 넣어 주면 된다 (주: os.File은 io.Writer를 구현하고 있다). 로그파일을 오픈할 때는 일반적으로 Write Only, Append Only 모드로 오픈한다.

아래 예제는 logfile.txt라는 로그파일에 로그를 출력하는 코드이다. 파일은 쓰기 및 추가모드로 오픈하고 log.New()의 첫 파라미터에 파일포인터를 지정하고 있다. 또한 로그플래그는 날짜/시간(log.Ldate|log.Ltime) 그리고 짧은 파일명/라인수(Lshortfile)를 함께 출력하도록 지정하였다. 아래 예제는 하나의 로그파일만을 사용하고 있지만, 여러 개의 Logger들을 만들어 그때 그때 다른 Logger를 사용할 수도 있다.
*/

var myLogger *log.Logger

func main3() {
	// 로그파일 오픈
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile) // 첫 파라미터에 파일 포인터를 저장하고 있다는데.... 왜 난 안보일까

	//....
	run()

	myLogger.Println("End of Program")
}

func run() {
	myLogger.Print("Test")
}

func main4() {
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 표준로거를 파일로그로 변경 -> 출력위치가 콘솔에서 파일로 변경된다.
	log.SetOutput(fpLog)

	run()
	log.Println("End of Program")
}

func run() {
	// 로그 메서드를 쓰면 파일에 출력됨
	log.Print("Test")
}

/*
복수 개의 로그 타겟에 동시에 로그를 출력하기 위해서 io.MultiWriter를 사용할 수 있다. io.MultiWriter()에 여러 io.Writer 들을 지정하여 복수 Writer를 만들고 이를 log.New() 혹은 log.SetOutput()에 지정하면, 복수 타겟에 로그를 쓰게 된다.
아래 예제는 파일과 콘솔에 동시에 로그를 출력하는 코드이다.
*/
func main5() {
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 파일과 화면에 같이 출력하기 위해 MultiWriter 생성
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)

	run()
	log.Println("End of Program")
}

func run() {
	log.Print("Test")
}
