package simple_factory

import "fmt"

type Api interface {
	Say(name string) string
}

func NewApi(t int) Api {
	if t == 1 {
		return &hiApi{}
	} else if t == 2 {
		return &helloApi{}
	}
	return nil
}

type hiApi struct {
}

func (*hiApi) Say(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}

type helloApi struct{}

func (*helloApi) Say(name string) string {
	return fmt.Sprintf("Hello,%s", name)
}
