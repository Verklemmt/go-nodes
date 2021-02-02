package lib

import "testing"

func TestAdd(t *testing.T) {
	n1 := new(Node)
	n1l := new(Node)
	n1r := new(Node)

	n1.Add(n1l)
	n1.Add(n1r)

	if n1.Right != n1r {
		t.Error()
	}
	if n1.Left != n1l {
		t.Error()
	}
}

func TestLength(t *testing.T) {
	n1 := new(Node)
	n1l := new(Node)
	n1r := new(Node)
	n2l := new(Node)

	if n1.Length() != 0 {
		t.Error()
	}

	n1.Add(n1l)

	if n1.Length() != 1 {
		t.Error()
	}

	n1.Add(n1r)
	n1.Add(n2l)

	if n1.Length() != 2 {
		t.Error()
	}

}
