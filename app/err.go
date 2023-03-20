package app

import "errors"

var (
	ErrorRpcConnect = errors.New("rpc connect err")
	ErrorCcCall     = errors.New("contract call err")
	ErrorTokenId    = errors.New("token id conversion failed")

	ErrorAccountInsufficient = errors.New("insufficient account balance")
	ErrorDbWrite             = errors.New("database write err")
	ErrorDbRead              = errors.New("database read err")

	ErrorOrderDone = errors.New("order is cancelled or finalized")
	ErrorOrderSign = errors.New("order signature verification failure")
)
