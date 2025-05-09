package ui

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Field(children ...g.Node) g.Node {
	return html.Div(html.Class("flex flex-col space-y-2"),
		g.Group(children),
	)
}
