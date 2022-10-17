package main

import (
	"bytes"
	"io/ioutil"
)

func main() {
	var b bytes.Buffer

	b.WriteString("skjdhdfgsi\n")
	b.WriteString("skjdhdfsfsfgsi\n")
	b.WriteString("skjdhdfgsdfsfsdfssi\n")
	ioutil.WriteFile("logio.txt", b.Bytes(), 0777)

}
