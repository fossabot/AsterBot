package main //docker controller

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ContainerStatus(id string) bool {
	status := false
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.ID[:len(id)] == id {
			status = true
		}
	}

	return status
}

func ContainerStart(id string) bool {
	status := true
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStart(ctx, id, types.ContainerStartOptions{})
	if err != nil {
		status = false
		//panic(err)
	}

	return status
}

func ContainerStop(id string) bool {
	status := true
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStop(ctx, id, nil)
	if err != nil {
		status = false
		//panic(err)
	}

	return status
}

func ContainerRestart(id string) bool {
	status := true
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	err = cli.ContainerRestart(ctx, id, nil)
	if err != nil {
		status = false
		//panic(err)
	}
	return status
}
