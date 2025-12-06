package task2

import "fmt"

// 指针
func modifyint(i *int) {
	*i += 10

}

func modifySlice(slice []int) {
	if len(slice) == 0 {
		return
	}
	for i := range slice {
		slice[i] *= 2
	}
}

func Point() {
	i := 2
	modifyint(&i)
	fmt.Printf("i: %v\n", i)

	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Printf("slice: %v\n", slice)
}
