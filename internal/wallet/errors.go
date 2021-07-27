package wallet

import (
	"errors"
)

var (
	// ErrEmptyCommandName is the error returned when a command name is empty.
	ErrEmptyCommandName = errors.New("empty command name")

	// ErrUnknownCommand is the error returned when a command name is not found.
	ErrUnknownCommand = errors.New("unknown command")

	// ErrNotEnoughArguments is returned when a command was not passed enough arguments.
	ErrNotEnoughArguments = errors.New("not enough arguments")

	// ErrMissingParam is returned when a parameter is missing.
	ErrMissingParam = errors.New("missing parameter")

	// ErrInvalidString is returned when a string is invalid.
	ErrInvalidString = errors.New("invalid string")

	// ErrEmptyParam is returned when a parameter is empty.
	ErrEmptyParam = errors.New("empty parameter")

	// ErrUnexpectedHashLength is returned when the passphrase hash length is incorrect
	ErrUnexpectedHashLength = errors.New("unexpected hash length")

	// ErrEmptyPassphrase is returned when the user supplies an empty passphrase
	ErrEmptyPassphrase = errors.New("passphrase cannot be empty")

	// ErrWalletExists is returned when trying to create a new wallet and it already exists
	ErrWalletExists = errors.New("wallet already exists")

	// ErrWalletClosed is returned when an open wallet is needed, but no wallet is open
	ErrWalletClosed = errors.New("no open wallet")
)
