/*
이메일을 발송하기 위한 한 방법으로 SMTP 메일을 이용할 수 있다. Go에서 SMTP를 이용해 이메일을 발송하기 위해서는 내장 패키지인 net/smtp 패키지를 사용할 수 있다. smtp의 SendMail() 함수는 SMTP 메일을 보내는 간단한 기능을 제공하는데, 이 함수는 다음과 같은 5개의 파라미터를 갖는다.
SMTP 서버의 주소와 포트번호를 적는다
smtp.Auth 인터페이스를 갖는 타입인데, SMTP 서버에 대한 로그인 정보를 갖는다. smtp.PlainAuth() 함수는 PLAIN authentication 메카니즘을 구현한 stmp.Auth 인터페이스를 리턴하므로, 여기서는 이 함수를 사용하였다.
송신자 이메일 주소를 적는다
수신자 이메일 주소를 적는다. 수신자는 복수일 수 있으므로 string 슬라이스를 사용한다.
메시지를 적는다. 메시지는 RFC 822 이메일 스타일이어야 하는데, 즉, 메시지 헤더가 먼저 오고, 빈칸 한 줄이 그 다음으로, 그리고 마지막으로 메시지 Body를 적는다. 이메일 헤더에는 "From", "To", "Subject", "Cc" 같은 필드들이 올 수 있다. 메시지의 모든 라인은 CRLF로 끝나야 한다.
smtp.SendMail() 함수의 첫번째 파라미터에는 SMTP 서버명과 포트번호를 지정하는데, SMTP 서버는 핫메일 (Live)과 같은 Public 서버 혹은 사설 SMTP 서버를 사용할 수 있다. 아래 예제는 Live 메일서버를 사용한 예제로서, Live STMP 서버를 사용하기 위해선 서버명 smtp.live.com 과 포트 587 을 사용한다. 그리고 Live SMTP 서버에 엑세스하기 위해 자신의 Live 계정과 암호를 smtp.PlainAuth() 함수에 지정한다. 메시지의 작성은 위에서 설명하였듯이 헤더와 공란 그리고 Body로 구성하는데, 여기서는 헤더에 Subject 만을 추가하였다.
*/

package main

import (
	"net/smtp"
)

func main() {
	// 메일서버 로그인 정보 설정
	auth := smtp.PlainAuth("", "sender@live.com", "pwd", "smtp.live.com")

	from := "sender@live.com"
	to := []string{"receiver@live.com"} // 복수 수신자 가능

	// 메시지 작성
	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "메일 테스트입니다\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	// 메일 보내기
	err := smtp.SendMail("smtp.live.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
