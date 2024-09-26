package puppy

import (
	"fmt"

	"github.com/NorthParis/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func From10() {
	fmt.Println("Hello, from v1.0 version")
}
