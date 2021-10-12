package array

import (
	"testing"
)

func TestNewArray(t *testing.T) {
	s1 := []int{1,2,3,4,5}
s2 := s1
s2 = append(s2, 1)
t.Log(s1,s2)
}

func Test_array_Add(t *testing.T) {

	a := NewArray(5)

	a.AddFirst(1)
	t.Log(a)
	a.AddLast(2)
	t.Log(a)
	a.Add(2,2)
	t.Log(a)
}




func Test_array_Remove(t *testing.T) {
	
}


func Test_array_RemoveElement(t *testing.T) {
	
}

func Test_array_Get(t *testing.T) {
	
}



func Test_array_Find(t *testing.T) {

}

func Test_array_Contains(t *testing.T) {
	
}

func Test_array_Set(t *testing.T) {

}

func Test_array_GetSize(t *testing.T) {
	tests := []struct {
		name     string
		a        *array
		wantSize int
	}{
		// TODO: Add test cases.
		{"1",NewArray(2),0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSize := tt.a.GetSize(); gotSize != tt.wantSize {
				t.Errorf("array.GetSize() = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func Test_array_GetCapaticy(t *testing.T) {
	tests := []struct {
		name string
		a    *array
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.GetCapaticy(); got != tt.want {
				t.Errorf("array.GetCapaticy() = %v, want %v", got, tt.want)
			}
		})
	}
}




