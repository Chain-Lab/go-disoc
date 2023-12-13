/**
  @author: decision
  @date: 2023/9/7
  @note: Simulate the node in DISOC
**/

package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/chain-lab/go-disoc/internal/ethereum"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	loadConfig()

	log.WithField("timestamp", time.Now().UnixMilli()).
		Infoln("Node start require data on DON")

	// Create ethereum client with provider and contract address
	client, err := ethereum.NewWDEthereumClient(
		config.String("ethereum.provider"),
		config.String("ethereum.contract"),
	)
	if err != nil {
		log.WithError(err).Errorln("Start ethereum client failed.")
		return
	}

	reqSource := make([]byte, 16)
	_, _ = rand.Read(reqSource)

	// Send node request, the process of subsection 4.2
	txHash, err := client.CallContractRequireData(reqSource)
	log.Infof("Transaction send time stamp: %d", time.Now().UnixMilli())
	if err != nil {
		log.WithError(err).Errorln("Send transaction failed.")
		return
	}
	log.WithFields(log.Fields{
		"transaction": *txHash,
		"timestamp":   time.Now().UnixMilli(),
	}).Infoln("Transaction send.")

	// Monitor contract address with encryptedKey
	// the process of subsection 4.4
	event, err := client.ListenContractEvent(reqSource, nil)
	log.Infof("Receive event time stamp: %d", time.Now().UnixMilli())
	if err != nil {
		log.WithError(err).Errorln("Listen event failed.")
		return
	}

	log.Infof("Receive event: (%s)\n",
		hex.EncodeToString(event.ReqSource))

	log.WithField("timestamp", time.Now().UnixMilli()).Infof(
		"Receive data: %s\n", hex.EncodeToString(event.SourceData))
}

func loadConfig() {
	config.WithOptions(config.ParseEnv)

	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("./config.yml")
	if err != nil {
		log.WithError(err).Errorln("Load config file failed.")
	}
}
