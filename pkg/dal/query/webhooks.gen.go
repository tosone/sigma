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

func newWebhook(db *gorm.DB, opts ...gen.DOOption) webhook {
	_webhook := webhook{}

	_webhook.webhookDo.UseDB(db, opts...)
	_webhook.webhookDo.UseModel(&models.Webhook{})

	tableName := _webhook.webhookDo.TableName()
	_webhook.ALL = field.NewAsterisk(tableName)
	_webhook.CreatedAt = field.NewInt64(tableName, "created_at")
	_webhook.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_webhook.DeletedAt = field.NewUint64(tableName, "deleted_at")
	_webhook.ID = field.NewInt64(tableName, "id")
	_webhook.NamespaceID = field.NewInt64(tableName, "namespace_id")
	_webhook.URL = field.NewString(tableName, "url")
	_webhook.Secret = field.NewString(tableName, "secret")
	_webhook.SslVerify = field.NewBool(tableName, "ssl_verify")
	_webhook.RetryTimes = field.NewInt(tableName, "retry_times")
	_webhook.RetryDuration = field.NewInt(tableName, "retry_duration")
	_webhook.Enable = field.NewBool(tableName, "enable")
	_webhook.EventNamespace = field.NewBool(tableName, "event_namespace")
	_webhook.EventRepository = field.NewBool(tableName, "event_repository")
	_webhook.EventTag = field.NewBool(tableName, "event_tag")
	_webhook.EventArtifact = field.NewBool(tableName, "event_artifact")
	_webhook.EventMember = field.NewBool(tableName, "event_member")

	_webhook.fillFieldMap()

	return _webhook
}

type webhook struct {
	webhookDo webhookDo

	ALL             field.Asterisk
	CreatedAt       field.Int64
	UpdatedAt       field.Int64
	DeletedAt       field.Uint64
	ID              field.Int64
	NamespaceID     field.Int64
	URL             field.String
	Secret          field.String
	SslVerify       field.Bool
	RetryTimes      field.Int
	RetryDuration   field.Int
	Enable          field.Bool
	EventNamespace  field.Bool
	EventRepository field.Bool
	EventTag        field.Bool
	EventArtifact   field.Bool
	EventMember     field.Bool

	fieldMap map[string]field.Expr
}

func (w webhook) Table(newTableName string) *webhook {
	w.webhookDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w webhook) As(alias string) *webhook {
	w.webhookDo.DO = *(w.webhookDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *webhook) updateTableName(table string) *webhook {
	w.ALL = field.NewAsterisk(table)
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.DeletedAt = field.NewUint64(table, "deleted_at")
	w.ID = field.NewInt64(table, "id")
	w.NamespaceID = field.NewInt64(table, "namespace_id")
	w.URL = field.NewString(table, "url")
	w.Secret = field.NewString(table, "secret")
	w.SslVerify = field.NewBool(table, "ssl_verify")
	w.RetryTimes = field.NewInt(table, "retry_times")
	w.RetryDuration = field.NewInt(table, "retry_duration")
	w.Enable = field.NewBool(table, "enable")
	w.EventNamespace = field.NewBool(table, "event_namespace")
	w.EventRepository = field.NewBool(table, "event_repository")
	w.EventTag = field.NewBool(table, "event_tag")
	w.EventArtifact = field.NewBool(table, "event_artifact")
	w.EventMember = field.NewBool(table, "event_member")

	w.fillFieldMap()

	return w
}

func (w *webhook) WithContext(ctx context.Context) *webhookDo { return w.webhookDo.WithContext(ctx) }

func (w webhook) TableName() string { return w.webhookDo.TableName() }

func (w webhook) Alias() string { return w.webhookDo.Alias() }

func (w webhook) Columns(cols ...field.Expr) gen.Columns { return w.webhookDo.Columns(cols...) }

func (w *webhook) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *webhook) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 16)
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["id"] = w.ID
	w.fieldMap["namespace_id"] = w.NamespaceID
	w.fieldMap["url"] = w.URL
	w.fieldMap["secret"] = w.Secret
	w.fieldMap["ssl_verify"] = w.SslVerify
	w.fieldMap["retry_times"] = w.RetryTimes
	w.fieldMap["retry_duration"] = w.RetryDuration
	w.fieldMap["enable"] = w.Enable
	w.fieldMap["event_namespace"] = w.EventNamespace
	w.fieldMap["event_repository"] = w.EventRepository
	w.fieldMap["event_tag"] = w.EventTag
	w.fieldMap["event_artifact"] = w.EventArtifact
	w.fieldMap["event_member"] = w.EventMember
}

