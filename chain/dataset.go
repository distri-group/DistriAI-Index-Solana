package chain

import (
	"context"
	"crypto/sha256"
	"distriai-index-solana/chain/distri_ai"
	"distriai-index-solana/common"
	"distriai-index-solana/model"
	"distriai-index-solana/utils/logs"
	"fmt"
	"github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// fetch all account data on the Solana blockchain and storage
func fetchAllDataset(out rpc.GetProgramAccountsResult) {
	var datasets []model.Dataset
	for _, keyedAcct := range out {
		acct := keyedAcct.Account
		d := new(distri_ai.Dataset)
		if err := d.UnmarshalWithDecoder(bin.NewBorshDecoder(acct.Data.GetBinary())); err != nil {
			continue
		}
		dataset := buildDatasetModel(d)
		datasets = append(datasets, dataset)
	}

	if len(datasets) > 0 {
		if dbResult := common.Db.Create(&datasets); dbResult.Error != nil {
			logs.Error(fmt.Sprintf("Database error: %s \n", dbResult.Error))
		}
	}
}

// Create a new  account
func addDataset(owner solana.PublicKey, name string) {
	nameHash := sha256.Sum256([]byte(name))
	address, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("dataset"),
			owner[:],
			nameHash[:],
		},
		distriProgramID,
	)
	if err != nil {
		logs.Error(fmt.Sprintf("FindProgramAddress error: %s \n", err))
		return
	}

	resp, err := rpcClient.GetAccountInfoWithOpts(
		context.TODO(),
		address,
		&rpc.GetAccountInfoOpts{
			Commitment: rpc.CommitmentConfirmed,
		},
	)
	if err != nil {
		logs.Warn(fmt.Sprintf("GetAccountInfoWithOpts error: %s \n", err))
		return
	}
	d := new(distri_ai.Dataset)
	if err := d.UnmarshalWithDecoder(bin.NewBorshDecoder(resp.Value.Data.GetBinary())); err != nil {
		return
	}

	aiModel := buildDatasetModel(d)
	if dbResult := common.Db.Create(&aiModel); dbResult.Error != nil {
		logs.Error(fmt.Sprintf("Database error: %s \n", dbResult.Error))
	}
}

// remove account
func removeDataset(owner solana.PublicKey, name string) {
	dbResult := common.Db.
		Where("owner = ?", owner.String()).
		Where("name = ?", name).
		Delete(&model.Dataset{})
	if dbResult.Error != nil {
		logs.Error(fmt.Sprintf("Database error: %s \n", dbResult.Error))
	}
}