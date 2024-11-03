package main

import (
	"sync"
)

type User struct {
	Name string
}

type IClient interface {
	GetUser(name string) (*User, error)
}

type Client struct{}

func (c Client) GetUser(name string) (*User, error) {
	// Симуляция дорогой операции
	return &User{Name: name}, nil
}

func NewClient() *Client {
	return &Client{}
}

type Decorator struct {
	client  IClient
	results map[string]*result
	mutex   sync.Mutex
}

type result struct {
	user  *User
	error error
	wg    sync.WaitGroup
}

func NewDecorator(client IClient) *Decorator {
	return &Decorator{
		client:  client,
		results: make(map[string]*result),
	}
}

func (d *Decorator) GetUser(name string) (*User, error) {
	d.mutex.Lock()
	res, exists := d.results[name]
	if !exists {
		res = &result{}
		res.wg.Add(1)
		d.results[name] = res
		d.mutex.Unlock()

		user, err := d.client.GetUser(name)
		println("calling api")
		res.user = user
		res.error = err

		res.wg.Done()
	} else {
		d.mutex.Unlock()
		res.wg.Wait()
	}

	if res.error != nil {
		return nil, res.error
	}
	return res.user, nil
}

func main() {
	client := NewClient()
	changedClient := NewDecorator(client)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value, err := changedClient.GetUser("Vasya")
			if err != nil {
				// Обработка ошибки
			}
			println(value.Name)
		}()
	}

	wg.Wait()
}
