package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

func traverseDir(roots []string) []os.FileInfo {
	var wg sync.WaitGroup
	var fileInfoCh = make(chan os.FileInfo)
	var filesInfo = make([]os.FileInfo, 0, filesAmount)

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileInfoCh)
	}

	go func() {
		wg.Wait()
		close(fileInfoCh)
	}()

loop:
	for {
		select {
		case <-done:
			for range fileInfoCh {
			}
			return nil
		case finfo, ok := <-fileInfoCh:
			if !ok {
				break loop
			}
			log.Printf("file %s added", finfo.Name())
			// filesInfo = append(filesInfo, finfo)
			// let's suppose we'd like to push every found file, not batch, emulating some real system
			handlerHTTP(finfo)
		}
	}
	return filesInfo
}

func walkDir(dir string, wg *sync.WaitGroup, fileInfoCh chan<- os.FileInfo) {
	defer wg.Done()
	if isDone() {
		return
	}
	entries, err := dirEntries(dir)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		// name := entry.Name()
		// if entry.IsDir() && (name != "." && name != "..") {
		if entry.IsDir() {
			log.Printf("directory: %s", entry.Name())
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, fileInfoCh)
		} else if entry.Size() > 0 {
			log.Printf("regular file: %s", entry.Name())
			fileInfoCh <- entry
		}
	}
}

func dirEntries(dir string) ([]os.FileInfo, error) {
	// blocking semaphore with 10 flows
	select {
	case semaphore <- struct{}{}:
	case <-done:
		return nil, nil
	}
	defer func() {
		<-semaphore
	}()

	file, err := os.Open(dir)
	defer file.Close()
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	files, err := file.Readdir(0)
	if err != nil {
		log.Printf("error readdir: %v", err)
	}
	return files, nil
}

func isDone() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
