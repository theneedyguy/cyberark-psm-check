package main

import (
	"log"
	"net/http"

	"github.com/kardianos/service"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

var logger service.Logger

type program struct{}

const (
	// Host name of the HTTP Server
	host = "0.0.0.0"
	// Port of the HTTP Server
	port = "80"
	// Cyber Ark PSM Name
	windowsService = "Cyber-Ark Privileged Session Manager"
)

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	http.HandleFunc("/", reportState)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func reportState(w http.ResponseWriter, r *http.Request) {
	manager, err := mgr.Connect()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("FAIL"))
	}
	serviceInstance, err := manager.OpenService(windowsService)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("FAIL"))
	} else {
		serviceState, _ := serviceInstance.Query()
		switch serviceState.State {
		case windows.SERVICE_RUNNING:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("PASS"))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("FAIL"))
		}
	}
}

func main() {
	svcConfig := &service.Config{
		Name:        "PSMSVCCheck",
		DisplayName: "CyberArk PSM Service Check",
		Description: "CyberArk PSM Service Check",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
