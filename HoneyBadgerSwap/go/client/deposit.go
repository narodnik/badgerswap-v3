package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/initc3/MP-SPDZ/Scripts/hbswap/go/utils"
	"math/big"
	"os"
)

func depositETH(conn *ethclient.Client, auth *bind.TransactOpts, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.FundETH(conn, auth.From, amt)

	auth.Value = amt
	utils.Deposit(conn, auth, utils.EthAddr, amt)

	utils.GetBalance(conn, utils.EthAddr, auth.From)
}

func depositTOK(conn *ethclient.Client, auth *bind.TransactOpts, _amt string) {
	amt := utils.StrToBig(_amt)
	if amt.Cmp(big.NewInt(0)) == 0 {
		return
	}

	utils.FundToken(conn, utils.TokenAddr, auth.From, amt)

	auth.Value = big.NewInt(0)
	utils.Approve(conn, auth, utils.HbswapAddr, amt)
	utils.Deposit(conn, auth, utils.TokenAddr, amt)

	utils.GetBalance(conn, utils.TokenAddr, auth.From)
}

func main() {
	_user := os.Args[1]
	amtETH, amtTOK := os.Args[2], os.Args[3]

	conn := utils.GetEthClient(utils.HttpEndpoint)

	user := utils.GetAccount(fmt.Sprintf("account_%s", _user))

	depositETH(conn, user, amtETH)
	depositTOK(conn, user, amtTOK)
}