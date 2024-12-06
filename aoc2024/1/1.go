package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func load_data() ([]int, []int) {
	f, err := os.Open("input_1.txt")
	// f, err := os.Open("test.txt")
	if nil != err {
		fmt.Println("Error: ", err.Error())
		os.Exit(-1)
	}
	defer f.Close()
	// wc -l input_1.txt => 1000
	list1 := make([]int, 1000)
	list2 := make([]int, 1000)
	// line: `<list_1> <list_2>`
	r := bufio.NewReader(f)
	for i := 0; err == nil && i < 1000; i++ {
		// for i := 0; err == nil && i < 6; i++ {
		s, e := r.ReadString(' ')
		if nil != e {
			fmt.Println("Error: ", err.Error())
			os.Exit(-1)
		}
		num, _ := strconv.Atoi(s[:len(s)-1])
		list1[i] = num

		s, err = r.ReadString('\n')
		if nil != err {
			fmt.Println("Error: ", err.Error())
			os.Exit(-1)
		}
		num, _ = strconv.Atoi(s[:len(s)-1])
		list2[i] = num
	}
	return list1, list2
}

type Num struct {
	val   int
	count int
}

func puzzle_1(list1, list2 []int) {
	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	diff := 0
	for i, v := range list1 {
		diff += int(math.Abs(float64(v - list2[i])))
	}
	fmt.Println("Ans 1:", diff)
}

func puzzle_2(list1, list2 []int) {
	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	nums := make(map[int]*Num)

	prev := -1
	j := 0
	for _, v := range list1 {
		if v == prev {
			nums[v].count++
			continue
		}
		prev = v
		count := 0
		for ; list2[j] <= v; j++ {
			if list2[j] == v {
				count++
			}
		}
		nums[v] = &Num{count, 1}
	}

	diff := 0
	for k, v := range nums {
		diff += k * v.val * v.count
	}
	fmt.Println("Ans 2:", diff)
}

func main() {
	list1, list2 := load_data()
	puzzle_1(list1, list2)
	puzzle_2(list1, list2)
}
