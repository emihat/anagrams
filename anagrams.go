package main

import (
	"fmt"
	"math/rand"
	"os"
)

var MaxPnum int = 65536

func usage() {
	fmt.Println("Usage: go run anagrams.go foobar")
	os.Exit(1)
}

func calcPnum(inRuneNum int) int {
	pnum := 1
	for i := 0; i < inRuneNum; i++ {
		pnum *= (i + 1)
	}

	if pnum < MaxPnum {
		return pnum
	} else {
		return MaxPnum
	}
}

func generateAnagram(inRune []rune, ch chan string) {
	word := ""
	check := make(map[int]bool)

	for i := 0; i < len(inRune); {
		j := rand.Intn(len(inRune))
		if _, flag := check[j]; flag == false {
			word = word + string(inRune[j])
			check[j] = true
			i++
		}
	}
	ch <- word
}

func generateAnagrams(in string) map[string]bool {
	var res = make(map[string]bool)
	inRune := []rune(in)
	pnum := calcPnum(len(inRune))
	//fmt.Println("pnum=", pnum)

	for i := 0; i < pnum; {
		ch := make(chan string)
		go generateAnagram(inRune, ch)
		word := <-ch
		if _, flag := res[word]; flag == false {
			res[word] = true
			i++
		}
	}
	return res
}

func main() {
	in := os.Args[1]

	if in == "" {
		usage()
	} else {
		res := generateAnagrams(in)
		for word, _ := range res {
			fmt.Println(word)
		}
	}
}
