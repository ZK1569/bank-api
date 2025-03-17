package service

import (
	"context"
	"sync"

	model "github.com/zk1569/test-mongodb/src/models"
	repository "github.com/zk1569/test-mongodb/src/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	lock *sync.Mutex
)

func init() {
	lock = &sync.Mutex{}
}

type Bank struct {
	bankRepository *repository.Bank
}

var singleBankInstance *Bank

func GetBankInstance() *Bank {
	if singleBankInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleBankInstance = &Bank{
			bankRepository: repository.GetBankInstance(),
		}
	}

	return singleBankInstance
}

func (self *Bank) CreateAccount(ctx context.Context, name string, firstname string, age int) (string, error) {
	user := model.User{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Firstname: firstname,
		Age:       age,
		History:   []model.History{},
	}

	err := self.bankRepository.CreateAccount(ctx, &user)
	if err != nil {
		return "", err
	}

	return user.ID.Hex(), nil
}

func (self *Bank) DeleteAccount(ctx context.Context, userID string) error {
	return self.bankRepository.DeleteAccount(ctx, userID)
}

func (self *Bank) GetAccount(ctx context.Context, userID string) (*model.User, error) {
	return self.bankRepository.GetAccount(ctx, userID)
}

func (self *Bank) AddMoney(ctx context.Context, userID string, amount int) error {
	user, err := self.GetAccount(ctx, userID)
	if err != nil {
		return err
	}
	user.History = append(user.History, model.History{Money: amount})

	return self.bankRepository.Update(ctx, user.ID.Hex(), user)
}

func (self *Bank) RemoveMoney(ctx context.Context, userID string, amount int) error {
	user, err := self.GetAccount(ctx, userID)
	if err != nil {
		return err
	}
	user.History = append(user.History, model.History{Money: -amount})

	return self.bankRepository.Update(ctx, user.ID.Hex(), user)
}
