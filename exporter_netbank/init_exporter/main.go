package init_exporter


import (
"flag"
"fmt"
"os"
"os/signal"
"syscall"
"time"

"linkwan.cn/linkExporter/src/config"
"linkwan.cn/linkExporter/src/log"
"linkwan.cn/linkExporter/src/manager"

"github.com/nats-io/nats.go"
)

func main() {
	flag.Parse()
	log.InitLogger()

	log.Info("linkExporter coming...")

	nc, err := nats.Connect(config.GetConfig().NatsURL(), nats.ReconnectHandler(func(conn *nats.Conn) {
		log.Debugf("nats reconnect")

	}), nats.ClosedHandler(func(conn *nats.Conn) {
		log.Debugf("nats close")

	}), nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
		log.Errorf("nats disconnect err: %v", err)

	}), nats.ErrorHandler(func(conn *nats.Conn, subscription *nats.Subscription, err error) {
		log.Errorf("nats error:%v", err)
	}), nats.Timeout(time.Duration(5)*time.Second))

	if err != nil {
		log.Fatalf("Failed to connect to NATS: %s", err)
		os.Exit(1)
	}

	if err := manager.Start(nc, "linkExporter"); err != nil {
		log.Fatalf("Failed to consume nats: %s", err)
	}

	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	log.Error(fmt.Sprintf("Tunnel Manager service terminated: %s", err))

}
