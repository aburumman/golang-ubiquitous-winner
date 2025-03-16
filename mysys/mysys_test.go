package mysys 

import (
	"testing"
)

func TestGetHelloWorld(t *testing.T) {
	retVal := GetHelloWorld()

	if retVal != "Hello World" {
		t.Fail()
	}

}

func TestGetuid(t *testing.T) {
	retVal = GetUserI
}