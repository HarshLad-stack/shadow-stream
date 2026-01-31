package steg

func EmbedBitsInStream(buf []byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = buf[i] & 0xFE
	}
}

// HideBit hides a single bit (0 or 1) inside a byte
func HideBit(pixelByte byte, bit byte) byte {
	pixelByte = pixelByte & 0xFE
	if bit == 1 {
		pixelByte = pixelByte | 0x01
	}
	return pixelByte
}

// ReadBit reads the hidden bit from a byte
func ReadBit(pixelByte byte) byte {
	return pixelByte & 0x01
}

// HideByte hides 1 byte (8 bits) into 8 pixel bytes
func HideByte(pixels []byte, secret byte) {
	for i := 0; i < 8; i++ {
		bit := (secret >> (7 - i)) & 1
		pixels[i] = HideBit(pixels[i], bit)
	}
}

// ReadByte extracts 1 byte from 8 pixel bytes
func ReadByte(pixels []byte) byte {
	var result byte = 0
	for i := 0; i < 8; i++ {
		bit := ReadBit(pixels[i])
		result |= bit << (7 - i)
	}
	return result
}

// HideBytes hides a byte slice into pixel bytes
func HideBytes(pixels []byte, data []byte) {
	for i := 0; i < len(data); i++ {
		start := i * 8
		HideByte(pixels[start:start+8], data[i])
	}
}

// ReadBytes extracts a byte slice of given length
func ReadBytes(pixels []byte, length int) []byte {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		start := i * 8
		result[i] = ReadByte(pixels[start : start+8])
	}
	return result
}
