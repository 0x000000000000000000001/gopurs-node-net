package main

import (
	"fmt"
	"time"
	
	"gopurs/output/Test.Main"
	"gopurs/output/gopurs_runtime"
)

func main() {
	gopurs_runtime.Apply(pkg_Test_Main.Get_main(), gopurs_runtime.Box[any](nil))
	fmt.Println("Main returned, sleeping for 2s to allow goroutines to finish...")
	time.Sleep(2 * time.Second)
}
