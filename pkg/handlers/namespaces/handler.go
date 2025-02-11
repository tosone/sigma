// Copyright 2023 sigma
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

package namespaces

import (
	"path"
	"reflect"

	"github.com/labstack/echo/v4"

	"github.com/go-sigma/sigma/pkg/auth"
	"github.com/go-sigma/sigma/pkg/consts"
	"github.com/go-sigma/sigma/pkg/dal/dao"
	"github.com/go-sigma/sigma/pkg/handlers"
	"github.com/go-sigma/sigma/pkg/middlewares"
	"github.com/go-sigma/sigma/pkg/modules/workq"
	"github.com/go-sigma/sigma/pkg/modules/workq/definition"
	"github.com/go-sigma/sigma/pkg/utils"
)

// Handler is the interface for the namespace handlers
type Handler interface {
	// PostNamespace handles the post namespace request
	PostNamespace(c echo.Context) error
	// ListNamespaces handles the list namespace request
	ListNamespaces(c echo.Context) error
	// GetNamespace handles the get namespace request
	GetNamespace(c echo.Context) error
	// DeleteNamespace handles the delete namespace request
	DeleteNamespace(c echo.Context) error
	// PutNamespace handles the put namespace request
	PutNamespace(c echo.Context) error
	// HotNamespace handles the hot namespace request
	HotNamespace(c echo.Context) error

	// AddNamespaceMember handles the add namespace member request
	AddNamespaceMember(c echo.Context) error
	// UpdateNamespaceMember handles the update namespace member request
	UpdateNamespaceMember(c echo.Context) error
	// DeleteNamespaceMember handles the delete namespace member request
	DeleteNamespaceMember(c echo.Context) error
	// ListNamespaceMembers handles the list namespace members request
	ListNamespaceMembers(c echo.Context) error
	// GetNamespaceMemberSelf handles the get self namespace member request
	GetNamespaceMemberSelf(c echo.Context) error
}

var _ Handler = &handler{}

type handler struct {
	authServiceFactory            auth.AuthServiceFactory
	auditServiceFactory           dao.AuditServiceFactory
	namespaceServiceFactory       dao.NamespaceServiceFactory
	namespaceMemberServiceFactory dao.NamespaceMemberServiceFactory
	repositoryServiceFactory      dao.RepositoryServiceFactory
	tagServiceFactory             dao.TagServiceFactory
	artifactServiceFactory        dao.ArtifactServiceFactory

	producerClient definition.WorkQueueProducer
}

type inject struct {
	authServiceFactory            auth.AuthServiceFactory
	auditServiceFactory           dao.AuditServiceFactory
	namespaceServiceFactory       dao.NamespaceServiceFactory
	namespaceMemberServiceFactory dao.NamespaceMemberServiceFactory
	repositoryServiceFactory      dao.RepositoryServiceFactory
	tagServiceFactory             dao.TagServiceFactory
	artifactServiceFactory        dao.ArtifactServiceFactory

	producerClient definition.WorkQueueProducer
}

// handlerNew creates a new instance of the distribution handlers
func handlerNew(injects ...inject) Handler {
	authServiceFactory := auth.NewAuthServiceFactory()
	auditServiceFactory := dao.NewAuditServiceFactory()
	namespaceServiceFactory := dao.NewNamespaceServiceFactory()
	namespaceMemberServiceFactory := dao.NewNamespaceMemberServiceFactory()
	repositoryServiceFactory := dao.NewRepositoryServiceFactory()
	tagServiceFactory := dao.NewTagServiceFactory()
	artifactServiceFactory := dao.NewArtifactServiceFactory()
	producerClient := workq.ProducerClient
	if len(injects) > 0 {
		ij := injects[0]
		if ij.authServiceFactory != nil {
			authServiceFactory = ij.authServiceFactory
		}
		if ij.auditServiceFactory != nil {
			auditServiceFactory = ij.auditServiceFactory
		}
		if ij.namespaceServiceFactory != nil {
			namespaceServiceFactory = ij.namespaceServiceFactory
		}
		if ij.namespaceMemberServiceFactory != nil {
			namespaceMemberServiceFactory = ij.namespaceMemberServiceFactory
		}
		if ij.repositoryServiceFactory != nil {
			repositoryServiceFactory = ij.repositoryServiceFactory
		}
		if ij.tagServiceFactory != nil {
			tagServiceFactory = ij.tagServiceFactory
		}
		if ij.artifactServiceFactory != nil {
			artifactServiceFactory = ij.artifactServiceFactory
		}
		if ij.producerClient != nil {
			producerClient = ij.producerClient
		}
	}
	return &handler{
		authServiceFactory:            authServiceFactory,
		auditServiceFactory:           auditServiceFactory,
		namespaceServiceFactory:       namespaceServiceFactory,
		namespaceMemberServiceFactory: namespaceMemberServiceFactory,
		repositoryServiceFactory:      repositoryServiceFactory,
		tagServiceFactory:             tagServiceFactory,
		artifactServiceFactory:        artifactServiceFactory,

		producerClient: producerClient,
	}
}

type factory struct{}

// Initialize initializes the namespace handlers
func (f factory) Initialize(e *echo.Echo) error {
	namespaceGroupWithoutAuth := e.Group(consts.APIV1 + "/namespaces")
	namespaceGroup := e.Group(consts.APIV1+"/namespaces", middlewares.AuthWithConfig(middlewares.AuthConfig{}))

	namespaceHandler := handlerNew()

	namespaceGroupWithoutAuth.GET("/", namespaceHandler.ListNamespaces)
	namespaceGroup.GET("/:id", namespaceHandler.GetNamespace)
	namespaceGroup.POST("/", namespaceHandler.PostNamespace)
	namespaceGroup.PUT("/:id", namespaceHandler.PutNamespace)
	namespaceGroup.DELETE("/:id", namespaceHandler.DeleteNamespace)
	namespaceGroup.GET("/hot", namespaceHandler.HotNamespace)

	namespaceGroup.GET("/:namespace_id/members/", namespaceHandler.ListNamespaceMembers)
	namespaceGroup.GET("/:namespace_id/members/self", namespaceHandler.GetNamespaceMemberSelf)
	namespaceGroup.POST("/:namespace_id/members/", namespaceHandler.AddNamespaceMember)
	namespaceGroup.PUT("/:namespace_id/members/:user_id", namespaceHandler.UpdateNamespaceMember)
	namespaceGroup.DELETE("/:namespace_id/members/:user_id", namespaceHandler.DeleteNamespaceMember)

	return nil
}

func init() {
	utils.PanicIf(handlers.RegisterRouterFactory(path.Base(reflect.TypeOf(factory{}).PkgPath()), &factory{}))
}
