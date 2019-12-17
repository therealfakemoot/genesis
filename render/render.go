package render

import (
	"io"

	geo "github.com/therealfakemoot/genesis/geo"
)

// RenderFunc will allow for a plug-n-play system.
//
// Given a populated geo.Map and an io.Writer, a
// RenderFunc writes out whatever is appropriate for that
// representation of the map data.
type RenderFunc func(w io.Writer, m geo.Map)

/*
type Renderer interface {
	Render(w io.Writer, m geo.Map)
}
*/
