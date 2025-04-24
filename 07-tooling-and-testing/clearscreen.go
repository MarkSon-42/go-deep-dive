package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const col = 30
	bar := fmt.Sprintf("[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf("\r"+bar, strings.Repeat("=", i)+">") // \r로 커서를 맨 앞으로 이동
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r"+bar+" Done!\n", strings.Repeat("=", col)) // 최종 출력
}
