package main

import (
	"fmt"
//	"hash"
	"log"
	"os"
	"time"
)

type pair struct{
	hash, path string
}

type fileList []string
type results map[string]fileList

func main(){
	start := time.Now()
	log.Println("Starting...")
	if len(os.Args) < 2 {
		log.Panic("Missing args")

	}
	log.Println("Starting scan...")
	if hashes, err := searchTree(os.Args[1]); err == nil{
		for hash, files := range hashes{
			if (len(files) > 1){
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files{
					fmt.Println("  ", file)
				}
			}
		}
	}
	log.Printf("Finished in %v", time.Since(start))
}