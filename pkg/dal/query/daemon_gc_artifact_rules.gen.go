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

func newDaemonGcArtifactRule(db *gorm.DB, opts ...gen.DOOption) daemonGcArtifactRule {
	_daemonGcArtifactRule := daemonGcArtifactRule{}

	_daemonGcArtifactRule.daemonGcArtifactRuleDo.UseDB(db, opts...)
	_daemonGcArtifactRule.daemonGcArtifactRuleDo.UseModel(&models.DaemonGcArtifactRule{})

	tableName := _daemonGcArtifactRule.daemonGcArtifactRuleDo.TableName()
	_daemonGcArtifactRule.ALL = field.NewAsterisk(tableName)
	_daemonGcArtifactRule.CreatedAt = field.NewInt64(tableName, "created_at")
	_daemonGcArtifactRule.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_daemonGcArtifactRule.DeletedAt = field.NewUint64(tableName, "deleted_at")
	_daemonGcArtifactRule.ID = field.NewInt64(tableName, "id")
	_daemonGcArtifactRule.NamespaceID = field.NewInt64(tableName, "namespace_id")
	_daemonGcArtifactRule.IsRunning = field.NewBool(tableName, "is_running")
	_daemonGcArtifactRule.RetentionDay = field.NewInt(tableName, "retention_day")
	_daemonGcArtifactRule.CronEnabled = field.NewBool(tableName, "cron_enabled")
	_daemonGcArtifactRule.CronRule = field.NewString(tableName, "cron_rule")
	_daemonGcArtifactRule.CronNextTrigger = field.NewTime(tableName, "cron_next_trigger")
	_daemonGcArtifactRule.Namespace = daemonGcArtifactRuleBelongsToNamespace{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Namespace", "models.Namespace"),
	}

	_daemonGcArtifactRule.fillFieldMap()

	return _daemonGcArtifactRule
}

type daemonGcArtifactRule struct {
	daemonGcArtifactRuleDo daemonGcArtifactRuleDo

	ALL             field.Asterisk
	CreatedAt       field.Int64
	UpdatedAt       field.Int64
	DeletedAt       field.Uint64
	ID              field.Int64
	NamespaceID     field.Int64
	IsRunning       field.Bool
	RetentionDay    field.Int
	CronEnabled     field.Bool
	CronRule        field.String
	CronNextTrigger field.Time
	Namespace       daemonGcArtifactRuleBelongsToNamespace

	fieldMap map[string]field.Expr
}

func (d daemonGcArtifactRule) Table(newTableName string) *daemonGcArtifactRule {
	d.daemonGcArtifactRuleDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d daemonGcArtifactRule) As(alias string) *daemonGcArtifactRule {
	d.daemonGcArtifactRuleDo.DO = *(d.daemonGcArtifactRuleDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *daemonGcArtifactRule) updateTableName(table string) *daemonGcArtifactRule {
	d.ALL = field.NewAsterisk(table)
	d.CreatedAt = field.NewInt64(table, "created_at")
	d.UpdatedAt = field.NewInt64(table, "updated_at")
	d.DeletedAt = field.NewUint64(table, "deleted_at")
	d.ID = field.NewInt64(table, "id")
	d.NamespaceID = field.NewInt64(table, "namespace_id")
	d.IsRunning = field.NewBool(table, "is_running")
	d.RetentionDay = field.NewInt(table, "retention_day")
	d.CronEnabled = field.NewBool(table, "cron_enabled")
	d.CronRule = field.NewString(table, "cron_rule")
	d.CronNextTrigger = field.NewTime(table, "cron_next_trigger")

	d.fillFieldMap()

	return d
}

func (d *daemonGcArtifactRule) WithContext(ctx context.Context) *daemonGcArtifactRuleDo {
	return d.daemonGcArtifactRuleDo.WithContext(ctx)
}

func (d daemonGcArtifactRule) TableName() string { return d.daemonGcArtifactRuleDo.TableName() }

func (d daemonGcArtifactRule) Alias() string { return d.daemonGcArtifactRuleDo.Alias() }

func (d daemonGcArtifactRule) Columns(cols ...field.Expr) gen.Columns {
	return d.daemonGcArtifactRuleDo.Columns(cols...)
}

func (d *daemonGcArtifactRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *daemonGcArtifactRule) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 11)
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["id"] = d.ID
	d.fieldMap["namespace_id"] = d.NamespaceID
	d.fieldMap["is_running"] = d.IsRunning
	d.fieldMap["retention_day"] = d.RetentionDay
	d.fieldMap["cron_enabled"] = d.CronEnabled
	d.fieldMap["cron_rule"] = d.CronRule
	d.fieldMap["cron_next_trigger"] = d.CronNextTrigger

}

func (d daemonGcArtifactRule) clone(db *gorm.DB) daemonGcArtifactRule {
	d.daemonGcArtifactRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d daemonGcArtifactRule) replaceDB(db *gorm.DB) daemonGcArtifactRule {
	d.daemonGcArtifactRuleDo.ReplaceDB(db)
	return d
}

type daemonGcArtifactRuleBelongsToNamespace struct {
	db *gorm.DB

	field.RelationField
}

