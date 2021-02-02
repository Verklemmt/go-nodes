package lib

import "testing"

func TestSearchAsync(t *testing.T) {

	network := new(Network)

	ln1 := new(Node)
	rn1 := new(Node)

	network.Add(ln1)
	network.Add(rn1)

	result := network.SearchAsync("Hallo")

	if result != nil {
		t.Error()
	}

	ln2 := new(Node)
	ln2.Data = "Hallo"

	network.Add(ln2)

	result = network.SearchAsync("Hallo")

	if result != ln2 {
		t.Error()
	}

}
