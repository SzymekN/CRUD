package producer

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "messages"
	brokerAddress = "192.168.33.50:9092"
)

var (
	KafkaCtx    context.Context
	KafkaWriter *kafka.Writer
)

func getKafkaWriter() *kafka.Writer {
	return KafkaWriter
}

func getKafkaCtx() context.Context {
	return KafkaCtx
}

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go produce(ctx)
	consume(ctx)
}

func ProduceMessage(k, val string) error {
	w := getKafkaWriter()
	ctx := getKafkaCtx()

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(k),
		// create an arbitrary message payload for the value
		Value: []byte(val),
		Time:  time.Now(),
	})
	if err != nil {
		fmt.Println("could not write message " + err.Error())
		return err
	}
	return nil
}

func Setup() {
	KafkaCtx = context.Background()

	l := log.New(os.Stdout, "kafka writer: ", 0)

	KafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})
}

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		// assign the logger to the writer
		Logger: l,
	})

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("this is message " + strconv.Itoa(i) + ", at time: " + time.Now().String()),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {
	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		// assign the logger to the reader
		Logger: l,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}