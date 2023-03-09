/*
database/sql 패키지에서 가장 중요한 Type은 sql.DB 인데, 일반적으로 sql.Open() 함수를 사용하여 sql.DB 객체를 얻는다. 즉, sql.Open(드라이버, Connection) 함수에서 어떤 DB 드라이버를 사용할 것인지 그리고 해당 DB의 연결 정보를 제공하면, 결과로 sql.DB 객체를 얻게 된다. 일단 이 sql.DB 객체를 얻은 후, sql.DB의 여러 메서드들을 사용하여 쿼리를 하고, SQL문을 실행한다.
예를 들어, 자주 사용되는 sql.DB 메서드로 쿼리를 하는 Query(), QueryRow(), 그리고 INSERT, UPDATE, DELETE등을 실행하는 Exec()를 들 수 있다.
*/

// sql.DB 객체 db 생성
db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")

// db 차후에 닫기
defer db.Close()

// SELECT 쿼리
rows, err := db.Query("SELECT id, name FROM test")

// INSERT 실행
db.Exec("INSERT INTO test(id, name) VALUES (1, 'Alex')")