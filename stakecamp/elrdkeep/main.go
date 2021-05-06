package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/urfave/cli/v2"
)

func run(host string) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	// First we get status
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/node/status", host), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	status := &statusReply{}

	err = json.Unmarshal(data, status)
	if err != nil {
		return err
	}

	if status.Code != "successful" {
		return fmt.Errorf("invalid response from node %s", status.Code)
	}

	pubKey := status.Data.Metrics.ErdPublicKeyBlockSign

	// Heartbeat
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

	nn := heatbeatsReply{}
	err = json.Unmarshal(data, &nn)
	if err != nil {
		return err
	}
	if nn.Code != "successful" {
		return fmt.Errorf("invalid response from node %s", nn.Code)
	}

	if len(nn.Data.Heartbeats) == 0 {
		return fmt.Errorf("no node was found")
	}

	var beat *Heartbeat
	for _, x := range nn.Data.Heartbeats {
		if x.PublicKey != pubKey {
			continue
		}
		beat = &x
	}
	if beat == nil {
		return fmt.Errorf("no public key was found with pears")
	}

	// 10min no hartbeat, that smells
	if beat.TimeStamp.UTC().Before(time.Now().Add(-time.Minute * 10).UTC()) {
		return fmt.Errorf("for 10 mins there was no hearbeat")
	}

	return nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Usage: "rest api host",
				Value: "localhost:8080",
			},
		},
		Action: func(c *cli.Context) error {
			err := run(c.String("host"))
			if err == nil {
				os.Exit(0)
			}
			log.Fatalf("healt failed check: %s", err.Error())
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
