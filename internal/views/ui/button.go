package ui

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type ButtonProps struct {
	Text string
	Type string
}

func Button(props ButtonProps, children ...g.Node) g.Node {
	return html.Button(
		html.Type(props.Type),
		g.Text(props.Text),
		g.Group(children))
}
