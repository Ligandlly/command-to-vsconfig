// Code generated from Command.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Command

import "github.com/antlr4-go/antlr/v4"

// CommandListener is a complete listener for a parse tree produced by CommandParser.
type CommandListener interface {
	antlr.ParseTreeListener

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterBin is called when entering the bin production.
	EnterBin(c *BinContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitBin is called when exiting the bin production.
	ExitBin(c *BinContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
