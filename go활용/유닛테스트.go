/*
Go는 간편하게 사용할 수 있는 테스트 프레임워크를 내장하고 있는데, "go test" 명령을 실행하여 테스트 코드들을 실행할 수 있다. "go test"는 현재 폴더에 있는 *_test.go 파일들을 테스트 코드로 인식하고, 이들을 일괄적으로 실행한다.

테스트 파일은 "testing" 이라는 표준패키지를 사용하는데, 먼저 testing 패키지를 import 하고, 테스트 메서드를 작성한다. 테스트 메서드는 TestXxx와 같은 특별한 메서드명을 갖는데, 앞의 Test는 해당 메서드가 테스트 메서드임을 알리는 것이고 Xxx는 임의의 메서드명으로 처음 글자는 항상 대문자이어야 한다. 메서드의 ProtoType은 아래와 같이 testing.T 포인터를 하나 입력으로 받으며 출력은 없다. 테스트 에러를 표시하기 위해 testing.T 의 Error(), Fail() 등의 메서드들을 사용한다.
*/


package calc
 
// 이코드와 테스트 코드는 동일 경로가 아니다. 이 코드를 테스트 한다고 가정해보자
func Sum(a ...int) int {
    sum := 0
    for _, i := range a {
        sum += i
    }
    return sum
}

// 같은 스크립트에 있지만 다른 경로라고 가정한다.
package calc_test
 
import (
    "calc"
    "testing"
)
 
func TestSum(t *testing.T) {
    s := calc.Sum(1, 2, 3)
 
    if s != 6 {
        t.Error("Wrong result")
    }
}