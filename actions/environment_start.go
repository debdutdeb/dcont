package actions

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Api struct {
	client *client.Client
}

func New() (*Api, error) {
	api, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &Api{client: api}, nil
}

func (a Api) IsActive(containerId string) bool {
	iResp, err := a.client.ContainerInspect(context.TODO(), containerId)
	if err != nil && client.IsErrNotFound(err) {
		return false
	}
	if !iResp.State.Running {
		return false
	}
	return true
}

// FIXME
func (a Api) BuildEnvironment(dockerfile string) (string, error) {
	return "", nil
}

func (a Api) Start(lang, image string, port int16, mounts []string) (string, error) {
	opts := languageServerOpts.Get(lang)
	mounts = append(mounts, opts.mounts...)
	log.Printf("mounts: %#v\n", mounts)
	exposedPorts, portBinds, _ := nat.ParsePortSpecs([]string{fmt.Sprintf("127.0.0.1:%d:%d", port, port)})
	containerConfig := container.Config{
		Image:        image,
		ExposedPorts: exposedPorts,
	}
	hostConfig := container.HostConfig{
		AutoRemove:   true,
		Binds:        mounts,
		Privileged:   true,
		PortBindings: portBinds,
	}
	con, err := a.client.ContainerCreate(context.TODO(), &containerConfig, &hostConfig, nil, "test")
	if err != nil {
		return "", err
	}
	err = a.client.ContainerStart(context.TODO(), con.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}
	resp, err := a.client.ContainerExecCreate(context.TODO(), con.ID, types.ExecConfig{
		Cmd:    []string{"sh", "-c", fmt.Sprintf(opts.startUpCommand, port)},
		Detach: false,
		Tty:    true,
	})

	if err != nil {
		return "", nil
	}
	err = a.client.ContainerExecStart(context.TODO(), resp.ID, types.ExecStartCheck{
		Tty:    true,
		Detach: false,
	})
	if err != nil {
		return "", nil
	}
	x, err := a.client.ContainerExecInspect(context.TODO(), resp.ID)
	log.Printf("e: %#v\n", x)
	return con.ID, nil
}

func (a Api) Destroy(id string) error {
	return a.client.ContainerStop(context.TODO(), id, nil)
}
