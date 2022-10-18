package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/aseemwangoo/golang-programs/structs"
)

// type UserCache interface {
// 	Set(key string, value *structs.Users)
// 	Get(key string) *structs.Users
// }

// func (c *Client) GetName(key string) (string, error) {
// 	cmd := c.client.Get(key)

// 	val, err := cmd.Result()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return val, nil
// }

// func (c *Client) SetName(key string, name string) error {
// 	return c.client.Set(key, name, 0*time.Second).Err()
// }

func (c *Client) GetUser(key string) (user *structs.Users) {
	val, err := c.client.Get(key).Result()
	if err != nil {
		return nil
	}

	resp := structs.Users{}
	err = json.Unmarshal([]byte(val), &resp)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	// Publish using Redis PubSub
	if err := c.client.Publish("send-user-name", payload).Err(); err != nil {
		log.Fatal(err)
	}

	return &resp
}

func (c *Client) SetUser(key string, user structs.Users) {
	json, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	c.client.Set(key, json, 20*time.Second)
}
