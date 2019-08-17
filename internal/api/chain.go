package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetChain returns the current blockchain content
func GetChain(w http.ResponseWriter, req *http.Request, pm httprouter.Params) {
	WriteJsonResponse(w, localPubSub.Blockchain.Chain)
}

// AppendToChain appends a block to the existing local blockchain
func AppendToChain(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	type appendBlockReq struct {
		Data string `json:"data"`
	}

	var reqBytes []byte
	var reqObj appendBlockReq

	if _, err := req.Body.Read(reqBytes); err != nil {
		WriteError(w, err)
		return
	}

	if err := json.Unmarshal(reqBytes, &reqObj); err != nil {
		WriteError(w, err)
		return
	}

	localPubSub.Blockchain.AppendBlock(reqObj.Data)

	localPubSub.BroadcastChain()

	WriteJsonResponse(w, localPubSub.Blockchain.Chain)
}
