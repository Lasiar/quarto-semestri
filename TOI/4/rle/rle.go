package rle

import "bytes"

func Encoding(data []byte) ([]byte, error) {
	var encData bytes.Buffer
	var buf byte
	var size int
	for i, d := range data {
		if len(data) == i+1 {
			break
		}
		if d == 255 {
			encData.Write([]byte{255, 255})
			continue
		}
		if size >= 253 {
			encData.Write(append([]byte{255, byte(size + 1)}, buf))
			size = 0
		}
		if d == data[i+1] {
			buf = d
			size++
			continue
		}
		if size != 0 {
			encData.Write(append([]byte{255, byte(size + 1)}, buf))
			size = 0
			continue
		}
		encData.Write([]byte{d})
	}
	if size != 0 {
		encData.Write(append([]byte{255, byte(size)}, buf))
	} else {
		encData.Write([]byte{data[len(data)-1]})
	}
	return encData.Bytes(), nil
}

func Decoding(data []byte) ([]byte, error) {
	var decData bytes.Buffer
	padding := true
	for i := 0; i < len(data)-1; i++ {
		if data[i] == 255 && data[i+1] == 255 {
			decData.Write([]byte{255})
			i++
			continue
		}
		if data[i] == 255 {
			if len(data)-3 == i {
				padding = false
			}
			for j := 0; j < int(data[i+1]); j++ {
				decData.Write([]byte{data[i+2]})
			}
			i += 2
			continue
		}
		decData.Write([]byte{data[i]})
	}
	if padding {
		decData.Write([]byte{data[len(data)-1]})
	}
	return decData.Bytes(), nil
}
