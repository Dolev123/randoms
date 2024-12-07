package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Node interface {
    Next(byte) Node
}

type singleNode struct {
    chr  byte
    next Node
}
func (n *singleNode) Next(b byte) Node {
    if n.chr == b {
	return n.next
    }
    return nil
}

type matchNode struct {
    nexts map[byte]Node
}
func (n *matchNode) Next(b byte) Node {
    if v, ok := n.nexts[b]; ok {
	return v
    }
    return nil
}

// final connects to start
type finalNode struct{
    next Node
}
func (n *finalNode) Next(b byte) Node {
    return n.next
}

// return (start_node, final_node)
func build_state_machine_1() (Node, Node) {
    // create nodes
    // m_node - checks for 'm' (start node)
    m_node := &singleNode{'m', nil}
    u_node := &singleNode{'u', nil}
    l_node := &singleNode{'l', nil}
    pstart_node := &singleNode{'(', nil}
    num1_node :=&matchNode{map[byte]Node{}} 
    comma_node :=&matchNode{map[byte]Node{}} 
    num2_node :=&matchNode{map[byte]Node{}}
    pend_node := &singleNode{')', nil}
    final_node := &finalNode{m_node}

    // setup connections
    m_node.next = u_node
    u_node.next = l_node
    l_node.next = pstart_node
    l_node.next = pstart_node
    pstart_node.next = num1_node
    num1_node.nexts[','] = comma_node
    for i:=byte(0); i < 10; i++ {
	num1_node.nexts['0'+i] = num1_node
    }
    for i:=byte(0); i < 10; i++ {
	comma_node.nexts['0'+i] = num2_node
    }
    num2_node.nexts[')'] = pend_node
    for i:=byte(0); i < 10; i++ {
	num2_node.nexts['0'+i] = num2_node
    }
    pend_node.next = final_node
 
    if false {
	fmt.Println("m_node \t\t:", m_node)
	fmt.Println("u_node \t\t:", u_node)
	fmt.Println("l_node \t\t:", l_node)
	fmt.Println("pstart_node \t:", pstart_node)
	fmt.Println("num1_node \t:", num1_node)
	fmt.Println("comma_node \t:", comma_node)
	fmt.Println("num2_node \t:", num2_node)
	fmt.Println("pend_node \t:", pend_node)
	fmt.Println("final_node \t:", final_node)
    }

    return m_node, pend_node
}

func exec_op(op string) int {
    // remove mul(
    op = op[4:]
    // left with "123,345"
    nums := strings.Split(op, ",")
    num1, _ := strconv.Atoi(nums[0])
    num2, _ := strconv.Atoi(nums[1])
    return num1 * num2
}

func puzzle_1(start, end Node) {
    // f, err := os.Open("test.txt")
    f, err := os.Open("input_3.txt")
    if nil != err {
	    fmt.Println("Error: ", err.Error())
	    os.Exit(-1)
    }
    defer f.Close()
    curr := start

    var ops []string
    var vals []int
    cache := make([]byte, 0)
    b := make([]byte, 1)
    for _, err = f.Read(b); nil == err; _,err = f.Read(b) {
	next := curr.Next(b[0])
	// fmt.Println("Read: ", b, string(b), "| Next:", next)
	if nil == next {
	    curr = start
	    cache = make([]byte, 0)
	} else if end == next {
	    op := string(cache)
	    val := exec_op(op)
	    // fmt.Println(op, "=>", val)
	    ops = append(ops, op)
	    vals = append(vals, val)
	    curr = start
	    cache = make([]byte, 0)
	} else {
	    cache = append(cache, b[0])
	    curr = next
	}
    }
    if err != nil {
	fmt.Println("Error: ", err.Error())
    }
    // fmt.Println("Ops:", ops)
    // fmt.Println("Vals:", vals)
    ans := 0
    for _, v := range vals {
	ans += v
    }
    fmt.Println("Ans 1:", ans)
}

func main() {
    start, end := build_state_machine_1()
    puzzle_1(start, end)
}
