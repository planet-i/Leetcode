// 演示如何使用 Go 语言中的通道来协调并发任务。
// 埃拉托斯特尼筛法 筛选素数
// 首先标记2为素数，然后将2的所有倍数标记为合数。
// 接下来，找到下一个未标记的数字（即3），将它标记为素数，然后将3的所有倍数标记为合数。
// 然后继续这个过程，依次找到未标记的下一个最小的数字，并标记它为素数，然后将它的倍数标记为合数，直到无法找到更多未标记的数字。
// 这个算法的效率很高，因为它避免了对每个数字都进行素数测试，而是通过逐步筛选合数来找到素数。最终，所有未标记为合数的数字都是素数。
package main

// Generate 从2开始产生自然数序列中的所有整数
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Filter 筛选出当前素数的倍数，将剩余数发送到out通道
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 { // 如果数不能被当前素数整除，就发送到out通道。否则忽略
			out <- i
		}
		// 当前素数的倍数就都被忽略了
	}
}

func main() {
	in := make(chan int)
	go Generate(in)
	for i := 0; i < 10; i++ {
		prime := <-in
		print(prime, "\n") // 在每次循环中从 ch 中读取一个素数，并将其存储在 prime 变量中
		out := make(chan int)
		go Filter(in, out, prime)
		in = out
	}
}
