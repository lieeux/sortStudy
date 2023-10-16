package tools

func Select(array *[100000]int) {
	for i := 0; i < len(array)-1; i++ {
		minid := i
		for j := i + 1; j < len(array); j++ { //array[i]不用参加比较
			if array[j] < array[minid] { //提升效率的关键在于这层循环时不交换数据，只记录id
				minid = j
			}
		}
		array[i], array[minid] = array[minid], array[i]
	}
	return
}
