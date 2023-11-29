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

package distribution

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/go-sigma/sigma/pkg/consts"
	"github.com/go-sigma/sigma/pkg/dal"
	"github.com/go-sigma/sigma/pkg/dal/dao"
	"github.com/go-sigma/sigma/pkg/dal/models"
	"github.com/go-sigma/sigma/pkg/tests"
	"github.com/go-sigma/sigma/pkg/types/enums"
	"github.com/go-sigma/sigma/pkg/utils/ptr"
)

func TestListRepositories(t *testing.T) {
	assert.NoError(t, tests.Initialize(t))
	assert.NoError(t, tests.DB.Init())
	defer func() {
		conn, err := dal.DB.DB()
		assert.NoError(t, err)
		assert.NoError(t, conn.Close())
		assert.NoError(t, tests.DB.DeInit())
	}()

	ctx := context.Background()

	const (
		namespaceName  = "test"
		repositoryName = "test/busybox"
	)

	userObj := &models.User{Username: "post-namespace", Password: ptr.Of("test"), Email: ptr.Of("test@gmail.com")}
	assert.NoError(t, dao.NewUserServiceFactory().New().Create(ctx, userObj))
	namespaceObj := &models.Namespace{Name: namespaceName, Visibility: enums.VisibilityPrivate}
	assert.NoError(t, dao.NewNamespaceServiceFactory().New().Create(ctx, namespaceObj))
	repositoryObj := &models.Repository{Name: repositoryName, NamespaceID: namespaceObj.ID, Visibility: enums.VisibilityPrivate}
	assert.NoError(t, dao.NewRepositoryServiceFactory().New().Create(ctx, repositoryObj, dao.AutoCreateNamespace{UserID: userObj.ID}))
	artifactObj := &models.Artifact{RepositoryID: repositoryObj.ID, Digest: "sha256:1234567890", Size: 1234, ContentType: "application/octet-stream", Raw: []byte("test"), PushedAt: time.Now()}
	assert.NoError(t, dao.NewArtifactServiceFactory().New().Create(ctx, artifactObj))
	tagObj := &models.Tag{Name: "latest", RepositoryID: repositoryObj.ID, ArtifactID: artifactObj.ID, PushedAt: time.Now()}
	assert.NoError(t, dao.NewTagServiceFactory().New().Create(ctx, tagObj))

	req := httptest.NewRequest(http.MethodGet, "/v2/_catalog", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.Set(consts.ContextUser, userObj)
	assert.NoError(t, (&factory{}).Initialize(c))
	assert.Equal(t, http.StatusOK, c.Response().Status)
}
