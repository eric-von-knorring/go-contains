package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
	"sync"
)

var channel = make(chan bool)
var anyMatch bool
var file string

var rootCommad = &cobra.Command {
	Use:                    "contains <words>",
	Short:                  "Check if standard input contains the words given",
	Long:                   "Reads standard input and returns error code 0 if the pattern was found or status code 1 if the words was not found.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		channels := make([]chan string, len(args))
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			waitForTextResult(len(args))
			wg.Done()
		}()

		for i, arg := range args {
			input := make(chan string, 10)
			go findText(arg, input)
			channels[i] = input
		}

		go propagateInput(channels)

		wg.Wait()
	},
}

func findText(text string, input chan string) {
	for in := range input {
		if strings.Contains(in, text) {
			channel <- true
			return
		}
	}
	channel <- false
}

func propagateInput(channels []chan string) {
	input := createInput()
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		for _, ch := range channels {
			ch <- scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for _, ch := range channels {
		close(ch)
	}
}

func createInput() *os.File {
	if file == "" {
		return os.Stdin
	} else {
		fileHandle, err := os.Open(file)
		if err != nil {
			log.Println(err)
			os.Exit(127)
		}
		return fileHandle
	}
}

func waitForTextResult(latch int) {
	for latch > 0 {
		if !<-channel && !anyMatch {
			os.Exit(1)
		} else if anyMatch {
			os.Exit(0)
		}
		latch--
	}
}

func Execute() {
	rootCommad.PersistentFlags().BoolVar(&anyMatch, "any", false, "Exit with error code 0 when any word matches")
	rootCommad.PersistentFlags().StringVarP(&file, "file", "f", "", "Read from a given file instead of standard in")
	if err := rootCommad.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(127)
	}
}
