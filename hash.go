package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func hashFile(path string)pair{
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	hash := md5.New()

	if _ , err := io.Copy(hash,file); err != nil {

			log.Fatal(err)
	}
	return pair{
			hash: fmt.Sprintf("%x", hash.Sum(nil)),
			path: fmt.Sprintf("%x", path),
		}
}


func searchTree_non_con(dir string)(results, error) {
	hashes:= make(results)	

	err := filepath.Walk(dir,func(p string, fi os.FileInfo, err error) error {
		if fi.Mode().IsRegular() && fi.Size() > 0{
			h :=hashFile(p)
			hashes[h.hash] = append(hashes[h.hash], h.path)
		}
		return nil

	})
	return hashes, err
}