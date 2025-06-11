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
	metadata, err := GetTokenMetadata(ctx, client, "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
	assert.Nil(err)
	d, _ := json.Marshal(metadata)
	t.Logf("metadata: %s", d)
	assert.Equal("USD Coin", metadata.Name)
	assert.Equal("USDC", metadata.Symbol)
	assert.Equal(6, metadata.Decimals)

	// ai16z
	metadata, err = GetTokenMetadata(ctx, client, "HeLp6NuQkmYB4pYWo2zYs22mESHXPQYzXbB8n4V98jwC")
	assert.Nil(err)
	d, _ = json.Marshal(metadata)
	t.Logf("metadata: %s", d)
}
