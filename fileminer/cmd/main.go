package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var path string

func main() {
	rootCmd := &cobra.Command{
		Use:   "start",
		Short: "start traversing",
		Run:   cmdHandler,
	}
	rootCmd.PersistentFlags().StringVar(&path, "path", "/", "Traverse from path")
	rootCmd.Execute()
}

func cmdHandler(cmd *cobra.Command, args []string) {
	traverseDir(path)
	err := os.Chdir(path)
	if err != nil {
		cmd.Println(fmt.Sprintf("error occured %v", err))
	}

}

func traverseDir(roots []string) {
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()
}

func traverseDir1() {
	chanStrm := make(chan chan struct{}, 10)
	//chanStrm
	bridge := func(done <-chan struct{}, chanStream <-chan <-chan struct{}) <-chan struct{} {
		valStream := make(chan struct{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan struct{}
				select {
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valStream <- val:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}
}
