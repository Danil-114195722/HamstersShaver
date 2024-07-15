package settings

import (
	"os"
	"errors"
	"encoding/json"
	"context"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"

	"github.com/tonkeeper/tongo/liteapi"

	tonapi "github.com/tonkeeper/tonapi-go"
)


type JsonWalletData struct {
	Hash		string	`json:"hash"`
	SeedPhrase	string	`json:"seed_phrase"`
}

// путь до JSON-файла с данными кошелька
var pathToWalletData string = "./settings/config/wallet.json"

// получение данных о кошельке из JSON-файла
func getJsonWallet() JsonWalletData {
	var jsonWallet JsonWalletData

	// открытие файла
	fileData, err := os.ReadFile(pathToWalletData)
	DieIf(err)

	// перевод данных из JSON в структуру JsonWalletData
	err = json.Unmarshal(fileData, &jsonWallet)
	DieIf(err)

	return jsonWallet
}


// создание клиента TON для tonutils-go
func getTonClientTonutils(conType string) ton.APIClientWrapped {
	var api ton.APIClientWrapped
	var tonConfigUrl string

	// тестовый конфиг
	if conType == "testnet" {
		tonConfigUrl = "https://ton-blockchain.github.io/testnet-global.config.json"
		InfoLog.Println("(tonutils) Connecting to testnet TON node...")
	// основной конфиг
	} else if conType == "mainnet" {
		tonConfigUrl = "https://ton.org/global.config.json"
		InfoLog.Println("(tonutils) Connecting to mainnet TON node...")
	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tonutils) Invalid conType parameter was given")
		DieIf(conTypeError)
		return api
	}

	client := liteclient.NewConnectionPool()

	// подключение с URL-конфигом
	err := client.AddConnectionsFromConfigUrl(context.Background(), tonConfigUrl)
	DieIf(err)

	// API-клиент с полной проверкой
	api = ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()

	InfoLog.Println("(tonutils) Successfully connected")
	return api
}

// создание клиента TON для tongo
func getTonClientTongo(conType string) *liteapi.Client {
	var client *liteapi.Client

	// тестовый конфиг
	if conType == "testnet" {
		client, err := liteapi.NewClientWithDefaultTestnet()
		DieIf(err)
		InfoLog.Println("(tongo) Connected to testnet TON node")
		return client
	// основной конфиг
	} else if conType == "mainnet" {
		client, err := liteapi.NewClientWithDefaultMainnet()
		DieIf(err)
		InfoLog.Println("(tongo) Connected to mainnet TON node")
		return client
	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tongo) Invalid conType parameter was given")
		DieIf(conTypeError)
		return client
	}
}

// создание клиента TON для tonapi-go
func getTonClientTonapi(conType string) *tonapi.Client {
	var client *tonapi.Client

	// тестовый конфиг
	if conType == "testnet" {
		client, err := tonapi.NewClient(tonapi.TestnetTonApiURL)
		DieIf(err)
		InfoLog.Println("(tonapi) Connected to testnet TON node")
		return client
	// основной конфиг
	} else if conType == "mainnet" {
		client, err := tonapi.New()
		DieIf(err)
		InfoLog.Println("(tonapi) Connected to mainnet TON node")
		return client
	// неправильный параметр конфига
	} else {
		conTypeError := errors.New("(tonapi) Invalid conType parameter was given")
		DieIf(conTypeError)
		return client
	}
}


// создание API клиента TON для tongo
// var TonAPI *liteapi.Client = getTonClient("testnet")
var TongoTonAPI *liteapi.Client = getTonClientTongo("mainnet")

// создание API клиента TON для tonutils-go
// var TonutilsTonAPI ton.APIClientWrapped = getTonClientTonutils("testnet")
// var TonutilsTonAPI ton.APIClientWrapped = getTonClientTonutils("mainnet")

// создание API клиента TON для tonapi-go
// var TonapiTonAPI *tonapi.Client = getTonClientTonapi("testnet")
var TonapiTonAPI *tonapi.Client = getTonClientTonapi("mainnet")

// данные кошелька из JSON-конфига
var JsonWallet JsonWalletData = getJsonWallet()