func (a daemonGcArtifactRuleBelongsToNamespace) Where(conds ...field.Expr) *daemonGcArtifactRuleBelongsToNamespace {
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

func (a daemonGcArtifactRuleBelongsToNamespace) WithContext(ctx context.Context) *daemonGcArtifactRuleBelongsToNamespace {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a daemonGcArtifactRuleBelongsToNamespace) Session(session *gorm.Session) *daemonGcArtifactRuleBelongsToNamespace {
	a.db = a.db.Session(session)
	return &a
}

func (a daemonGcArtifactRuleBelongsToNamespace) Model(m *models.DaemonGcArtifactRule) *daemonGcArtifactRuleBelongsToNamespaceTx {
	return &daemonGcArtifactRuleBelongsToNamespaceTx{a.db.Model(m).Association(a.Name())}
}

type daemonGcArtifactRuleBelongsToNamespaceTx struct{ tx *gorm.Association }

func (a daemonGcArtifactRuleBelongsToNamespaceTx) Find() (result *models.Namespace, err error) {
	return result, a.tx.Find(&result)
}

func (a daemonGcArtifactRuleBelongsToNamespaceTx) Append(values ...*models.Namespace) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a daemonGcArtifactRuleBelongsToNamespaceTx) Replace(values ...*models.Namespace) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a daemonGcArtifactRuleBelongsToNamespaceTx) Delete(values ...*models.Namespace) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a daemonGcArtifactRuleBelongsToNamespaceTx) Clear() error {
	return a.tx.Clear()
}

func (a daemonGcArtifactRuleBelongsToNamespaceTx) Count() int64 {
	return a.tx.Count()
}

type daemonGcArtifactRuleDo struct{ gen.DO }

func (d daemonGcArtifactRuleDo) Debug() *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Debug())
}

func (d daemonGcArtifactRuleDo) WithContext(ctx context.Context) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d daemonGcArtifactRuleDo) ReadDB() *daemonGcArtifactRuleDo {
	return d.Clauses(dbresolver.Read)
}

func (d daemonGcArtifactRuleDo) WriteDB() *daemonGcArtifactRuleDo {
	return d.Clauses(dbresolver.Write)
}

func (d daemonGcArtifactRuleDo) Session(config *gorm.Session) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Session(config))
}

func (d daemonGcArtifactRuleDo) Clauses(conds ...clause.Expression) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d daemonGcArtifactRuleDo) Returning(value interface{}, columns ...string) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d daemonGcArtifactRuleDo) Not(conds ...gen.Condition) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d daemonGcArtifactRuleDo) Or(conds ...gen.Condition) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d daemonGcArtifactRuleDo) Select(conds ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d daemonGcArtifactRuleDo) Where(conds ...gen.Condition) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d daemonGcArtifactRuleDo) Order(conds ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d daemonGcArtifactRuleDo) Distinct(cols ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d daemonGcArtifactRuleDo) Omit(cols ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d daemonGcArtifactRuleDo) Join(table schema.Tabler, on ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d daemonGcArtifactRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d daemonGcArtifactRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d daemonGcArtifactRuleDo) Group(cols ...field.Expr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d daemonGcArtifactRuleDo) Having(conds ...gen.Condition) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d daemonGcArtifactRuleDo) Limit(limit int) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d daemonGcArtifactRuleDo) Offset(offset int) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d daemonGcArtifactRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d daemonGcArtifactRuleDo) Unscoped() *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Unscoped())
}

func (d daemonGcArtifactRuleDo) Create(values ...*models.DaemonGcArtifactRule) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d daemonGcArtifactRuleDo) CreateInBatches(values []*models.DaemonGcArtifactRule, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d daemonGcArtifactRuleDo) Save(values ...*models.DaemonGcArtifactRule) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d daemonGcArtifactRuleDo) First() (*models.DaemonGcArtifactRule, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonGcArtifactRule), nil
	}
}

func (d daemonGcArtifactRuleDo) Take() (*models.DaemonGcArtifactRule, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonGcArtifactRule), nil
	}
}

func (d daemonGcArtifactRuleDo) Last() (*models.DaemonGcArtifactRule, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonGcArtifactRule), nil
	}
}

func (d daemonGcArtifactRuleDo) Find() ([]*models.DaemonGcArtifactRule, error) {
	result, err := d.DO.Find()
	return result.([]*models.DaemonGcArtifactRule), err
}

func (d daemonGcArtifactRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DaemonGcArtifactRule, err error) {
	buf := make([]*models.DaemonGcArtifactRule, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d daemonGcArtifactRuleDo) FindInBatches(result *[]*models.DaemonGcArtifactRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d daemonGcArtifactRuleDo) Attrs(attrs ...field.AssignExpr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d daemonGcArtifactRuleDo) Assign(attrs ...field.AssignExpr) *daemonGcArtifactRuleDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d daemonGcArtifactRuleDo) Joins(fields ...field.RelationField) *daemonGcArtifactRuleDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d daemonGcArtifactRuleDo) Preload(fields ...field.RelationField) *daemonGcArtifactRuleDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d daemonGcArtifactRuleDo) FirstOrInit() (*models.DaemonGcArtifactRule, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonGcArtifactRule), nil
	}
}

func (d daemonGcArtifactRuleDo) FirstOrCreate() (*models.DaemonGcArtifactRule, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DaemonGcArtifactRule), nil
	}
}

func (d daemonGcArtifactRuleDo) FindByPage(offset int, limit int) (result []*models.DaemonGcArtifactRule, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d daemonGcArtifactRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d daemonGcArtifactRuleDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d daemonGcArtifactRuleDo) Delete(models ...*models.DaemonGcArtifactRule) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *daemonGcArtifactRuleDo) withDO(do gen.Dao) *daemonGcArtifactRuleDo {
	d.DO = *do.(*gen.DO)
	return d
}
