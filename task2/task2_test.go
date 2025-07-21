package task2

import (
	"fmt"
	"testing"
)

func TestTask2(t *testing.T) {
	var p1 *uint
	i := uint(10)
	p1 = &i
	IntSpliceAdd(p1)
	fmt.Println(*p1)
}

func TestGoRoutineTask(t *testing.T) {
	GoRoutineTask()
}

func TestRectangle_Perimeter(t *testing.T) {

	r := Rectangle{
		Length: 10,
		Width:  20,
		Height: 30,
	}
	r.Area()
	fmt.Println(r.Perimeter())
}

func TestEmployee_PrintInfo(t *testing.T) {
	emp := &Employee{
		Person: &Person{
			Name: "BEN",
			Age:  18,
		},
		EmployeeID: 1,
	}
	fmt.Println(emp.PrintInfo())
}

func TestChanelBaseFunc(t *testing.T) {
	ChanelBaseFunc()
}

func TestGoRoutineTask2(t *testing.T) {
	ChannelBufferFunc()
}

func TestSyncMutexFunc(t *testing.T) {
	SyncMutexFunc()
}

func TestSyncMutexAtomicFunc(t *testing.T) {
	SyncMutexAtomicFunc()
}
