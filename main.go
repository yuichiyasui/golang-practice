package main

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			println("奇数:", i)
		} else {
			println("偶数:", i)
		}
	}
}
