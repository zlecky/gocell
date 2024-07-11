package main

import (
	CutWords "gocell/cutwords"
	FileExam "gocell/example"
)

func main() {
	FileExam.CompressFile()
	CutWords.CutWords()
}
