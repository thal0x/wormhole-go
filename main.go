package main

import (
	"context"
	"fmt"

	"wormhole-go/generated/wormhole"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	client := rpc.New("https://mainnet.helius-rpc.com/?api-key=cd780b8b-89d7-4efa-be29-299a6e6f757a")

	txSig := solana.MustSignatureFromBase58("4JeezAnzy7tyUNV9uDnpUFb94ybtacC8tRqMtodt5i1TWefRj229DN5ZGHv96RpLpJNjovsroeDUqx6kradk2Ba9")

	parsedTx, err := client.GetParsedTransaction(context.Background(), txSig, &rpc.GetParsedTransactionOpts{})
	if err != nil {
		panic(err)
	}

	accountMetaList := make([]*solana.AccountMeta, len(parsedTx.Transaction.Message.AccountKeys))
	for i, key := range parsedTx.Transaction.Message.AccountKeys {
		accountMetaList[i] = &solana.AccountMeta{
			PublicKey:  key.PublicKey,
			IsSigner:   key.Signer,
			IsWritable: key.Writable,
		}
	}

	for _, ix := range parsedTx.Transaction.Message.Instructions {
		if ix.ProgramId != solana.MustPublicKeyFromBase58("wormDTUJ6AWPNvk59vGQbDvGJmqbDTdgWgAqcLBCgUb") {
			continue
		}

		// fmt.Println(ix.Data)

		instruction, err := wormhole.DecodeInstruction(accountMetaList, ix.Data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(instruction)
	}
}
