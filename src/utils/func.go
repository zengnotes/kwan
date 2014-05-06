package utils

import (
	//"fmt"
	"math/rand"
	"math/big"
	"strconv"
	"encoding/base64"
)

func RandomString(size int) string {
	if size <= 0 {
		size = 5
	}
	alpha := `abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789` //better to define a const
	buf := make([]byte, size)
	alpha_len := len(alpha)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(alpha_len)]
	}
	return string(buf)

}

func unicode(input string) []int32 {
	var result []int32
	for _, c:= range input {
		result = append(result, c)
	}
	return result
}

func Crypt(input , key string) string {
	key2 := unicode(key)
	key_len := len(key2)
	result := ""
	i := 0
	for _, c := range input {
		k := key2[i % key_len]
		ck := new(big.Int).Xor(big.NewInt(int64(c)),big.NewInt(int64(k)))
		cki, _ := strconv.Atoi(ck.String())
		result += string(rune(cki))
		i++
	}
	return result
}

func Base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64_decode(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str)
	return string(data)
}
//http://play.golang.org/p/gmYKrUcx4D