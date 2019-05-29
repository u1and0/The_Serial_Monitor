package main

import (
	"fmt"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyACM0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open port
	port, err := serial.Open(options)
	time.Sleep(1 * time.Second) // 古いArduinoは接続に時間かかるので、少し待つのが定石
	check(err)
	fmt.Printf("Port opened %v\n", time.Now()) // シリアル接続時刻

	// Close port
	// 遅延実行、必ずポートを閉じる
	defer port.Close()
	defer fmt.Printf("Port closed %v\n", time.Now()) // シリアル切断時刻

	buf := make([]byte, 100000) // 読み取る文字列の数だけバッファを多く取っておく
	_, err = port.Read(buf)
	// b or _
	check(err)

	// fmt.Printf("buffer %v\n", b)
	// fmt.Println("Read line: ")
	fmt.Println(string(buf))
}
