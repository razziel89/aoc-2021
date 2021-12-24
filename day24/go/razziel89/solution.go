package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	smallest = 10000000000000
	largest  = 99999999999999
)

// *    inp a - Read an input value and write it to variable a.
// *    add a b - Add the value of a to the value of b, then store the result in variable a.
// *    mul a b - Multiply the value of a by the value of b, then store the result in variable a.
// *    div a b - Divide the value of a by the value of b, truncate the result to an integer, then
//			store the result in variable a. (Here, "truncate" means to round the value toward zero.)
// *    mod a b - Divide the value of a by the value of b, then store the remainder in variable a.
//			(This is also called the modulo operation.)
// *    eql a b - If the value of a and b are equal, then store the value 1 in variable a.
//			Otherwise, store the value 0 in variable a.

// (Program authors should be especially cautious; attempting to execute div with b=0 or attempting
// to execute mod with a<0 or b<=0 will cause the program to crash and might even damage the ALU.
// These operations are never intended in any serious ALU program.)

const (
	// Ops.
	opInp = 'i'
	opAdd = 'a'
	opMul = 'm'
	opDiv = 'd'
	opMod = 'o'
	opEql = 'e'
	// Registers.
	wReg    = 'w'
	xReg    = 'x'
	yReg    = 'y'
	zReg    = 'z'
	allRegs = "wxyz"
)

type acl struct {
	w, x, y, z int
	num        int
}

func (a *acl) set(reg rune, num int) {
	switch reg {
	case wReg:
		a.w = num
	case xReg:
		a.x = num
	case yReg:
		a.y = num
	case zReg:
		a.z = num
	default:
		log.Fatal("unknown register")
	}
}

func (a acl) get(reg rune) int {
	switch reg {
	case wReg:
		return a.w
	case xReg:
		return a.x
	case yReg:
		return a.y
	case zReg:
		return a.z
	}
	log.Fatal("unknown register")
	return 0 // Will never be reached.
}

func (a *acl) setNum(num int) {
	a.num = num
}

func (a acl) getNum() int {
	return a.num
}

type op struct {
	act rune
	reg rune
	dat interface{}
}

// String gets a string rep.
func (o op) String() string {
	str := fmt.Sprintf("%c -> %c", o.act, o.reg)
	switch t := o.dat.(type) {
	case int:
		str += fmt.Sprintf(" : %d", t)
	case rune:
		str += fmt.Sprintf(" : %c", t)
	}
	if o.dat != nil {
	}
	return str
}

var reader = bufio.NewReader(os.Stdin)

func readLine() (string, error) {
	return reader.ReadString('\n')
}

func assertNotNil(i interface{}) {
	if i == nil {
		log.Fatal("nil interface")
	}
}

func parseInput() []op {
	result := []op{}
	for {
		line, err := readLine()
		if err == io.EOF {
			// Success case, no more input to read.
			return result
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		if len(fields[1]) != 1 {
			log.Fatal("cannot extract register")
		}
		if !strings.Contains(allRegs, fields[1]) {
			log.Fatal("unknown register")
		}
		reg := rune(fields[1][0])
		var data interface{}
		if len(fields) == 3 {
			if strings.Contains(allRegs, fields[2]) {
				// Is a register.
				data = rune(fields[2][0])
			} else {
				// Is a number.
				num, err := strconv.Atoi(fields[2])
				if err != nil {
					log.Fatal(err.Error())
				}
				data = num
			}
		}
		var o op
		switch fields[0] {
		case "inp":
			o = op{act: opInp, reg: reg}
		case "add":
			assertNotNil(data)
			o = op{act: opAdd, reg: reg, dat: data}
		case "mul":
			assertNotNil(data)
			o = op{act: opMul, reg: reg, dat: data}
		case "div":
			assertNotNil(data)
			o = op{act: opDiv, reg: reg, dat: data}
		case "mod":
			assertNotNil(data)
			o = op{act: opMod, reg: reg, dat: data}
		case "eql":
			assertNotNil(data)
			o = op{act: opEql, reg: reg, dat: data}
		}
		result = append(result, o)
	}
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

// *    inp a - Read an input value and write it to variable a.
// *    add a b - Add the value of a to the value of b, then store the result in variable a.
// *    mul a b - Multiply the value of a by the value of b, then store the result in variable a.
// *    div a b - Divide the value of a by the value of b, truncate the result to an integer, then
//			store the result in variable a. (Here, "truncate" means to round the value toward zero.)
// *    mod a b - Divide the value of a by the value of b, then store the remainder in variable a.
//			(This is also called the modulo operation.)
// *    eql a b - If the value of a and b are equal, then store the value 1 in variable a.
//			Otherwise, store the value 0 in variable a.

// (Program authors should be especially cautious; attempting to execute div with b=0 or attempting
// to execute mod with a<0 or b<=0 will cause the program to crash and might even damage the ALU.
// These operations are never intended in any serious ALU program.)

func findNum(inState acl, ops []op, startDigs []int) (int, bool) {
	for dig := startDigs[0]; dig > 0; dig-- {
		state := inState
		for opIdx, op := range ops {
			switch op.act {
			case opInp:
				var val int
				switch d := op.dat.(type) {
				case rune:
				case int:
				default:
					log.Fatal("unknown data type")
				}
				state.set(op.reg, val)
			case opAdd:
			case opMul:
			case opDiv:
			case opMod:
			case opEql:
			default:
				log.Fatal("unknown op")
			}
		}
	}
	return 0, false
}

func main() {
	ops := parseInput()
	for _, o := range ops {
		fmt.Println(o)
	}
	startNum := largest
	state := acl{} // Automatically zero initiated.
	num, valid := findNum(state, ops, numToDigs(startNum))
	fmt.Println(num, valid)
}
