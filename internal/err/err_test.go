package err

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestErr(t *testing.T) {
	err := Foo("a")
	switch {
	case errors.Is(err, ErrInvalidParameters):
		t.Error(err) // will match if raised
	case errors.Is(err, ErrPermissionDenied):
		t.Error("2") // will match if raised
	}
	t.Error("done")
}
func Foo(s string) error {
	if s != "bar" && s != "baz" {
		err := errors.New("s not in [bar baz]")
		return ErrInvalidParameters.wrap(err)
	}
	return ErrPermissionDenied
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(50)

	for i := 0; i < 50; i++ {
		go func(j int, f fn) {
			if !checkEven(j) {
				fmt.Println(j) // Good. Read local copy of the loop counter.
				fmt.Println(f(j))
			}
			wg.Done()
		}(i, returnx100)

	}

	wg.Wait()
	t.Error("done")
}
func returnx100(x int) int {
	return x * 100
}

func checkEven(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false

}

type fn func(int) int
