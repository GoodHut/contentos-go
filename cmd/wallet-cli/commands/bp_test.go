package commands

import (
	"github.com/coschain/contentos-go/cmd/wallet-cli/commands/utils/mock"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet/mock"
	"github.com/coschain/contentos-go/rpc/mock_grpcpb"
	"github.com/coschain/contentos-go/rpc/pb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBpRegisterWithoutFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"register", "initminer", "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil)
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestBpRegisterWithUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"register", "initminer", "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv", "--url", "http://example.com"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		bp_op := op.GetOp3()
		myassert.Equal(bp_op.Owner.Value, "initminer")
		myassert.Equal(bp_op.Url, "http://example.com")
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestBpRegisterWithDesc(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"register", "initminer", "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv", "--desc", "hello world"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		bp_op := op.GetOp3()
		myassert.Equal(bp_op.Owner.Value, "initminer")
		myassert.Equal(bp_op.Desc, "hello world")
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestBpRegisterWithFee(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"register", "initminer", "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv", "--fee", "100"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		bp_op := op.GetOp3()
		myassert.Equal(bp_op.Owner.Value, "initminer")
		myassert.Equal(bp_op.Props.AccountCreationFee.Value, uint64(100))
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestBpRegisterWithBlockSize(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	myassert := assert.New(t)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"register", "initminer", "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv", "--blocksize", "1000"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		bp_op := op.GetOp3()
		myassert.Equal(bp_op.Owner.Value, "initminer")
		myassert.Equal(bp_op.Props.MaximumBlockSize, uint32(1000))
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestBpUnRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	myassert := assert.New(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"unregister", "initminer"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		bp_op := op.GetOp4()
		myassert.Equal(bp_op.Owner.Value, "initminer")
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestBpVoteWithoutFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	myassert := assert.New(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"vote", "initminer", "initminer"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		bp_op := op.GetOp5()
		myassert.Equal(bp_op.Cancel, false)
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}

func TestBpVoteCancel(t *testing.T) {
	ctrl := gomock.NewController(t)
	myassert := assert.New(t)
	client := mock_grpcpb.NewMockApiServiceClient(ctrl)
	mywallet := mock_wallet.NewMockWallet(ctrl)
	passwordReader := mock_utils.NewMockPasswordReader(ctrl)
	cmd := BpCmd()
	cmd.SetContext("wallet", mywallet)
	cmd.SetContext("rpcclient", client)
	cmd.SetContext("preader", passwordReader)
	for _, child := range cmd.Commands() {
		child.Context = cmd.Context
	}
	cmd.SetArgs([]string{"vote", "initminer", "initminer", "-c"})
	priv_account := &wallet.PrivAccount{
		Account: wallet.Account{
			Name:   "initminer",
			PubKey: "COS6oKUcS7jNfPk48SEHENfeHbkWWjH7QAJt6C5tzGyL46yTWWBBv",
		},
		PrivKey: "27Pah3aJ8XbaQxgU1jxmYdUzWaBbBbbxLbZ9whSH9Zc8GbPMhw",
	}
	mywallet.EXPECT().GetUnlockedAccount("initminer").Return(priv_account, true)
	resp := &grpcpb.BroadcastTrxResponse{Status: 1, Msg: "success"}
	client.EXPECT().BroadcastTrx(gomock.Any(), gomock.Any()).Return(resp, nil).Do(func(context interface{}, req *grpcpb.BroadcastTrxRequest) {
		op := req.Transaction.Trx.Operations[0]
		bp_op := op.GetOp5()
		myassert.Equal(bp_op.Cancel, true)
	})
	_, err := cmd.ExecuteC()
	if err != nil {
		t.Error(err)
	}
}