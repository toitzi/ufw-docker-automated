package ufwhandler

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog/log"
)

func Sync(ctx *context.Context, createChannel chan *types.ContainerJSON, client *client.Client) {
	containers, err := client.ContainerList(*ctx, container.ListOptions{Filters: filters.NewArgs(filters.Arg("label", "UFW_MANAGED=TRUE"))})
	if err != nil {
		log.Error().Err(err).Msg("ufw-docker-automated: Couldn't retrieve existing containers.")
	}

	for _, c := range containers {
		cont, err := client.ContainerInspect(*ctx, c.ID)
		if err != nil {
			log.Error().Err(err).Msg("ufw-docker-automated: Couldn't inspect existing container.")
			continue
		}
		createChannel <- &cont
	}
}
