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

package models

import (
	"database/sql"
	"time"

	"gorm.io/gen"
	"gorm.io/plugin/soft_delete"
)

// Blob represents a blob
type Blob struct {
	CreatedAt int64                 `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64                 `gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
	ID        int64                 `gorm:"primaryKey"`

	Digest      string
	Size        int64
	ContentType string

	LastPull  sql.NullTime
	PushedAt  time.Time `gorm:"autoCreateTime"`
	PullTimes uint      `gorm:"default:0"`

	Artifacts []*Artifact `gorm:"many2many:artifact_blobs;"`
}

// BlobAssociateWithArtifact ...
type BlobAssociateWithArtifact interface {
	// SELECT blob_id FROM artifact_blobs WHERE blob_id in (@ids)
	BlobAssociateWithArtifact(ids []int64) (gen.M, error)
}
