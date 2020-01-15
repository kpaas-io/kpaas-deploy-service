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

package task

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/kpaas-io/kpaas/pkg/deploy/action"
	"github.com/kpaas-io/kpaas/pkg/deploy/consts"
)

func init() {
	RegisterProcessor(TaskTypeDeployConfig, new(deployConfigProcessor))
}

// deployConfigProcessor implements the specific logic to config the deploy
type deployConfigProcessor struct {
}

// Spilt the task into multiple deploy config actions
func (p *deployConfigProcessor) SplitTask(t Task) error {
	if err := p.verifyTask(t); err != nil {
		logrus.Errorf("Invalid task: %s", err)
		return err
	}

	logger := logrus.WithFields(logrus.Fields{
		consts.LogFieldTask: t.GetName(),
	})

	logger.Debug("Start to split deploy config task")

	configTask := t.(*DeployConfigTask)

	// split task into actions: will create an action for every node, the action type
	// is ActionTypeDeployConfig
	actions := make([]action.Action, 0, len(configTask.NodeConfigs))
	for _, nodeCfg := range configTask.NodeConfigs {
		actionCfg := &action.DeployConfigActionConfig{
			NodeConfig:      nodeCfg,
			MasterNodes:     configTask.MasterNodes,
			ClusterConfig:   configTask.ClusterConfig,
			LogFileBasePath: configTask.LogFileDir,
		}
		act, err := action.NewDeployConfigAction(actionCfg)
		if err != nil {
			return err
		}
		actions = append(actions, act)
	}
	configTask.Actions = actions

	logger.Debugf("Finish to split deploy config task: %d actions", len(actions))

	return nil
}

// Verify if the task is valid.
func (p *deployConfigProcessor) verifyTask(t Task) error {
	if t == nil {
		return consts.ErrEmptyTask
	}

	_, ok := t.(*DeployConfigTask)
	if !ok {
		return fmt.Errorf("%s: %T", consts.MsgTaskTypeMismatched, t)
	}

	return nil
}
