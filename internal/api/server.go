package api

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/Anirudh9794/myblockchain/internal/bchain"
	"github.com/Anirudh9794/myblockchain/internal/redis"
)

// local instance of pubsub
var localPubSub redis.ChainPubSub

// StartServer creates a blockchain api
func StartServer() {
	router := httprouter.New()

	b := bchain.CreateBlockchain()
	localPubSub = redis.CreateBPS(b)

	defer localPubSub.CloseChannel()

	go localPubSub.HandleMessage()

	//experimental code
	time.Sleep(1 * time.Second)
	localPubSub.Blockchain.AppendBlock("ani")
	if err := localPubSub.BroadcastChain(); err != nil {
		log.Println("Error occured:", err.Error())
	}
	time.Sleep(5 * time.Second)
	log.Println(localPubSub.Blockchain.Chain)

	router.GET(BasePath+Chain, GetChain)
	router.POST(BasePath+Mining, AppendToChain)

	log.Println("Listening on", EndPoint+":"+Port)
	log.Println(http.ListenAndServe(EndPoint+":"+Port, router))
}
