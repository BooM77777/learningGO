package main

import "bytes"

type Codec struct {
}

func msb(b byte) bool {
	return b>>7 == 1
}

func getLen(b byte) int {
	return int(b & 127)
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	buf := bytes.Buffer{}
	if len(strs) == 0 {
		buf.Write([]byte{0})
	} else {
		buf.Write([]byte{1})
		for _, word := range strs {
			if len(word) == 0 {
				buf.Write([]byte{0})
			}
			for i := 0; i < len(word); i += 127 {
				if len(word)-i < 128 {
					buf.Write([]byte{byte(len(word) - i)})
					buf.Write([]byte(word[i:]))
				} else {
					buf.Write([]byte{byte(255)})
					buf.Write([]byte(word[i : i+127]))
				}
			}
		}
	}
	return string(buf.Bytes())
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) (ret []string) {

	byteSeq := []byte(strs)
	if byteSeq[0] == 1 {
		i, N, length, nextIsFragmentation := 1, len(byteSeq), 0, true
		for i < N {
			word := bytes.Buffer{}
			for {
				length, nextIsFragmentation, i = getLen(byteSeq[i]), msb(byteSeq[i]), i+1
				word.Write(byteSeq[i : i+length])
				i += length
				if nextIsFragmentation == false {
					if word.Len() == 0 {
						ret = append(ret, "")
					} else {
						ret = append(ret, string(word.Bytes()))
					}
					break
				}
			}
		}
	}
	return
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));

func main() {

}
