package main

import "container/heap"

// 1488. 避免洪水泛滥
func avoidFlood(rains []int) []int {

	rainsLen := len(rains)       //后面会用到长度
	ans := make([]int, rainsLen) //结果

	// 湖泊下雨日记录
	rainyDaysInLakes := map[int][]int{}

	//初始化下雨记录
	for i, v := range rains {
		rainyDaysInLakes[v] = append(rainyDaysInLakes[v], i)
		if v >= 1 { //下雨
			ans[i] = -1 //默认不能干活
		}
		if v == 0 { //不下雨并且可以干活
			ans[i] = 1 //默认不干活
		}
	}

	// 最近下雨日期小根堆
	dateHeap := &latestDateHeap{}
	heap.Init(dateHeap)

	// 有水的湖泊
	lakeWithWater := map[int]bool{}

	for i, v := range rains {

		if v >= 1 { //下雨
			if ok, v := lakeWithWater[v]; ok && v {
				return []int{}
			}
			lakeWithWater[v] = true           //添加有水的湖泊
			if len(rainyDaysInLakes[v]) > 1 { //防止下标越界
				rainyDaysInLakes[v] = rainyDaysInLakes[v][1:] //下过雨移除记录
			} else if len(rainyDaysInLakes[v]) == 1 {
				rainyDaysInLakes[v] = []int{} //之下一天清除记录
			}

			//有水的湖泊再记录,同时防止下标越界
			if ok, ishaveWater := lakeWithWater[v]; ok && ishaveWater && len(rainyDaysInLakes[v]) > 0 {
				//更新最近下雨日期
				heap.Push(dateHeap, itemContent{v, rainyDaysInLakes[v][0]})
			}

		} else if v == 0 && len(*dateHeap) > 0 { //不下雨并且可以干活
			workLake := (heap.Pop(dateHeap)).(itemContent) //抽即将下雨并且有水得的湖泊
			if len(rainyDaysInLakes[workLake.lakeNum]) > 0 && i < workLake.dateWillRain {
				lakeWithWater[workLake.lakeNum] = false //移除有水的湖泊
				// rainyDaysInLakes[workLake.lakeNum] = rainyDaysInLakes[workLake.lakeNum][1:] //移除即将下雨的日期
				ans[i] = workLake.lakeNum
			}
		}
	}

	return ans[:]
}

// 最近下雨日期小根堆--------------------------------------------------------------
type latestDateHeap []itemContent

// 湖泊下雨结构体
type itemContent struct {
	lakeNum      int //第几个湖泊
	dateWillRain int //下次将要下雨的日期
}

// 实现heap.Interface接口------------------
// Len is the number of elements in the collection.
func (_latestDateHeap latestDateHeap) Len() int { return len(_latestDateHeap) }

// 实现sort.Iterface------------------
// Less reports whether the element with
// index i should sort before the element with index j.
func (_latestDateHeap latestDateHeap) Less(i, j int) bool {
	return _latestDateHeap[i].dateWillRain < _latestDateHeap[j].dateWillRain
}

// Swap swaps the elements with indexes i and j.
func (_latestDateHeap latestDateHeap) Swap(i, j int) {
	_latestDateHeap[i], _latestDateHeap[j] = _latestDateHeap[j], _latestDateHeap[i]
}

// 实现sort.Iterface-------------------
func (_latestDateHeap *latestDateHeap) Push(item interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*_latestDateHeap = append(*_latestDateHeap, item.(itemContent))
}

func (_latestDateHeap *latestDateHeap) Pop() interface{} {
	old := *_latestDateHeap
	lenOld := len(old)                //切片长度
	endOne := old[lenOld-1]           //最后一盒元素
	*_latestDateHeap = old[:lenOld-1] //弹出最后一个元素
	return endOne
}

// // 最近下雨日期有开关的大/小根堆
// type dateHeapSTRUCT struct {
// 	heap  []itemContent `json:"heap,omitempty"`
// 	onOff bool          `json:"onOff,omitempty"` //true:小根堆,false:大跟堆
// }

// type itemContent struct {
// 	lakeNum      int `json:"lakeNum,omitempty"`      //第几个湖泊
// 	dateWillRain int `json:"dateWillRain,omitempty"` //下次将要下雨的日期
// }

// // Len is the number of elements in the collection.
// func (_dateHeap dateHeapSTRUCT) Len() int { return len(_dateHeap.heap) }

// // Less reports whether the element with
// // index i should sort before the element with index j.
// func (_dateHeap dateHeapSTRUCT) Less(i, j int) bool {
// 	if _dateHeap.onOff {
// 		return _dateHeap.heap[i].dateWillRain < _dateHeap.heap[j].dateWillRain
// 	}
// 	return _dateHeap.heap[i].dateWillRain > _dateHeap.heap[j].dateWillRain
// }

// // Swap swaps the elements with indexes i and j.
// func (_dateHeap dateHeapSTRUCT) Swap(i, j int) {
// 	_dateHeap.heap[i], _dateHeap.heap[j] = _dateHeap.heap[j], _dateHeap.heap[i]
// }

// func (_dateHeap *dateHeapSTRUCT) Push(item itemContent) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*&_dateHeap.heap = append(_dateHeap.heap, item)
// }

// func (_dateHeap *dateHeapSTRUCT) Pop() itemContent {
// 	old := _dateHeap.heap
// 	lenOld := len(old)                 //切片长度
// 	endOne := old[lenOld-1]            //最后一盒元素
// 	_dateHeap.heap = old[0 : lenOld-1] //弹出最后一个元素
// 	return endOne
// }

// This example inserts several ints into an latestDateHeap, checks the minimum,
// and removes them in order of priority.
// func Example_intHeap() {
// 	latestDateHeap := &latestDateHeap{2, 1, 5}
// 	heap.Init(latestDateHeap)
// 	heap.Push(latestDateHeap, 3)
// 	fmt.Printf("minimum: %d\n", (*latestDateHeap)[0])
// 	for latestDateHeap.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(latestDateHeap))
// 	}
// 	// Output:
// 	// minimum: 1
// 	// 1 2 3 5
// }

/**
 * @description: 判断是否有
 * @param {[]int} intSplice
 * @param {int} v
 * @return {bool} *
 */
func isHaveInSplice(intSplice []int, item int) bool {
	isHave := false
	for _, v := range intSplice {
		if v == item {
			isHave = false
			return isHave
		}
	}
	return isHave
}

/**
 * @description: 获取int切片最大值--------------------------------------------
 * @param {[]int} numbers
 * @return {int} maxInt
 */
func getIntsMaxOneValue(numbers []int) (maxInt int) {
	for _, v := range numbers {
		if v > maxInt {
			maxInt = v
		}
	}
	return maxInt
}
