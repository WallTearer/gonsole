package gonsole_test

import (
	"github.com/walltearer/gonsole"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	gon := gonsole.New()
	if reflect.TypeOf(gon).String() != "*gonsole.Gonsole" {
		t.Error("Failed to create new gonsole instance")
	}
}

func ExampleGonsole_ReadInt() {
	// using fake input source to simulate user's input
	fakeInput, _ := ioutil.TempFile("", "")
	defer fakeInput.Close()

	io.WriteString(fakeInput, "wrongone\n")
	io.WriteString(fakeInput, "wrong2\n")
	io.WriteString(fakeInput, "123\n")
	fakeInput.Seek(0, os.SEEK_SET)

	gon := gonsole.New()
	gon.SetSource(fakeInput)

	gon.ReadInt("Prompt msg: ")
	// Output:
	// Prompt msg: Converting Error: strconv.ParseInt: parsing "wrongone": invalid syntax
	// Prompt msg: Converting Error: strconv.ParseInt: parsing "wrong2": invalid syntax
	// Prompt msg:
}
