package stores

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	MONGOClient struct {
		*mongo.Client
		*mongo.Database
	}

	Option struct {
		ServiceName string
		URI         string
		Addr        string
		Username    string
		Password    string
		Database    string
	}
)

func MongoConnect(opt Option) (*MONGOClient, error) {
	opts := options.Client()
	if len(opt.URI) > 0 {
		opts.ApplyURI(opt.URI)
	} else {
		opts.SetDirect(true)
		opts.SetAppName(opt.ServiceName)
		opts.SetServerSelectionTimeout(time.Second * 1)
		opts.SetConnectTimeout(time.Second * 2)
		opts.SetMaxConnIdleTime(time.Second * 1)
		opts.SetMaxPoolSize(10)
		opts.ApplyURI("mongodb://" + opt.Addr)
		if (len(opt.Username) > 0) && (len(opt.Password) > 0) {
			opts.SetAuth(options.Credential{
				Username: opt.Username,
				Password: opt.Password,
			})
		}
	}

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	var ctx = context.Background()
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &MONGOClient{Client: client, Database: client.Database(opt.Database)}, nil
}
