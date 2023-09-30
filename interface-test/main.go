package main

type class interface {
	display()
	read()
	delete()
}

type struct1 struct {
	name       string
	age        string
	department string
}

type struct2 struct {
	name string
	age  string
}

func (s2 struct2) display() {

}

func (s2 struct2) read() {

}

func (s2 struct2) delete() {

}

func (s1 struct1) display() {

}

func (s1 struct1) read() {

}

func (s1 struct1) delete() {

}

func main() {
	class1 := make([]class, 2)

	class1[0] = struct1{}
	class1[1] = struct2{}

	class1[0].display()
	class1[1].read()

}
