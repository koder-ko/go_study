package go_study

import "fmt"

// fmt print basic
func usePrint(){

	/*
	기본적인 프린트 println print
	 */
	fmt.Println("Hello, world!") // Hello, world! + 다음라인으로 전환
	fmt.Print("Hello, world!") // Hello, world! +다음라인 전환 x


	/*
	c 언어의 프린트 방식 이용 printf
	 */
	fmt.Printf("%d is my number", 7) // 7 is my number + 다음라인 전환x
	fmt.Printf("go is the best = %b", true) // go is the best = true + 다음라인 전환 x
	// printf => print format 즉 뒤에 들어온 인자를 문자열로 포맷하여 넣어준다는 뜻

	// %d, %b 는 서식문자로 꽤나 종류가 다양합니다.

	/*
	서식문자	출력 형태
	%t	bool
	%b	2진수 정수
	%c	문자
	%d	10진수 정수
	%o	8진수 정수
	%x	16진수 정수, 소문자
	%X	16진수 정수, 대문자
	%f	10진수 방식의 고정 소수점 실수
	%F	10진수 방식의 고정 소수점 실수
	%e	지수 표현 실수, e
	%E

	지수 표현 실수, E
	%g	간단한 10진수 실수
	%G	간단한 10진수 실수
	%s	문자열
	%p	포인터
	%U	유니코드
	%T	타입
	%v	모든 형식
	%#v	#을 이용해구분할 수 있는 형식 표현
	 */
}

// fmt scan basic
func useScan(){
}