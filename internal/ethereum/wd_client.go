/**
  @author: decision
  @date: 2023/12/8
  @note:
**/

package ethereum

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gookit/config/v2"
	log "github.com/sirupsen/logrus"
	"math/big"
)

type WDClient struct {
	w3Client     *ethclient.Client
	contractAddr string
	chainID      *big.Int
}

// NewWDEthereumClient Create new client with provider http/ws address
func NewWDEthereumClient(providerAddr, contractAddr string) (*WDClient, error) {
	ethClient, err := ethclient.Dial(providerAddr)

	// todo: Set as params
	id := config.Int("ethereum.id", 1)
	chainId := big.NewInt(int64(id))
	if err != nil {
		log.WithError(err).Errorln("Start ethereum client failed.")
		return nil, err
	}

	return &WDClient{
		w3Client:     ethClient,
		contractAddr: contractAddr,
		chainID:      chainId,
	}, nil
}

func (c *WDClient) CallContractRequireData(reqSource []byte) (*string, error) {
	prv := config.String("ethereum.private")
	privateKey, err := crypto.HexToECDSA(prv)

	if err != nil {
		return nil, fmt.Errorf("load private key from config file failed")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {

		return nil, fmt.Errorf("get public key with private key failed")
	}

	// Obtain address nonce
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.w3Client.PendingNonceAt(context.Background(), address)

	if err != nil {
		return nil, fmt.Errorf("get nonce with address %s failed", address)
	}

	// Obtain network suggest gas price
	//gasPrice, err := c.w3Client.SuggestGasPrice(context.Background())
	gasPrice := big.NewInt(10_000_000_000)

	if err != nil {
		return nil, fmt.Errorf("get network suggest fee failed")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)

	if err != nil {
		return nil, fmt.Errorf("create auth failed")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	//auth.Nonce = big.NewInt(int64(2))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300_000)
	auth.GasPrice = gasPrice

	contractAddr := common.HexToAddress(c.contractAddr)
	instance, err := NewProxy(contractAddr, c.w3Client)

	if err != nil {
		return nil, fmt.Errorf("create new ethereum contract address failed")
	}

	transaction, err := instance.RequireOracleData(auth, reqSource)

	if err != nil {
		return nil, fmt.Errorf("send transaction failed with error: %s", err)
	}

	txHash := transaction.Hash().String()
	return &txHash, nil
}

func (c *WDClient) CallContractResponseData(encryptedKey,
	encryptedData []byte) (*string, error) {
	prv := config.String("ethereum.private")
	privateKey, err := crypto.HexToECDSA(prv)

	if err != nil {
		return nil, fmt.Errorf("load private key from config file failed")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {

		return nil, fmt.Errorf("get public key with private key failed")
	}

	// Obtain address nonce
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.w3Client.PendingNonceAt(context.Background(), address)

	if err != nil {
		return nil, fmt.Errorf("get nonce with address %s failed", address)
	}

	// Obtain network suggest gas price
	//gasPrice, err := c.w3Client.SuggestGasPrice(context.Background())
	gasPrice := big.NewInt(10_000_000_000)

	if err != nil {
		return nil, fmt.Errorf("get network suggest fee failed")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)

	if err != nil {
		return nil, fmt.Errorf("create auth failed")
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300_000)
	auth.GasPrice = gasPrice

	contractAddr := common.HexToAddress(c.contractAddr)
	instance, err := NewProxy(contractAddr, c.w3Client)

	if err != nil {
		return nil, fmt.Errorf("create new ethereum contract address failed")
	}

	transaction, err := instance.RespondOracleData(auth, encryptedData,
		encryptedKey)

	if err != nil {
		return nil, fmt.Errorf("send transaction failed")
	}

	txHash := transaction.Hash().String()
	return &txHash, nil
}

func (c *WDClient) ListenContractEvent(reqSource []byte, topics [][]common.Hash) (
	*ProxyDataResponse, error) {
	contractAddr := common.HexToAddress(c.contractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		Topics:    topics,
	}

	events := make(chan types.Log)
	sub, err := c.w3Client.SubscribeFilterLogs(context.Background(), query, events)

	if err != nil {
		return nil, fmt.Errorf("listen to address %s failed", contractAddr)
	}

	log.Infof("Start listen address %s ...", contractAddr)
	instance, err := NewProxy(contractAddr, c.w3Client)

	if err != nil {
		return nil, fmt.Errorf("start ethereum instance failed")
	}

	for {
		select {
		case err := <-sub.Err():
			log.WithError(err).Errorln("Log filter errored.")
		case event := <-events:
			response, err := instance.ParseDataResponse(event)

			if err != nil {
				log.WithError(err).Errorln(
					"Parse event to respond struct failed.")
				continue
			}

			if !bytes.Equal(reqSource, response.ReqSource) {
				log.Warningln("Encrypted key not equal.")
				continue
			}

			return response, nil
		}
	}
}

func (c *WDClient) OracleListenContractEvent(
	eventChan chan *ProxyDataRequest) (
	*ProxyDataResponse, error) {
	contractAddr := common.HexToAddress(c.contractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	events := make(chan types.Log)
	sub, err := c.w3Client.SubscribeFilterLogs(context.Background(), query, events)

	if err != nil {
		return nil, fmt.Errorf("listen to address %s failed", contractAddr)
	}

	log.Infof("Start listen address %s ...", contractAddr)
	instance, err := NewProxy(contractAddr, c.w3Client)

	if err != nil {
		return nil, fmt.Errorf("start ethereum instance failed")
	}

	for {
		select {
		case err := <-sub.Err():
			log.WithError(err).Errorln("Log filter errored.")
		case event := <-events:
			request, err := instance.ParseDataRequest(event)

			if err != nil {
				log.WithError(err).Errorln(
					"Parse event to respond struct failed.")
				continue
			}

			eventChan <- request
		}
	}
}
