package layouts

import (
	g "maragu.dev/gomponents"
)

type AuthLayoutProps struct {
	Title string
}

func AuthLayout(props AuthLayoutProps, children ...g.Node) g.Node {
	return BaseLayout(BaseLayoutProps{
		Title: props.Title,
	}, children...)
}
