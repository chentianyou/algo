////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2024/9/9 23:05
////////////////////////////////////////////////////////////////////////////////

package sort

// BubbleSort
/*
冒泡排序

平均O(n^2) 最好O(n) 最坏O(n^2) 稳定
比较相邻元素，如果第一个比第二个大就交换
*/
func BubbleSort(arr []int) []int {
	var l = len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// SelectionSort
/*
选择排序

平均O(n^2) 最好O(n^2) 最坏O(n^2) 非稳定
1. 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
2. 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
3. 重复第二步，直到所有元素均排序完毕。
*/
func SelectionSort(arr []int) []int {
	var l = len(arr)
	for i := 0; i < l; i++ {
		minIdx := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[i] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

// InsertSort
/*
插入排序

平均O(n^2) 最好O(n^2) 最坏O(n^2) 稳定
1. 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
2. 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
*/
func InsertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

// ShellSort
/*
希尔排序
插入排序的优化

平均O(nlogn) 最好O(nlogn) 最坏O(nlogn) 非稳定

1. 选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
2. 按增量序列个数 k，对序列进行 k 趟排序；
3. 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/
func ShellSort(arr []int) []int {
	length := len(arr)
	gap := length / 2
	for gap > 0 {
		for i := gap; i < length; i++ {
			// 取间隔gap的子序列做插入排序
			for j := i - gap; j >= 0 && arr[j] > arr[j+gap]; j -= gap {
				arr[j+gap], arr[j] = arr[j], arr[j+gap]
			}
		}
		gap = gap / 2
	}
	return arr
}

// MergeSort
/*
归并排序

平均O(nlogn) 最好O(nlogn) 最坏O(nlogn) 非稳定

将数组不断的分为两份直到元素个数为1，然后按大小合并
*/
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func merge(left []int, right []int) (result []int) {
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}
	return
}

// QuickSort
/*
快速排序

平均O(nlogn) 最好O(nlogn) 最坏O(nlogn) 非稳定
1. 从数列中挑出一个元素，称为 "基准"（pivot）;
2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
   在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
*/
func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		idx := partition(arr, left, right)
		quickSort(arr, left, idx-1)
		quickSort(arr, idx+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	arch := left
	idx := left + 1
	for i := idx; i <= right; i++ {
		if arr[i] < arr[arch] {
			arr[idx], arr[i] = arr[i], arr[idx]
			idx++
		}
	}
	arr[idx-1], arr[arch] = arr[arch], arr[idx-1]
	return idx - 1
}

// HeapSort
/*
堆排序

1. 创建一个堆 H[0……n-1]；
2. 把堆首（最大值）和堆尾互换；
3. 把堆的尺寸缩小 1，并调用 shift_down(0)，目的是把新的数组顶端数据调整到相应位置；
4. 重复步骤 2，直到堆的尺寸为 1。
*/
func HeapSort(arr []int) []int {
	length := len(arr)
	// 创建大顶堆
	for i := length/2 - 1; i >= 0; i-- {
		heapify(arr, i, length)
	}

	// 调整结构
	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
	return arr
}

func heapify(arr []int, i, length int) {
	temp := arr[i]                              // 取出root节点
	for k := i*2 + 1; k < length; k = k*2 + 1 { // 从i节点左子开始
		if k+1 < length && arr[k] < arr[k+1] { // 如果左子小于右子，k指向右子
			k++
		}
		if arr[k] > temp { // 子节点大于父节点，将子节点值赋给父节点（不用进行交换）
			arr[i] = arr[k]
			i = k // 找当前节点的子节点
		} else {
			break
		}
	}
	arr[i] = temp // 将temp值放到最终的位置
}

// CountingSort
/*
计数排序

计数排序是用来排序0到100之间的数字的最好的算法，但是它不适合按字母顺序排序人名
但是，计数排序可以用在基数排序中的算法来排序数据范围很大的数组。

算法的步骤如下：

（1）找出待排序的数组中最大和最小的元素
（2）统计数组中每个值为i的元素出现的次数，存入数组C的第i项
（3）对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
（4）反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1
*/
func CountingSort(arr []int, maxVal int) []int {
	bucketLen := maxVal + 1
	bucket := make([]int, bucketLen)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	idx := 0
	for i := 0; i < bucketLen; i++ {
		for bucket[i] > 0 {
			arr[idx] = i
			idx++
			bucket[i]--
		}
	}
	return arr
}

// BucketSort
/*
桶排序

1. 在额外空间充足的情况下，尽量增大桶的数量
2. 使用的映射函数能够将输入的 N 个数据均匀的分配到 K 个桶中
*/
func BucketSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	maxVal, minVal := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if maxVal < arr[i] {
			maxVal = arr[i]
		} else if minVal > arr[i] {
			minVal = arr[i]
		}
	}

	bucketSize := 3
	bucketCount := (maxVal-minVal)/bucketSize + 1
	buckets := make([][]int, bucketCount)
	for i, v := range arr {
		idx := (v - minVal) / bucketSize
		buckets[idx] = append(buckets[idx], arr[i])
	}

	idx := 0
	for i := 0; i < bucketCount; i++ {
		InsertSort(buckets[i])
		for _, v := range buckets[i] {
			arr[idx] = v
			idx++
		}
	}
	return arr
}
