package helper

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 生成随机字符串
func RandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	b := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

// 错误终止
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// 解析string为float64
func ParseStr2Float64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

// 解析string为int64
func ParseStr2Int64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

// 解析科学计数法为int64
func ParseBigDecimal2Int64(str string) int64 {
	decimalNum, _ := decimal.NewFromString(str)
	return decimalNum.BigInt().Int64()
}

// 移除Bom头
func RemoveBom(b []byte) []byte {
	return bytes.TrimPrefix(
		b,
		[]byte{
			239,
			187,
			191,
		},
	)
}

// GBK转Utf8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// Utf8转GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
