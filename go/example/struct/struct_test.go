package struct_test

import (
	"fmt"
	"testing"
)

type Ta struct {
	age  int
	name string
}

func TestXxx(t *testing.T) {
	sn1 := Ta{age: 11, name: "qq"}
	sn2 := Ta{age: 11, name: "qq"}
	if sn1 == sn2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

type P struct {
	name string
}

func TestP(t *testing.T) {
	var p *P
	p = &P{name: "test"}
	fmt.Println(p.name)
	fmt.Println(&p)
	fmt.Println((*p).name)
}

type Myint1 int
type Myint2 = int

func (this Myint1) Set(i int) {
}
func TestVar(t *testing.T) {
	var i int = 1
	var i1 Myint1 = Myint1(i)
	fmt.Println(i1)
	var i2 Myint2 = i
	fmt.Println(i2)
}

func TestPointer(t *testing.T) {
	person := &Person{28}
	defer fmt.Println(person.age)

	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)
	defer func() {
		fmt.Println(person.age)
	}()
	person.age = 29
}

type Person struct {
	age int
}
