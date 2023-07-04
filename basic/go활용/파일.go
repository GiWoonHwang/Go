package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// 입력파일 열기
	fi, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close() // 무조건 실행

	// 출력파일 생성
	fo, err := os.Create("C:\\temp\\2.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024) // 이게 뭘까 ?

	// 루프
	for {
		// 읽기
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}

		// 끝이면 루프 종료
		if cnt == 0 {
			break
		}

		// 쓰기
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}

/*
Go 표준 패키지인 ioutil 패키지는 I/O 관련한 편리한 유틸러티를 제공하는 패키지이다. 입력 파일이 매우 크지 않은 경우, 이 패키지의 ReadFile, WriteFile 함수를 이용하면 편리하게 파일을 읽고 쓸 수 있다. 아래 예제는 ioutil을 사용하여 파일을 그대로 복사하는 코드이다.
*/

func main2() {
	//파일 읽기
	bytes, err := ioutil.ReadFile("C:\\temp\\1.txt")
	if err != nil {
		panic(err)
	}
	//파일 쓰기
	err = ioutil.WriteFile("C:\\temp\\2.txt", bytes, 0)
	if err != nil {
		panic(err)
	}
}
