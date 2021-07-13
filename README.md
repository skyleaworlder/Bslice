# Bslice

一个无用的数据结构。

动机是看到一道层次遍历但需要记忆上一级节点的题。写的时候必须维护两个队列，这让我很难受。于是我就想着，如果我都对一个数据结构操作就好了，所以就有了这个。

## API

```go
func (bs *Bslice) Top() (elem interface{}, err error)

func (bs *Bslice) Front() (elem interface{}, err error)

func (bs *Bslice) EnqueueMultiElem(elems []interface{}, changeAfterOp bool) (size int)

func (bs *Bslice) DequeueMultiElem(num int) (elems []interface{}, size int, err error)

func (bs *Bslice) DequeueAll() (elems []interface{}, size int, err error)
```

通过 `EnqueueMultiElem` 函数的 `changeAfterOp` 参数，决定下一次 `enqueue` 的队列是否改变。

## Example

```go
func main() {
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

        // print out:
        // after 1 enqueue operation: &{true 0 [1 2 46 18 21 8 215] []}
        // after 2 enqueue operation: &{false 1 [1 2 46 18 21 8 215 2 4 5 6] []}
        // after 3 enqueue operation: &{false 1 [1 2 46 18 21 8 215 2 4 5 6] [3 4 4 7 0]}
        fmt.Println("after ", i+1, " enqueue operation:", bs)
    }

    deqarr, n, err := bs.DequeueMultiElem(2)
    // print out: &{false 1 [46 18 21 8 215 2 4 5 6] [3 4 4 7 0]} [1 2] 2 <nil>
    fmt.Println(bs, deqarr, n, err)

    deqarr, n, err = bs.DequeueMultiElem(len(testEnqueueElements) - 2)
    // print out: &{false 1 [18 21 8 215 2 4 5 6] [3 4 4 7 0]} [46] 1 <nil>
    fmt.Println(bs, deqarr, n, err)

    deqarr, n, err = bs.DequeueAll()
    // print out: &{false 1 [] [3 4 4 7 0]} [18 21 8 215 2 4 5 6] 8 <nil>
    fmt.Println(bs, deqarr, n, err)
}
```
