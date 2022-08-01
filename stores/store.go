package stores

import (
	"fg/x"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store[T any] interface {
	FindALL() ([]*T, error)
	FindKeyValue(kay string, value any) ([]*T, error)
	FindSkip(offset, limit int64) ([]*T, int64, error)
	FindByID(id primitive.ObjectID) (*T, error)
	Create(t T) error
	CreateMany(t []T) error
	Update(id primitive.ObjectID, t T) error
	Delete(id primitive.ObjectID) error
}

type mgo[T any] struct {
	c   *mongo.Collection
	ctx x.Context
}

func NewStore[T any](mgc *MONGOClient) Store[T] {
	var t T
	var name = reflect.TypeOf(t).Name()
	name = strings.ToLower(name)

	return &mgo[T]{c: mgc.Collection(name), ctx: x.ContextTimeoutDefault()}
}

func (m *mgo[T]) FindALL() ([]*T, error) {
	ctx, cancel := m.ctx()
	defer cancel()

	cur, err := m.c.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer func() { err = cur.Close(ctx) }()

	var ts []*T
	if err = cur.All(ctx, &ts); err != nil {
		return nil, err
	}
	if ts == nil {
		ts = []*T{}
	}

	return ts, nil
}

func (m *mgo[T]) FindKeyValue(kay string, value any) ([]*T, error) {
	ctx, cancel := m.ctx()
	defer cancel()

	cur, err := m.c.Find(ctx, bson.D{primitive.E{Key: kay, Value: value}})
	if err != nil {
		return nil, err
	}

	defer func() { err = cur.Close(ctx) }()

	var ts []*T
	if err = cur.All(ctx, &ts); err != nil {
		return nil, err
	}
	if ts == nil {
		ts = []*T{}
	}

	return ts, nil
}

func (m *mgo[T]) FindSkip(offset, limit int64) ([]*T, int64, error) {
	ctx, cancel := m.ctx()
	defer cancel()
	var count int64

	cur, err := m.c.Find(ctx, bson.D{}, &options.FindOptions{Sort: bson.M{"created_at": -1}, Skip: &offset, Limit: &limit})
	if err != nil {
		return nil, count, err
	}

	defer func() { err = cur.Close(ctx) }()

	var ts []*T
	if err = cur.All(ctx, &ts); err != nil {
		return nil, count, err
	}
	if ts == nil {
		ts = []*T{}
	}

	count, err = m.c.CountDocuments(ctx, bson.D{})
	if err != nil {
		return ts, count, err
	}

	return ts, count, nil
}

func (m *mgo[T]) FindByID(id primitive.ObjectID) (*T, error) {
	ctx, cancel := m.ctx()
	defer cancel()

	var t T
	if err := m.c.FindOne(ctx, bson.M{"_id": id}).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *mgo[T]) Create(t T) error {
	ctx, cancel := m.ctx()
	defer cancel()

	_, err := m.c.InsertOne(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func (m *mgo[T]) CreateMany(ts []T) error {
	ctx, cancel := m.ctx()
	defer cancel()

	_, err := m.c.InsertMany(ctx, []interface{}{ts})
	if err != nil {
		return err
	}

	return nil
}

func (m *mgo[T]) Update(id primitive.ObjectID, t T) error {
	ctx, cancel := m.ctx()
	defer cancel()

	_, err := m.c.UpdateByID(ctx, id, t)
	if err != nil {
		return err
	}

	return nil
}

func (m *mgo[T]) UpdateBy(id primitive.ObjectID, update any) error {
	ctx, cancel := m.ctx()
	defer cancel()

	//var upsert = true
	//var after = options.After
	//var opt = options.FindOneAndUpdateOptions{
	//	ReturnDocument: &after,
	//	Upsert:         &upsert,
	//}

	err := m.c.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.D{{Key: "$set", Value: update}}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (m *mgo[T]) Delete(id primitive.ObjectID) error {
	ctx, cancel := m.ctx()
	defer cancel()

	_, err := m.c.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
