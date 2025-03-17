package repository

import (
	"context"
	"sync"

	errs "github.com/zk1569/test-mongodb/src/errors"
	model "github.com/zk1569/test-mongodb/src/models"
	"github.com/zk1569/test-mongodb/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	lock *sync.Mutex
)

func init() {
	lock = &sync.Mutex{}
}

type Bank struct {
	db *utils.Database
}

var singleBankInstance *Bank

func GetBankInstance() *Bank {
	if singleBankInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleBankInstance = &Bank{
			db: utils.ConnectionToBd(),
		}
	}

	return singleBankInstance
}

func (self *Bank) CreateAccount(ctx context.Context, user *model.User) error {

	_, err := self.db.Coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (self *Bank) DeleteAccount(ctx context.Context, userID string) error {
	objectId, err1 := primitive.ObjectIDFromHex(userID)
	if err1 != nil {
		return errs.BadId
	}

	filter := bson.D{{"_id", objectId}}

	_, err := self.db.Coll.DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errs.NotFound
		}
		return err
	}

	return nil
}

func (self *Bank) GetAccount(ctx context.Context, userID string) (*model.User, error) {
	objectId, err1 := primitive.ObjectIDFromHex(userID)
	if err1 != nil {
		return nil, errs.BadId
	}

	filter := bson.D{{"_id", objectId}}

	var result model.User

	err := self.db.Coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NotFound
		}
		return nil, err
	}

	return &result, nil
}

func (self *Bank) Update(ctx context.Context, userID string, newUser *model.User) error {
	objectId, err1 := primitive.ObjectIDFromHex(userID)
	if err1 != nil {
		return errs.BadId
	}

	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", newUser}}

	_, err := self.db.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errs.NotFound
		}
		return err
	}

	return nil

}
