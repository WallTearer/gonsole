package gonsole

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Gonsole struct {
	reader *bufio.Reader
}

func New() *Gonsole {
	gon := &Gonsole{}

	// setting standard input source for reader
	gon.SetSource(os.Stdin)

	return gon
}

// Set custom input source (this separate method is mainly here for testing purposes)
func (gon *Gonsole) SetSource(input *os.File) {
	gon.reader = bufio.NewReader(input)
}

// Asks user to enter a valid integer value until he enters correct value.
func (gon *Gonsole) ReadInt(msg string) int {
	var readSuccess bool = false
	var inputInt int
	var errConv error

	for readSuccess == false {
		fmt.Printf("%s", msg)

		inputStr, err := gon.reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Reading Error: %v\n", err)
			continue
		}

		inputStr = strings.TrimSpace(inputStr)
		inputInt, errConv = strconv.Atoi(inputStr)
		if errConv != nil {
			fmt.Printf("Converting Error: %v\n", errConv)
			continue
		}

		readSuccess = true
	}

	return inputInt
}
