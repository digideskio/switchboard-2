package switchboard

import (
	"errors"
	"fmt"
	"net"

	"github.com/pivotal-golang/lager"
)

var BridgesProvider = NewBridges
var Dialer = net.Dial

type Backend interface {
	HealthcheckUrl() string
	Bridge(clientConn net.Conn) error
	SeverConnections()
	AsJSON() BackendJSON
}

type backend struct {
	ipAddress       string
	port            uint
	healthcheckPort uint
	logger          lager.Logger
	bridges         Bridges
}

type BackendJSON struct {
	IP string `json:"ip"`
}

func NewBackend(ipAddress string, port uint, healthcheckPort uint, logger lager.Logger) Backend {
	return &backend{
		ipAddress:       ipAddress,
		port:            port,
		healthcheckPort: healthcheckPort,
		logger:          logger,
		bridges:         BridgesProvider(),
	}
}

func (b backend) HealthcheckUrl() string {
	return fmt.Sprintf("http://%s:%d", b.ipAddress, b.healthcheckPort)
}

func (b backend) Bridge(clientConn net.Conn) error {
	backendConn, err := Dialer("tcp", fmt.Sprintf("%s:%d", b.ipAddress, b.port))
	if err != nil {
		return errors.New(fmt.Sprintf("Error connection to backend: %v", err))
	}

	go func() {
		bridge := b.bridges.Create(clientConn, backendConn)
		bridge.Connect()
		b.bridges.Remove(bridge)
	}()

	return nil
}

func (b backend) SeverConnections() {
	b.bridges.RemoveAndCloseAll()
}

func (b backend) AsJSON() BackendJSON {
	return BackendJSON{
		IP: b.ipAddress,
	}
}
