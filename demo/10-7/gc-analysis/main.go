package main

func main() {
	const numElements = 10000000

	var data []int

	for i := 0; i < numElements; i++ {
		data = append(data, i)
		processData(data)
	}
}

func processData(data []int) {
	_ = len(data)
}
