// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-sigma/sigma/pkg/dal/models"
)

func newArtifact(db *gorm.DB, opts ...gen.DOOption) artifact {
	_artifact := artifact{}

	_artifact.artifactDo.UseDB(db, opts...)
	_artifact.artifactDo.UseModel(&models.Artifact{})

	tableName := _artifact.artifactDo.TableName()
	_artifact.ALL = field.NewAsterisk(tableName)
	_artifact.CreatedAt = field.NewTime(tableName, "created_at")
	_artifact.UpdatedAt = field.NewTime(tableName, "updated_at")
	_artifact.DeletedAt = field.NewUint(tableName, "deleted_at")
	_artifact.ID = field.NewInt64(tableName, "id")
	_artifact.RepositoryID = field.NewInt64(tableName, "repository_id")
	_artifact.Digest = field.NewString(tableName, "digest")
	_artifact.Size = field.NewInt64(tableName, "size")
	_artifact.BlobsSize = field.NewInt64(tableName, "blobs_size")
	_artifact.ContentType = field.NewString(tableName, "content_type")
	_artifact.Raw = field.NewBytes(tableName, "raw")
	_artifact.LastPull = field.NewField(tableName, "last_pull")
	_artifact.PushedAt = field.NewTime(tableName, "pushed_at")
	_artifact.PullTimes = field.NewInt64(tableName, "pull_times")
	_artifact.Tags = artifactHasManyTags{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Tags", "models.Tag"),
		Repository: struct {
			field.RelationField
			Namespace struct {
				field.RelationField
			}
			Tags struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Tags.Repository", "models.Repository"),
			Namespace: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Tags.Repository.Namespace", "models.Namespace"),
			},
			Tags: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Tags.Repository.Tags", "models.RepositoryTag"),
			},
		},
		Artifact: struct {
			field.RelationField
			Repository struct {
				field.RelationField
			}
			Tags struct {
				field.RelationField
			}
			ArtifactIndexes struct {
				field.RelationField
			}
			Blobs struct {
				field.RelationField
				Artifacts struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Tags.Artifact", "models.Artifact"),
			Repository: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Tags.Artifact.Repository", "models.Repository"),
			},
			Tags: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Tags.Artifact.Tags", "models.Tag"),
			},
			ArtifactIndexes: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Tags.Artifact.ArtifactIndexes", "models.Artifact"),
			},
			Blobs: struct {
				field.RelationField
				Artifacts struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Tags.Artifact.Blobs", "models.Blob"),
				Artifacts: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Tags.Artifact.Blobs.Artifacts", "models.Artifact"),
				},
			},
		},
	}

	_artifact.Repository = artifactBelongsToRepository{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Repository", "models.Repository"),
	}

	_artifact.ArtifactIndexes = artifactManyToManyArtifactIndexes{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ArtifactIndexes", "models.Artifact"),
	}

	_artifact.Blobs = artifactManyToManyBlobs{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Blobs", "models.Blob"),
	}

	_artifact.fillFieldMap()

	return _artifact
}

type artifact struct {
	artifactDo artifactDo

	ALL          field.Asterisk
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Uint
	ID           field.Int64
	RepositoryID field.Int64
	Digest       field.String
	Size         field.Int64
	BlobsSize    field.Int64
	ContentType  field.String
	Raw          field.Bytes
	LastPull     field.Field
	PushedAt     field.Time
	PullTimes    field.Int64
	Tags         artifactHasManyTags

	Repository artifactBelongsToRepository

	ArtifactIndexes artifactManyToManyArtifactIndexes

	Blobs artifactManyToManyBlobs

	fieldMap map[string]field.Expr
}

func (a artifact) Table(newTableName string) *artifact {
	a.artifactDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a artifact) As(alias string) *artifact {
	a.artifactDo.DO = *(a.artifactDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *artifact) updateTableName(table string) *artifact {
	a.ALL = field.NewAsterisk(table)
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewUint(table, "deleted_at")
	a.ID = field.NewInt64(table, "id")
	a.RepositoryID = field.NewInt64(table, "repository_id")
	a.Digest = field.NewString(table, "digest")
	a.Size = field.NewInt64(table, "size")
	a.BlobsSize = field.NewInt64(table, "blobs_size")
	a.ContentType = field.NewString(table, "content_type")
	a.Raw = field.NewBytes(table, "raw")
	a.LastPull = field.NewField(table, "last_pull")
	a.PushedAt = field.NewTime(table, "pushed_at")
	a.PullTimes = field.NewInt64(table, "pull_times")

	a.fillFieldMap()

	return a
}

func (a *artifact) WithContext(ctx context.Context) *artifactDo { return a.artifactDo.WithContext(ctx) }

func (a artifact) TableName() string { return a.artifactDo.TableName() }

func (a artifact) Alias() string { return a.artifactDo.Alias() }

func (a artifact) Columns(cols ...field.Expr) gen.Columns { return a.artifactDo.Columns(cols...) }

func (a *artifact) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *artifact) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 17)
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["id"] = a.ID
	a.fieldMap["repository_id"] = a.RepositoryID
	a.fieldMap["digest"] = a.Digest
	a.fieldMap["size"] = a.Size
	a.fieldMap["blobs_size"] = a.BlobsSize
	a.fieldMap["content_type"] = a.ContentType
	a.fieldMap["raw"] = a.Raw
	a.fieldMap["last_pull"] = a.LastPull
	a.fieldMap["pushed_at"] = a.PushedAt
	a.fieldMap["pull_times"] = a.PullTimes

}

