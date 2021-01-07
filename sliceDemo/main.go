package main

import (
	"flag"

	"fmt"

	"log"

	"os"

	"runtime/pprof"
)

var (

	//定义外部输入文件名字

	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file.")
)

func main() {

	log.Println("begin")

	flag.Parse()

	if *cpuprofile != "" {

		f, err := os.Create(*cpuprofile)

		if err != nil {

			log.Fatal(err)

		}

		pprof.StartCPUProfile(f)

		defer pprof.StopCPUProfile()

	}

	for i := 0; i < 30; i++ {

		nums := fibonacci(i)

		fmt.Println(nums)

	}

}

//递归实现的斐波纳契数列

func fibonacci(num int) int {

	if num < 2 {

		return 1

	}

	return fibonacci(num-1) + fibonacci(num-2)

}
