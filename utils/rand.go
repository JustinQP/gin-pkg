package utils

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

const letterBytes = "0123456789ABCDEF"

// 16进制字符串 转 整形字符串
func HexTo10String(s string) (string, error) {
	bigInt := new(big.Int)
	c, isOk := bigInt.SetString(s, 0)
	if !isOk {
		return "", fmt.Errorf("HexString(%s) to DecimalString err", s)
	}
	return c.String(), nil
}

// 大整形字符串  转  16进制字符串
func HexFrom10String(s string) (string, error) {
	bigInt, isOk := big.NewInt(0).SetString(s, 0)
	if !isOk {
		return "", fmt.Errorf("DecimalString(%s to HexString err", s)
	}
	return fmt.Sprintf("0x%064x", bigInt), nil
}

func Rand64String() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 64)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	s, _ := HexTo10String("0x" + string(b))
	return s
}

func Rand16String() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	s, _ := HexTo10String("0x" + string(b))
	return s
}

func Rand6Int() string {
	str := "0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}

	return string(b)
}

// 代币数值转换智能合约的数值类型，带单位  1.1  ==> 1100000000000000000
func TokenValueConv(value float64, decimals uint8) string {
	decimalsValue := decimal.NewFromFloat(math.Pow10(int(decimals)))
	s := decimal.NewFromFloat(value).Mul(decimalsValue)
	return s.String()
}

// 智能合约的数值类型转浮点型 1100000000000000000 ==>  1.1
func TokenValueTo(amount string, decimals uint8) (float64, error) {
	amountVaule, err := decimal.NewFromString(amount)
	if err != nil {
		return 0, fmt.Errorf("TokenValueTo(%s) is err", err)
	}
	decimalsValue := decimal.NewFromFloat(math.Pow10(int(decimals)))

	s := amountVaule.Div(decimalsValue)
	sF, _ := s.Float64()
	return sF, nil
}
