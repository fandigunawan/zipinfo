/*
Tool to sort file size in zip from small to big
by github.com/fandigunawan

Background :
I got a problem while delivering package to my customer. There are a lot of files
that I must submit, however I need to check from the final zip output which file
that make the largest contributor to the zip package.

Usage : path_to_zip_file

The tool will skip directory type file
*/
package main

import (
	"archive/zip"
	"fmt"
	"log"
	"sort"
	"os"
)
type KeyValue struct {
	Key string
	Value int64
}
/* These interface is used together with sort function */
type KeyValueList []KeyValue
// Swap values
func (kvl KeyValueList) Swap(i, j int) {kvl[i], kvl[j] =  kvl[j], kvl[i]}
// Len of list
func (kvl KeyValueList) Len() int { return len(kvl)}
// Custom comparer
func (kvl KeyValueList) Less(i, j int) bool {return kvl[i].Value < kvl[j].Value}

func sortByValue(data map[string]int64) KeyValueList {
	p := make(KeyValueList, len(data))
	i := 0
	// Copy data to KeyValueList storage
	for k, v := range data {
		p[i] = KeyValue{k,v}
		i++
	}
	// Sort the KeyValueList based on Value on KeyValue
	sort.Sort(p)
	return p
	
}
func main() {
	if(len(os.Args) != 2) {
		fmt.Println("Invalid parameter")
		fmt.Println("Usage : file_path")
		return;
	}
	
	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err){
		fmt.Println("File %s does not exist", os.Args[1])
		return
	}
	r, err := zip.OpenReader(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	
	var zipInfo = make(map[string]int64)
	
	// Copy the zip file info to map
	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			zipInfo[f.Name] = f.FileInfo().Size()	
		}
	}
	// Sort the value
	kvl := sortByValue(zipInfo)
	// Display it
	for _, kv := range kvl {
		fmt.Printf("%d bytes : %s", kv.Value, kv.Key)
		fmt.Println();
	}
}
