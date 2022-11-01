package network

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinary4BytesHeader(t *testing.T) {
	t.Run("Pack returns binary encoded length", func(t *testing.T) {
		header := NewBinary4BytesHeader()

		header.SetLength(319)
		var buf bytes.Buffer
		n, err := header.WriteTo(&buf)

		require.NoError(t, err)
		require.Equal(t, 4, n)

		// len 319 encoded in bytes
		require.Equal(t, []byte{0x00, 0x00, 0x01, 0x3F}, buf.Bytes())
	})

	t.Run("Read reads 4 bytes and decode length", func(t *testing.T) {
		header := NewBinary4BytesHeader()

		// len 319 encoded in binary
		packed := []byte{0x00, 0x00, 0x01, 0x3F}
		_, err := header.ReadFrom(bytes.NewReader(packed))

		require.NoError(t, err)
		require.Equal(t, 319, header.Length())
	})
}
