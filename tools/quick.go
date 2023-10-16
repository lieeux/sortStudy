package tools

func Quick(left int, right int, array *[100000]int) {
	l := left                      //left表示数组左边的下标
	r := right                     //right表示数组右边的下标
	pivot := array[(left+right)/2] //pivot表示中轴，这里会自动取整
	for l < r {                    //把比pivot小的数放在左边，把比pivot大的数放在右边
		for array[l] < pivot { //从左开始找一个>=pivot的值
			l++
		}
		for array[r] > pivot { //从右开始找一个<=pivot的值
			r--
		}
		if l >= r { //如果>=pivot的值已经不在<=pivot的值的左边了，就不用再交换了
			break
		}
		array[l], array[r] = array[r], array[l]

		if array[l] == pivot { //这俩判断可以少一个，但都没有就卡死了
			r--
		}
		if array[r] == pivot { //这俩判断可以少一个，但都没有就卡死了
			l++
		}

	}

	if l == r { //跳过pivot
		l++
		r--
	}

	if left < r { //左边递归
		Quick(left, r, array)
	}
	if l < right { //右边递归
		Quick(l, right, array)
	}
	return
}
