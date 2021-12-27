package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// smallest = 10000000000000
	largest = 99999999999999
	last    = 13
)

var t = time.Now()

type acl struct {
	// w, x, y int
	z int
}

func numToDigs(num int) []int {
	str := fmt.Sprint(num)
	result := []int{}
	for _, digAsStr := range strings.Split(str, "") {
		dig, err := strconv.Atoi(digAsStr)
		if err != nil {
			log.Fatal("internal error")
		}
		result = append(result, dig)
	}
	fmt.Println(num, result)
	return result
}

func neq(a, b int) int {
	if a == b {
		return 0
	}
	return 1
}

type fn = func(acl, int) acl

var funcs = []fn{fn01, fn02, fn03, fn04, fn05, fn06, fn07, fn08, fn09, fn10, fn11, fn12, fn13, fn14}

//nolint:gomnd
func fn01(state acl, dig int) acl {
	fmt.Println(1, dig)
	return acl{
		// w: dig,
		// x: 1,
		// y: dig + 9,
		z: dig + 9,
	}
}

//nolint:gomnd
func fn02(s acl, dig int) acl {
	fmt.Println(2, dig)
	return acl{
		// w: dig,
		// x: 1,
		// y: dig + 4,
		z: s.z*26 + dig + 4,
	}
}

//nolint:gomnd
func fn03(s acl, dig int) acl {
	fmt.Println(3, dig)
	return acl{
		// w: dig,
		// x: 1,
		// y: dig + 2,
		z: s.z*26 + dig + 2,
	}
}

//nolint:gomnd
func fn04(s acl, dig int) acl {
	fmt.Println(time.Since(t))
	t = time.Now()
	fmt.Println(4, dig)
	val := neq(s.z%26-9, dig)
	return acl{
		// w: dig,
		// x: val,
		// y: (dig + 5) * val,
		z: s.z/26*(25*val+1) + (dig+5)*val,
	}
}

//nolint:gomnd
func fn05(s acl, dig int) acl {
	val := neq(s.z%26-9, dig)
	return acl{
		// w: dig,
		// x: val,
		// y: (dig + 1) * val,
		z: s.z/26*(25*val+1) + (dig+1)*val,
	}
}

//nolint:gomnd
func fn06(s acl, dig int) acl {
	return acl{
		// w: dig,
		// x: 1,
		// y: dig + 6,
		z: 26 * s.z * (dig + 6),
	}
}

//nolint:gomnd
func fn07(s acl, dig int) acl {
	return acl{
		// w: dig,
		// x: 1,
		// y: dig + 11,
		z: 26*s.z + dig + 11,
	}
}

//nolint:gomnd
func fn08(s acl, dig int) acl {
	val := neq(s.z%26-10, dig)
	return acl{
		// w: dig,
		// x: val,
		// y: (dig + 15) * val,
		z: s.z/26*(25*val+1) + (dig+15)*val,
	}
}

//nolint:gomnd
func fn09(s acl, dig int) acl {
	return acl{
		// w: dig,
		// x: 1,
		// y: dig + 7,
		z: 26*s.z + dig + 7,
	}
}

//nolint:gomnd
func fn10(s acl, dig int) acl {
	val := neq(s.z%26-2, dig)
	return acl{
		// w: dig,
		// x: val,
		// y: (dig + 12) * val,
		z: s.z/26*(25*val+1) + (dig+12)*val,
	}
}

//nolint:gomnd
func fn11(s acl, dig int) acl {
	return acl{
		// w: dig,
		// x: 1,
		// y: dig + 15,
		z: 26*s.z + dig + 15,
	}
}

//nolint:gomnd
func fn12(s acl, dig int) acl {
	val := neq(s.z%26-15, dig)
	return acl{
		// w: dig,
		// x: val,
		// y: (dig + 9) * val,
		z: s.z/26*(25*val+1) + (dig+9)*val,
	}
}

//nolint:gomnd
func fn13(s acl, dig int) acl {
	val := neq(s.z%26-9, dig)
	return acl{
		// w: dig,
		// x: val,
		// y: (dig + 12) * val,
		z: s.z/26*(25*val+1) + (dig+12)*val,
	}
}

//nolint:gomnd
func fn14(s acl, dig int) acl {
	val := neq(s.z%26-3, dig)
	return acl{
		// w: dig,
		// x: val,
		// y: (dig + 12) * val,
		z: s.z/26*(25*val+1) + (dig+12)*val,
	}
}

const ten = 10

func pow10(exp int) int {
	result := 1
	for count := 0; count < exp; count++ {
		result *= ten
	}
	return result
}

//nolint:funlen
func findNum(inState acl, startDigs *[]int, level int) (int, bool) {
	myFn := funcs[level]
	for dig := (*startDigs)[level]; dig > 0; dig-- {
		state := myFn(inState, dig)
		if level == last {
			if state.z == 0 {
				fmt.Println("LAST", dig)
				return dig, true
			}
		} else {
			num, valid := findNum(state, startDigs, level+1)
			if valid {
				fmt.Println(dig)
				return pow10(last-level)*dig + num, true
			}
		}
	}
	return 0, false
}

func main() {
	myFirstDig, err := strconv.Atoi(os.Getenv("MY_FIRST_DIG"))
	if err != nil {
		log.Fatal("MY_FIRST_DIG not set")
	}
	digs := numToDigs(largest)
	digs[0] = myFirstDig
	fmt.Println(digs)
	state := acl{}
	num, valid := findNum(state, &digs, 0)
	fmt.Println(num, valid)
}
