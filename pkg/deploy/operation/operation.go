// Copyright 2019 Shanghai JingDuo Information Technology co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operation

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/kpaas-io/kpaas/pkg/deploy/command"
	"github.com/kpaas-io/kpaas/pkg/deploy/utils"
)

type Operation interface {
	AddCommands(commands ...command.Command)
	Do() ([]byte, []byte, error)
	DoWithLogWriter(writer io.Writer) error
}

type BaseOperation struct {
	Commands []command.Command
}

func (op *BaseOperation) AddCommands(commands ...command.Command) {
	op.Commands = append(op.Commands, commands...)
}

func (op *BaseOperation) Do() (stdout, stderr []byte, err error) {
	for _, cmd := range op.Commands {
		stdout, stderr, err = cmd.Execute()
		if err != nil {
			err = fmt.Errorf("run cmd %v error: %v", cmd, err)
			return
		}
	}

	return
}

func (op *BaseOperation) DoWithLogWriter(writer io.Writer) error {
	for _, cmd := range op.Commands {
		startTime := time.Now()
		stdout, stderr, err := cmd.Execute()
		endTime := time.Now()
		// if writer nil, just write log to stdout
		if writer == nil {
			writer = os.Stdout
		}
		utils.WriteExecuteLog(writer, &utils.ExecuteLogItem{
			StartTime: startTime,
			EndTime:   endTime,
			Command:   cmd.GetCommand(),
			Stdout:    stdout,
			Stderr:    stderr,
			Err:       err,
		})
		if err != nil {
			err = fmt.Errorf("run cmd %v error: %w", cmd, err)
			return err
		}
	}
	return nil
}

func (op *BaseOperation) ResetCommands() {
	op.Commands = []command.Command{}
}
