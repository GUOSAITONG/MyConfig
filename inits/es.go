package inits

import (
	"fmt"
	"github.com/GUOSAITONG/MyConfig/config"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func InitEs() {
	var err error
	es := config.Config.Es
	addr := fmt.Sprintf("http://%s:%d", es.Host, es.Port)
	cfg := elasticsearch.Config{
		Addresses: []string{
			addr,
		},
		// ...
	}
	config.ES, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	log.Println("es 连接成功")
}
