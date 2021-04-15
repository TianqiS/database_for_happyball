package request

import (
	"github.com/TianqiS/database_for_happyball/configs"
	databaseGrpc "github.com/TianqiS/database_for_happyball/database_grpc"
	"github.com/TianqiS/database_for_happyball/framework"
	"github.com/TianqiS/database_for_happyball/internal/event/info"
	"go.mongodb.org/mongo-driver/bson/primitive"
	proto2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
)

type AddItem struct {
	framework.BaseEvent
	AddItem int32 // 添加的Item的类型
	Item interface{}
}

func (addItemReq *AddItem) ToMessage() interface{} {
	return nil
}

func (addItemReq *AddItem) FromMessage(message interface{}) {
	addItemReqPb := message.(*databaseGrpc.AddItem)
	addItemReq.AddItem = int32(addItemReqPb.AddItem)
	addItemReq.Item = itemDecoder(addItemReqPb.Item, addItemReq.AddItem)
}

func itemDecoder(item *anypb.Any, addItem int32) interface{}{
	defer func() {
		err := recover()
		if err != nil {
			log.Println("itemDecoder error:", err)
		}
	}()
	switch addItem {
	case configs.ItemPlayer:
		result := &databaseGrpc.Account{}
		err := anypb.UnmarshalTo(item, result, proto2.UnmarshalOptions{})
		if err != nil {
			panic(err)
		}
		accountEvent := &info.AccountEvent{}
		accountEvent.FromMessage(result)
		accountEvent.ObjectId = primitive.NewObjectID().Hex()
		accountModel, err := accountEvent.ToModel()
		if err != nil {
			panic(err)
		}
		return accountModel
		break
	default:
	}
	return nil
}
