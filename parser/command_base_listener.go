// Code generated from Command.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Command

import "github.com/antlr4-go/antlr/v4"

// BaseCommandListener is a complete listener for a parse tree produced by CommandParser.
type BaseCommandListener struct{}

var _ CommandListener = &BaseCommandListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCommandListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCommandListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCommandListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCommandListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseCommandListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseCommandListener) ExitCommand(ctx *CommandContext) {}

// EnterBin is called when production bin is entered.
func (s *BaseCommandListener) EnterBin(ctx *BinContext) {}

// ExitBin is called when production bin is exited.
func (s *BaseCommandListener) ExitBin(ctx *BinContext) {}

// EnterFile is called when production file is entered.
func (s *BaseCommandListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseCommandListener) ExitFile(ctx *FileContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseCommandListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseCommandListener) ExitArg(ctx *ArgContext) {}

// EnterValue is called when production value is entered.
func (s *BaseCommandListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseCommandListener) ExitValue(ctx *ValueContext) {}
