/*
if 문은 해당 조건이 맞으면 {} 블럭안의 내용을 실행한다. go의 if 조건문은 아래 예제에서 보듯이 조건식을 괄호()로 둘러 싸지 않아도 된다. 그리고 반드시 조건 블럭 시작 브레이스({)를 if문과 같은 라인에 두어야 한다.
주목해야 할 점은 if문의 조건식은 반드시 boolean 식으로 표현되어야 한다는 것이다. 이점은 c/cpp 같은 다른 언어들이 조건식이 1,0과 같은 숫자를 쓸 수 이쓴ㄴ 것과 대조적이다.
if k == 1 { // 같은 라인
	println("one")
}


*/

// if 문은 else if, 혹은 else문을 함께 가질 수 있다.
if k == 1{
	println("one")
}
else if k == 2 {
	println("two")
}
else{
	println("other")
}

// if 문에서 조건식을 사용하기 전에 간단한 문장을 함께 실행할 수 있다. 이때 정의된 val는 if문 블럭에서 선언된 것이기 때문에 scope를 벗어나면 에러를 출력한다.
if val := i * 2; val < max {
	println(val)
}

val ++ // scope 를 벗어나 에러 출력


// 여러 값을 비교해야 할 경우 혹은 다수의 조건식을 체크해야 하는 경우 switch문을 사용한다. 다른 언어들과 비슷하게 switch 문 뒤에 하나의 변수를 지정하고 case문에 해당 변수가 가질 수 있는 값들을 지정하여, 각 경우에 다른 문장 블럭들을 실행할 수 있다.

package main
 
func main() {
	// 다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만, Go는 다음 case로 가지 않는다
    var name string
    var category = 1
 
    switch category {
    case 1:
        name = "Paper Book"
    case 2:
        name = "eBook"
    case 3, 4:
        name = "Blog"
    default:
        name = "Other"
    }
    println(name)
     
    // Expression을 사용한 경우
    switch x := category << 2; x - 1 {
        //...
    }   
}

/*
C 혹은 C# 과 같은 언어에서 case 문은 case 블럭 마지막에 break 문을 명시하여 switch 문을 빠져나온다. 만약 break 문이 없으면, case 문 밑의 모든 문장들을 실행해 버린다. Go는 case문 마지막에 break 문을 적든 break 문을 생략하든,
항상 break 하여 switch 문을 빠져나온다. 이는 Go 컴파일러가 자동으로 break 문을 각 case문 블럭 마지막에 추가하기 때문이다. Go에서 만약 이러한 디폴트 break 문을 사용하지 않고, C나 C#처럼 계속 밑의 문장들(다음 case문 코드 블럭들)을 실행하게 하려면,
fallthrough 문을 명시해 주면 된다. fallthrough 문을 사용한 아래 예제를 실행하면, "2 이하/3 이하/default 도달"을 모두 출력하게 된다. 즉, 일단 case 2 에 도착한 후 fallthrough 가 있으므로, (val 값이 3이 아님에도) case 3의 코드 블럭을 계속 실행하고,
case 3에도 fallthrough 가 있으므로 default 블럭을 계속 실행한다.
*/

package main
 
import "fmt"
 
func main() {
    check(2)
}
 
func check(val int) {
    switch val {
    case 1:
        fmt.Println("1 이하")
        fallthrough
    case 2:
        fmt.Println("2 이하")
        fallthrough
    case 3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}