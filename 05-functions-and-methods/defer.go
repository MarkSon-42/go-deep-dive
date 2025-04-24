package main

import (
	"fmt"
	// "os"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	defer fmt.Println("third defer")
	fmt.Println("end")
}

// func readFile(filename string) error {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close() // 함수가 끝날 때 파일을 닫음 -> 함수가 어떤 이유로 종료되든, 반드시 파일을 닫아줌.
// 	return nil
// }
