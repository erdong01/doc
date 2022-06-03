package abstractfactory

import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

func TestA(t *testing.T) {
	ExampleRdbFactory()
}

func runFuncName() string {

	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func getMainAndDetail(factory DAOFactory) {
	fmt.Printf("Func Name:%s \n", runFuncName())

	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func ExampleRdbFactory() {
	log.Printf("Func Name:%s \n", runFuncName())
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func ExampleXmlFactory() {
	log.Printf("Func Name:%s \n", runFuncName())
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)

}
