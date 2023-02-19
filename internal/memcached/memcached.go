package memcached

import (
	"bufio"
	"einride_test/internal"
	"fmt"
	"net"
)

type MutexMapMemcachedAdapter struct {
	mutexMap internal.MutexMapPort
}

func (m *MutexMapMemcachedAdapter) Get(key string) string {

	return "get " + key + "\r\n"
}

func (m *MutexMapMemcachedAdapter) Set(key string, value string) string {
	return "set " + key + " 0 0 " + fmt.Sprint((len(value))) + "\r\n" + value + "\r\n"
}

func NewApplication(mutexMap internal.MutexMapPort) *MutexMapMemcachedAdapter {
	return &MutexMapMemcachedAdapter{mutexMap: mutexMap}
}

func (m *MutexMapMemcachedAdapter) handleRequests(conn net.Conn, request string) {
	fmt.Fprint(conn, request)
	print("Sent: " + request)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)
}

func (m *MutexMapMemcachedAdapter) Listen() {
	conn, err := net.Dial("tcp", "localhost:11211")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	println("Connected to memcached server")

	m.handleRequests(conn, m.Set("key", "value"))
	m.handleRequests(conn, m.Get("key"))
}
