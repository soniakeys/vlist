// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// VList is an ordered list data structure.
//
// References:
// http://en.wikipedia.org/wiki/VList,
// http://infoscience.epfl.ch/record/64410/files/techlists.pdf "Fast Functional
// Lists, Hash-Lists, Deques and Variable Length Arrays" Phil Bagwell.
package vlist

import "fmt"

// A VList holds an ordered list of VEle's.
//
// The zero value for the type is an empty list.
type VList struct {
	base   *vSeg
	offset int
}

type vSeg struct {
	next *vSeg
	ele  []VEle
}

// Element type VEle could be anything.
// Recompile with different type definition here to specialize as desired.

// A VEle is an element of a VList.
type VEle interface{}

// Cons adds an element to the front of the VList.
func (v VList) Cons(a VEle) VList {
	if v.base == nil {
		return VList{base: &vSeg{ele: []VEle{a}}}
	}
	if v.offset == 0 {
		l2 := len(v.base.ele) * 2
		ele := make([]VEle, l2)
		ele[l2-1] = a
		return VList{&vSeg{v.base, ele}, l2 - 1}
	}
	v.offset--
	v.base.ele[v.offset] = a
	return v
}

// Index locates the i'th element.
//
// Consistent with Go slices it is 0 based and panics on index out of range.
func (v VList) Index(i int) (r VEle) {
	if i >= 0 {
		i += v.offset
		for sg := v.base; sg != nil; sg = sg.next {
			if i < len(sg.ele) {
				return sg.ele[i]
			}
			i -= len(sg.ele)
		}
	}
	panic("Index out of range")
}

// Car returns the first element of a VList.
//
// Consistent with the Index method, it panics on an empty list.
// (Note this is not the usual convention of the Lisp language.)
func (v VList) Car() (r VEle) {
	if v.base == nil {
		panic("Car on empty vList")
	}
	return v.base.ele[v.offset]
}

// Cdr returns a VList beginning at the second element the receiver.
//
// It panics on an empty list.
func (v VList) Cdr() VList {
	if v.base == nil {
		panic("Cdr on empty vList")
	}
	v.offset++
	if v.offset < len(v.base.ele) {
		return v
	}
	return VList{v.base.next, 0}
}

// Len returns the length of the list.  It is constant time.
func (v VList) Len() int {
	if v.base == nil {
		return 0
	}
	return len(v.base.ele)*2 - v.offset - 1
}

// String satisfies a stringer interface such as fmt.Stringer.
func (v VList) String() string {
	if v.base == nil {
		return "[]"
	}
	r := fmt.Sprintf("[%v", v.base.ele[v.offset])
	for sg, sl := v.base, v.base.ele[v.offset+1:]; ; {
		for _, e := range sl {
			r = fmt.Sprintf("%s %v", r, e)
		}
		sg = sg.next
		if sg == nil {
			break
		}
		sl = sg.ele
	}
	return r + "]"
}
