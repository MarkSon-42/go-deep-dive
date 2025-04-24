# 05: Functions and Methods

defer : 함수가 종료되기 직전에 특정 코드를 실행하도록 예약


~~~go
func readFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()  // 함수가 끝날 때 파일을 닫음 -> 함수가 어떤 이유로 종료되든, 반드시 파일을 닫아줌.
    return nil
}
~~~