package main

import "time"

type Heartbeat struct {
	PublicKey        string    `json:"publicKey"`
	TimeStamp        time.Time `json:"timeStamp"`
	MaxInactiveTime  string    `json:"maxInactiveTime"`
	IsActive         bool      `json:"isActive"`
	ReceivedShardID  int       `json:"receivedShardID"`
	ComputedShardID  int       `json:"computedShardID"`
	TotalUpTimeSec   int       `json:"totalUpTimeSec"`
	TotalDownTimeSec int       `json:"totalDownTimeSec"`
	VersionNumber    string    `json:"versionNumber"`
	NodeDisplayName  string    `json:"nodeDisplayName"`
	Identity         string    `json:"identity"`
	PeerType         string    `json:"peerType"`
	Nonce            int       `json:"nonce"`
	NumInstances     int       `json:"numInstances"`
	PidString        string    `json:"pidString"`
	PeerSubType      int       `json:"peerSubType"`
}

type HeartbeatResponse struct {
	Data struct {
		Heartbeats []Heartbeat `json:"heartbeats"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}
