package main

import (
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"github.com/Shopify/sarama"
)
func main(){
	config := cluster.NewConfig()
	config.Group.Mode = cluster.ConsumerModePartitions
	config.Consumer.Offsets.Initial = sarama.OffsetOldest//从下次消费的地方开发 这个参数会记录上次消费到的地方
	consumer, err := cluster.NewConsumer([]string{"127.0.0.1:9092"}, "myTopic", []string{"myTopic"},config)
	partition, ok := <-consumer.Partitions()
	if err != nil && ok{
		log.Fatal( err )
	}
	for{
		select {
			case msg := <-partition.Messages():
				log.Print(string(msg.Value), msg.Offset)
			case err := <-partition.Errors():
				log.Print(err)
		}
	}
}
