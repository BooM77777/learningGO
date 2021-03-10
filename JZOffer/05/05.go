package main

func replaceSpace(s string) string {
	retVal := make([]byte, 0, 3*len(s))
	for i := range s {
		if s[i] == ' ' {
			retVal = append(retVal, []byte{'%', '2', '0'}...)
		} else {
			retVal = append(retVal, s[i])
		}
	}
	return string(retVal)
}

func main() {

}
