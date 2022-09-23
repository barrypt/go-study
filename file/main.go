package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// 创建文件
	f, err := os.Create("asong.txt")
	if err != nil {
		log.Fatalf("create file failed err=%s\n", err)
	}
	// 获取文件信息
	fileInfo, err := f.Stat()
	if err != nil {
		log.Fatalf("get file info failed err=%s\n", err)
	}

	log.Printf("File Name is %s\n", fileInfo.Name())
	log.Printf("File Permissions is %s\n", fileInfo.Mode())
	log.Printf("File ModTime is %s\n", fileInfo.ModTime())

	// 改变文件权限
	err = f.Chmod(0777)
	if err != nil {
		log.Fatalf("chmod file failed err=%s\n", err)
	}

	// 改变拥有者
	err = f.Chown(os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatalf("chown file failed err=%s\n", err)
	}

	// 再次获取文件信息 验证改变是否正确
	fileInfo, err = f.Stat()
	if err != nil {
		log.Fatalf("get file info second failed err=%s\n", err)
	}
	log.Printf("File change Permissions is %s\n", fileInfo.Mode())

	// 关闭文件
	err = f.Close()
	if err != nil {
		log.Fatalf("close file failed err=%s\n", err)
	}

	// 删除文件
	err = os.Remove("asong.txt")
	if err != nil {
		log.Fatalf("remove file failed err=%s\n", err)
	}
}

func WriteAll(filename string) error {
	err := os.WriteFile("asong.txt", []byte("Hi asong\n"), 0666)
	if err != nil {
		return err
	}
	return nil
}

// 直接操作IO
func WriteLine(filename string) error {
	data := []string{
		"asong",
		"test",
		"123",
	}
	f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	for _, line := range data {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	f.Close()
	return nil
}

// 使用缓存区写入
func WriteLine2(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	for i := 0; i < 2; i++ {
		// 写字符串到buffer
		bytesWritten, err := bufferedWriter.WriteString(
			"asong真帅\n",
		)
		if err != nil {
			return err
		}
		log.Printf("Bytes written: %d\n", bytesWritten)
	}
	// 写内存buffer到硬盘
	err = bufferedWriter.Flush()
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

///某些场景我们想根据给定的偏移量写入数据，可以使用os中的writeAt方法，例子如下：
func WriteAt(filename string) error {
	data := []byte{
		0x41, // A
		0x73, // s
		0x20, // space
		0x20, // space
		0x67, // g
	}
	f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	replaceSplace := []byte{
		0x6F, // o
		0x6E, // n
	}
	_, err = f.WriteAt(replaceSplace, 2)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

///os库中的方法对文件都是直接的IO操作，频繁的IO操作会增加CPU的中断频率，所以我们可以使用内存缓存区来减少IO操作，在写字节到硬盘前使用内存缓存，当内存缓存区的容量到达一定数值时在写内存数据buffer到硬盘，bufio就是这样示一个库，来个例子我们看一下怎么使用：
func WriteBuffer(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字符串到buffer
	bytesWritten, err := bufferedWriter.WriteString(
		"asong真帅\n",
	)
	if err != nil {
		return err
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		return err
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
	// 写内存buffer到硬盘
	err = bufferedWriter.Flush()
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

func ReadAll(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	log.Printf("read %s content is %s", filename, data)
	return nil
}

func ReadAll2(filename string) error {
	file, err := os.Open("asong.txt")
	if err != nil {
		return err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("read %s content is %s\n", filename, content)
	}

	file.Close()
	return nil

}

///os库中提供了Read方法是按照字节长度读取，如果我们想要按行读取文件需要配合bufio一起使用，bufio中提供了三种方法ReadLine、ReadBytes("\n")、ReadString("\n")可以按行读取数据，下面我使用ReadBytes("\n")来写个例子：
func ReadLine(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	bufferedReader := bufio.NewReader(file)
	for {
		// ReadLine is a low-level line-reading primitive. Most callers should use
		// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		lineBytes, err := bufferedReader.ReadBytes('\n')
		bufferedReader.ReadLine()
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		log.Printf("readline %s every line data is %s\n", filename, line)
	}
	file.Close()
	return nil
}

///有些场景我们想按照字节长度读取文件，这时我们可以如下方法：

//os库的Read方法

//os库配合bufio.NewReader调用Read方法

//os库配合io库的ReadFull、ReadAtLeast方法
// use bufio.NewReader
func ReadByte(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	// 创建 Reader
	r := bufio.NewReader(file)

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// use os
func ReadByte2(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// use os and io.ReadAtLeast
func ReadByte3(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := io.ReadAtLeast(file, buf, 0)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

///bufio包中提供了Scanner扫描器模块，它的主要作用是把数据流分割成一个个标记并除去它们之间的空格，他支持我们定制Split函数做为分隔函数，分隔符可以不是一个简单的字节或者字符，我们可以自定义分隔函数，在分隔函数实现分隔规则以及指针移动多少，返回什么数据，如果没有定制Split函数，那么就会使用默认ScanLines作为分隔函数，也就是使用换行作为分隔符，bufio中还提供了默认方法ScanRunes、ScanWrods，下面我们用SacnWrods方法写个例子，获取用空格分隔的文本：
func ReadScanner(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	// 可以定制Split函数做分隔函数
	// ScanWords 是scanner自带的分隔函数用来找空格分隔的文本字
	scanner.Split(bufio.ScanWords)
	for {
		success := scanner.Scan()
		if !success {
			// 出现错误或者EOF是返回Error
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
				break
			} else {
				return err
			}
		}
		// 得到数据，Bytes() 或者 Text()
		log.Printf("readScanner get data is %s", scanner.Text())
	}
	file.Close()
	return nil
}
