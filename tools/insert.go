package tools

func Insert(array *[100000]int) {
	for i := 1; i < len(array); i++ { //从第二个元素开始做插入
		value := array[i]                        //待插入值
		index := i - 1                           //开始比较的值的下标
		for index >= 0 && array[index] > value { //没到头且被比较的值更大
			array[index+1] = array[index] //把被比较的值后移一位
			index--                       //再往前比一次
		}
		if index != i-1 { //如果发生过后移
			array[index+1] = value //把值插入没比过的那个值的后面
		}
	}
	return
}
