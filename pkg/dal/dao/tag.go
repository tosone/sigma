// The MIT License (MIT)
//
// Copyright © 2023 Tosone <i@tosone.cn>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dao

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/ximager/ximager/pkg/dal/models"
	"github.com/ximager/ximager/pkg/dal/query"
	"github.com/ximager/ximager/pkg/types"
)

// TagService is the interface that provides the tag service methods.
type TagService interface {
	// Save save a new tag if conflict update.
	Save(ctx context.Context, tag *models.Tag) (*models.Tag, error)
	// Get gets the tag with the specified tag ID.
	GetByID(ctx context.Context, tagID uint64) (*models.Tag, error)
	// GetByName gets the tag with the specified tag name.
	GetByName(ctx context.Context, repository, tag string) (*models.Tag, error)
	// DeleteByName deletes the tag with the specified tag name.
	DeleteByName(ctx context.Context, repository string, tag string) error
	// Incr increases the pull times of the artifact.
	Incr(ctx context.Context, id uint64) error
	// ListByDtPagination lists the tags by the specified repository and pagination.
	ListByDtPagination(ctx context.Context, repository string, limit int, lastID ...uint64) ([]*models.Tag, error)
	// ListTag lists the tags by the specified request.
	ListTag(ctx context.Context, req types.ListTagRequest) ([]*models.Tag, error)
	// CountArtifact counts the artifacts by the specified request.
	CountTag(ctx context.Context, req types.ListTagRequest) (int64, error)
	// DeleteByID deletes the tag with the specified tag ID.
	DeleteByID(ctx context.Context, id uint64) error
	// CountByArtifact counts the tags by the specified artifact.
	CountByArtifact(ctx context.Context, artifactIDs []uint64) (map[uint64]int64, error)
}

type tagService struct {
	tx *query.Query
}

// NewTagService creates a new tag service.
func NewTagService(txs ...*query.Query) TagService {
	tx := query.Q
	if len(txs) > 0 {
		tx = txs[0]
	}
	return &tagService{
		tx: tx,
	}
}

// Save save a new tag if conflict update.
func (s *tagService) Save(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	err := s.tx.Tag.WithContext(ctx).Save(tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// Get gets the tag with the specified tag ID.
func (s *tagService) GetByID(ctx context.Context, tagID uint64) (*models.Tag, error) {
	tag, err := s.tx.Tag.WithContext(ctx).Where(s.tx.Tag.ID.Eq(tagID)).First()
	if err != nil {
		return nil, err
	}
	return tag, nil
}

// GetByName gets the tag with the specified tag name.
func (s *tagService) GetByName(ctx context.Context, repository, tag string) (*models.Tag, error) {
	tagObj, err := s.tx.Tag.WithContext(ctx).
		LeftJoin(s.tx.Repository, s.tx.Tag.RepositoryID.EqCol(s.tx.Repository.ID)).
		Where(s.tx.Tag.Name.Eq(tag)).
		Where(s.tx.Repository.Name.Eq(repository)).
		Preload(s.tx.Tag.Artifact).
		First()
	if err != nil {
		return nil, err
	}
	return tagObj, nil
}

// DeleteByName deletes the tag with the specified tag name.
func (s *tagService) DeleteByName(ctx context.Context, repository, tag string) error {
	matched, err := s.tx.Tag.WithContext(ctx).DeleteByName(repository, tag)
	if err != nil {
		return err
	}
	if matched == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Incr increases the pull times of the artifact.
func (s *tagService) Incr(ctx context.Context, id uint64) error {
	_, err := s.tx.Tag.WithContext(ctx).Where(s.tx.Tag.ID.Eq(id)).
		UpdateColumns(map[string]interface{}{
			"pull_times": gorm.Expr("pull_times + ?", 1),
			"last_pull":  time.Now(),
		})
	if err != nil {
		return err
	}
	return nil
}

// ListByDtPagination lists the tags by the specified repository and pagination.
func (s *tagService) ListByDtPagination(ctx context.Context, repository string, limit int, lastID ...uint64) ([]*models.Tag, error) {
	do := s.tx.Tag.WithContext(ctx).
		LeftJoin(s.tx.Repository, s.tx.Tag.RepositoryID.EqCol(s.tx.Repository.ID)).
		Where(s.tx.Repository.Name.Eq(repository))
	if len(lastID) > 0 {
		do = do.Where(s.tx.Tag.ID.Gt(lastID[0]))
	}
	tags, err := do.Order(s.tx.Tag.ID).Limit(limit).Find()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// ListTag lists the tags by the specified request.
func (s *tagService) ListTag(ctx context.Context, req types.ListTagRequest) ([]*models.Tag, error) {
	return s.tx.Tag.WithContext(ctx).
		LeftJoin(s.tx.Repository, s.tx.Tag.RepositoryID.EqCol(s.tx.Repository.ID)).
		Where(s.tx.Repository.Name.Eq(req.Repository)).
		Offset(req.PageSize * (req.PageNum - 1)).Limit(req.PageSize).Find()
}

// CountArtifact counts the artifacts by the specified request.
func (s *tagService) CountTag(ctx context.Context, req types.ListTagRequest) (int64, error) {
	return s.tx.Tag.WithContext(ctx).
		LeftJoin(s.tx.Repository, s.tx.Tag.RepositoryID.EqCol(s.tx.Repository.ID)).
		Where(s.tx.Repository.Name.Eq(req.Repository)).
		Count()
}

// DeleteByID deletes the tag with the specified tag ID.
func (s *tagService) DeleteByID(ctx context.Context, id uint64) error {
	matched, err := s.tx.Tag.WithContext(ctx).Where(s.tx.Tag.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	if matched.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// CountByArtifact counts the tags by the specified artifact.
func (s *tagService) CountByArtifact(ctx context.Context, artifactIDs []uint64) (map[uint64]int64, error) {
	tagCount := make(map[uint64]int64)
	var count []struct {
		ArtifactID uint64 `gorm:"column:artifact_id"`
		Count      int64  `gorm:"column:count"`
	}
	err := s.tx.Tag.WithContext(ctx).
		LeftJoin(s.tx.Artifact, s.tx.Tag.ArtifactID.EqCol(s.tx.Artifact.ID)).
		Where(s.tx.Artifact.ID.In(artifactIDs...)).
		Group(s.tx.Artifact.ID).
		Select(s.tx.Artifact.ID.As("artifact_id"), s.tx.Tag.ID.Count().As("count")).
		Scan(&count)
	if err != nil {
		return nil, err
	}
	for _, c := range count {
		tagCount[c.ArtifactID] = c.Count
	}
	return tagCount, nil
}
