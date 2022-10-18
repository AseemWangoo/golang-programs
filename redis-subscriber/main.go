package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		DB:          0, // use default DB
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})
	err := redisClient.Ping().Err()
	if err != nil {
		// Sleep for 3 seconds and wait for Redis to initialize
		time.Sleep(3 * time.Second)
		err := redisClient.Ping().Err()
		if err != nil {
			panic(err)
		}
	}
	topic := redisClient.Subscribe("send-user-name")
	channel := topic.Channel()

	for msg := range channel {
		u := &User{}
		// Unmarshal the data into the user
		err := u.UnmarshalBinary([]byte(msg.Payload))
		if err != nil {
			panic(err)
		}

		fmt.Printf("User: %v having age: %v and id: %v\n", u.Name, u.Age, u.ID)
		log.Println("Received message from " + msg.Channel + " channel.")
	}
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	CreatedTime string    `json:"created_time"`
	UpdatedTime string    `json:"updated_time"`
	Source      string    `json:"source"`
}

// UnmarshalBinary decodes the struct into a User
func (u *User) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, u); err != nil {
		return err
	}
	return nil
}