func (a artifact) clone(db *gorm.DB) artifact {
	a.artifactDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a artifact) replaceDB(db *gorm.DB) artifact {
	a.artifactDo.ReplaceDB(db)
	return a
}

type artifactHasManyTags struct {
	db *gorm.DB

	field.RelationField

	Repository struct {
		field.RelationField
		Namespace struct {
			field.RelationField
		}
		Tags struct {
			field.RelationField
		}
	}
	Artifact struct {
		field.RelationField
		Repository struct {
			field.RelationField
		}
		Tags struct {
			field.RelationField
		}
		ArtifactIndexes struct {
			field.RelationField
		}
		Blobs struct {
			field.RelationField
			Artifacts struct {
				field.RelationField
			}
		}
	}
}

func (a artifactHasManyTags) Where(conds ...field.Expr) *artifactHasManyTags {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a artifactHasManyTags) WithContext(ctx context.Context) *artifactHasManyTags {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a artifactHasManyTags) Session(session *gorm.Session) *artifactHasManyTags {
	a.db = a.db.Session(session)
	return &a
}

func (a artifactHasManyTags) Model(m *models.Artifact) *artifactHasManyTagsTx {
	return &artifactHasManyTagsTx{a.db.Model(m).Association(a.Name())}
}

type artifactHasManyTagsTx struct{ tx *gorm.Association }

func (a artifactHasManyTagsTx) Find() (result []*models.Tag, err error) {
	return result, a.tx.Find(&result)
}

func (a artifactHasManyTagsTx) Append(values ...*models.Tag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a artifactHasManyTagsTx) Replace(values ...*models.Tag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a artifactHasManyTagsTx) Delete(values ...*models.Tag) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a artifactHasManyTagsTx) Clear() error {
	return a.tx.Clear()
}

func (a artifactHasManyTagsTx) Count() int64 {
	return a.tx.Count()
}

type artifactBelongsToRepository struct {
	db *gorm.DB

	field.RelationField
}

func (a artifactBelongsToRepository) Where(conds ...field.Expr) *artifactBelongsToRepository {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a artifactBelongsToRepository) WithContext(ctx context.Context) *artifactBelongsToRepository {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a artifactBelongsToRepository) Session(session *gorm.Session) *artifactBelongsToRepository {
	a.db = a.db.Session(session)
	return &a
}

func (a artifactBelongsToRepository) Model(m *models.Artifact) *artifactBelongsToRepositoryTx {
	return &artifactBelongsToRepositoryTx{a.db.Model(m).Association(a.Name())}
}

type artifactBelongsToRepositoryTx struct{ tx *gorm.Association }

func (a artifactBelongsToRepositoryTx) Find() (result *models.Repository, err error) {
	return result, a.tx.Find(&result)
}

func (a artifactBelongsToRepositoryTx) Append(values ...*models.Repository) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a artifactBelongsToRepositoryTx) Replace(values ...*models.Repository) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a artifactBelongsToRepositoryTx) Delete(values ...*models.Repository) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a artifactBelongsToRepositoryTx) Clear() error {
	return a.tx.Clear()
}

func (a artifactBelongsToRepositoryTx) Count() int64 {
	return a.tx.Count()
}

type artifactManyToManyArtifactIndexes struct {
	db *gorm.DB

	field.RelationField
}

func (a artifactManyToManyArtifactIndexes) Where(conds ...field.Expr) *artifactManyToManyArtifactIndexes {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a artifactManyToManyArtifactIndexes) WithContext(ctx context.Context) *artifactManyToManyArtifactIndexes {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a artifactManyToManyArtifactIndexes) Session(session *gorm.Session) *artifactManyToManyArtifactIndexes {
	a.db = a.db.Session(session)
	return &a
}

func (a artifactManyToManyArtifactIndexes) Model(m *models.Artifact) *artifactManyToManyArtifactIndexesTx {
	return &artifactManyToManyArtifactIndexesTx{a.db.Model(m).Association(a.Name())}
}

type artifactManyToManyArtifactIndexesTx struct{ tx *gorm.Association }

func (a artifactManyToManyArtifactIndexesTx) Find() (result []*models.Artifact, err error) {
	return result, a.tx.Find(&result)
}

func (a artifactManyToManyArtifactIndexesTx) Append(values ...*models.Artifact) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a artifactManyToManyArtifactIndexesTx) Replace(values ...*models.Artifact) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a artifactManyToManyArtifactIndexesTx) Delete(values ...*models.Artifact) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a artifactManyToManyArtifactIndexesTx) Clear() error {
	return a.tx.Clear()
}

