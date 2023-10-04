// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bitmagnet-io/bitmagnet/internal/model"
)

func newContent(db *gorm.DB, opts ...gen.DOOption) content {
	_content := content{}

	_content.contentDo.UseDB(db, opts...)
	_content.contentDo.UseModel(&model.Content{})

	tableName := _content.contentDo.TableName()
	_content.ALL = field.NewAsterisk(tableName)
	_content.Type = field.NewField(tableName, "type")
	_content.Source = field.NewString(tableName, "source")
	_content.ID = field.NewString(tableName, "id")
	_content.Title = field.NewString(tableName, "title")
	_content.ReleaseDate = field.NewTime(tableName, "release_date")
	_content.ReleaseYear = field.NewField(tableName, "release_year")
	_content.Adult = field.NewField(tableName, "adult")
	_content.OriginalLanguage = field.NewField(tableName, "original_language")
	_content.OriginalTitle = field.NewField(tableName, "original_title")
	_content.Overview = field.NewField(tableName, "overview")
	_content.Runtime = field.NewField(tableName, "runtime")
	_content.Popularity = field.NewField(tableName, "popularity")
	_content.VoteAverage = field.NewField(tableName, "vote_average")
	_content.VoteCount = field.NewField(tableName, "vote_count")
	_content.SearchString = field.NewString(tableName, "search_string")
	_content.CreatedAt = field.NewTime(tableName, "created_at")
	_content.UpdatedAt = field.NewTime(tableName, "updated_at")
	_content.Collections = contentManyToManyCollections{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Collections", "model.ContentCollection"),
		MetadataSource: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Collections.MetadataSource", "model.MetadataSource"),
		},
	}

	_content.Attributes = contentHasManyAttributes{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Attributes", "model.ContentAttribute"),
		MetadataSource: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Attributes.MetadataSource", "model.MetadataSource"),
		},
	}

	_content.MetadataSource = contentBelongsToMetadataSource{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("MetadataSource", "model.MetadataSource"),
	}

	_content.fillFieldMap()

	return _content
}

type content struct {
	contentDo

	ALL              field.Asterisk
	Type             field.Field
	Source           field.String
	ID               field.String
	Title            field.String
	ReleaseDate      field.Time
	ReleaseYear      field.Field
	Adult            field.Field
	OriginalLanguage field.Field
	OriginalTitle    field.Field
	Overview         field.Field
	Runtime          field.Field
	Popularity       field.Field
	VoteAverage      field.Field
	VoteCount        field.Field
	SearchString     field.String
	CreatedAt        field.Time
	UpdatedAt        field.Time
	Collections      contentManyToManyCollections

	Attributes contentHasManyAttributes

	MetadataSource contentBelongsToMetadataSource

	fieldMap map[string]field.Expr
}

func (c content) Table(newTableName string) *content {
	c.contentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c content) As(alias string) *content {
	c.contentDo.DO = *(c.contentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *content) updateTableName(table string) *content {
	c.ALL = field.NewAsterisk(table)
	c.Type = field.NewField(table, "type")
	c.Source = field.NewString(table, "source")
	c.ID = field.NewString(table, "id")
	c.Title = field.NewString(table, "title")
	c.ReleaseDate = field.NewTime(table, "release_date")
	c.ReleaseYear = field.NewField(table, "release_year")
	c.Adult = field.NewField(table, "adult")
	c.OriginalLanguage = field.NewField(table, "original_language")
	c.OriginalTitle = field.NewField(table, "original_title")
	c.Overview = field.NewField(table, "overview")
	c.Runtime = field.NewField(table, "runtime")
	c.Popularity = field.NewField(table, "popularity")
	c.VoteAverage = field.NewField(table, "vote_average")
	c.VoteCount = field.NewField(table, "vote_count")
	c.SearchString = field.NewString(table, "search_string")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *content) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *content) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 20)
	c.fieldMap["type"] = c.Type
	c.fieldMap["source"] = c.Source
	c.fieldMap["id"] = c.ID
	c.fieldMap["title"] = c.Title
	c.fieldMap["release_date"] = c.ReleaseDate
	c.fieldMap["release_year"] = c.ReleaseYear
	c.fieldMap["adult"] = c.Adult
	c.fieldMap["original_language"] = c.OriginalLanguage
	c.fieldMap["original_title"] = c.OriginalTitle
	c.fieldMap["overview"] = c.Overview
	c.fieldMap["runtime"] = c.Runtime
	c.fieldMap["popularity"] = c.Popularity
	c.fieldMap["vote_average"] = c.VoteAverage
	c.fieldMap["vote_count"] = c.VoteCount
	c.fieldMap["search_string"] = c.SearchString
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt

}

