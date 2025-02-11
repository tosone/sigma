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

package dao_test

import (
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/go-sigma/sigma/pkg/dal"
	"github.com/go-sigma/sigma/pkg/dal/dao"
	"github.com/go-sigma/sigma/pkg/dal/models"
	"github.com/go-sigma/sigma/pkg/dal/query"
	"github.com/go-sigma/sigma/pkg/logger"
	"github.com/go-sigma/sigma/pkg/tests"
	"github.com/go-sigma/sigma/pkg/types"
	"github.com/go-sigma/sigma/pkg/types/enums"
	"github.com/go-sigma/sigma/pkg/utils/ptr"
)

func TestNamespaceServiceFactory(t *testing.T) {
	f := dao.NewNamespaceServiceFactory()
	namespaceService := f.New()
	assert.NotNil(t, namespaceService)
	namespaceService = f.New(query.Q)
	assert.NotNil(t, namespaceService)
}

func TestNamespaceService(t *testing.T) {
	viper.SetDefault("log.level", "debug")
	logger.SetLevel("debug")
	err := tests.Initialize(t)
	assert.NoError(t, err)
	err = tests.DB.Init()
	assert.NoError(t, err)
	defer func() {
		conn, err := dal.DB.DB()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
		err = tests.DB.DeInit()
		assert.NoError(t, err)
	}()

	ctx := log.Logger.WithContext(context.Background())

	f := dao.NewNamespaceServiceFactory()
	userServiceFactory := dao.NewUserServiceFactory()
	err = query.Q.Transaction(func(tx *query.Query) error {
		userService := userServiceFactory.New(tx)
		userObj := &models.User{Username: "namespace-service", Password: ptr.Of("test"), Email: ptr.Of("test@gmail.com")}
		err = userService.Create(ctx, userObj)
		assert.NoError(t, err)

		namespaceService := f.New(tx)

		namespaceObj := &models.Namespace{
			Name:       "test",
			Visibility: enums.VisibilityPrivate,
		}
		err := namespaceService.Create(ctx, namespaceObj)
		assert.NoError(t, err)

		ns, err := namespaceService.Get(ctx, namespaceObj.ID)
		assert.NoError(t, err)
		assert.Equal(t, ns.ID, namespaceObj.ID)
		assert.Equal(t, ns.Name, namespaceObj.Name)

		ns1, err := namespaceService.GetByName(ctx, "test")
		assert.NoError(t, err)
		assert.Equal(t, ns1.ID, namespaceObj.ID)
		assert.Equal(t, ns1.Name, namespaceObj.Name)

		namespaceList, _, err := namespaceService.ListNamespace(ctx, ptr.Of("t"), types.Pagination{
			Limit: ptr.Of(int(100)),
			Page:  ptr.Of(int(0)),
		}, types.Sortable{})
		assert.NoError(t, err)
		assert.Equal(t, len(namespaceList), int(1))

		namespaceList, _, err = namespaceService.ListNamespace(ctx, ptr.Of("t"), types.Pagination{
			Limit: ptr.Of(int(100)),
			Page:  ptr.Of(int(0)),
		}, types.Sortable{
			Sort:   ptr.Of("created_at"),
			Method: ptr.Of(enums.SortMethodDesc),
		})
		assert.NoError(t, err)
		assert.Equal(t, len(namespaceList), int(1))

		count, err := namespaceService.CountNamespace(ctx, ptr.Of("t"))
		assert.NoError(t, err)
		assert.Equal(t, count, int64(1))

		err = namespaceService.UpdateByID(ctx, namespaceObj.ID, map[string]interface{}{query.Namespace.Description.ColumnName().String(): "test"})
		assert.NoError(t, err)

		err = namespaceService.UpdateByID(ctx, 10, map[string]interface{}{query.Namespace.Description.ColumnName().String(): "test"})
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

		err = namespaceService.DeleteByID(ctx, namespaceObj.ID)
		assert.NoError(t, err)

		err = namespaceService.DeleteByID(ctx, 10)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

		return nil
	})
	assert.NoError(t, err)
}

func TestNamespaceServiceQuota(t *testing.T) {
	viper.SetDefault("log.level", "debug")
	logger.SetLevel("debug")
	err := tests.Initialize(t)
	assert.NoError(t, err)
	err = tests.DB.Init()
	assert.NoError(t, err)
	defer func() {
		conn, err := dal.DB.DB()
		assert.NoError(t, err)
		err = conn.Close()
		assert.NoError(t, err)
		err = tests.DB.DeInit()
		assert.NoError(t, err)
	}()

	ctx := log.Logger.WithContext(context.Background())

	namespaceServiceFactory := dao.NewNamespaceServiceFactory()
	userServiceFactory := dao.NewUserServiceFactory()

	err = query.Q.Transaction(func(tx *query.Query) error {
		userService := userServiceFactory.New(tx)
		userObj := &models.User{Username: "artifact-service", Password: ptr.Of("test"), Email: ptr.Of("test@gmail.com")}
		err = userService.Create(ctx, userObj)
		assert.NoError(t, err)

		namespaceService := namespaceServiceFactory.New(tx)

		namespaceObj := &models.Namespace{
			Name: "test",
		}
		err := namespaceService.Create(ctx, namespaceObj)
		assert.NoError(t, err)

		err = namespaceService.UpdateQuota(ctx, namespaceObj.ID, 100)
		assert.NoError(t, err)

		err = namespaceService.UpdateQuota(ctx, 10, 100)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)

		return nil
	})
	assert.NoError(t, err)
}
