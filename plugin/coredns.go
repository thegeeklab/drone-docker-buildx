package plugin

import (
	"io"
	"net"
	"os"
	"os/exec"
)

func (p Plugin) startCoredns() {
	cmd := exec.Command("coredns", "-conf", "/etc/coredns/Corefile")
	if p.settings.Daemon.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = io.Discard
		cmd.Stderr = io.Discard
	}

	go func() {
		trace(cmd)
		_ = cmd.Run()
	}()
}

func getContainerIP() (string, error) {
	netInterfaceAddrList, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, netInterfaceAddr := range netInterfaceAddrList {
		netIP, ok := netInterfaceAddr.(*net.IPNet)
		if ok && !netIP.IP.IsLoopback() && netIP.IP.To4() != nil {
			return netIP.IP.String(), nil
		}
	}

	return "", nil
}
