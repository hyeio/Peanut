// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package runtime

import (
	"github.com/clivern/peanut/core/model"
)

// DockerCompose type
type DockerCompose struct {
}

// NewDockerCompose creates a new instance
func NewDockerCompose() *DockerCompose {
	instance := new(DockerCompose)
	return instance
}

// Deploy deploys services
func (d *DockerCompose) Deploy(service *model.ServiceRecord) (*model.ServiceRecord, error) {
	return service, nil
}

// Destroy destroys services
func (d *DockerCompose) Destroy(service *model.ServiceRecord) (*model.ServiceRecord, error) {
	return service, nil
}
