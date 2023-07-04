/*
Go에서 CSV(Comma Separated Values) 파일을 처리하기 위해 "encoding/csv" 패키지를 사용한다. csv 패키지는 CSV 파일을 읽거나 쓰는 기능을 제공한다.

CSV 파일을 읽기 위해선 먼저 csv.NewReader() 를 실행하여 csv 파일 포맷을 읽을 수 있는 Reader를 생성하고, 이 리더로부터 Read() 혹은 ReadAll() 메서드를 호출하여 데이타를 읽어들인다. Read()는 한 라인을 읽어 들이고, ReadAll()은 전체를 한꺼번에 읽어들인다.

아래 예제는 test.csv 파일을 ReadAll()을 사용하여 모두 읽어 들여 [][]string 슬라이스에 넣은 후, for 루프를 사용하여 각 Cell의 내용을 출력하는 코드이다.
*/

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 파일 오픈
	file, _ := os.Open("./test.csv")

	// csv reader 생성
	rdr := csv.NewReader(bufio.NewReader(file)) // csv 파일 포맷을 읽을수 있게 한다.

	// csv 내용 모두 읽기
	rows, _ := rdr.ReadAll() // Read 혹은 ReadAll 메서드를 호출하여 데이타를 읽어듼다. (한 라인/ 전체 라인)

	// 행,열 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}
}

/*
CSV 파일을 쓰기 위해서는 csv.NewWriter() 를 실행하여 CSV 포맷으로 쓸 수 있는 Writer를 생성하고, 이 Writer로부터 Write() 혹은 WriteAll() 메서드를 호출하여 데이타를 쓰게된다.
Write()는 하나의 레코드를 쓰고, WriteAll()은 전체 레코드를 한꺼번에 쓰고 Flush를 호출한다.
*/

func main2() {
	// 파일 생성
	file, err := os.Create("./output.csv")
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"A", "0.25"})
	wr.Write([]string{"B", "55.70"})
	wr.Flush()
}
