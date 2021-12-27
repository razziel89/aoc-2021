#!/bin/bash
run() {
  MY_FIRST_DIG="$1" go run . < real.dat | tee "$1.dat"
}

run 9 &
echo $!
run 7 &
echo $!
run 5 &
echo $!
run 3 &
echo $!
run 1 &
echo $!

wait
