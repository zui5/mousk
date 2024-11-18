package ui

import (
	"bytes"
	"embed"
	"fmt"
	"os"
)

var (
	blueIconBytes  []byte
	whiteIconBytes []byte
)

// //go:embed ../../build/icons/iconwhite.png
// var assets2 embed.FS
//

//go:embed icons/iconwhite.png
var iconwhite embed.FS

func init() {
	// 打开PNG文件
	blueIconBytes = readIcon("./build/icons/iconcircle.ico")
	whiteIconBytes = readIcon("./build/icons/iconcircle.ico")
}

func readIcon(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	defer file.Close()

	// 读取文件内容并转换为 []byte
	icoBytes, err := fileToBytes(file)
	if err != nil {
		fmt.Println("Error converting file to bytes:", err)
		return nil
	}
	return icoBytes

	// // 解码图像
	// img, _, err := image.Decode(file)
	// if err != nil {
	// 	fmt.Println("Error decoding image:", err)
	// 	return nil
	// }

	// var buf bytes.Buffer
	// // 使用 PNG 编码器将 image.Image 写入缓冲区
	// err = png.Encode(&buf, img)
	// if err != nil {
	// 	return nil
	// }

	// // 返回缓冲区中的字节切片
	// return buf.Bytes()

	// // 将 image.Image 转换为 []byte
	// imgBytes, err := imageToBytes(img)
	// if err != nil {
	// 	fmt.Println("Error converting image to bytes:", err)
	// 	return
	// }

	// fileData, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return nil
	// }
	// return fileData

	// // 解码PNG文件
	// img, err := png.Decode(file)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return nil
	// }

	// // 创建一个字节缓冲区
	// var buf bytes.Buffer

	// // 将图像编码为PNG格式并写入字节缓冲区
	// err = png.Encode(&buf, img)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return nil
	// }

	// 将缓冲区中的内容转换为字节切片
	// byteData := buf.Bytes()
	// return byteData

}

func fileToBytes(file *os.File) ([]byte, error) {
	// 获取文件的所有内容
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GetTrayIcon(mode int) []byte {
	if mode == 0 {
		return getBlueTrayIcon()
	}
	return getWhiteTrayIcon()

}

func getBlueTrayIcon() []byte {
	return blueIconBytes
}

func getWhiteTrayIcon() []byte {
	return whiteIconBytes
}
