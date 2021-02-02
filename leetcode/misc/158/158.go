package main

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

// Buffer 缓冲区
type Buffer struct {
	buf      []byte
	readIter int
	bufSize  int
	read4    func([]byte) int
}

// NewBuffer 构造函数
func NewBuffer(read4 func([]byte) int) *Buffer {
	return &Buffer{
		buf:      make([]byte, 4),
		readIter: 0,
		bufSize:  0,
		read4:    read4,
	}
}

func (b *Buffer) load() {
	b.bufSize = b.read4(b.buf)
	b.readIter = 0
}

func (b *Buffer) read(outBuf []byte, n int) int {
	if b.readIter == b.bufSize {
		b.load()
	}
	var cnt int
	// 清空缓冲区
	if b.readIter != 0 {
		if b.bufSize-b.readIter >= n {
			copy(outBuf, b.buf[b.readIter:b.readIter+n])
			b.readIter += n
			return n
		}
		cnt = b.bufSize - b.readIter
		copy(outBuf[:cnt], b.buf[b.readIter:b.bufSize])
		b.load()
	}
	for b.bufSize > 0 && cnt < n {
		if cnt+b.bufSize <= n {
			copy(outBuf[cnt:cnt+b.bufSize], b.buf)
			cnt += b.bufSize
			b.load()
		} else {
			b.readIter = n - cnt
			copy(outBuf[cnt:], b.buf[:b.readIter])
			cnt = n
		}
	}
	return cnt
}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	buffer := NewBuffer(read4)
	// implement read below.
	return func(buf []byte, n int) int {
		return buffer.read(buf, n)
	}
}

func main() {

}
