package util

import "golang.org/x/text/encoding/unicode"

func DetectAndConvertToString(data []byte) string {
	if len(data) < 2 {
		return string(data)
	}

	// UTF-8 BOM (EF BB BF)
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return string(data[3:])
	}

	// UTF-16 BE BOM (FE FF)
	if data[0] == 0xFE && data[1] == 0xFF {
		decoder := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder()
		result, _ := decoder.String(string(data[2:]))
		return result
	}

	// UTF-16 LE BOM (FF FE)
	if data[0] == 0xFF && data[1] == 0xFE {
		decoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
		result, _ := decoder.String(string(data[2:]))
		return result
	}

	// UTF-32 BE BOM (00 00 FE FF)
	if len(data) >= 4 && data[0] == 0x00 && data[1] == 0x00 && data[2] == 0xFE && data[3] == 0xFF {
		// UTF-32BE는 직접 변환 필요
		return convertUTF32BEToString(data[4:])
	}

	// UTF-32 LE BOM (FF FE 00 00)
	if len(data) >= 4 && data[0] == 0xFF && data[1] == 0xFE && data[2] == 0x00 && data[3] == 0x00 {
		return convertUTF32LEToString(data[4:])
	}

	return string(data)
}

func convertUTF32BEToString(data []byte) string {
	var result []rune
	for i := 0; i < len(data); i += 4 {
		if i+3 >= len(data) {
			break
		}
		r := rune(uint32(data[i])<<24 | uint32(data[i+1])<<16 | uint32(data[i+2])<<8 | uint32(data[i+3]))
		result = append(result, r)
	}
	return string(result)
}

func convertUTF32LEToString(data []byte) string {
	var result []rune
	for i := 0; i < len(data); i += 4 {
		if i+3 >= len(data) {
			break
		}
		r := rune(uint32(data[i+3])<<24 | uint32(data[i+2])<<16 | uint32(data[i+1])<<8 | uint32(data[i]))
		result = append(result, r)
	}
	return string(result)
}
