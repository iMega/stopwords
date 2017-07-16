/*
TF - простой способ оценить важность термина для какого-либо текста относительно остальных.
Принцип: если слово встречается часто в исследуемом тексте, при этом встречаясь редко во всех остальных текстах,
то это слово имеет больший вес для первого.

Для чего нужно
Составить список стоп-слов
*/
package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/leveldb"
	"github.com/golang/leveldb/db"
	"github.com/imega/stopwords/filters"
)

// TF struct
type TF struct {
	Rune      string
	Frequency float64
	Count     int64
}

type tfdb struct {
	Frequency float64
	Count     int64
}

func main() {
	var docCount = 0

	filename := flag.String("file", "", "Parse file")
	flag.Parse()

	bdata, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	text := string(bdata)
	text = strings.ToLower(text)

	bracket := filters.NewBrackets()
	text = bracket.Filter(text)

	controlChar := filters.NewControlCharacters()
	text = controlChar.Filter(text)

	punktuation := filters.NewPunctuation()
	text = punktuation.Filter(text)

	doubles := filters.NewDoublecates()
	text = doubles.Filter(text)

	text = strings.TrimSpace(text)

	bag := strings.Split(text, " ")

	words := make(map[string]int)
	for i := range bag {
		words[bag[i]]++
	}

	fmt.Println("не", words["не"])
	fmt.Println("порох", words["порох"])
	fmt.Println("пятнадцатилетние", words["пятнадцатилетние	"])

	tfs := []TF{}

	for i, v := range words {
		tfs = append(tfs, TF{
			Rune:      i,
			Frequency: .5 * (1 + float64(v)/float64(len(bag))),
			Count:     int64(v),
		})
	}

	opts := &db.Options{}
	ldb, err := leveldb.Open("stop-words.leveldb", opts)
	if err != nil {
		fmt.Errorf("Could not open db, %s", err)
	}

	readOpts := &db.ReadOptions{}
	writeOpts := &db.WriteOptions{}
	data, err := ldb.Get([]byte("#doc_count"), readOpts)
	if err != nil {
		fmt.Errorf("Could not read from db, %s", err)
	}
	if len(data) > 0 {
		docCount = int(binary.BigEndian.Uint64(data))
	}

	docCount++
	fmt.Printf("Document: %d\n", docCount)

	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(docCount))
	err = ldb.Set([]byte("#doc_count"), bs, writeOpts)
	if err != nil {
		fmt.Errorf("Could not write to db, %s", err)
	}

	buf := &bytes.Buffer{}

	for _, i := range tfs {
		data, err := ldb.Get([]byte(i.Rune), readOpts)
		if err != nil {
			fmt.Errorf("Could not read from db, %s", err)
		}
		t := tfdb{
			Count: 1,
		}

		if len(data) > 0 {
			buf.Reset()
			buf.Write(data)
			err = binary.Read(buf, binary.BigEndian, &t)
			if err != nil {
				fmt.Errorf("Could not read from binary, %s", err)
			}
			t.Count++
		}

		tt := struct {
			Frequency float64
			Count     int64
		}{
			Frequency: (t.Frequency + i.Frequency) / float64(t.Count),
			Count:     t.Count,
		}

		fmt.Println(i.Rune, tt)
		buf.Reset()
		err = binary.Write(
			buf,
			binary.BigEndian,
			tt,
		)
		if err != nil {
			fmt.Errorf("Could not write to binary, %s", err)
		}
		err = ldb.Set([]byte(i.Rune), buf.Bytes(), writeOpts)
		if err != nil {
			fmt.Errorf("Could not write to db, %s", err)
		}
	}

	err = ldb.Close()
	if err != nil {
		fmt.Errorf("Could not close db, %s", err)
	}
}
