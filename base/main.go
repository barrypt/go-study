package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

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

func main() {

	hhh := `[{"name":"123"},{"name":"2565"}]`

	var jjj b
	json.Unmarshal([]byte(hhh), &jjj)

	for i, v := range jjj {
		fmt.Println(i, v)
	}

	GGG := &a{Name: "133"}

	fmt.Println(GGG)

	fmt.Println(GGG.B[0].Name)

	for i, v := range GGG.B {
		fmt.Println(i, v)
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
