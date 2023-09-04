package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type postFunc func(c context.Context) error

// interface 实现
func PostFuncImpl(c context.Context) error {
	return nil
}

func postParm(p postFunc) {

	p(nil)

}

// 购买商品
type buyGoods struct {
	Num   int             `json:"num"`   // 数量',
	Price decimal.Decimal `json:"price"` // 单价'
}

type a struct {
	Name string `json:"name"`
	B    b
}
type b []struct {
	Name string `json:"name"`
}

func IndexOf[T comparable](collection []T, element T) int {
	for i, item := range collection {
		if item == element {
			return i
		}
	}

	return -1
}

func main() {

	var pf = postFunc(PostFuncImpl)

	pf(nil)
	postParm(PostFuncImpl)

	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		bar <- 12
	}()
	go func() {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		defer wg.Done()
		select {
		case t := <-bar:
			foo <- t
			println("t", t)
		default:
			println("default")
		}
	}()
	//wg.Wait()

	//Demo2()

	hhh := `[{"name":"123"},{"name":"2565"}]`

	var jjj b
	json.Unmarshal([]byte(hhh), &jjj)

	for i, v := range jjj {
		fmt.Println(i, v)
	}

	GGG := &a{Name: "133"}

	fmt.Println(GGG)

	//fmt.Println(GGG.B[0].Name)

	for i, v := range GGG.B {
		fmt.Println("y", i, v)
	}

	e := buyGoods{Num: 123}
	sellNum, _ := decimal.NewFromString(strconv.Itoa(e.Num))
	totalMoney := MulDecimal(e.Price, sellNum)

	fmt.Printf("monry%s", totalMoney)
}

// 主要用于处理浮点数据精度

// 加法
func AddDecimal(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Add(d2)
}

// 减法
func SubDecimal(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Sub(d2)
}

// 乘法
func MulDecimal(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Mul(d2)
}

// 除法
func DivDecimal(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Div(d2)
}

// int
func IntDecimal(d decimal.Decimal) int64 {
	return d.IntPart()
}

// float
func FloatDecimal(d decimal.Decimal) float64 {
	f, exact := d.Float64()
	if !exact {
		return f
	}
	return 0
}

func Demo2() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i // 或者直接 return 效果相同
}
