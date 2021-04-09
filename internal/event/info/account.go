package info

import (
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
	"github.com/TianqiS/database_for_happyball/model"
)

type AccountEvent struct {
	*framework.BaseEvent
	Name string
	LoginPassword string // 登录密码
	AccountAvatar string // 头像
	Level int64 // 当前等级
	Delete bool // 当前账号是否注销
	Region string // 用户的地区
	Phone string // 电话
	MaxScore int64 //最大分数
	CreateAt int64
	UpdateAt int64
}

func (accountEvent *AccountEvent) ToMessage() interface{} {
	accountPb := &databaseGrpc.Account{}
	accountPb.Name = accountEvent.Name
	accountPb.LoginPassword = accountEvent.LoginPassword
	accountPb.AccountAvatar = accountEvent.AccountAvatar
	accountPb.Level = accountEvent.Level
	accountPb.Delete = accountEvent.Delete
	accountPb.Region = accountEvent.Region
	accountPb.Phone = accountEvent.Phone
	accountPb.MaxScore = accountEvent.MaxScore
	accountPb.CreateAt = accountEvent.CreateAt
	accountPb.UpdateAt = accountEvent.UpdateAt
	return accountPb
}

func (accountEvent *AccountEvent) FromMessage(message interface{}) {
	messagePb := message.(*databaseGrpc.Account)
	accountEvent.Name = messagePb.Name
	accountEvent.LoginPassword = messagePb.LoginPassword
	accountEvent.AccountAvatar = messagePb.AccountAvatar
	accountEvent.Level = messagePb.Level
	accountEvent.Delete = messagePb.Delete
	accountEvent.Region = messagePb.Region
	accountEvent.Phone = messagePb.Phone
	accountEvent.MaxScore = messagePb.MaxScore
	accountEvent.CreateAt = messagePb.CreateAt
	accountEvent.UpdateAt = messagePb.UpdateAt
}

func (accountEvent *AccountEvent) ToModel() interface{} {
	account := &model.Account{}
	account.Name = accountEvent.Name
	account.LoginPassword = accountEvent.LoginPassword
	account.Level = accountEvent.Level
	account.Delete = accountEvent.Delete
	account.Region = accountEvent.Region
	account.Phone = accountEvent.Phone
	account.MaxScore = accountEvent.MaxScore
	account.CreateAt = accountEvent.CreateAt
	account.UpdateAt = accountEvent.UpdateAt
	return account
}
