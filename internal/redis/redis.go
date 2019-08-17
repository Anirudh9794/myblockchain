package redis

import (
	"encoding/json"

	r "github.com/go-redis/redis"

	"github.com/Anirudh9794/myblockchain/internal/bchain"
	"github.com/Anirudh9794/myblockchain/internal/pkg/utils"
)

var (
	redisEndpoint string
	redisPort     string
	redisClient   *r.Client
)

// ChainPubSub represents a blockchain pubsub implementation
type ChainPubSub struct {
	Blockchain *bchain.Blockchain
	ps         *r.PubSub
}

func init() {
	redisEndpoint = utils.GetEnvOrDefault("REDIS_ENDPOINT", RedisDefaultEndpoint)
	redisPort = utils.GetEnvOrDefault("REDIS_PORT", RedisDefaultPort)

	redisURL := redisEndpoint + ":" + redisPort

	redisClient = r.NewClient(&r.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})
}

// CreateBPS creates a pubsub subscribed for a channel
func CreateBPS(blockchain *bchain.Blockchain) ChainPubSub {
	pubsub := redisClient.Subscribe(BlockchainChannel)

	return ChainPubSub{
		Blockchain: blockchain,
		ps:         pubsub,
	}
}

// BroadcastChain broadcasts the chain for all the nodes
func (chainps ChainPubSub) BroadcastChain() error {
	b, err := json.Marshal(chainps.Blockchain.Chain)
	if err != nil {
		return err
	}

	return redisClient.Publish(BlockchainChannel, string(b)).Err()
}

// HandleMessage fetches the message broadcasted and updates he chain
func (chainps ChainPubSub) HandleMessage() {
	ch := chainps.ps.Channel()

	for msg := range ch {

		var chain []bchain.Block
		if err := json.Unmarshal([]byte(msg.Payload), &chain); err != nil {
			continue
		}
		chainps.Blockchain.ReplaceChain(chain)
	}
}

// CloseChannel closes the channel associated with the ChainPubSub
func (chainps ChainPubSub) CloseChannel() error {
	return chainps.ps.Close()
}
