// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:       db,
		Tag:      newTag(db, opts...),
		Todo:     newTodo(db, opts...),
		TodoTag:  newTodoTag(db, opts...),
		UserAuth: newUserAuth(db, opts...),
		UserInfo: newUserInfo(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Tag      tag
	Todo     todo
	TodoTag  todoTag
	UserAuth userAuth
	UserInfo userInfo
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:       db,
		Tag:      q.Tag.clone(db),
		Todo:     q.Todo.clone(db),
		TodoTag:  q.TodoTag.clone(db),
		UserAuth: q.UserAuth.clone(db),
		UserInfo: q.UserInfo.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:       db,
		Tag:      q.Tag.replaceDB(db),
		Todo:     q.Todo.replaceDB(db),
		TodoTag:  q.TodoTag.replaceDB(db),
		UserAuth: q.UserAuth.replaceDB(db),
		UserInfo: q.UserInfo.replaceDB(db),
	}
}

type queryCtx struct {
	Tag      *tagDo
	Todo     *todoDo
	TodoTag  *todoTagDo
	UserAuth *userAuthDo
	UserInfo *userInfoDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Tag:      q.Tag.WithContext(ctx),
		Todo:     q.Todo.WithContext(ctx),
		TodoTag:  q.TodoTag.WithContext(ctx),
		UserAuth: q.UserAuth.WithContext(ctx),
		UserInfo: q.UserInfo.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}