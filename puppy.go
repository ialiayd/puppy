package puppy

import (
	"github.com/ialiayd/dog"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof Woof Woof"
}

func BigBark() string {
	return dog.WhenGrownUp("WWWooofff WWWooofff WWWooofff")
}
