package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

var (
	lower       = "abcdefghijklmnopqrstuvwxyz"
	upper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num         = "0123456789"
	specialEasy = "!#%&/()=.:-_@"
	specialHard = "!\"#¤%&/()=?@£$€¥{[]}\\~'*,.-;:_§<>|"
)

func randInt(max int) int {
	n, err := (rand.Int(rand.Reader, big.NewInt(int64(max))))
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}
func randString(s string) string {
	i := randInt(len(s))
	return s[i : i+1]
}

func main() {
	var easy bool
	flag.BoolVar(&easy, "easy", true, "Easy human readable/typable")
	var pwLen int
	flag.IntVar(&pwLen, "len", 10, "Character length of password")
	flag.Parse()

	chars := lower + upper + num
	if easy {
		chars = chars + specialEasy
		chars = strings.Replace(chars, "l", "", -1)
		chars = strings.Replace(chars, "I", "", -1)
		chars = strings.Replace(chars, "1", "", -1)
		chars = strings.Replace(chars, "|", "", -1)
		chars = strings.Replace(chars, "O", "", -1)
		chars = strings.Replace(chars, "0", "", -1)
	} else {
		chars = chars + specialHard
	}

	var pass string
	for i := 0; i < pwLen; i++ {
		pass = pass + randString(chars)
	}

	fmt.Println(pass)
}
