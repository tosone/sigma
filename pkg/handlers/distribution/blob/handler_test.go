// Copyright 2023 XImager
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

package blob

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	daomock "github.com/ximager/ximager/pkg/dal/dao/mocks"
	"github.com/ximager/ximager/pkg/handlers/distribution"
)

func TestHandlerNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	daoMockBlobService := daomock.NewMockBlobServiceFactory(ctrl)

	handler := handlerNew(inject{blobServiceFactory: daoMockBlobService})
	assert.NotNil(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v2/test-none-exist", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	f := &factory{}
	err := f.Initialize(c)
	assert.ErrorIs(t, err, distribution.ErrNext)
}
