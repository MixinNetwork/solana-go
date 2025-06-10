package web3kit

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func TestGetToken2022Metadata(t *testing.T) {
	ctx := context.Background()
	client := rpc.New(rpc.MainNetBeta_RPC)
	metadata, decimal, fee, err := GetToken2022Metadata(ctx, client, solana.MPK("BUSKVVHpTYSxX1pHqi16MY5n1T3U1RQ4xxr9B3Dcpump"), solana.TokenProgramID, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("metadata: %+v %d", metadata, decimal)
	if fee != nil {
		t.Logf("fee: %+v", fee)
	}
}
