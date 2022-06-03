package observer

import "testing"

func TestXxx(t *testing.T) {
	ExampleObserver()
}
func ExampleObserver() {
	subject := NewSubject()
	reader := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	reader4 := NewReader("reader4")
	subject.Attach(reader)
	subject.Attach(reader2)
	subject.Attach(reader3)
	subject.Attach(reader4)

	subject.UpdateContext("observer mode")
}
