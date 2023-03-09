package struct_test

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	if sn1 == sn2 {
		fmt.Println("true")
	}
}

func TestXxxx(t *testing.T) {

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

func TestTTTTT(t *testing.T) {
	MTestF = make(map[int]ITestF)
	for i := 1; i < 110; i++ {
		MTestF[i] = &TestF{}
	}
	InitSTestF()
	iiTS := time.Now()
	for ii := 0; ii < 10000000; ii++ {
		MTestF[65].FF()
	}
	iiTE := time.Now()
	fmt.Println(iiTE.Sub(iiTS))
	iiiTS := time.Now()
	for iii := 0; iii < 10000000; iii++ {
		STestF.ITestF65.FF()
	}
	iiiTE := time.Now()
	fmt.Println(iiiTE.Sub(iiiTS))
}

func InitSTestF() {
	STestF.ITestF1 = &TestF{}
	STestF.ITestF2 = &TestF{}
	STestF.ITestF3 = &TestF{}
	STestF.ITestF4 = &TestF{}
	STestF.ITestF5 = &TestF{}
	STestF.ITestF6 = &TestF{}
	STestF.ITestF7 = &TestF{}
	STestF.ITestF8 = &TestF{}
	STestF.ITestF9 = &TestF{}
	STestF.ITestF10 = &TestF{}
	STestF.ITestF11 = &TestF{}
	STestF.ITestF12 = &TestF{}
	STestF.ITestF13 = &TestF{}
	STestF.ITestF14 = &TestF{}
	STestF.ITestF15 = &TestF{}
	STestF.ITestF16 = &TestF{}
	STestF.ITestF17 = &TestF{}
	STestF.ITestF18 = &TestF{}
	STestF.ITestF19 = &TestF{}
	STestF.ITestF20 = &TestF{}
	STestF.ITestF21 = &TestF{}
	STestF.ITestF22 = &TestF{}
	STestF.ITestF23 = &TestF{}
	STestF.ITestF24 = &TestF{}
	STestF.ITestF25 = &TestF{}
	STestF.ITestF26 = &TestF{}
	STestF.ITestF27 = &TestF{}
	STestF.ITestF28 = &TestF{}
	STestF.ITestF29 = &TestF{}
	STestF.ITestF30 = &TestF{}
	STestF.ITestF41 = &TestF{}
	STestF.ITestF42 = &TestF{}
	STestF.ITestF43 = &TestF{}
	STestF.ITestF44 = &TestF{}
	STestF.ITestF45 = &TestF{}
	STestF.ITestF46 = &TestF{}
	STestF.ITestF47 = &TestF{}
	STestF.ITestF48 = &TestF{}
	STestF.ITestF49 = &TestF{}
	STestF.ITestF50 = &TestF{}
	STestF.ITestF51 = &TestF{}
	STestF.ITestF52 = &TestF{}
	STestF.ITestF53 = &TestF{}
	STestF.ITestF54 = &TestF{}
	STestF.ITestF55 = &TestF{}
	STestF.ITestF56 = &TestF{}
	STestF.ITestF57 = &TestF{}
	STestF.ITestF58 = &TestF{}
	STestF.ITestF59 = &TestF{}
	STestF.ITestF60 = &TestF{}
	STestF.ITestF61 = &TestF{}
	STestF.ITestF62 = &TestF{}
	STestF.ITestF63 = &TestF{}
	STestF.ITestF64 = &TestF{}
	STestF.ITestF65 = &TestF{}
	STestF.ITestF66 = &TestF{}
	STestF.ITestF67 = &TestF{}
	STestF.ITestF68 = &TestF{}
	STestF.ITestF69 = &TestF{}
	STestF.ITestF70 = &TestF{}
	STestF.ITestF71 = &TestF{}
	STestF.ITestF72 = &TestF{}
	STestF.ITestF73 = &TestF{}
	STestF.ITestF74 = &TestF{}
	STestF.ITestF75 = &TestF{}
	STestF.ITestF76 = &TestF{}
	STestF.ITestF77 = &TestF{}
	STestF.ITestF78 = &TestF{}
	STestF.ITestF79 = &TestF{}
	STestF.ITestF80 = &TestF{}
	STestF.ITestF81 = &TestF{}
	STestF.ITestF82 = &TestF{}
	STestF.ITestF83 = &TestF{}
	STestF.ITestF84 = &TestF{}
	STestF.ITestF85 = &TestF{}
	STestF.ITestF86 = &TestF{}
	STestF.ITestF87 = &TestF{}
	STestF.ITestF88 = &TestF{}
	STestF.ITestF89 = &TestF{}
	STestF.ITestF90 = &TestF{}
	STestF.ITestF91 = &TestF{}
	STestF.ITestF92 = &TestF{}
	STestF.ITestF93 = &TestF{}
	STestF.ITestF94 = &TestF{}
	STestF.ITestF95 = &TestF{}
	STestF.ITestF96 = &TestF{}
	STestF.ITestF97 = &TestF{}
	STestF.ITestF99 = &TestF{}
	STestF.ITestF98 = &TestF{}
	STestF.ITestF100 = &TestF{}
	STestF.ITestF101 = &TestF{}
	STestF.ITestF102 = &TestF{}
	STestF.ITestF103 = &TestF{}
	STestF.ITestF104 = &TestF{}
	STestF.ITestF105 = &TestF{}
	STestF.ITestF106 = &TestF{}
	STestF.ITestF107 = &TestF{}
	STestF.ITestF108 = &TestF{}
	STestF.ITestF109 = &TestF{}
}

type TestF struct {
}

func (this *TestF) FF() {

}

type ITestF interface {
	FF()
}

var MTestF map[int]ITestF
var STestF struct {
	ITestF1   ITestF
	ITestF2   ITestF
	ITestF3   ITestF
	ITestF4   ITestF
	ITestF5   ITestF
	ITestF6   ITestF
	ITestF7   ITestF
	ITestF8   ITestF
	ITestF9   ITestF
	ITestF10  ITestF
	ITestF11  ITestF
	ITestF12  ITestF
	ITestF13  ITestF
	ITestF14  ITestF
	ITestF15  ITestF
	ITestF16  ITestF
	ITestF17  ITestF
	ITestF18  ITestF
	ITestF19  ITestF
	ITestF20  ITestF
	ITestF21  ITestF
	ITestF22  ITestF
	ITestF23  ITestF
	ITestF24  ITestF
	ITestF25  ITestF
	ITestF26  ITestF
	ITestF27  ITestF
	ITestF28  ITestF
	ITestF29  ITestF
	ITestF30  ITestF
	ITestF41  ITestF
	ITestF42  ITestF
	ITestF43  ITestF
	ITestF44  ITestF
	ITestF45  ITestF
	ITestF46  ITestF
	ITestF47  ITestF
	ITestF48  ITestF
	ITestF49  ITestF
	ITestF50  ITestF
	ITestF51  ITestF
	ITestF52  ITestF
	ITestF53  ITestF
	ITestF54  ITestF
	ITestF55  ITestF
	ITestF56  ITestF
	ITestF57  ITestF
	ITestF58  ITestF
	ITestF59  ITestF
	ITestF60  ITestF
	ITestF61  ITestF
	ITestF62  ITestF
	ITestF63  ITestF
	ITestF64  ITestF
	ITestF65  ITestF
	ITestF66  ITestF
	ITestF67  ITestF
	ITestF68  ITestF
	ITestF69  ITestF
	ITestF70  ITestF
	ITestF71  ITestF
	ITestF72  ITestF
	ITestF73  ITestF
	ITestF74  ITestF
	ITestF75  ITestF
	ITestF76  ITestF
	ITestF77  ITestF
	ITestF78  ITestF
	ITestF79  ITestF
	ITestF80  ITestF
	ITestF81  ITestF
	ITestF82  ITestF
	ITestF83  ITestF
	ITestF84  ITestF
	ITestF85  ITestF
	ITestF86  ITestF
	ITestF87  ITestF
	ITestF88  ITestF
	ITestF89  ITestF
	ITestF90  ITestF
	ITestF91  ITestF
	ITestF92  ITestF
	ITestF93  ITestF
	ITestF94  ITestF
	ITestF95  ITestF
	ITestF96  ITestF
	ITestF97  ITestF
	ITestF99  ITestF
	ITestF98  ITestF
	ITestF100 ITestF
	ITestF101 ITestF
	ITestF102 ITestF
	ITestF103 ITestF
	ITestF104 ITestF
	ITestF105 ITestF
	ITestF106 ITestF
	ITestF107 ITestF
	ITestF108 ITestF
	ITestF109 ITestF
}
