package state

import (
	"encoding/json"
	"io/ioutil"
)

type VolumeMount struct {
	HostPath   string `json:"hostPath,omitempty"`
	TargetPath string `json:"targetPath,omitempty"`
}

type ProjectState struct {
	Image        string        `json:"image,omitempty"`
	DockerFile   string        `json:"dockerfile,omitempty"`
	ContainerId  string        `json:"containerId,omitempty"`
	VolumeMounts []VolumeMount `json:"volumeMounts,omitempty"`
	LsPort       int16         `json:"languageServerPort,omitempty"`
}

func Load() (*ProjectState, error) {
	p := &ProjectState{}
	file, err := ioutil.ReadFile("./.dcont/state.json")
	if err != nil {
		return nil, err
	}
	return p, json.Unmarshal(file, p) // yeah don't scream I'm using a pointer here
}

func (p ProjectState) Save() error {
	content, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("./.dcont/state.json", content, 0644)
}


func New() *ProjectState {
	return &ProjectState{}
}

func (p *ProjectState) SetPort(port int16) *ProjectState {
	p.LsPort = port
	return p
}

func (p *ProjectState) SetVolumes(volumes []VolumeMount) *ProjectState {
	p.VolumeMounts = volumes
	return p
}

func (p *ProjectState) SetImage(image string) *ProjectState {
	p.Image = image
	return p
}

func (p *ProjectState) SetContainerId(id string) *ProjectState {
	p.ContainerId = id
	return p
}
