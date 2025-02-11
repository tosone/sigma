// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-sigma/sigma/pkg/dal/models"
)

func newBlob(db *gorm.DB, opts ...gen.DOOption) blob {
	_blob := blob{}

	_blob.blobDo.UseDB(db, opts...)
	_blob.blobDo.UseModel(&models.Blob{})

	tableName := _blob.blobDo.TableName()
	_blob.ALL = field.NewAsterisk(tableName)
	_blob.CreatedAt = field.NewInt64(tableName, "created_at")
	_blob.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_blob.DeletedAt = field.NewUint64(tableName, "deleted_at")
	_blob.ID = field.NewInt64(tableName, "id")
	_blob.Digest = field.NewString(tableName, "digest")
	_blob.Size = field.NewInt64(tableName, "size")
	_blob.ContentType = field.NewString(tableName, "content_type")
	_blob.LastPull = field.NewField(tableName, "last_pull")
	_blob.PushedAt = field.NewTime(tableName, "pushed_at")
	_blob.PullTimes = field.NewUint(tableName, "pull_times")
	_blob.Artifacts = blobManyToManyArtifacts{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Artifacts", "models.Artifact"),
		Repository: struct {
			field.RelationField
			Namespace struct {
				field.RelationField
			}
			Builder struct {
				field.RelationField
				Repository struct {
					field.RelationField
				}
				CodeRepository struct {
					field.RelationField
					User3rdParty struct {
						field.RelationField
						User struct {
							field.RelationField
						}
					}
					Branches struct {
						field.RelationField
					}
				}
			}
		}{
			RelationField: field.NewRelation("Artifacts.Repository", "models.Repository"),
			Namespace: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Artifacts.Repository.Namespace", "models.Namespace"),
			},
			Builder: struct {
				field.RelationField
				Repository struct {
					field.RelationField
				}
				CodeRepository struct {
					field.RelationField
					User3rdParty struct {
						field.RelationField
						User struct {
							field.RelationField
						}
					}
					Branches struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("Artifacts.Repository.Builder", "models.Builder"),
				Repository: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Artifacts.Repository.Builder.Repository", "models.Repository"),
				},
				CodeRepository: struct {
					field.RelationField
					User3rdParty struct {
						field.RelationField
						User struct {
							field.RelationField
						}
					}
					Branches struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Artifacts.Repository.Builder.CodeRepository", "models.CodeRepository"),
					User3rdParty: struct {
						field.RelationField
						User struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("Artifacts.Repository.Builder.CodeRepository.User3rdParty", "models.User3rdParty"),
						User: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("Artifacts.Repository.Builder.CodeRepository.User3rdParty.User", "models.User"),
						},
					},
					Branches: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Artifacts.Repository.Builder.CodeRepository.Branches", "models.CodeRepositoryBranch"),
					},
				},
			},
		},
		Referrer: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Artifacts.Referrer", "models.Artifact"),
		},
		Vulnerability: struct {
			field.RelationField
			Artifact struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Artifacts.Vulnerability", "models.ArtifactVulnerability"),
			Artifact: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Artifacts.Vulnerability.Artifact", "models.Artifact"),
			},
		},
		Sbom: struct {
			field.RelationField
			Artifact struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Artifacts.Sbom", "models.ArtifactSbom"),
			Artifact: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Artifacts.Sbom.Artifact", "models.Artifact"),
			},
		},
		Tags: struct {
			field.RelationField
			Repository struct {
				field.RelationField
			}
			Artifact struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Artifacts.Tags", "models.Tag"),
			Repository: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Artifacts.Tags.Repository", "models.Repository"),
			},
			Artifact: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Artifacts.Tags.Artifact", "models.Artifact"),
			},
		},
		ArtifactIndexes: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Artifacts.ArtifactIndexes", "models.Artifact"),
		},
		Blobs: struct {
			field.RelationField
			Artifacts struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Artifacts.Blobs", "models.Blob"),
			Artifacts: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Artifacts.Blobs.Artifacts", "models.Artifact"),
			},
		},
	}

	_blob.fillFieldMap()

	return _blob
}

