package geo

import (
	"bytes"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	ikea "github.com/ikkerens/ikeapack"
)

// Encode does stuff
func (m *Map) Pack(w io.Writer) error {
	b := &bytes.Buffer{}
	if err := ikea.Pack(b, m); err != nil {
		return fmt.Errorf("map packing failed: %w", err)

	}

	return nil
}

func Unpack(r io.Reader) (Map, error) {
	var m Map

	l := logrus.New()
	if err := ikea.Unpack(r, &m); err != nil {
		return m, fmt.Errorf("map unpacking failed: %w", err)
	}
	m.Log = l
	return m, nil
}