func (w webhook) clone(db *gorm.DB) webhook {
	w.webhookDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w webhook) replaceDB(db *gorm.DB) webhook {
	w.webhookDo.ReplaceDB(db)
	return w
}

type webhookDo struct{ gen.DO }

func (w webhookDo) Debug() *webhookDo {
	return w.withDO(w.DO.Debug())
}

func (w webhookDo) WithContext(ctx context.Context) *webhookDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w webhookDo) ReadDB() *webhookDo {
	return w.Clauses(dbresolver.Read)
}

func (w webhookDo) WriteDB() *webhookDo {
	return w.Clauses(dbresolver.Write)
}

func (w webhookDo) Session(config *gorm.Session) *webhookDo {
	return w.withDO(w.DO.Session(config))
}

func (w webhookDo) Clauses(conds ...clause.Expression) *webhookDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w webhookDo) Returning(value interface{}, columns ...string) *webhookDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w webhookDo) Not(conds ...gen.Condition) *webhookDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w webhookDo) Or(conds ...gen.Condition) *webhookDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w webhookDo) Select(conds ...field.Expr) *webhookDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w webhookDo) Where(conds ...gen.Condition) *webhookDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w webhookDo) Order(conds ...field.Expr) *webhookDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w webhookDo) Distinct(cols ...field.Expr) *webhookDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w webhookDo) Omit(cols ...field.Expr) *webhookDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w webhookDo) Join(table schema.Tabler, on ...field.Expr) *webhookDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w webhookDo) LeftJoin(table schema.Tabler, on ...field.Expr) *webhookDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w webhookDo) RightJoin(table schema.Tabler, on ...field.Expr) *webhookDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w webhookDo) Group(cols ...field.Expr) *webhookDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w webhookDo) Having(conds ...gen.Condition) *webhookDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w webhookDo) Limit(limit int) *webhookDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w webhookDo) Offset(offset int) *webhookDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w webhookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *webhookDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w webhookDo) Unscoped() *webhookDo {
	return w.withDO(w.DO.Unscoped())
}

func (w webhookDo) Create(values ...*models.Webhook) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w webhookDo) CreateInBatches(values []*models.Webhook, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w webhookDo) Save(values ...*models.Webhook) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w webhookDo) First() (*models.Webhook, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Webhook), nil
	}
}

func (w webhookDo) Take() (*models.Webhook, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Webhook), nil
	}
}

func (w webhookDo) Last() (*models.Webhook, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Webhook), nil
	}
}

func (w webhookDo) Find() ([]*models.Webhook, error) {
	result, err := w.DO.Find()
	return result.([]*models.Webhook), err
}

func (w webhookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Webhook, err error) {
	buf := make([]*models.Webhook, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w webhookDo) FindInBatches(result *[]*models.Webhook, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w webhookDo) Attrs(attrs ...field.AssignExpr) *webhookDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w webhookDo) Assign(attrs ...field.AssignExpr) *webhookDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w webhookDo) Joins(fields ...field.RelationField) *webhookDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w webhookDo) Preload(fields ...field.RelationField) *webhookDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w webhookDo) FirstOrInit() (*models.Webhook, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Webhook), nil
	}
}

func (w webhookDo) FirstOrCreate() (*models.Webhook, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Webhook), nil
	}
}

func (w webhookDo) FindByPage(offset int, limit int) (result []*models.Webhook, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w webhookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w webhookDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w webhookDo) Delete(models ...*models.Webhook) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *webhookDo) withDO(do gen.Dao) *webhookDo {
	w.DO = *do.(*gen.DO)
	return w
}
