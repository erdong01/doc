package hashTable

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var hashTable HashTable
	hashTable.Insert(&Emp{Id: 1, Name: "员工1"})
	hashTable.Insert(&Emp{Id: 2, Name: "员工2"})
	hashTable.Insert(&Emp{Id: 3, Name: "员工3"})
	hashTable.Insert(&Emp{Id: 4, Name: "员工4"})
	hashTable.Insert(&Emp{Id: 14, Name: "员工14"})
	hashTable.Insert(&Emp{Id: 28, Name: "员工28"})
	hashTable.Insert(&Emp{Id: 30, Name: "员工30"})
	emp := hashTable.FundById(30)
	hashTable.ShowAll()
	fmt.Println(*emp)
	hashTable.FundById(44)
}

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head   *Emp
	HashId int
}

func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp
	if cur == nil {
		this.Head = emp
		return
	}
	for {
		if cur != nil {
			//比较
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next

		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}
func (this *EmpLink) ShowLink() {
	if this.Head == nil {
		fmt.Println(this.HashId, "链表是空的。")
		return
	}
	cur := this.Head
	for {
		fmt.Printf("hash表id:%d 雇员id:%d | ", this.HashId, cur.Id)
		cur = cur.Next
		if cur == nil {
			break
		}

	}
	fmt.Println()
}
func (this *EmpLink) FindById(id int) *Emp {
	if this.Head == nil {
		fmt.Println(this.HashId, "链表是空的。")
		return nil
	}
	cur := this.Head
	for {
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
		if cur == nil {
			break
		}

	}
	fmt.Println()
	return nil
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {

	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].HashId = i
		this.LinkArr[i].ShowLink()
	}
}
func (this *HashTable) FundById(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func (this *HashTable) HashFun(id int) int {
	return id % 7
}
