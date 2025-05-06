package layouts

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type AuthLayoutProps struct {
	Title string
}

func AuthLayout(props AuthLayoutProps, children ...g.Node) g.Node {
	return html.Doctype(html.HTML(
		html.Lang("en"),
		html.Head(
			html.TitleEl(g.Text(props.Title)),
			html.Meta(html.Charset("utf-8")),
			html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
			html.Link(html.Rel("stylesheet"), html.Href("/styles.css")),
		),
		html.Body(
			g.Group(children),
		),
	))
}
