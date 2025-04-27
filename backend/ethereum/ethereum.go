package ethereum

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func FetchEthereumBlock() string {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/98b64647179e4384b066688c1c1d6a03")
	if err != nil {
		log.Println("ETH error:", err)
		return ""
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println("ETH header error:", err)
		return ""
	}
	return "Ethereum Block: " + header.Number.String()
}
