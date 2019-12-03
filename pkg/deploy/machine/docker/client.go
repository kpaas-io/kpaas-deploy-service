// Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"fmt"

	"github.com/docker/docker/client"
)

func NewTunneledClient(hostName string) (*client.Client, error) {
	localDockerSocketFile := composeLocalDockerSocketFile(hostName)
	opt := client.WithHost(localDockerSocketFile)

	cli, err := client.NewClientWithOpts(opt)
	if err != nil {
		return nil, fmt.Errorf("failed to create docker client, error: %v", err)
	}

	return cli, nil
}
