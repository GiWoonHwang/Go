/*
MySql 드라이버가 설치된 후, 아래와 같이 database/sql과 MySql 드라이버를 import하는데,
MySql 드라이버 패키지는 _ 로 alias를 주어 개발자가 드라이버 패키지를 직접 사용하지 않게 한다.
이 경우 드라이버 패키지는 database/sql 패키지가 내부적으로 사용하게 되며, 개발자는 database/sql를 통해서 모든 SQL 프로세싱을 진행하게 된다.
*/

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//...(db 사용)....
}

/*
패키지 database/sql에서 가장 먼저 사용하는 것은 sql.Open()으로 이는 sql.DB 객체를 리턴한다. sql.Open()의 첫번째 파라미터는 드라이버명으로 여기서는 mysql로 적으며, 두번째 파라미터는 각 드라이버의 Connection String으로서 MySql Connection을 적으면 된다.
위의 Connection은 로컬서버에 TCP 3306 포트로 root와 그 암호를 사용하여 접속하며, testdb라는 Database에 접속할 것을 나타낸다. 여기서 한가지 주목할 것은 sql.Open()은 실제 DB Connection을 Open하지 않는다는 점이다.
즉, sql.DB는 드라이버종류와 Connection 정보를 가지고는 있지만, 실제 DB를 연결하지 않으며, 많은 경우 Connection 정보조차 체크하지도 않는다. 실제 DB Connection은 Query 등과 같이 실제 DB 연결이 필요한 싯점에 이루어지게 된다.

MySql에서 쿼리를 위해 2종류의 메서드를 사용한다. 즉, 하나의 Row만을 리턴할 경우 QueryRow() 메서드를, 복수개의 Row를 리턴할 경우 Query() 메서드를 사용한다. 하나의 Row에서 실제 데이타를 읽어 로컬 변수에 할당하기 위해 Scan() 메서드를 사용하며,
복수 Row에서 다음 Row로 이동하기 위해 Next() 메서드를 사용한다. 아래 예제에서 QueryRow()는 SQL 문을 실행해서 리턴된 하나의 ROW의 데이타를 Scan() 안의 파라미터(name)에 넣고 있음을 볼 수 있다.
*/

func main2() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name) // 하나의 row에서 실제 데이터를 읽어 로컬 변수에 할당하기 위해 scan을 사용한다.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func main3() {
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 복수 Row를 갖는 SQL 쿼리
	var id int
	var name string
	rows, err := db.Query("SELECT id, name FROM test1 where id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() { // 복수의 쿼리를 읽기 위해 row를 사용한다
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
}

/*
여기서 한가지 주목할 것은 SQL 쿼리에서 ? (Placeholder)를 사용하여 Parameterized Query를 사용하고 있다는 점이다. 이는 SQL Injection과 같은 문제를 방지하기 위해 파라미터를 문자열 결합이 아닌 별도의 파라미터로 대입시키는 방식이다.
위의 예제에서 Placeholder ? 에는 1이 대입된다. Placeholder는 데이타베이스의 종류에 따라 다르게 사용하는데, 예를 들어 MySql은 ? 를 사용하고, Oracle은 :val1, :val2 등을 사용하고, PostgreSQL은 $1, $2 등을 사용한다.
*/
