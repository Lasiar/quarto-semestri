package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Lasiar/quarto-semestri/TOI/4/huffman"
	"github.com/Lasiar/quarto-semestri/TOI/4/lzw"
	"github.com/Lasiar/quarto-semestri/TOI/4/rle"
)

var path, huffmanStr string

func init() {
	flag.StringVar(&path, "file", "", "")
	flag.Parse()
}

func main() {
	if path == "" {
		fmt.Print("Напишите путь до файла -> ")
		if _, err := fmt.Scan(&path); err != nil {
			log.Fatalf("Ошибка чтения: %v", err)
		}
	}
	// Чтение файла
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}
	// Сжатие с использованием lzw
	lzwC := lzw.Compress(file)
	// Сжатие с использование rle
	rleC, err := rle.Encoding(file)
	if err != nil {
		log.Fatalf("Ошибка при работе с rle алгоритмом: %v", err)
	}
	fmt.Printf("lzw длина: %v\nrle длина: %v\nоригинал: %v\n", len(lzwC), len(rleC), len(file))
	fmt.Print("Напишите текст для хафмана: ->")
	if _, err := fmt.Scanln(&huffmanStr); err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}
	h := huffman.New(strings.NewReader(huffmanStr))
	h.Traverse()
}
