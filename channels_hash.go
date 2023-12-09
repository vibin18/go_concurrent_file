package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func collectHashes(pairs <-chan pair, result chan<- results){
	hashes := make(results)

	for p := range pairs{
		hashes[p.hash] = append(hashes[p.hash],p.path)
	}
	result <- hashes
}

func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool){
	for path := range paths{
		pairs <- hashFile(path)
	}
	done <- true
}

func searchTree(dir string)(results,error) {
	worker := 2 * runtime.GOMAXPROCS(0)
	paths := make(chan string)
	pairs := make(chan pair,10)
	done := make(chan bool)
	result := make(chan results)

	for i:=0; i < worker; i++{
		go processFiles(paths, pairs,done)
	}

	go collectHashes(pairs,result)



	err := filepath.Walk(dir,func(p string, fi os.FileInfo, err error) error {
		if fi.Mode().IsRegular() && fi.Size() > 0{
			paths <- p
		}
		return nil

	})

	if err != nil {
		log.Fatal(err)
	}

	close(paths)

	for i:=0; i < worker; i++{
		<- done
	}

	close(pairs)

	hashes := <-result

	return hashes, err
}