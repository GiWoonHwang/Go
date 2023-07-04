/*
복수 개의 SQL 문을 하나의 트랜잭션으로 묶기 위하여 sql.DB의 Begin() 메서드를 사용한다. 트랜잭션은 복수 개의 SQL 문을 실행하다 중간에 어떤 한 SQL문에서라도 에러가 발생하면 전체 SQL문을 취소하게 되고 (이를 롤백이라 한다), 모두 성공적으로 실행되어야 전체를 커밋하게 된다.

Begin() 메서드는 sql.Tx 객체를 리턴하는데, 이 Tx 객체로부터 Tx.Exec() 등을 실행하여 트랜잭션을 수행한 후, 마지막에 최종 Commit을 위해 Tx.Commit() 메서드를 호출한다.

트랜잭션을 취소하는 롤백을 위해서는 Tx.Rollback() 메서드를 호출하는데, 통상 Tx 객체를 얻은 직후 defer tx.Rollback() 을 호출하여 이후 문장들에서 에러가 발생하면 롤백하도록 defer로 지정해 준다.
*/

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // defer

	// 트랜잭션 시작
	tx, err := db.Begin() // 복수 개의 sql문을 실행하다 중간에 하나라도 실패하게 되면 전체를 취소하게 되고 전체가 다 성공해야 커밋하게 된다
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback() //중간에 에러시 롤백하도록 defer 한다

	// INSERT 문 실행
	_, err = tx.Exec("INSERT INTO test1 VALUES (?, ?)", 15, "Jack")
	if err != nil {
		//에러메시지를 출력하고 panic() 호출.
		//panic()은 defer를 실행한다.
		log.Panic(err)
	}

	_, err = tx.Exec("INSERT INTO test2 VALUES (?, ?)", 15, "Data")
	if err != nil {
		log.Panic(err)
	}

	// 트랜잭션 커밋
	err = tx.Commit() // 최종 커밋
	if err != nil {
		log.Panic(err)
	}
}
