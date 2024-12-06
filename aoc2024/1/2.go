package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type report struct {
	levels []int
	isSafe bool
}

func checkSafe(r *report) bool {
	inc := r.levels[0] < r.levels[1]
	for i := 1; i < len(r.levels); i++ {
		curr := r.levels[i]
		prev := r.levels[i-1]
		if inc && curr <= prev {
			return false
		} else if !inc && curr >= prev {
			return false
		} else if math.Abs(float64(prev-curr)) > 3 {
			return false
		}
	}
	return true
}

func NewReport(line string) *report {
	levels := strings.Split(line, " ")
	r := new(report)
	r.levels = make([]int, len(levels))
	for i, l := range levels {
		r.levels[i], _ = strconv.Atoi(l)
	}
	r.isSafe = checkSafe(r)
	return r
}

func load_data() []*report {
	f, err := os.Open("input_2.txt")
	// f, err := os.Open("test.txt")
	if nil != err {
		fmt.Println("Error: ", err.Error())
		os.Exit(-1)
	}
	defer f.Close()
	reports := make([]*report, 0)
	// wc -l input_2.txt => 1000
	r := bufio.NewReader(f)
	for i := 0; i < 1000; i++ {
		s, e := r.ReadString('\n')
		if nil != e {
			fmt.Println("Error: ", err.Error())
			os.Exit(-1)
		}
		reports = append(reports, NewReport(s[:len(s)-1]))
	}
	return reports
}

func puzzle_1(reports []*report) {
	count := 0
	for _, v := range reports {
		if v != nil && v.isSafe {
			count++
		}
	}
	fmt.Println("Ans 1:", count)
}

func puzzle_2(reports []*report) {
	count := 0
	for _, v := range reports {
		if v.isSafe {
			count++
		} else {
			for i, _ := range v.levels {
				new_report := &report{
					levels: []int{},
				}
				new_report.levels = append(new_report.levels, v.levels[:i]...)
				new_report.levels = append(new_report.levels, v.levels[i+1:]...)

				if checkSafe(new_report) {
					count++
					// break inner loop
					break
				}
			}
		}
	}
	fmt.Println("Ans 2:", count)
}

func main() {
	reports := load_data()
	puzzle_1(reports)
	puzzle_2(reports)
}
