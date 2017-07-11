package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := &bytes.Buffer{}

	i, err := buf.Write([]byte("afgsgsfgdfG"))
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	//buf := bytes.NewBuffer([]byte("afgsgsfgdfG"))

	fmt.Println(buf.ReadByte())
}
