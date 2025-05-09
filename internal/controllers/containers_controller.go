package controllers

import (
	"net/http"

	"github.com/alexisbcz/panache/internal/views/pages"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainersController struct {
}

func NewContainersController() *ContainersController {
	return &ContainersController{}
}

func (c *ContainersController) Index(w http.ResponseWriter, r *http.Request) error {
	// Docker client setup
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	summaries, err := cli.ContainerList(r.Context(), container.ListOptions{})
	if err != nil {
		return err
	}

	return pages.ListContainersPage(summaries).Render(w)
}