type blob struct {
	blobDo blobDo

	ALL         field.Asterisk
	CreatedAt   field.Int64
	UpdatedAt   field.Int64
	DeletedAt   field.Uint64
	ID          field.Int64
	Digest      field.String
	Size        field.Int64
	ContentType field.String
	LastPull    field.Field
	PushedAt    field.Time
	PullTimes   field.Uint
	Artifacts   blobManyToManyArtifacts

	fieldMap map[string]field.Expr
}

func (b blob) Table(newTableName string) *blob {
	b.blobDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b blob) As(alias string) *blob {
	b.blobDo.DO = *(b.blobDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *blob) updateTableName(table string) *blob {
	b.ALL = field.NewAsterisk(table)
	b.CreatedAt = field.NewInt64(table, "created_at")
	b.UpdatedAt = field.NewInt64(table, "updated_at")
	b.DeletedAt = field.NewUint64(table, "deleted_at")
	b.ID = field.NewInt64(table, "id")
	b.Digest = field.NewString(table, "digest")
	b.Size = field.NewInt64(table, "size")
	b.ContentType = field.NewString(table, "content_type")
	b.LastPull = field.NewField(table, "last_pull")
	b.PushedAt = field.NewTime(table, "pushed_at")
	b.PullTimes = field.NewUint(table, "pull_times")

	b.fillFieldMap()

	return b
}

func (b *blob) WithContext(ctx context.Context) *blobDo { return b.blobDo.WithContext(ctx) }

func (b blob) TableName() string { return b.blobDo.TableName() }

func (b blob) Alias() string { return b.blobDo.Alias() }

func (b blob) Columns(cols ...field.Expr) gen.Columns { return b.blobDo.Columns(cols...) }

func (b *blob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *blob) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 11)
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["id"] = b.ID
	b.fieldMap["digest"] = b.Digest
	b.fieldMap["size"] = b.Size
	b.fieldMap["content_type"] = b.ContentType
	b.fieldMap["last_pull"] = b.LastPull
	b.fieldMap["pushed_at"] = b.PushedAt
	b.fieldMap["pull_times"] = b.PullTimes

}

func (b blob) clone(db *gorm.DB) blob {
	b.blobDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b blob) replaceDB(db *gorm.DB) blob {
	b.blobDo.ReplaceDB(db)
	return b
}

type blobManyToManyArtifacts struct {
	db *gorm.DB

	field.RelationField

	Repository struct {
		field.RelationField
		Namespace struct {
			field.RelationField
		}
		Builder struct {
			field.RelationField
			Repository struct {
				field.RelationField
			}
			CodeRepository struct {
				field.RelationField
				User3rdParty struct {
					field.RelationField
					User struct {
						field.RelationField
					}
				}
				Branches struct {
					field.RelationField
				}
			}
		}
	}
	Referrer struct {
		field.RelationField
	}
	Vulnerability struct {
		field.RelationField
		Artifact struct {
			field.RelationField
		}
	}
	Sbom struct {
		field.RelationField
		Artifact struct {
			field.RelationField
		}
	}
	Tags struct {
		field.RelationField
		Repository struct {
			field.RelationField
		}
		Artifact struct {
			field.RelationField
		}
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

func (a blobManyToManyArtifacts) Where(conds ...field.Expr) *blobManyToManyArtifacts {
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

func (a blobManyToManyArtifacts) WithContext(ctx context.Context) *blobManyToManyArtifacts {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a blobManyToManyArtifacts) Session(session *gorm.Session) *blobManyToManyArtifacts {
	a.db = a.db.Session(session)
	return &a
}

func (a blobManyToManyArtifacts) Model(m *models.Blob) *blobManyToManyArtifactsTx {
	return &blobManyToManyArtifactsTx{a.db.Model(m).Association(a.Name())}
}

type blobManyToManyArtifactsTx struct{ tx *gorm.Association }

func (a blobManyToManyArtifactsTx) Find() (result []*models.Artifact, err error) {
	return result, a.tx.Find(&result)
}

func (a blobManyToManyArtifactsTx) Append(values ...*models.Artifact) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a blobManyToManyArtifactsTx) Replace(values ...*models.Artifact) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a blobManyToManyArtifactsTx) Delete(values ...*models.Artifact) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a blobManyToManyArtifactsTx) Clear() error {
	return a.tx.Clear()
}

