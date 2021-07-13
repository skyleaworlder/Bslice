package bslice

import "errors"

// Bslice is for some special circumstances
// when now is true,
//  elements will be enqueued to sa, and be dequeued from sb
// when now is false,
//  elements will be enqueued to sb, and be dequeued from sa
type Bslice struct {
	now     bool
	version int
	sa      []interface{}
	sb      []interface{}
}

// CreateBslice is to create a bslice
func CreateBslice() (bs *Bslice) {
	return &Bslice{
		now: true, version: 0,
		sa: make([]interface{}, 0),
		sb: make([]interface{}, 0),
	}
}

// Top is to get the last enqueued element
func (bs *Bslice) Top() (elem interface{}, err error) {
	if bs.now {
		if len(bs.sa) == 0 {
			errmsg := "bs.sa cannot get top element, the len of bs.sa is 0 now"
			return elem, errors.New(errmsg)
		}
		return bs.sa[len(bs.sa)-1], nil
	}
	if len(bs.sb) == 0 {
		errmsg := "bs.sb cannot get top element, the len of bs.sb is 0 now"
		return elem, errors.New(errmsg)
	}
	return bs.sb[len(bs.sb)-1], nil
}

// Front is to get the first element that might dequeue
// if bs.now is true, return the first element of bs.sb
// if bs.now is false, return the first element of bs.sa
func (bs *Bslice) Front() (elem interface{}, err error) {
	if bs.now {
		if len(bs.sb) == 0 {
			errmsg := "bs.sb cannot get top element, the len of bs.sb is 0 now"
			return elem, errors.New(errmsg)
		}
		return bs.sb[0], nil
	}
	if len(bs.sa) == 0 {
		errmsg := "bs.sa cannot get top element, the len of bs.sa is 0 now"
		return elem, errors.New(errmsg)
	}
	return bs.sa[0], nil
}

// EnqueueMultiElem is a function to enqueue multiple elements into queue
func (bs *Bslice) EnqueueMultiElem(elems []interface{}, changeAfterOp bool) (size int) {
	if changeAfterOp {
		defer func() { bs.now = !bs.now; bs.version++ }()
	}

	for _, elem := range elems {
		size = bs.enqueue(elem, false)
	}
	return
}

// DequeueMultiElem is to return multiple elements
func (bs *Bslice) DequeueMultiElem(num int) (elems []interface{}, size int, err error) {
	for i := 0; i < num; i++ {
		elem, err := bs.dequeue()
		if err != nil {
			errmsg := "bs.DequeueMultiElem error, iter num is smaller than given 'num' now:"
			return elems, i, errors.New(errmsg + err.Error())
		}
		elems = append(elems, elem)
	}
	return elems, num, nil
}

// DequeueAll is to dequeue all the elements in corresponding slice
func (bs *Bslice) DequeueAll() (elems []interface{}, size int, err error) {
	if bs.now {
		// flush bs.sb
		elems, bs.sb = bs.sb, bs.sb[len(bs.sb):len(bs.sb)]
		return elems, len(elems), nil
	}
	elems, bs.sa = bs.sa, bs.sa[len(bs.sa):len(bs.sa)]
	return elems, len(elems), nil
}

func (bs *Bslice) enqueue(elem interface{}, changeAfterOp bool) (size int) {
	// bs.now will be changed when changeAfterOp is true
	if changeAfterOp {
		defer func() { bs.now = !bs.now; bs.version++ }()
	}

	// append new elem into sa when bs.now is true
	if bs.now {
		bs.sa = append(bs.sa, elem)
		return len(bs.sa)
	}
	bs.sb = append(bs.sb, elem)
	return len(bs.sb)
}

func (bs *Bslice) dequeue() (elem interface{}, err error) {
	if bs.now {
		if len(bs.sb) == 0 {
			errmsg := "bs.sb cannot dequeue, the len of bs.sb is 0 now"
			return elem, errors.New(errmsg)
		}
		elem, bs.sb = bs.sb[0], bs.sb[1:]
		return
	}

	if len(bs.sa) == 0 {
		errmsg := "bs.sa cannot dequeue, the len of bs.sa is 0 now"
		return elem, errors.New(errmsg)
	}
	elem, bs.sa = bs.sa[0], bs.sa[1:]
	return
}
