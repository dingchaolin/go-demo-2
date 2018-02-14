package main

import (
	"context"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"github.com/Shopify/sarama"
	elastic "gopkg.in/olivere/elastic.v5"
	"time"
)

func indexName()string{
	//这个时间必须是20060102
	date := time.Now().Format("20060102")
	return "topic-" + date
}

func main(){
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	config.Consumer.Offsets.Initial = sarama.OffsetOldest//从下次消费的地方开发 这个参数会记录上次消费到的地方
	consumer, err := cluster.NewConsumer([]string{"127.0.0.1:9092"}, "myTopic", []string{"myTopic"},config)
	partition, ok := <-consumer.Partitions()
	if err != nil && ok{
		log.Fatal( err )
	}

	esclient, err := elastic.NewClient( elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Fatal( err )
	}

	for{
		select {
			case msg := <-partition.Messages():
				log.Print(string(msg.Value), msg.Offset)

				// elasticsearch写入
				// 所有 indexname type 必须全小写 可以有- _ 等字符
				// 第二个Index相当于库名 Type相当于表名
				// 最后一个参数写context.TODO() 即可
				_, err := esclient.Index().
					               Index(indexName()).
					               Type("topic").
					               BodyString(string(msg.Value)).
					               Do(context.TODO())
				if err != nil {
					log.Fatal(err)
				}
			case err := <-partition.Errors():
				log.Print(err)
		}
	}
}
