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
	_, ok := m.mutexMap.GetMutexMap(key)
	if ok {
		return "get " + key + "\r\n"
	}

	return ""
}

func (m *MutexMapMemcachedAdapter) Set(key string, value string) string {
	m.mutexMap.SetMutexMap(key, []string{value})
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
	listener, err := net.Dial("tcp", "host.docker.internal:11211")

	if err != nil {
		panic(err)
	}
	defer listener.Close()
	println("Connected to memcached server")

	m.handleRequests(listener, m.Set("key", "value"))
	m.handleRequests(listener, m.Get("key"))
}
