package ui

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type LabelProps struct {
	CommonProps
	Text string
	For  string
}

const LabelBaseClasses = "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"

func Label(props LabelProps, children ...g.Node) g.Node {
	return html.Label(
		html.Class(mergeClasses(LabelBaseClasses, props.Class)),
		html.For(props.For),
		g.Text(props.Text),
		g.Group(children),
	)
}
