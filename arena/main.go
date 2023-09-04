package  main

import (
	"arena"
   )

   type T struct {
	val int
   }
   func main() {
	a := arena.New()
	var ptrT *T
	a.New(&ptrT)
	ptrT.val = 1
	var sliceT []T
	a.NewSlice(&sliceT, 100)
	sliceT[99].val = 4
	a.Free()
   }