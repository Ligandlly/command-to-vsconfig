// Code generated from Command.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CommandLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var commandlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func commandlexerLexerInit() {
	staticData := &commandlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'='",
	}
	staticData.symbolicNames = []string{
		"", "", "Key", "Whitespace", "String",
	}
	staticData.ruleNames = []string{
		"T__0", "Key", "Whitespace", "String",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 30, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 15, 8, 1, 1, 1, 1, 1, 1, 2, 4, 2, 20,
		8, 2, 11, 2, 12, 2, 21, 1, 2, 1, 2, 1, 3, 4, 3, 27, 8, 3, 11, 3, 12, 3,
		28, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 2, 2, 0, 9, 10, 32, 32, 7, 0,
		34, 34, 39, 39, 45, 58, 65, 91, 93, 93, 95, 95, 97, 122, 32, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0,
		0, 0, 3, 14, 1, 0, 0, 0, 5, 19, 1, 0, 0, 0, 7, 26, 1, 0, 0, 0, 9, 10, 5,
		61, 0, 0, 10, 2, 1, 0, 0, 0, 11, 15, 5, 45, 0, 0, 12, 13, 5, 45, 0, 0,
		13, 15, 5, 45, 0, 0, 14, 11, 1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 15, 16, 1,
		0, 0, 0, 16, 17, 3, 7, 3, 0, 17, 4, 1, 0, 0, 0, 18, 20, 7, 0, 0, 0, 19,
		18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0,
		0, 22, 23, 1, 0, 0, 0, 23, 24, 6, 2, 0, 0, 24, 6, 1, 0, 0, 0, 25, 27, 7,
		1, 0, 0, 26, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28,
		29, 1, 0, 0, 0, 29, 8, 1, 0, 0, 0, 5, 0, 14, 21, 26, 28, 1, 6, 0, 0,
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

// CommandLexerInit initializes any static state used to implement CommandLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCommandLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CommandLexerInit() {
	staticData := &commandlexerLexerStaticData
	staticData.once.Do(commandlexerLexerInit)
}

// NewCommandLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCommandLexer(input antlr.CharStream) *CommandLexer {
	CommandLexerInit()
	l := new(CommandLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &commandlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Command.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CommandLexer tokens.
const (
	CommandLexerT__0       = 1
	CommandLexerKey        = 2
	CommandLexerWhitespace = 3
	CommandLexerString_    = 4
)
