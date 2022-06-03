package sort_test

import (
	"fmt"
	"reflect"
	"testing"
)

//插入排序
func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{args: args{
			head: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
				},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := InsertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort2(t *testing.T) {
	arr := [5]int{23, 0, 12, 56, 34}
	InsertSort(&arr)
	fmt.Println(arr)

}

func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		fmt.Println(insertIndex)
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}
