package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://47.97.163.47:9200"),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func () {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item saver error %v : %v\n",item,err)
			}
		}
	}()
	return out, nil
}


func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
    if item.Id != "" {
    	indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}