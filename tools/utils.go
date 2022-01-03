package tools

import (
	"container/list"
	"fmt"
	"strconv"

	"github.com/nixomose/nixomosegotools/tools/thousands"
)

func Stringtouint32(str string) (error, uint32) {
	var val, err = strconv.ParseUint(str, 10, 32)
	if err != nil {
		return err, 0
	}
	return nil, uint32(val)
}

func Uint32tostring(v uint32) string {
	return fmt.Sprint(v)
}

func Inttostring(v int) string {
	return fmt.Sprint(v)
}

// Max returns the larger of x or y.
func Maxint64(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}
func Maxint(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Minint64(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

func Minint(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Powint(b int, e int) int {
	// IntPow calculates b to the eth power. Since the result is an int, it is assumed that e is a positive power

	if e == 0 {
		return 1
	}
	result := e
	for i := 2; i <= e; i++ {
		result *= b
	}
	return result
}

func Getlistitematpos(list *list.List, pos int) *list.Element {
	// return the item a list position pos, nil if out of range

	var count int = 0

	for item := list.Front(); item != nil; item = item.Next() {
		if count == pos {
			return item
		}
		count = count + 1
	}
	return nil

}

func Prettylargenumber_uint64(n uint64) string {
	//	nn := strconv.FormatInt(int64(n), 10)
	s, err := thousands.Separate(int64(n), "en")
	if err != nil {
		return "0"
	}
	return s
}
func Prettylargenumber_int64(n int64) string {
	//	nn := strconv.FormatInt(int64(n), 10)
	s, err := thousands.Separate(n, "en")
	if err != nil {
		return "0"
	}
	return s
}
