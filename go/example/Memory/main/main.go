package main

type smallStruct struct {
	a, b int64
	c, d int64
}

func main() {
	smallAllocation()
}

func smallAllocation() *smallStruct {
	return &smallStruct{}
}
