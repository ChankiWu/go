package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

// kafka consumer
func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}

	partitionList, err := consumer.Partitions("test") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("partitionlist is: ", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		// int32(partition) is 0
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetOldest)

		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}

		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		var wg sync.WaitGroup

		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				// fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
				fmt.Println(string(msg.Value))
			}
		}(pc)

		wg.Wait()
	}
}
