package main

import "testing"

type (
	User struct {
		Name string
		Age  int
	}
)


func (u User) UserInit() {
	u.Name = "name"
	u.Age  = 20
}

func main() {
	u := User{}
	u.UserInit()
	t := testing.T{}

	if u.Name != "name" {
		t.Error("unexpected Name")
	}
	if u.Age != 20 {
		t.Error("unexpected Age")
	}
}
