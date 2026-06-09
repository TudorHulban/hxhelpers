package hxhelpers

import (
	"fmt"
	"testing"
)

type person struct {
	Age  *uint
	Name string
}

func (p *person) info() string {
	return fmt.Sprintf(
		"%s:%d",
		p.Name,
		p.Age,
	)
}

func TestErrorsTernary(t *testing.T) {
	ageAlice := uint(23)

	alice := person{
		Name: "Alice",
		Age:  &ageAlice,
	}

	john := person{
		Name: "John",
	}

	fmt.Println(
		Ternary(
			true,

			alice.info(),
			john.info(),
		),
	)

	fmt.Println(
		Ternary(
			false,

			alice.info(),
			john.info(),
		),
	)
}
