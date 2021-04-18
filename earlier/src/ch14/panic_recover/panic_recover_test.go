package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("defer")
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")
	//os.Exit(-1)
	panic(errors.New("something wrong"))
}
