package main

func poorPigs(buckets int, minutesToDie int, minutesToTest int) (ret int) {
	state := minutesToTest/minutesToDie + 1
	for i := 1; i < buckets; i *= state {
		ret++
	}
	return
}
func main() {

}
