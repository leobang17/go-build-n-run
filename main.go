package main

func main() {
	// 1. 현재 프로그램 runtime의 directory

	// 2. go-build-n-run을 실행할 수 있는 standalone go 환경인지 check (main.go file이 있는가)

	// 3. go build의 build output 이름 구하기  (우선순위)
	// 	- a) go.mod file이 있다면, module 이름 
	//	- b) go.mod가 없다면 현재 directory의 마지막 이름

	// 4. script command 실행 (go build && ./{goProgName})
}