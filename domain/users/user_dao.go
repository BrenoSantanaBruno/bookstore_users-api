package users

import (
	"fmt"
	"time"

	"github.com/LinuxLoverCoder/bookstore_users-api/utils/date_utils"
	"github.com/LinuxLoverCoder/bookstore_users-api/utils/errors"
)

//https://github.com/nvisibleinc/go-ari-library/blob/8d54a15dc2a4620195d7e8e74720919f754e1977/rabbitmq.go#L71-L100
//
//func (r *RabbitMQ) StartConsumer(topic string) (chan []byte, error) {
//	c := make(chan []byte)
//	channel, err := r.consumerConn.Channel()
//	if err != nil {
//		return nil, err
//	}
//	queue, err := channel.QueueDeclare(
//		topic, // name of queue
//		true,  // durable
//		false, // delete when unused
//		false, // exclusive
//		true,  // nowait
//		nil)   // arguments
//
//	if err != nil {
//		return nil, err
//	}
//	deliveries, err := channel.Consume(queue.Name, "", false, false, true, true, nil)
//	if err != nil {
//		return nil, err
//	}
//	go func(deliveries <-chan amqp.Delivery, c chan []byte) {
//		for d := range deliveries {
//			c <- d.Body
//			d.Ack(false) // false does *not* mean don't acknowledge, see library docs for details
//		}
//	}(deliveries, c)
//
//	return c, nil
//}

var (
	usersDB = make(map[int64]*User)
)

func something() {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.FirstName)
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %d already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d does not exist", user.Id))
	}
	now := time.Now().UTC()
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
