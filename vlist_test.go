package vlist

import "testing"

func TestVList(t *testing.T) {
	var v VList
	t.Log("zero value for type.  empty vList:", v)
	logStructure(t, v)
	if v.String() != "[]" {
		t.Fatal("zero value fail")
	}

	for a := 6; a >= 1; a-- {
		v = v.Cons(VEle(a))
	}
	t.Log("demonstrate cons. 6 elements added:", v)
	logStructure(t, v)
	if v.String() != "[1 2 3 4 5 6]" {
		t.Fatal("Cons fail")
	}

	v = v.Cdr()
	t.Log("demonstrate cdr. 1 element removed:", v)
	logStructure(t, v)
	if v.String() != "[2 3 4 5 6]" {
		t.Fatal("Cdr fail")
	}

	t.Log("demonstrate car.  Car =", v.Car())
	if v.Car() != 2 {
		t.Fatal("Car fail")
	}

	t.Log("demonstrate length.  Len =", v.Len())
	if v.Len() != 5 {
		t.Fatal("Len fail")
	}

	t.Log("demonstrate element access. v[3] =", v.Index(3))
	if v.Index(3) != 5 {
		t.Fatal("Index fail")
	}

	v = v.Cdr().Cdr()
	t.Log("show cdr releasing segment. 2 elements removed:", v)
	logStructure(t, v)
	if v.String() != "[4 5 6]" {
		t.Fatal("Segment release fail")
	}
}

func logStructure(t *testing.T, v VList) {
	t.Log("offset:", v.offset)
	for sg := v.base; sg != nil; sg = sg.next {
		t.Log(" ", sg.ele)
	}
}
