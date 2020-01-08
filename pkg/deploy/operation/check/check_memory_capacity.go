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

package check

import (
	"fmt"

	"github.com/kpaas-io/kpaas/pkg/deploy/command"
	"github.com/kpaas-io/kpaas/pkg/deploy/machine"
	"github.com/kpaas-io/kpaas/pkg/deploy/operation"
	pb "github.com/kpaas-io/kpaas/pkg/deploy/protos"
)

type CheckMemoryOperation struct {
	operation.BaseOperation
	CheckOperations
	Machine machine.IMachine
}

func (ckops *CheckMemoryOperation) CreateCommandAndRun(config *pb.NodeCheckConfig) (stdOut, stdErr []byte, err error) {
	ops := &CheckMemoryOperation{}

	m, err := machine.NewMachine(config.Node)
	if err != nil {
		return nil, nil, err
	}
	ckops.Machine = m

	// close ssh client if machine is not nil
	if ckops.Machine != nil {
		defer ckops.Machine.Close()
	}

	ops.AddCommands(command.NewShellCommand(m, "free", "-b | awk '/Mem/{print $2}'"))

	if len(ops.Commands) == 0 {
		return nil, nil, fmt.Errorf("check memory command is empty")
	}

	// run commands
	stdOut, stdErr, err = ops.Do()

	return
}

// check if memory capacity satisfied with minimal requirement
func CheckMemoryCapacity(comparedMemory string, desiredMemory float64) error {
	err := operation.CheckEntity(comparedMemory, desiredMemory)
	if err != nil {
		return err
	}

	return nil
}
