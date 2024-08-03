package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"bws/pkg/db/errors"
)

type FindResultInterface interface {
	One(val any) (err error)
	All(val any) (err error)
	GetCol() (col *Col)
	GetSingleResult() (res *mongo.SingleResult)
	GetCursor() (cur *mongo.Cursor)
	GetError() (err error)
}

type FindResult struct {
	col *Col
	res *mongo.SingleResult
	cur *mongo.Cursor
	err error
}

func NewFindResult() (fr *FindResult) {
	return &FindResult{}
}

func NewFindResultWithError(err error) (fr *FindResult) {
	return &FindResult{
		err: err,
	}
}

func (fr *FindResult) GetError() (err error) {
	if fr.err != nil {
		return fr.err
	}
	return nil
}

func (fr *FindResult) One(val any) (err error) {
	if fr.err != nil {
		return fr.err
	}

	if fr.cur != nil {
		if !fr.cur.TryNext(fr.col.ctx) {
			return mongo.ErrNoDocuments
		}
		return fr.cur.Decode(val)
	}
	return fr.res.Decode(val)
}

func (fr *FindResult) All(val any) (err error) {
	if fr.err != nil {
		return fr.err
	}
	var ctx context.Context
	if fr.col == nil {
		ctx = context.Background()
	} else {
		ctx = fr.col.ctx
	}
	if fr.cur == nil {
		return errors.ErrNoCursor
	}
	if !fr.cur.TryNext(ctx) {
		return ctx.Err()
	}
	return fr.cur.All(ctx, val)
}

func (fr *FindResult) GetCol() (col *Col) {
	return fr.col
}

func (fr *FindResult) GetSingleResult() (res *mongo.SingleResult) {
	return fr.res
}

func (fr *FindResult) GetCursor() (cur *mongo.Cursor) {
	return fr.cur
}
