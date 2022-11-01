package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

var _ Header = (*Binary4Bytes)(nil)

type Binary4Bytes struct {
	Len uint32
}

func NewBinary4BytesHeader() *Binary4Bytes {
	return &Binary4Bytes{}
}

func (h *Binary4Bytes) SetLength(length int) error {
	if length > math.MaxUint32 {
		return fmt.Errorf("length %d exceeds max length for 4 bytes header %d", length, math.MaxUint32)
	}

	h.Len = uint32(length)

	return nil
}

func (h *Binary4Bytes) Length() int {
	return int(h.Len)
}

func (h *Binary4Bytes) WriteTo(w io.Writer) (int, error) {
	err := binary.Write(w, binary.BigEndian, h.Len)
	if err != nil {
		return 0, fmt.Errorf("writint uint32 into writer: %w", err)
	}

	return binary.Size(h.Len), nil
}

func (h *Binary4Bytes) ReadFrom(r io.Reader) (int, error) {
	err := binary.Read(r, binary.BigEndian, &h.Len)
	if err != nil {
		return 0, fmt.Errorf("reading uint32 from reader: %w", err)
	}

	return binary.Size(h.Len), nil
}