func (a blobManyToManyArtifactsTx) Count() int64 {
	return a.tx.Count()
}

type blobDo struct{ gen.DO }

// SELECT blob_id FROM artifact_blobs WHERE blob_id in (@ids)
func (b blobDo) BlobAssociateWithArtifact(ids []int64) (result map[string]interface{}, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, ids)
	generateSQL.WriteString("SELECT blob_id FROM artifact_blobs WHERE blob_id in (?) ")

	result = make(map[string]interface{})
	var executeSQL *gorm.DB
	executeSQL = b.UnderlyingDB().Raw(generateSQL.String(), params...).Take(result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (b blobDo) Debug() *blobDo {
	return b.withDO(b.DO.Debug())
}

func (b blobDo) WithContext(ctx context.Context) *blobDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b blobDo) ReadDB() *blobDo {
	return b.Clauses(dbresolver.Read)
}

func (b blobDo) WriteDB() *blobDo {
	return b.Clauses(dbresolver.Write)
}

func (b blobDo) Session(config *gorm.Session) *blobDo {
	return b.withDO(b.DO.Session(config))
}

func (b blobDo) Clauses(conds ...clause.Expression) *blobDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b blobDo) Returning(value interface{}, columns ...string) *blobDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b blobDo) Not(conds ...gen.Condition) *blobDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b blobDo) Or(conds ...gen.Condition) *blobDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b blobDo) Select(conds ...field.Expr) *blobDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b blobDo) Where(conds ...gen.Condition) *blobDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b blobDo) Order(conds ...field.Expr) *blobDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b blobDo) Distinct(cols ...field.Expr) *blobDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b blobDo) Omit(cols ...field.Expr) *blobDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b blobDo) Join(table schema.Tabler, on ...field.Expr) *blobDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b blobDo) LeftJoin(table schema.Tabler, on ...field.Expr) *blobDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b blobDo) RightJoin(table schema.Tabler, on ...field.Expr) *blobDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b blobDo) Group(cols ...field.Expr) *blobDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b blobDo) Having(conds ...gen.Condition) *blobDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b blobDo) Limit(limit int) *blobDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b blobDo) Offset(offset int) *blobDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b blobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *blobDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b blobDo) Unscoped() *blobDo {
	return b.withDO(b.DO.Unscoped())
}

func (b blobDo) Create(values ...*models.Blob) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b blobDo) CreateInBatches(values []*models.Blob, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b blobDo) Save(values ...*models.Blob) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b blobDo) First() (*models.Blob, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blob), nil
	}
}

func (b blobDo) Take() (*models.Blob, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blob), nil
	}
}

func (b blobDo) Last() (*models.Blob, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blob), nil
	}
}

func (b blobDo) Find() ([]*models.Blob, error) {
	result, err := b.DO.Find()
	return result.([]*models.Blob), err
}

func (b blobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Blob, err error) {
	buf := make([]*models.Blob, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b blobDo) FindInBatches(result *[]*models.Blob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b blobDo) Attrs(attrs ...field.AssignExpr) *blobDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b blobDo) Assign(attrs ...field.AssignExpr) *blobDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b blobDo) Joins(fields ...field.RelationField) *blobDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b blobDo) Preload(fields ...field.RelationField) *blobDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b blobDo) FirstOrInit() (*models.Blob, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blob), nil
	}
}

func (b blobDo) FirstOrCreate() (*models.Blob, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blob), nil
	}
}

func (b blobDo) FindByPage(offset int, limit int) (result []*models.Blob, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b blobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b blobDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b blobDo) Delete(models ...*models.Blob) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *blobDo) withDO(do gen.Dao) *blobDo {
	b.DO = *do.(*gen.DO)
	return b
}