func (c content) clone(db *gorm.DB) content {
	c.contentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c content) replaceDB(db *gorm.DB) content {
	c.contentDo.ReplaceDB(db)
	return c
}

type contentManyToManyCollections struct {
	db *gorm.DB

	field.RelationField

	MetadataSource struct {
		field.RelationField
	}
}

func (a contentManyToManyCollections) Where(conds ...field.Expr) *contentManyToManyCollections {
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

func (a contentManyToManyCollections) WithContext(ctx context.Context) *contentManyToManyCollections {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a contentManyToManyCollections) Session(session *gorm.Session) *contentManyToManyCollections {
	a.db = a.db.Session(session)
	return &a
}

func (a contentManyToManyCollections) Model(m *model.Content) *contentManyToManyCollectionsTx {
	return &contentManyToManyCollectionsTx{a.db.Model(m).Association(a.Name())}
}

type contentManyToManyCollectionsTx struct{ tx *gorm.Association }

func (a contentManyToManyCollectionsTx) Find() (result []*model.ContentCollection, err error) {
	return result, a.tx.Find(&result)
}

func (a contentManyToManyCollectionsTx) Append(values ...*model.ContentCollection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a contentManyToManyCollectionsTx) Replace(values ...*model.ContentCollection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a contentManyToManyCollectionsTx) Delete(values ...*model.ContentCollection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a contentManyToManyCollectionsTx) Clear() error {
	return a.tx.Clear()
}

func (a contentManyToManyCollectionsTx) Count() int64 {
	return a.tx.Count()
}

type contentHasManyAttributes struct {
	db *gorm.DB

	field.RelationField

	MetadataSource struct {
		field.RelationField
	}
}

func (a contentHasManyAttributes) Where(conds ...field.Expr) *contentHasManyAttributes {
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

func (a contentHasManyAttributes) WithContext(ctx context.Context) *contentHasManyAttributes {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a contentHasManyAttributes) Session(session *gorm.Session) *contentHasManyAttributes {
	a.db = a.db.Session(session)
	return &a
}

func (a contentHasManyAttributes) Model(m *model.Content) *contentHasManyAttributesTx {
	return &contentHasManyAttributesTx{a.db.Model(m).Association(a.Name())}
}

type contentHasManyAttributesTx struct{ tx *gorm.Association }

func (a contentHasManyAttributesTx) Find() (result []*model.ContentAttribute, err error) {
	return result, a.tx.Find(&result)
}

func (a contentHasManyAttributesTx) Append(values ...*model.ContentAttribute) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a contentHasManyAttributesTx) Replace(values ...*model.ContentAttribute) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a contentHasManyAttributesTx) Delete(values ...*model.ContentAttribute) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a contentHasManyAttributesTx) Clear() error {
	return a.tx.Clear()
}

func (a contentHasManyAttributesTx) Count() int64 {
	return a.tx.Count()
}

type contentBelongsToMetadataSource struct {
	db *gorm.DB

	field.RelationField
}

func (a contentBelongsToMetadataSource) Where(conds ...field.Expr) *contentBelongsToMetadataSource {
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

func (a contentBelongsToMetadataSource) WithContext(ctx context.Context) *contentBelongsToMetadataSource {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a contentBelongsToMetadataSource) Session(session *gorm.Session) *contentBelongsToMetadataSource {
	a.db = a.db.Session(session)
	return &a
}

func (a contentBelongsToMetadataSource) Model(m *model.Content) *contentBelongsToMetadataSourceTx {
	return &contentBelongsToMetadataSourceTx{a.db.Model(m).Association(a.Name())}
}

type contentBelongsToMetadataSourceTx struct{ tx *gorm.Association }

func (a contentBelongsToMetadataSourceTx) Find() (result *model.MetadataSource, err error) {
	return result, a.tx.Find(&result)
}

func (a contentBelongsToMetadataSourceTx) Append(values ...*model.MetadataSource) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a contentBelongsToMetadataSourceTx) Replace(values ...*model.MetadataSource) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a contentBelongsToMetadataSourceTx) Delete(values ...*model.MetadataSource) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a contentBelongsToMetadataSourceTx) Clear() error {
	return a.tx.Clear()
}

