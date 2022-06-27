package prefix

import (
	"encoding/binary"
	"fmt"
	"strings"
)

var Binary = Prefixers{
	Fixed: &binaryFixedPrefixer{},
	L:     &binaryVarPrefixer{1},
	LL:    &binaryVarPrefixer{2},
	LLL:   &binaryVarPrefixer{3},
	LLLL:  &binaryVarPrefixer{4},
}

type binaryVarPrefixer struct {
	Digits int
}

func (p *binaryVarPrefixer) EncodeLength(maxLen, dataLen int) ([]byte, error) {

	if dataLen > maxLen {
		return nil, fmt.Errorf("field length: %d is larger than maximum: %d", dataLen, maxLen)
	}

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(dataLen))

	for _, v := range buf[:8-p.Digits] {
		if v != 0 {
			return nil, fmt.Errorf("number of digits in length: %d exceeds: %d", dataLen, p.Digits)
		}
	}

	return buf[8-p.Digits:], nil
}

func (p *binaryVarPrefixer) DecodeLength(maxLen int, data []byte) (int, int, error) {
	length := p.Digits
	if len(data) < length {
		return 0, 0, fmt.Errorf("length mismatch: want to read %d bytes, get only %d", length, len(data))
	}

	buf := make([]byte, 8)
	copy(buf[8-length:], data)
	dataLen := binary.BigEndian.Uint64(buf)

	if int(dataLen) > maxLen {
		return 0, 0, fmt.Errorf("data length %d is larger than maximum %d", dataLen, maxLen)
	}

	return int(dataLen), length, nil
}

func (p *binaryVarPrefixer) Inspect() string {
	return fmt.Sprintf("Bin.%s", strings.Repeat("L", p.Digits))
}

type binaryFixedPrefixer struct {
}

func (p *binaryFixedPrefixer) EncodeLength(fixLen, dataLen int) ([]byte, error) {
	if dataLen != fixLen {
		return nil, fmt.Errorf("field length: %d should be fixed: %d", dataLen, fixLen)
	}

	return []byte{}, nil
}

func (p *binaryFixedPrefixer) DecodeLength(fixLen int, data []byte) (int, int, error) {
	return fixLen, 0, nil
}

func (p *binaryFixedPrefixer) Inspect() string {
	return "Binary.Fixed"
}
