package mongo

import (
	"context"

	"bws/pkg/trace"

	"go.mongodb.org/mongo-driver/mongo"
)

func RunTransaction(fn func(mongo.SessionContext) error) (err error) {
	return RunTransactionWithContext(context.Background(), fn)
}

func RunTransactionWithContext(ctx context.Context, fn func(mongo.SessionContext) error) (err error) {
	c, err := GetMongoClient()
	if err != nil {
		return err
	}

	s, err := c.StartSession()
	if err != nil {
		return trace.TraceError(err)
	}

	if err := s.StartTransaction(); err != nil {
		return trace.TraceError(err)
	}

	if err := mongo.WithSession(ctx, s, func(sc mongo.SessionContext) error {
		if err := fn(sc); err != nil {
			return trace.TraceError(err)
		}
		if err = s.CommitTransaction(sc); err != nil {
			return trace.TraceError(err)
		}
		return nil
	}); err != nil {
		return trace.TraceError(err)
	}

	return nil
}
