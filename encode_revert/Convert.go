package main

import (
	"os"
	"fmt"
	"bufio"
	"code.google.com/p/mahonia"
	"io"
)

func FromGbkToUtf_8(fileName string)error {
	file, err := os.Open(fileName)
	if nil != err {
		fmt.Println(err.Error())
		return err
	}
	rd := bufio.NewReader(file)
	resultFileName := fmt.Sprintf("result_%s", fileName)
	fmt.Println("结果文件名", resultFileName)
	fileForWrite, err := os.Create(resultFileName)
	if nil != err {
		fmt.Println(err.Error())
		return err
	}
	w := bufio.NewWriter(fileForWrite)
	dec := mahonia.NewDecoder("gbk")
	for {
		line, err := rd.ReadString('\n')
		if ret, ok := dec.ConvertStringOK(line); ok {
			w.WriteString(ret)
		}
		if err != nil || io.EOF == err {
			break
		}
	}
	file.Close()
	w.Flush()
	fileForWrite.Close()
	return nil
}
