package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time module.")

	//current time
	presentTime := time.Now()
	fmt.Println(presentTime)

	//formatting
	formatTime := presentTime.Format("01-02-2006 15:04:05 Monday") // We have to send these fixed hardcoded values to format time.
	fmt.Println(formatTime)

	// other time than current time and formatting
	createDate := time.Date(2029, time.June, 10, 14, 25, 51, 0, time.UTC)
	fmt.Println(createDate)
	formatcreateDate := createDate.Format("01-02-2006 15:04:05 Monday") // We have to send these fixed hardcoded values to format time.
	fmt.Print(formatcreateDate)

	/*
		go build - used to create .exe file which can be used to run a main file automatically in a folder or project
		.exe file is normally used to build an application
		An ".exe" file is an executable file format primarily associated with Microsoft Windows operating systems.
		The term "exe" stands for "executable," and these files contain machine code that can be executed by the
		computer's operating system.
	*/
}
