package tools

func Bubble(array *[100000]int) {
	for i := len(array) - 1; i > 0; i-- { //控制比到哪
		noexchange := true
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				noexchange = false
			}
		}
		if noexchange { //没交换就说明排好了
			break
		}
	}
	return
}
