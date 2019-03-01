package render

import (
	"io"

	geo "github.com/therealfakemoot/genesis/geo"
)

type RenderFunc func(w io.Writer, m geo.Map)

/*
type Renderer interface {
	Render(w io.Writer, m geo.Map)
}
*/
