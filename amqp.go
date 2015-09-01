// Package amqputil provides some abstractions for AMQP for easy the testing.
// The best way to test
package amqputil

type (
	Conn interface{
		Channel() (Channel, error)
		Dial(amqpuri string) error
		AutoRedial(errChan chan error, cbk func())
		Close() error
	}

	Channel interface{
		Ack(tag uint64, multiple bool) error
		Nack(tag uint64, multiple bool, requeue bool) error
		Reject(tag uint64, requeue bool) error

		Cancel(consumer string, noWait bool) error 
		ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args interface{}) error
		QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args interface{}) (Queue, error)
		QueueBind(name, key, exchange string, noWait bool, args interface{}) error
		Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args interface{}) (<-chan Delivery, error)
		Publisher
	}

	Queue interface{
		Name() string
		Messages() int
		Consumers() int
	}

	Publisher interface{
		Publish(exc, route string, msg []byte) error
	}

	Delivery interface {
		Ack(multiple bool) error
		Nack(multiple, request bool) error
		Reject(requeue bool) error

		Body() []byte
		DeliveryTag() uint64
	}
)
