package pages

import (
	"github.com/alexisbcz/panache/internal/views/layouts"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SignUpPage() g.Node {
	return layouts.AuthLayout(layouts.AuthLayoutProps{
		Title: "Sign Up",
	}, html.Div(html.Class("w-full max-w-md mx-auto"),
		html.H1(html.Class("text-3xl font-bold text-center mb-8"), g.Text("Create your account")),
		html.Form(html.Action("/auth/sign_up"), html.Method("POST"), html.Class("space-y-6"),
			// Email field
			html.Div(html.Class("space-y-2"),
				html.Label(html.For("email"), html.Class("block text-sm font-medium text-gray-700"), g.Text("Email")),
				html.Input(html.Type("email"), html.ID("email"), html.Name("email"),
					html.Required(),
					html.Placeholder("john.doe@example.com"),
					html.Class("block w-full h-10 px-3 py-1.5 rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"),
				),
			),
			// Password field
			html.Div(html.Class("space-y-2"),
				html.Label(html.For("password"), html.Class("block text-sm font-medium text-gray-700"), g.Text("Password")),
				html.Input(html.Type("password"), html.ID("password"), html.Name("password"),
					html.Required(),
					html.Class("block w-full h-10 px-3 py-1.5 rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"),
					html.Placeholder("••••••••••••••••"),
				),
			),
			// Submit button
			html.Button(html.Type("submit"),
				html.Class("w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"),
				g.Text("Sign up"),
			),
		),
	))
}
