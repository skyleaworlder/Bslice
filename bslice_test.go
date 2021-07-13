package bslice

import (
	"fmt"
	"testing"
)

func Test_CreateBslice(t *testing.T) {
	bs := CreateBslice()
	fmt.Println(bs)
}

func Test_Enqueue(t *testing.T) {
	bs := CreateBslice()
	testEnqueueElements := [][]int{
		{1, 2, 46, 18, 21, 8, 215},
		{2, 4, 5, 6},
		{3, 4, 4, 7, 0},
	}
	testEnqueueChangeOp := []bool{
		false, true, false,
	}

	for i, elems := range testEnqueueElements {
		enqarr := make([]interface{}, len(elems))
		for j, v := range elems {
			enqarr[j] = v
		}
		bs.EnqueueMultiElem(enqarr, testEnqueueChangeOp[i])
		fmt.Println("after ", i+1, " enqueue operation:", bs)
	}
}

func Test_Dequeue(t *testing.T) {
	bs := CreateBslice()
	testEnqueueElements := [][]int{
		{1, 2, 46, 18, 21, 8, 215},
		{2, 4, 5, 6},
		{3, 4, 4, 7, 0},
	}
	testEnqueueChangeOp := []bool{
		false, true, false,
	}

	for i, elems := range testEnqueueElements {
		enqarr := make([]interface{}, len(elems))
		for j, v := range elems {
			enqarr[j] = v
		}
		bs.EnqueueMultiElem(enqarr, testEnqueueChangeOp[i])
	}

	deqarr, n, err := bs.DequeueMultiElem(2)
	fmt.Println(bs, deqarr, n, err)
	deqarr, n, err = bs.DequeueMultiElem(len(testEnqueueElements) - 2)
	fmt.Println(bs, deqarr, n, err)
	deqarr, n, err = bs.DequeueAll()
	fmt.Println(bs, deqarr, n, err)
}
