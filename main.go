package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type servicePollFunction = func() (bool, error)

var services = map[string](servicePollFunction){
	"mysql": pollMySQL,
	"redis": pollRedis,
}

func main() {
	sleepInterval := getEnvVar("INTERVAL", "2")
	sleepIntervalInt, err := strconv.Atoi(sleepInterval)
	if err != nil {
		fmt.Printf("INTERVAL is not a valid integer: %s\n", err)
		os.Exit(1)
	}

	desiredServices := getEnvVar("SERVICES", "")
	if desiredServices == "" {
		fmt.Printf("SERVICES is not set\n")
		os.Exit(1)
	}

	servicesArray := strings.Split(desiredServices, ",")
	for _, serviceName := range servicesArray {
		servicePollFunc, ok := services[serviceName]
		if !ok {
			fmt.Printf("Service '%s' is not supported\n", serviceName)
			os.Exit(1)
		}

		fmt.Printf("Waiting for %s...\n", serviceName)

		err := pollService(sleepIntervalInt, servicePollFunc)
		if err != nil {
			os.Exit(1)
		}
	}
}

func pollService(interval int, servicePollFunc servicePollFunction) error {
	for {
		retry, err := servicePollFunc()
		if retry == false && err == nil {
			return nil
		}

		fmt.Printf("Error: %s\n", err)
		if retry == true {
			time.Sleep(time.Duration(interval) * time.Second)
		} else {
			return err
		}
	}
}

func getEnvVar(name string, def string) string {
	value := os.Getenv(name)
	if value == "" {
		return def
	}
	return value
}
