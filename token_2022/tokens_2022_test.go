package web3kit

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func TestGetTokenMetadata(t *testing.T) {
	ctx := context.Background()
	client := rpc.New(rpc.MainNetBeta_RPC)
	metadata, decimal, fee, err := GetTokenMetadata(ctx, client, solana.MPK("Ey59PH7Z4BFU4HjyKnyMdWt5GGN76KazTAwQihoUXRnk"), solana.Token2022ProgramID, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("metadata: %+v %d", metadata, decimal)
	if fee != nil {
		t.Logf("fee: %+v", fee)
	}
}
