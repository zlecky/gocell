package FileExam

import (
	"compress/gzip"
	"io"
	"os"
)

func CompressFile() {
	// 打开源文件
	srcFile, err := os.Open("./data/source.txt")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	// 创建压缩文件
	destFile, err := os.Create("./data/compressed.gz")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// 创建 gzip 压缩器
	gzipWriter := gzip.NewWriter(destFile)
	defer gzipWriter.Close()

	// 将源文件内容写入压缩器中
	_, err = io.Copy(gzipWriter, srcFile)
	if err != nil {
		panic(err)
	}
}
