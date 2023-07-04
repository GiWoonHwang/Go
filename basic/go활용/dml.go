/*
데이타를 INSERT, UPDATE, DELETE (DML Operation)하기 위해서 sql.DB 객체의 Exec() 메서드를 사용한다. Query/QueryRow 메서드는 데이타를 리턴할 때 사용하는 반면, DML과 같이 리턴되는 데이타가 없는 경우는 Exec 메서드를 사용해야 한다.
*/

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// INSERT 문 실행
	result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "Jack") // 파라미터를 넣어줌
	if err != nil {
		log.Fatal(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected() // 갱신된 레코드 수
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
}

/*
Prepared Statement는 데이타베이스 서버에 Placeholder를 가진 SQL문을 미리 준비시키는 것으로, 차후 해당 Statement를 호출할 때 준비된 SQL문을 빠르게 실행하도록 하는 기법이다.
Go에서 Prepared Statement를 사용하기 위해서는 sql.DB의 Prepare() 메서드를 써서 Placeholder를 가진 SQL문을 미리 준비시키고, sql.Stmt 객체를 리턴받는다.
차후 이 sql.Stmt 객체의 Exec (혹은 Query/QueryRow) 메서드를 사용하여 준비된 SQL문을 실행한다.
*/

func main2() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE test1 SET name=? WHERE id=?") // dml 같이 여러 조건을 넣기 위해 사용하는건가 ??
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	_, err = stmt.Exec("Tom", 1) //Placeholder 파라미터 순서대로 전달
	checkError(err)
	_, err = stmt.Exec("Jack", 2)
	checkError(err)
	_, err = stmt.Exec("Shawn", 3)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
