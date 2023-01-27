package main

import "fmt"

/*
type NameOfStruct struct{
  ...fields
}
*/

type Vertex struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MyIntegerDataType int

func SomeFunc(x MyIntegerDataType) MyIntegerDataType {
	return 0
}
func main() {
	/*
		var myData MyIntegerDataType

		fmt.Println(myData)
		// declare the variable of struct type
		v := Vertex{X: 1, Y: 2} // x=1, y=2

		n := Vertex{1, 2} //x=1, y=2

		g := Vertex{} // x=0, y=0
	*/
	newVarWithValue := NewVertexWithValue(1, 2)
	newVar := NewVertex()
	newVar.SetX(3) // newVar.X=3
	fmt.Println(newVar, newVarWithValue)
	// x := newVar.X  // accessing the field
	// y := newVar.Y

}

// getter - gets the data
// setter - sets the data
// function with a reciever = method
// two functions never have the same name
// two methods can have the same name as long as the recievers are different
func (v Vertex) GetX() int {
	return v.X
}
func (v MyIntegerDataType) GetX() int {
	return 0
}

// setter - sets the value to the field

func (v *Vertex) SetX(i int) {
	v.X = i
}

func NewVertexWithValue(x, y int) Vertex {
	// default value 3
	if x == 0 {
		x = 3
	}
	if y == 0 {
		y = 3
	}
	return Vertex{x, y}
}

func NewVertex() Vertex {
	var x, y int
	// default value 3
	if x == 0 {
		x = 3
	}
	if y == 0 {
		y = 3
	}
	return Vertex{x, y}
}
