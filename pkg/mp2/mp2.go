package mp2

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
)

type Map struct {
	*Header       // 0x0
	Tiles   Tiles // 0x1AC (428)

	// *Addons

	// Uniq uint32 // EOF - 0x4
	// EOF
}

func LoadMap(r io.Reader) (*Map, error) {
	header, err := LoadHeader(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse info")
	}

	tiles := make(Tiles, header.Width()*header.Height())
	if err := binary.Read(r, binary.LittleEndian, tiles); err != nil {
		return nil, errors.Wrap(err, "failed to parse tiles")
	}

	m := &Map{
		Header: header,
		Tiles:  tiles,
	}

	return m, nil
}
