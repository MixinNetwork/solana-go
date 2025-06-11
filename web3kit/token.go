package web3kit

import (
	"context"
	"fmt"

	bin "github.com/gagliardetto/binary"
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

type Metadata struct {
	Name     string
	Symbol   string
	Uri      string
	Decimals int
}

func GetTokenMetadata(ctx context.Context, client *rpc.Client, mint string) (*Metadata, error) {
	assetAccountInfo, err := client.GetAccountInfo(ctx, solana.MPK(mint))
	if err != nil {
		return nil, fmt.Errorf("get account info: %w", err)
	}
	owner := assetAccountInfo.Value.Owner.String()
	if owner == solana.TokenProgramID.String() {
		mintAccount := &token.Mint{}
		if err := mintAccount.UnmarshalWithDecoder(bin.NewBinDecoder(assetAccountInfo.GetBinary())); err != nil {
			return nil, fmt.Errorf("unmarshal mint account: %w", err)
		}
		decimals := int(mintAccount.Decimals)
		metadataAddress, _, err := solana.FindTokenMetadataAddress(solana.MPK(mint))
		if err != nil {
			return nil, fmt.Errorf("derive metadata address: %w", err)
		}
		var meta token_metadata.Metadata
		err = client.GetAccountDataBorshInto(ctx, metadataAddress, &meta)
		if err != nil {
			return nil, fmt.Errorf("metadata address: %w", err)
		}
		return &Metadata{
			Name:     meta.Data.Name,
			Symbol:   meta.Data.Symbol,
			Uri:      meta.Data.Uri,
			Decimals: decimals,
		}, nil

	} else if owner == solana.Token2022ProgramID.String() {
		meta, decimals, _, err := GetToken2022Metadata(ctx, client, solana.MPK(mint), solana.Token2022ProgramID, nil)
		if err != nil {
			return nil, fmt.Errorf("get token metadata: %w", err)
		}
		return &Metadata{
			Name:     meta.Name,
			Symbol:   meta.Symbol,
			Uri:      meta.Uri,
			Decimals: int(decimals),
		}, nil

	} else {
		return nil, fmt.Errorf("unsupported token program: %s", owner)
	}
}
