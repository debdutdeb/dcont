package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/debdutdeb/devcontainer-lite/actions"
	"github.com/debdutdeb/devcontainer-lite/state"
	"github.com/spf13/cobra"
)

// maybe I should use viper
func environmentCmd() *cobra.Command {
	var (
		port    int16
		mounts  []string
		image   string
		destroy bool
	)
	var environmentCommand = &cobra.Command{
		Use:   "environment",
		Short: "Start a development environment (container)",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := actions.New()
			if err != nil {
				return err
			}
			pstate, err := state.Load()
			if err != nil {
				//
			}
			if destroy {
				if pstate == nil {
					log.Println("no state file found, skipping ..")
					return nil
				}
				err := api.Destroy(pstate.ContainerId)
				if err != nil {
					return err
				}
				return state.Destroy()
			}
			switch {
			case actions.IsGolang():
				log.Println("starting go environment")
				id, err := api.Start("golang", image, port, mounts)
				if err != nil {
					return err
				}
				return state.New().
					SetImage(image).
					SetVolumes(parseMounts(mounts)).
					SetPort(port).
					SetContainerId(id).
					Save()
			default:
				return errors.New("unknown project detected")
			}
		},
	}
	environmentCommand.Flags().Int16Var(&port, "port", 3001, "port the language server wil wait for connections on")
	environmentCommand.Flags().StringArrayVar(&mounts, "mounts", []string{}, "volumes to share")
	environmentCommand.Flags().StringVar(&image, "image", "", "image to use")
	environmentCommand.Flags().BoolVar(&destroy, "destroy", false, "destroy current environment")
	return environmentCommand
}

// [[maybe_unused]]
func parseMounts(mountArgs []string) []state.VolumeMount {
	var mounts []state.VolumeMount
	for _, mountStr := range mountArgs {
		// hostPath:targetPath
		paths := strings.Split(mountStr, ":")
		mounts = append(mounts, state.VolumeMount{
			HostPath:   paths[0],
			TargetPath: paths[1],
		})
	}
	return mounts
}

func startGolang() {
}
