package web3kit

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/assert"
)

func TestGetTokenMetadata(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	client := rpc.New(rpc.MainNetBeta_RPC)
	// FAG
	metadata, err := GetTokenMetadata(ctx, client, "BUSKVVHpTYSxX1pHqi16MY5n1T3U1RQ4xxr9B3Dcpump")
	assert.Nil(err)
	d, _ := json.Marshal(metadata)
	t.Logf("metadata: %s", d)

	// ai16z
	metadata, err = GetTokenMetadata(ctx, client, "HeLp6NuQkmYB4pYWo2zYs22mESHXPQYzXbB8n4V98jwC")
	assert.Nil(err)
	d, _ = json.Marshal(metadata)
	t.Logf("metadata: %s", d)
}
