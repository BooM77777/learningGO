package main

import "fmt"

func main() {
	s := "cdadabcc"
	fmt.Println(removeDuplicateLetters(s))
}

func removeDuplicateLetters(s string) string {

	var counter [26]int
	for _, c := range s {
		counter[c-'a']++
	}
	ret := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(ret) > 0 && ch < ret[len(ret)-1] {
				tmpCh := ret[len(ret)-1] - 'a'
				if counter[tmpCh] == 0 {
					break
				}
				inStack[tmpCh] = false
				ret = ret[:len(ret)-1]
			}
			ret = append(ret, ch)
			inStack[ch-'a'] = true
		}
		counter[ch-'a']--
	}

	return string(ret)
}
