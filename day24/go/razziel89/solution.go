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
	// smallest = 10000000000000
	largest = 99999999999999
)

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

// func (a *acl) setNum(num int) {
// 	a.num = num
// }
//
// func (a acl) getNum() int {
// 	return a.num
// }

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

//nolint:gomnd
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

//nolint:funlen
func findNum(inState acl, inReg rune, ops []op, startDigs []int) (int, bool) {
	if inState.num%10 != 0 {
		log.Fatal("read-in number modulo 10 not equal zero")
	}
DIGLOOP:
	for dig := startDigs[0]; dig > 0; dig-- {
		// Copy state.
		state := inState
		// Add my number to the correct input register.
		state.set(inReg, dig)
		for opIdx, op := range ops {
			// Read out registers and optional payload.
			reg := op.reg
			regVal := state.get(reg)
			var data int
			switch d := op.dat.(type) {
			case int:
				data = d
			case rune:
				data = state.get(d)
				// In other cases, this is the empty interface. Don't do anything.
			}
			// Process op.
			switch op.act {
			case opInp:
				// inp a - Read an input value and write it to variable a.
				// The writing to a register will happen in the next call.
				state.num *= 10
				num, valid := findNum(state, reg, ops[opIdx+1:], startDigs[1:])
				if valid {
					return num, true
				}
				// No valid number could be found for dig as our digit. Move on to the next one.
				continue DIGLOOP
			case opAdd:
				// add a b - Add the value of a to the value of b, then store the result in
				// variable a.
				val := regVal + data
				state.set(reg, val)
			case opMul:
				// mul a b - Multiply the value of a by the value of b, then store the result in
				// variable a.
				val := regVal * data
				state.set(reg, val)
			case opDiv:
				// Cation when attempting to execute div with b=0. Such a case allows us to skip
				// large swathes of numbers.
				if data == 0 {
					return 0, false
				}
				// div a b - Divide the value of a by the value of b, truncate the result to an
				// integer, then store the result in variable a. (Here, "truncate" means to round
				// the value toward zero.)
				val := regVal / data // Go also rounds negative numbers towards zero here. Nice!
				state.set(reg, val)
			case opMod:
				// Caution when attempting to execute mod with a<0 or b<=0.  Such a case allows us
				// to skip large swathes of numbers.
				if regVal < 0 || data <= 0 {
					return 0, false
				}
				// mod a b - Divide the value of a by the value of b, then store the remainder in
				// variable a. (This is also called the modulo operation.)
				val := regVal % data
				state.set(reg, val)
			case opEql:
				// eql a b - If the value of a and b are equal, then store the value 1 in variable
				// a. Otherwise, store the value 0 in variable a.
				if regVal == data {
					state.set(reg, 1)
				} else {
					state.set(reg, 0)
				}
			default:
				log.Fatal("unknown op")
			}
		}
		// End, op loop.
		// If we arrive here, we found a number, but only if the z register contains the value 0.
		// Otherwise, this is no valid number.
		if state.z == 0 {
			return state.num + dig, true
		}
		return 0, false
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
	if ops[0].act != opInp {
		log.Fatal("need to start reading in")
	}
	startReg := ops[0].reg
	num, valid := findNum(state, startReg, ops[1:], numToDigs(startNum))
	fmt.Println(num, valid)
}