func (a artifactManyToManyArtifactIndexesTx) Count() int64 {
	return a.tx.Count()
}

type artifactManyToManyBlobs struct {
	db *gorm.DB

	field.RelationField
}

func (a artifactManyToManyBlobs) Where(conds ...field.Expr) *artifactManyToManyBlobs {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a artifactManyToManyBlobs) WithContext(ctx context.Context) *artifactManyToManyBlobs {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a artifactManyToManyBlobs) Session(session *gorm.Session) *artifactManyToManyBlobs {
	a.db = a.db.Session(session)
	return &a
}

func (a artifactManyToManyBlobs) Model(m *models.Artifact) *artifactManyToManyBlobsTx {
	return &artifactManyToManyBlobsTx{a.db.Model(m).Association(a.Name())}
}

type artifactManyToManyBlobsTx struct{ tx *gorm.Association }

func (a artifactManyToManyBlobsTx) Find() (result []*models.Blob, err error) {
	return result, a.tx.Find(&result)
}

func (a artifactManyToManyBlobsTx) Append(values ...*models.Blob) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a artifactManyToManyBlobsTx) Replace(values ...*models.Blob) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a artifactManyToManyBlobsTx) Delete(values ...*models.Blob) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a artifactManyToManyBlobsTx) Clear() error {
	return a.tx.Clear()
}

func (a artifactManyToManyBlobsTx) Count() int64 {
	return a.tx.Count()
}

type artifactDo struct{ gen.DO }

func (a artifactDo) Debug() *artifactDo {
	return a.withDO(a.DO.Debug())
}

func (a artifactDo) WithContext(ctx context.Context) *artifactDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a artifactDo) ReadDB() *artifactDo {
	return a.Clauses(dbresolver.Read)
}

func (a artifactDo) WriteDB() *artifactDo {
	return a.Clauses(dbresolver.Write)
}

func (a artifactDo) Session(config *gorm.Session) *artifactDo {
	return a.withDO(a.DO.Session(config))
}

func (a artifactDo) Clauses(conds ...clause.Expression) *artifactDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a artifactDo) Returning(value interface{}, columns ...string) *artifactDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a artifactDo) Not(conds ...gen.Condition) *artifactDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a artifactDo) Or(conds ...gen.Condition) *artifactDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a artifactDo) Select(conds ...field.Expr) *artifactDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a artifactDo) Where(conds ...gen.Condition) *artifactDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a artifactDo) Order(conds ...field.Expr) *artifactDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a artifactDo) Distinct(cols ...field.Expr) *artifactDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a artifactDo) Omit(cols ...field.Expr) *artifactDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a artifactDo) Join(table schema.Tabler, on ...field.Expr) *artifactDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a artifactDo) LeftJoin(table schema.Tabler, on ...field.Expr) *artifactDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a artifactDo) RightJoin(table schema.Tabler, on ...field.Expr) *artifactDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a artifactDo) Group(cols ...field.Expr) *artifactDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a artifactDo) Having(conds ...gen.Condition) *artifactDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a artifactDo) Limit(limit int) *artifactDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a artifactDo) Offset(offset int) *artifactDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a artifactDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *artifactDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a artifactDo) Unscoped() *artifactDo {
	return a.withDO(a.DO.Unscoped())
}

func (a artifactDo) Create(values ...*models.Artifact) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a artifactDo) CreateInBatches(values []*models.Artifact, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a artifactDo) Save(values ...*models.Artifact) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a artifactDo) First() (*models.Artifact, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Artifact), nil
	}
}

func (a artifactDo) Take() (*models.Artifact, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Artifact), nil
	}
}

func (a artifactDo) Last() (*models.Artifact, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Artifact), nil
	}
}

func (a artifactDo) Find() ([]*models.Artifact, error) {
	result, err := a.DO.Find()
	return result.([]*models.Artifact), err
}

func (a artifactDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Artifact, err error) {
	buf := make([]*models.Artifact, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a artifactDo) FindInBatches(result *[]*models.Artifact, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a artifactDo) Attrs(attrs ...field.AssignExpr) *artifactDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a artifactDo) Assign(attrs ...field.AssignExpr) *artifactDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a artifactDo) Joins(fields ...field.RelationField) *artifactDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a artifactDo) Preload(fields ...field.RelationField) *artifactDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a artifactDo) FirstOrInit() (*models.Artifact, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Artifact), nil
	}
}

func (a artifactDo) FirstOrCreate() (*models.Artifact, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Artifact), nil
	}
}

func (a artifactDo) FindByPage(offset int, limit int) (result []*models.Artifact, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a artifactDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a artifactDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a artifactDo) Delete(models ...*models.Artifact) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *artifactDo) withDO(do gen.Dao) *artifactDo {
	a.DO = *do.(*gen.DO)
	return a
}