func (a contentBelongsToMetadataSourceTx) Count() int64 {
	return a.tx.Count()
}

type contentDo struct{ gen.DO }

type IContentDo interface {
	gen.SubQuery
	Debug() IContentDo
	WithContext(ctx context.Context) IContentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IContentDo
	WriteDB() IContentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IContentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IContentDo
	Not(conds ...gen.Condition) IContentDo
	Or(conds ...gen.Condition) IContentDo
	Select(conds ...field.Expr) IContentDo
	Where(conds ...gen.Condition) IContentDo
	Order(conds ...field.Expr) IContentDo
	Distinct(cols ...field.Expr) IContentDo
	Omit(cols ...field.Expr) IContentDo
	Join(table schema.Tabler, on ...field.Expr) IContentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IContentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IContentDo
	Group(cols ...field.Expr) IContentDo
	Having(conds ...gen.Condition) IContentDo
	Limit(limit int) IContentDo
	Offset(offset int) IContentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDo
	Unscoped() IContentDo
	Create(values ...*model.Content) error
	CreateInBatches(values []*model.Content, batchSize int) error
	Save(values ...*model.Content) error
	First() (*model.Content, error)
	Take() (*model.Content, error)
	Last() (*model.Content, error)
	Find() ([]*model.Content, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Content, err error)
	FindInBatches(result *[]*model.Content, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Content) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IContentDo
	Assign(attrs ...field.AssignExpr) IContentDo
	Joins(fields ...field.RelationField) IContentDo
	Preload(fields ...field.RelationField) IContentDo
	FirstOrInit() (*model.Content, error)
	FirstOrCreate() (*model.Content, error)
	FindByPage(offset int, limit int) (result []*model.Content, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IContentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c contentDo) Debug() IContentDo {
	return c.withDO(c.DO.Debug())
}

func (c contentDo) WithContext(ctx context.Context) IContentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c contentDo) ReadDB() IContentDo {
	return c.Clauses(dbresolver.Read)
}

func (c contentDo) WriteDB() IContentDo {
	return c.Clauses(dbresolver.Write)
}

func (c contentDo) Session(config *gorm.Session) IContentDo {
	return c.withDO(c.DO.Session(config))
}

func (c contentDo) Clauses(conds ...clause.Expression) IContentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c contentDo) Returning(value interface{}, columns ...string) IContentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c contentDo) Not(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c contentDo) Or(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c contentDo) Select(conds ...field.Expr) IContentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c contentDo) Where(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c contentDo) Order(conds ...field.Expr) IContentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c contentDo) Distinct(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c contentDo) Omit(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c contentDo) Join(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c contentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c contentDo) RightJoin(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c contentDo) Group(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c contentDo) Having(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c contentDo) Limit(limit int) IContentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c contentDo) Offset(offset int) IContentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c contentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c contentDo) Unscoped() IContentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c contentDo) Create(values ...*model.Content) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c contentDo) CreateInBatches(values []*model.Content, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c contentDo) Save(values ...*model.Content) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c contentDo) First() (*model.Content, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Take() (*model.Content, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Last() (*model.Content, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Find() ([]*model.Content, error) {
	result, err := c.DO.Find()
	return result.([]*model.Content), err
}

func (c contentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Content, err error) {
	buf := make([]*model.Content, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c contentDo) FindInBatches(result *[]*model.Content, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c contentDo) Attrs(attrs ...field.AssignExpr) IContentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c contentDo) Assign(attrs ...field.AssignExpr) IContentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c contentDo) Joins(fields ...field.RelationField) IContentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c contentDo) Preload(fields ...field.RelationField) IContentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c contentDo) FirstOrInit() (*model.Content, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) FirstOrCreate() (*model.Content, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) FindByPage(offset int, limit int) (result []*model.Content, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c contentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c contentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c contentDo) Delete(models ...*model.Content) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *contentDo) withDO(do gen.Dao) *contentDo {
	c.DO = *do.(*gen.DO)
	return c
}