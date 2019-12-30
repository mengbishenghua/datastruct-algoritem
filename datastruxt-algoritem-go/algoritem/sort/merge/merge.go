package merge

// Create by czx on 2019/12/30

func Merge(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2
		sort(arr, l, m)
		sort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	temp := make([]int, r-l+1)
	p1 := l
	p2 := m + 1
	i := 0
	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			temp[i] = arr[p1]
			i++
			p1++
		} else {
			temp[i] = arr[p2]
			i++
			p2++
		}
	}

	for p1 <= m {
		temp[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= r {
		temp[i] = arr[p2]
		i++
		p2++
	}

	copy(arr[l:], temp)
}
