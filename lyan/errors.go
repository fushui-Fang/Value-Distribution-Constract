package lyan

import "errors"

var (
	ErrInvalidArgs				= errors.New("invalid args")
	ErrInvalidFunction			= errors.New("invalid function")
	ErrEmptyAddr				= errors.New("the addr is empty")
	ErrNoAccount				= errors.New("no account found")
	ErrKeyNoData				= errors.New("state key found, but no data")
	ErrInvalidAddr				= errors.New("the addr is invalid(contains '_' or other error)")
	ErrAlreadyRegistered		= errors.New("account already registered")
	ErrAddrWithPubkey			= errors.New("addr not matches this pubkey")
	ErrTimeOut					= errors.New("the tx maps timestamp is time out(two minutes or more)")
	ErrInvalidTX				= errors.New("some transaction invalid")
	ErrInit						= errors.New("something has wrong with the init function")
	ErrInitArgNum				= errors.New("Arg's number has wrong in the init function")
	ErrAccountNotEnoughBalance	= errors.New("account has not enough balance")
	ErrNotEnoughFee				= errors.New("the tx has not enough fee")
	ErrKindExists				= errors.New("this kind is already exists")
)