package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func run(host string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// node status
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/node/status", host), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// read the body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var sr StatusResponse
	err = json.Unmarshal(data, &sr)
	if err != nil {
		return err
	}

	if sr.Code != "successful" {
		return fmt.Errorf("invalid response from node %s", sr.Code)
	}

	if sr.Data.Metrics.Error == "node is starting" {
		fmt.Println("node is starting")
		return nil
	}

	// node heartbeat status
	req, err = http.NewRequest("GET", fmt.Sprintf("http://%s/node/heartbeatstatus", host), nil)
	if err != nil {
		return err
	}

	resp, err = http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var hbr HeartbeatResponse
	err = json.Unmarshal(data, &hbr)
	if err != nil {
		return err
	}

	if hbr.Code != "successful" {
		return fmt.Errorf("invalid response from node %s", hbr.Code)
	}

	if len(hbr.Data.Heartbeats) == 0 {
		return fmt.Errorf("no node was found")
	}

	var hb *Heartbeat
	for _, b := range hbr.Data.Heartbeats {
		if b.PublicKey != sr.Data.Metrics.ErdPublicKeyBlockSign {
			continue
		}
		hb = &b
	}

	if hb == nil {
		return fmt.Errorf("no node was found")
	}

	if hb.TimeStamp.UTC().Before(time.Now().Add(-time.Minute * 10).UTC()) {
		return fmt.Errorf("node is not responding")
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "elrdkeep",
		Usage: "binary for node health check",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Usage: "node host address",
				Value: "0.0.0.0:8080",
			},
		},
		Action: func(c *cli.Context) error {
			err := run(c.String("host"))

			if err == nil {
				os.Exit(0)
			}

			log.Fatalf("[elrdkeep] healthcheck failed: %s", err.Error())

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
