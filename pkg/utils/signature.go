package utils

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignRecover(signedMessage string, message string) (string, error) {
	// Hash the unsigned message using EIP-191
	msg := accounts.TextHash([]byte(message))

	// Get the bytes of the signed message
	decodedMessage := hexutil.MustDecode(signedMessage)

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[crypto.RecoveryIDOffset] == 27 || decodedMessage[crypto.RecoveryIDOffset] == 28 {
		decodedMessage[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(msg, decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("could not get a public get from the message signature")
	}
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex(), nil
}

func VerifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	return from == recoveredAddr.Hex()
}
