package main

import (
	"fmt"
	"runtime/debug"
	"xanthing/cmd"
)

func main() {
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Print("系统异常:", r)
				fmt.Println("Stack trace:")
				fmt.Println(string(debug.Stack()))
			}
		}()
		cmd.Execute()
	}()
}
