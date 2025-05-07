package pages

import (
	"github.com/alexisbcz/panache/internal/views/layouts"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SignInPage() g.Node {
	return layouts.AuthLayout(layouts.AuthLayoutProps{
		Title: "Reset Password",
	}, html.Div(
		html.H1(html.Class("text-red-600"), g.Text("Hello, Refresher")),
	))
}
