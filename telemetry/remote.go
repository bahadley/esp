package telemetry

import (
	"fmt"
	"net"

	"github.com/bahadley/esp/log"
)

func Ingest() {

	saddr, err := net.ResolveUDPAddr("udp", ":22221")
	if err != nil {
		log.Logfatalerror(err)
	}

	conn, err := net.ListenUDP("udp", saddr)
	if err != nil {
		log.Logoutput(log.ErrPrefix, err.Error())
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		log.Logoutput(log.InfoPrefix,
			fmt.Sprintf("Recv(%s): %s", caddr, string(buf[0:n])))

		if err != nil {
			log.Logoutput(log.ErrPrefix, err.Error())
		}
	}
}
