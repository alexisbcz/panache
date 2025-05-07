package layouts

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

type BaseLayoutProps struct {
	Title string
}

func BaseLayout(props BaseLayoutProps, children ...g.Node) g.Node {
	return html.Doctype(html.HTML(
		html.Lang("en"),
		html.Head(
			html.TitleEl(g.Text(props.Title)),
			html.Meta(html.Charset("utf-8")),
			html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
			html.Link(html.Rel("stylesheet"), html.Href("/styles.css")),
			html.Script(html.Type("module"), html.Src("https://cdn.jsdelivr.net/gh/starfederation/datastar@v1.0.0-beta.11/bundles/datastar.js")),
		),
		html.Body(
			g.Attr("data-on-load", "@get('/hotreload', {retryInterval: 100})"),
			g.Group(children),
		),
	))
}
