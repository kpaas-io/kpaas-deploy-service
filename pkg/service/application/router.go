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

package application

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/kpaas-io/kpaas/pkg/service/api/v1/deploy"
	_ "github.com/kpaas-io/kpaas/pkg/service/swaggerdocs"
)

func (a *app) setRoutes() {
	a.httpHandler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := a.httpHandler.Group("/api/v1")
	wizardGroup := v1.Group("/deploy/wizard")

	wizardGroup.GET("/progresses", deploy.GetWizardProgress)
	wizardGroup.DELETE("/progresses", deploy.ClearWizard)

	wizardGroup.GET("/clusters", deploy.GetCluster)
	wizardGroup.POST("/clusters", deploy.SetCluster)

	wizardGroup.GET("/nodes", deploy.GetNodeList)
	wizardGroup.GET("/nodes/{ip}", deploy.GetNode)
	wizardGroup.POST("/nodes", deploy.AddNode)
	wizardGroup.PUT("/nodes/{ip}", deploy.UpdateNode)
	wizardGroup.DELETE("/nodes/{ip}", deploy.DeleteNode)

	wizardGroup.POST("/checks", deploy.CheckNodeList)
	wizardGroup.GET("/checks", deploy.GetCheckingNodeListResult)

	wizardGroup.POST("/deploys", deploy.Deploy)
	wizardGroup.GET("/deploys", deploy.GetDeployReport)

	wizardGroup.GET("/logs/{id}", deploy.DownloadLog)

	v1.POST("/ssh/tests", deploy.TestConnectNode)

	v1.POST("/ssh_certificates", deploy.AddSSHCertificate)
	v1.GET("/ssh_certificates", deploy.GetCertificateList)
}
