package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	TimeStamp time.Time
	Task      Task
}

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}

type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	Cmd           []string
	Image         string
	Memory        int64
	Disk          int64
	Env           []string
	RestartPolicy string
}

// type Docker struct {
// 	Client      *client.Client
// 	Config      Config
// 	ContainerId string
// }

// type DockerResult struct {
// 	Error       error
// 	Action      string
// 	ContainerId string
// 	Result      string
// }

// func (d *Docker) Run() DockerResult {
// 	ctx := context.Background()
// 	reader, err := d.Client.ImagePull(ctx, d.Config.Image, types.ImagePullOptions{})
// 	if err != nil {
// 		log.Printf("Error pulling image %s: %v\n", d.Config.Image, err)
// 		return DockerResult{Error: err}
// 	}
// 	io.Copy(os.Stdout, reader)
// }



