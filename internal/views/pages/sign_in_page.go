package pages

import (
	"github.com/alexisbcz/panache/internal/views/layouts"
	"github.com/alexisbcz/panache/internal/views/ui"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SignInPage() g.Node {
	return layouts.AuthLayout(layouts.AuthLayoutProps{
		Title: "Log in to your account",
	}, html.Div(html.Class("w-full max-w-md mx-auto"),
		html.H1(html.Class("text-3xl font-bold text-center mb-8"), g.Text("Log in to your account")),
		html.Form(html.Action("/sign-in"), html.Method("POST"), html.Class("space-y-6"),
			// Email field
			ui.Field(
				ui.Label(ui.LabelProps{For: "email", Text: "Email Address"}),
				ui.Input(ui.InputProps{Type: "email", Id: "email", Name: "email", Placeholder: "john.doe@example.com"}),
			),
			// Password field
			ui.Field(
				ui.Label(ui.LabelProps{For: "password", Text: "Password"}),
				ui.Input(ui.InputProps{Type: "password", Id: "password", Name: "password", Placeholder: "••••••••••••••••"}),
			),
			// Submit button
			ui.Button(ui.ButtonProps{Text: "Log In", Type: "submit"}),
		),
	))
}
