////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2024/9/9 23:09
////////////////////////////////////////////////////////////////////////////////

package sort

import "testing"

var arr = []int{1, 5, 7, 2, 3, 9, 6, 4, 8, 10, 0}

func TestBubbleSort(t *testing.T) {
	t.Log(BubbleSort(arr))
}

func TestSelectionSort(t *testing.T) {
	t.Log(SelectionSort(arr))
}

func TestInsertionSort(t *testing.T) {
	t.Log(InsertSort(arr))
}

func TestMergeSort(t *testing.T) {
	t.Log(ShellSort(arr))
}

func TestMergeSort2(t *testing.T) {
	t.Log(MergeSort(arr))
}

func TestQuickSort(t *testing.T) {
	t.Log(QuickSort(arr))
}

func TestHeapSort(t *testing.T) {
	t.Log(HeapSort(arr))
}

func TestCountingSort(t *testing.T) {
	t.Log(CountingSort(arr, 12))
}

func TestBucketSort(t *testing.T) {
	t.Log(BucketSort(arr))
}
