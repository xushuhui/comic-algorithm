package array

import (
	"bytes"
	"fmt"
)

type array struct {
	size int
	data []int
}

func NewArray(capacity int) *array {
	return &array{
		size: 0,
		data: make([]int, capacity),
	}
}
func (a *array) Add(index int, element int) {
	if index < 0 || index > a.size {
		panic("index out of range")
	}
	for i := a.size-1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	//todo resize
	a.data[index] = element
	a.size++
}
func (a *array) AddFirst(element int) {
	a.Add(0,element)
}
func (a *array) AddLast(element int) {
	a.Add(a.GetSize()-1,element)
}
func (a *array) Remove(index int) {
	if index < 0 || index > a.size {
		panic("index out of range")
	}

}
func (a *array) RemoveFirst() {
	a.Remove(0)
}
func (a *array) RemoveLast() {
	a.Remove(0)
}
func (a *array) RemoveElement(element int) {

}
func (a *array) Get(index int) int {
	if index < 0 || index > a.size {
		panic("index out of range")
	}
	return a.data[index]
}
func (a *array) GetFirst() (element int) {
	return a.Get(0)
}
func (a *array) GetLast() (element int) {
	return a.Get(a.size-1)
}
func (a *array) Find(element int) (index int) {
	return index
}
func (a *array) Contains(element int) (is bool) {
	return is
}

func (a *array) Set(index int, element int) {
	if index < 0 || index > a.size {
		panic("index out of range")
	}
	return
}
func (a *array) GetSize() (size int) {
	return a.size
}
func (a *array) GetCapaticy() int {
	return len(a.data)
}
func (a *array) resize(newCapaticy int) {
	newArray := NewArray(newCapaticy)

	for i := 0; i < a.size; i++ {
		newArray.data[i] = a.data[i]
	}
	a.data = newArray.data
	newArray = nil
	return
}
func (a *array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 int 类型转换为字符串
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
