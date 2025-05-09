package pages

import (
	"github.com/alexisbcz/panache/internal/views/layouts"
	"github.com/docker/docker/api/types/container"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func ListContainersPage(summaries []container.Summary) g.Node {
	return layouts.BaseLayout(layouts.BaseLayoutProps{
		Title: "Containers",
	}, html.Div(html.Class("w-full max-w-md mx-auto"),
		html.H1(html.Class("text-3xl font-bold text-center mb-8"), g.Text("Containers")),
		html.Div(
			g.Map(summaries, func(s container.Summary) g.Node {
				return html.A(
				// g.Text(s.HostConfig.Annotations[]),
				)
			}),
		),
	))
}
