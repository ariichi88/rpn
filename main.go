package main

import (
	"strconv"
	"fmt"
	"os"
	"strings"
)

func main() {

	var c calc

	exec := func(input string) float64 {
		var r float64
		f, err := strconv.ParseFloat(input, 64)
		if err != nil {
			if c.isCalculable() {
				r = c.calculation(input)
				c.delStack()
				c.addStack(r)
			} 
		} else {
			c.addStack(f)
			return f
		}
		return r
	}

	if len(os.Args) == 1 {
		var i string
		var f float64
		fmt.Println("exit with q")
		for {
			fmt.Scan(&i)
			f = exec(i)
			if strings.Contains(i, "q") {
				fmt.Println("exit")
				os.Exit(0)
			}
			fmt.Println(f)
		}
	} else {
		var f float64
		for _, arg := range(os.Args) {
			f = exec(arg)
		}
		fmt.Println(f)
	}
}
