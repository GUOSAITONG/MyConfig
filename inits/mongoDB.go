package inits

import (
	"context"
	"fmt"
	"github.com/GUOSAITONG/MyConfig/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitMongoDB() *mongo.Collection {
	data := config.Config.MongoDB
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", data.User, data.Password, data.Host, data.Port))
	//连接到mongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	log.Println("mongodb init success")
	//由于都是同一个集合进行操作，所以在初始化连接时就选择集合，避免出现大量重复代码
	return client.Database("mongodb").Collection("user")
}
