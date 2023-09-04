package main

import (
	"fmt"
	"gotest/utest"
)


func main(){

  var d = utest.Abs(6)
  
  fmt.Println(d)

  var a= utest.Max(6,1)
  
  fmt.Println(a)

}