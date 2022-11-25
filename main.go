// goTimer - Timer using tempfile
// First run save time now
// Next run show duratin time and deletes tempfile
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	fileprefix  = "gotimer-"
	filepostfix = ".tmp"
	timeformat  = "2006-01-02T15:04:05.999999-07:00"
)

func main() {
	var (
		flagPname       = flag.String("name", "one", "Name of the timer")
		timerfilename   = os.TempDir() + fileprefix + *flagPname + filepostfix
		timerfileExists = doesFileExist(timerfilename)
	)
	flag.Usage = func() { usage() }
	flag.Parse()

	if timerfileExists {
		if err := readTimeNowFile(timerfilename); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else {
		if err := createTimeNowFile(timerfilename); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

// doesFileExist checks if file exists
func doesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)
	// check if error is "file not exists"
	return !os.IsNotExist(error)
}

// createTimeNowFile create a tempfile with time now
func createTimeNowFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	t := time.Now()
	_, err = f.WriteString(t.Format(timeformat)) // store in file
	if err != nil {
		return err
	}
	return err
}

// readTimeNowFile read a tempfile with time and show past time
// then deletes temp file
func readTimeNowFile(filename string) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	savedtime, err := time.Parse(timeformat, string(fileData))
	if err != nil {
		return err
	}
	err = os.Remove(filename)
	if err != nil {
		return err
	}
	fmt.Println(time.Since(savedtime))
	return nil
}

func usage() {
	fmt.Printf("Chronograph that is independent from running in ")
	fmt.Printf("the same shell, it is using a temp file to store start time. ")
	fmt.Printf("Dependent on the system there will be a delay time ")
	fmt.Printf("on test system it about 57ms\n")
	fmt.Printf("To get the delay run %[1]s & %[1]s to test\n", os.Args[0])
	fmt.Printf("%s by David Nilsson\n", os.Args[0])
	fmt.Printf("Usage:\n")
	fmt.Printf("	%s\n", os.Args[0])
	fmt.Printf("Starts standard timer\n")
	fmt.Printf("	%s\n", os.Args[0])
	fmt.Printf("Stops timer\n")
	flag.PrintDefaults()
}
