package main

func main() {
	consumer := NewKafka()
	consumer.Connect()
	for {

		NewProducer()
	}
}
