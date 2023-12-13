/**
  @author: decision
  @date: 2023/9/7
  @note: Simulate the oracle and DON in DISOC
**/

package main

import (
	"github.com/chain-lab/go-disoc/internal/ethereum"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	eventChanSize = 128
)

func main() {
	loadConfig()

	eventChan := make(chan *ethereum.ProxyDataRequest, eventChanSize)

	//hexPrv := config.String("source.private")

	// Create ethereum client with provider and contract address
	client, err := ethereum.NewWDEthereumClient(
		config.String("ethereum.provider"),
		config.String("ethereum.contract"))
	if err != nil {
		log.WithError(err).Errorln("Start ethereum client failed.")
		return
	}

	// Start monitor contract events
	go client.OracleListenContractEvent(eventChan)
	log.Infoln("Start monitor events...")

	for {
		select {
		// Get event from event channel
		case event := <-eventChan:
			log.WithField("timestamp", time.Now().UnixMilli()).Infoln(
				"Contract request event received.")
			//key := event.ReqSource
			//if err != nil {
			//	log.WithError(err).Errorln("Decrypto key failed.")
			//}

			// Simulate obtain data from DON
			data := strconv.Itoa(time.Now().Nanosecond())
			log.Infof("Waiting for message %s send.", data)

			// Send transaction, the data on-chaining process in subsection 4.4
			txHash, err := client.CallContractResponseData(event.ReqSource, []byte(data))
			if err != nil {
				log.Warningln("Send contract response failed.")
				continue
			}

			log.WithFields(log.Fields{
				"transaction": *txHash,
				"timestamp":   time.Now().UnixMilli(),
			}).Infoln("Respond transaction send.")
		}
	}
}

func loadConfig() {
	config.WithOptions(config.ParseEnv)

	config.AddDriver(yaml.Driver)
	err := config.LoadFiles("./config.yml")
	if err != nil {
		log.WithError(err).Errorln("Load config file failed.")
	}
}
