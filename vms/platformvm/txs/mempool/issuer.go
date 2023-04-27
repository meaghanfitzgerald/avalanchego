// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package mempool

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ txs.Visitor = (*issuer)(nil)

	errCantIssueAdvanceTimeTx     = errors.New("can not issue an advance time tx")
	errCantIssueRewardValidatorTx = errors.New("can not issue a reward validator tx")
)

type issuer struct {
	m         *mempool
	tx        *txs.Tx
	timestamp time.Time
}

func (*issuer) AdvanceTimeTx(*txs.AdvanceTimeTx) error {
	return errCantIssueAdvanceTimeTx
}

func (*issuer) RewardValidatorTx(*txs.RewardValidatorTx) error {
	return errCantIssueRewardValidatorTx
}

func (i *issuer) AddValidatorTx(*txs.AddValidatorTx) error {
	if i.m.cfg.IsContinuousStakingActivated(i.timestamp) {
		i.m.addDecisionTx(i.tx)
	} else {
		i.m.addStakerTx(i.tx)
	}
	return nil
}

func (i *issuer) AddSubnetValidatorTx(*txs.AddSubnetValidatorTx) error {
	if i.m.cfg.IsContinuousStakingActivated(i.timestamp) {
		i.m.addDecisionTx(i.tx)
	} else {
		i.m.addStakerTx(i.tx)
	}
	return nil
}

func (i *issuer) AddDelegatorTx(*txs.AddDelegatorTx) error {
	if i.m.cfg.IsContinuousStakingActivated(i.timestamp) {
		i.m.addDecisionTx(i.tx)
	} else {
		i.m.addStakerTx(i.tx)
	}
	return nil
}

func (i *issuer) RemoveSubnetValidatorTx(*txs.RemoveSubnetValidatorTx) error {
	i.m.addDecisionTx(i.tx)
	return nil
}

func (i *issuer) CreateChainTx(*txs.CreateChainTx) error {
	i.m.addDecisionTx(i.tx)
	return nil
}

func (i *issuer) CreateSubnetTx(*txs.CreateSubnetTx) error {
	i.m.addDecisionTx(i.tx)
	return nil
}

func (i *issuer) ImportTx(*txs.ImportTx) error {
	i.m.addDecisionTx(i.tx)
	return nil
}

func (i *issuer) ExportTx(*txs.ExportTx) error {
	i.m.addDecisionTx(i.tx)
	return nil
}

func (i *issuer) TransformSubnetTx(*txs.TransformSubnetTx) error {
	i.m.addDecisionTx(i.tx)
	return nil
}

func (i *issuer) AddPermissionlessValidatorTx(*txs.AddPermissionlessValidatorTx) error {
	if i.m.cfg.IsContinuousStakingActivated(i.timestamp) {
		i.m.addDecisionTx(i.tx)
	} else {
		i.m.addStakerTx(i.tx)
	}
	return nil
}

func (i *issuer) AddPermissionlessDelegatorTx(*txs.AddPermissionlessDelegatorTx) error {
	if i.m.cfg.IsContinuousStakingActivated(i.timestamp) {
		i.m.addDecisionTx(i.tx)
	} else {
		i.m.addStakerTx(i.tx)
	}
	return nil
}

func (i *issuer) StopStakerTx(*txs.StopStakerTx) error {
	i.m.addDecisionTx(i.tx)
	return nil
}
