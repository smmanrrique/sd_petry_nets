package utils

import "strings"

var LocalIP3s = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002"}
var LocalIPs = []string{"127.0.1.1:5000", "127.0.1.1:5001", "127.0.1.1:5002", "127.0.1.1:5003",
	"127.0.1.1:5004", "127.0.1.1:5005", "127.0.1.1:5006"}
var RemoteIP1s = []string{"155.210.154.199:1400", "155.210.154.197:1400", "155.210.154.196:1400"}
var RemoteIP2s = []string{"155.210.154.207:1400", "155.210.154.208:1400", "155.210.154.209:1400"}
var Command = make(map[string]string)

// Connections is an array of connection
type Connections []*Connect

// Connection is a interface
type Connection interface {
	GetIDSubRed() string
	GetIp() string
	GetPort() string
	GetIDs() []string
	GetAccept() bool
}

// Connect is a struct that contains information about connection
type Connect struct {
	IDSubRed, IP, Port string
	IDs                []string
	Delays             []int
	Accept             bool
}

func (c *Connect) GetIDSubRed() string {
	return c.IDSubRed
}

func (c *Connect) GetIp() string {
	return c.IP
}

func (c *Connect) GetPort() string {
	return c.Port
}

func (c *Connect) GetAccept() bool {
	return c.Accept
}

func (c *Connect) GetIds() []string {
	return c.IDs
}

// NewConnec will create slice of Connect
func NewConnec(IPs []string) Connections {
	var connections Connections
	for _, val := range IPs {
		addr := strings.Split(val, ":")
		conn := new(Connect)
		conn.IP = addr[0]
		conn.Port = addr[1]
		conn.Accept = false
		conn.IDSubRed = val
		connections = append(connections, conn)
	}
	return connections
}

// GetConnection return connection by Index in slices
func (c Connections) GetConnection(n int) *Connect {
	for i, connect := range c {
		if i == n {
			return connect
		}
	}
	return &Connect{}
}