// Code generated from Command.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Command

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CommandParser struct {
	*antlr.BaseParser
}

var CommandParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func commandParserInit() {
	staticData := &CommandParserStaticData
	staticData.LiteralNames = []string{
		"", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "Key", "Whitespace", "String",
	}
	staticData.RuleNames = []string{
		"command", "bin", "file", "arg", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 32, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		1, 0, 1, 0, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 28, 8, 3, 1, 4, 1, 4, 1, 4,
		0, 0, 5, 0, 2, 4, 6, 8, 0, 0, 29, 0, 10, 1, 0, 0, 0, 2, 17, 1, 0, 0, 0,
		4, 19, 1, 0, 0, 0, 6, 27, 1, 0, 0, 0, 8, 29, 1, 0, 0, 0, 10, 11, 3, 2,
		1, 0, 11, 13, 3, 4, 2, 0, 12, 14, 3, 6, 3, 0, 13, 12, 1, 0, 0, 0, 14, 15,
		1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 1, 1, 0, 0, 0,
		17, 18, 5, 4, 0, 0, 18, 3, 1, 0, 0, 0, 19, 20, 5, 4, 0, 0, 20, 5, 1, 0,
		0, 0, 21, 22, 5, 2, 0, 0, 22, 28, 3, 8, 4, 0, 23, 24, 5, 2, 0, 0, 24, 25,
		5, 1, 0, 0, 25, 28, 3, 8, 4, 0, 26, 28, 5, 2, 0, 0, 27, 21, 1, 0, 0, 0,
		27, 23, 1, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 7, 1, 0, 0, 0, 29, 30, 5, 4,
		0, 0, 30, 9, 1, 0, 0, 0, 2, 15, 27,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CommandParserInit initializes any static state used to implement CommandParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCommandParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CommandParserInit() {
	staticData := &CommandParserStaticData
	staticData.once.Do(commandParserInit)
}

// NewCommandParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCommandParser(input antlr.TokenStream) *CommandParser {
	CommandParserInit()
	this := new(CommandParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CommandParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Command.g4"

	return this
}

// CommandParser tokens.
const (
	CommandParserEOF        = antlr.TokenEOF
	CommandParserT__0       = 1
	CommandParserKey        = 2
	CommandParserWhitespace = 3
	CommandParserString_    = 4
)

// CommandParser rules.
const (
	CommandParserRULE_command = 0
	CommandParserRULE_bin     = 1
	CommandParserRULE_file    = 2
	CommandParserRULE_arg     = 3
	CommandParserRULE_value   = 4
)

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Bin() IBinContext
	File() IFileContext
	AllArg() []IArgContext
	Arg(i int) IArgContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommandParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Bin() IBinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinContext)
}

func (s *CommandContext) File() IFileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileContext)
}

func (s *CommandContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *CommandParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CommandParserRULE_command)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Bin()
	}
	{
		p.SetState(11)
		p.File()
	}
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == CommandParserKey {
		{
			p.SetState(12)
			p.Arg()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinContext is an interface to support dynamic dispatch.
type IBinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() antlr.TerminalNode

	// IsBinContext differentiates from other interfaces.
	IsBinContext()
}

type BinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinContext() *BinContext {
	var p = new(BinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_bin
	return p
}

func InitEmptyBinContext(p *BinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_bin
}

func (*BinContext) IsBinContext() {}

func NewBinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinContext {
	var p = new(BinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommandParserRULE_bin

	return p
}

func (s *BinContext) GetParser() antlr.Parser { return s.parser }

func (s *BinContext) String_() antlr.TerminalNode {
	return s.GetToken(CommandParserString_, 0)
}

func (s *BinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.EnterBin(s)
	}
}

func (s *BinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.ExitBin(s)
	}
}

func (p *CommandParser) Bin() (localctx IBinContext) {
	localctx = NewBinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CommandParserRULE_bin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(17)
		p.Match(CommandParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() antlr.TerminalNode

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommandParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) String_() antlr.TerminalNode {
	return s.GetToken(CommandParserString_, 0)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *CommandParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CommandParserRULE_file)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(CommandParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() antlr.TerminalNode
	Value() IValueContext

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommandParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Key() antlr.TerminalNode {
	return s.GetToken(CommandParserKey, 0)
}

func (s *ArgContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *CommandParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CommandParserRULE_arg)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Match(CommandParserKey)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(22)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Match(CommandParserKey)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.Match(CommandParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Match(CommandParserKey)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CommandParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommandParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) String_() antlr.TerminalNode {
	return s.GetToken(CommandParserString_, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommandListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *CommandParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CommandParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(CommandParserString_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
