// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package bytefeeder

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

type preludioLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var preludiolexerLexerStaticData struct {
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

func preludiolexerLexerInit() {
	staticData := &preludiolexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'func'", "'prql'", "'ret'", "'->'", "'='", "':='", "'+'", "'-'",
		"'*'", "'^'", "'/'", "'%'", "'~'", "'=='", "'!='", "'>='", "'>'", "'<='",
		"'<'", "'|'", "':'", "','", "'.'", "'$'", "'..'", "'['", "']'", "'('",
		"')'", "'{'", "'}'", "", "", "'_'", "'`'", "'\"'", "'''", "'\"\"\"'",
		"'''''", "'and'", "'or'", "'not'", "'??'", "'na'", "'@'", "'!'",
	}
	staticData.symbolicNames = []string{
		"", "FUNC", "PRQL", "RET", "ARROW", "ASSIGN", "DECLARE", "PLUS", "MINUS",
		"STAR", "EXP", "DIV", "MOD", "MODEL", "EQ", "NE", "GE", "GT", "LE",
		"LT", "BAR", "COLON", "COMMA", "DOT", "DOLLAR", "RANGE", "LBRACKET",
		"RBRACKET", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LANG", "RANG",
		"UNDERSCORE", "BACKTICK", "DOUBLE_QUOTE", "SINGLE_QUOTE", "TRIPLE_DOUBLE_QUOTE",
		"TRIPLE_SINGLE_QUOTE", "AND", "OR", "NOT", "COALESCE", "NA", "INDEXING",
		"FUNCTION_CALL", "WHITESPACE", "NEWLINE", "SINGLE_LINE_COMMENT", "BOOLEAN_LIT",
		"IDENT", "IDENT_START", "IDENT_NEXT", "INTEGER_LIT", "RANGE_LIT", "FLOAT_LIT",
		"STRING_CHAR", "STRING_LIT", "STRING_INTERP_LIT", "STRING_RAW_LIT",
		"STRING_PATH_LIT", "REGEX_LIT", "DATE_LIT", "DURATION_LIT",
	}
	staticData.ruleNames = []string{
		"FUNC", "PRQL", "RET", "ARROW", "ASSIGN", "DECLARE", "PLUS", "MINUS",
		"STAR", "EXP", "DIV", "MOD", "MODEL", "EQ", "NE", "GE", "GT", "LE",
		"LT", "BAR", "COLON", "COMMA", "DOT", "DOLLAR", "RANGE", "LBRACKET",
		"RBRACKET", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LANG", "RANG",
		"UNDERSCORE", "BACKTICK", "DOUBLE_QUOTE", "SINGLE_QUOTE", "TRIPLE_DOUBLE_QUOTE",
		"TRIPLE_SINGLE_QUOTE", "AND", "OR", "NOT", "COALESCE", "NA", "INDEXING",
		"FUNCTION_CALL", "WHITESPACE", "NEWLINE", "SINGLE_LINE_COMMENT", "BOOLEAN_LIT",
		"IDENT", "IDENT_START", "IDENT_NEXT", "INTEGER_LIT", "RANGE_LIT", "FLOAT_LIT",
		"STRING_CHAR", "STRING_LIT", "STRING_INTERP_LIT", "STRING_RAW_LIT",
		"STRING_PATH_LIT", "REGEX_LIT", "DATE_LIT", "DURATION_LIT", "DIGIT",
		"LETTER", "EXPONENT", "ESC", "UNICODE_ESCAPE", "OCTAL_ESCAPE", "HEX_ESCAPE",
		"HEXDIGIT", "STRING_INTERP_START_SINGLE", "STRING_INTERP_START_DOUBLE",
		"STRING_RAW_START_SINGLE", "STRING_RAW_START_DOUBLE", "STRING_PATH_START_SINGLE",
		"STRING_PATH_START_DOUBLE", "REGEX_START_SINGLE", "REGEX_START_DOUBLE",
		"DATE_START_SINGLE", "DATE_START_DOUBLE", "REGEX_FIRST_CHAR", "REGEX_CHAR",
		"REGEX_CLASS_CHAR", "REGEX_BACK_SEQ", "DURATION_KIND",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 64, 718, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47,
		3, 47, 299, 8, 47, 1, 47, 1, 47, 1, 48, 1, 48, 5, 48, 305, 8, 48, 10, 48,
		12, 48, 308, 9, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 3, 49, 321, 8, 49, 1, 50, 1, 50, 1, 50, 1, 50,
		5, 50, 327, 8, 50, 10, 50, 12, 50, 330, 9, 50, 1, 51, 1, 51, 3, 51, 334,
		8, 51, 1, 51, 1, 51, 1, 51, 5, 51, 339, 8, 51, 10, 51, 12, 51, 342, 9,
		51, 1, 52, 1, 52, 3, 52, 346, 8, 52, 1, 53, 4, 53, 349, 8, 53, 11, 53,
		12, 53, 350, 1, 54, 1, 54, 3, 54, 355, 8, 54, 1, 54, 1, 54, 1, 54, 3, 54,
		360, 8, 54, 1, 54, 1, 54, 1, 54, 3, 54, 365, 8, 54, 3, 54, 367, 8, 54,
		1, 55, 4, 55, 370, 8, 55, 11, 55, 12, 55, 371, 1, 55, 1, 55, 5, 55, 376,
		8, 55, 10, 55, 12, 55, 379, 9, 55, 1, 55, 3, 55, 382, 8, 55, 1, 55, 4,
		55, 385, 8, 55, 11, 55, 12, 55, 386, 1, 55, 3, 55, 390, 8, 55, 1, 55, 1,
		55, 4, 55, 394, 8, 55, 11, 55, 12, 55, 395, 1, 55, 3, 55, 399, 8, 55, 3,
		55, 401, 8, 55, 1, 56, 1, 56, 3, 56, 405, 8, 56, 1, 57, 1, 57, 5, 57, 409,
		8, 57, 10, 57, 12, 57, 412, 9, 57, 1, 57, 1, 57, 1, 57, 1, 57, 5, 57, 418,
		8, 57, 10, 57, 12, 57, 421, 9, 57, 1, 57, 1, 57, 3, 57, 425, 8, 57, 1,
		58, 1, 58, 5, 58, 429, 8, 58, 10, 58, 12, 58, 432, 9, 58, 1, 58, 1, 58,
		1, 58, 1, 58, 5, 58, 438, 8, 58, 10, 58, 12, 58, 441, 9, 58, 1, 58, 1,
		58, 3, 58, 445, 8, 58, 1, 59, 1, 59, 5, 59, 449, 8, 59, 10, 59, 12, 59,
		452, 9, 59, 1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 458, 8, 59, 10, 59, 12,
		59, 461, 9, 59, 1, 59, 1, 59, 3, 59, 465, 8, 59, 1, 60, 1, 60, 5, 60, 469,
		8, 60, 10, 60, 12, 60, 472, 9, 60, 1, 60, 1, 60, 1, 60, 1, 60, 5, 60, 478,
		8, 60, 10, 60, 12, 60, 481, 9, 60, 1, 60, 1, 60, 3, 60, 485, 8, 60, 1,
		61, 1, 61, 1, 61, 1, 61, 5, 61, 491, 8, 61, 10, 61, 12, 61, 494, 9, 61,
		1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 5, 61, 502, 8, 61, 10, 61, 12,
		61, 505, 9, 61, 1, 61, 1, 61, 3, 61, 509, 8, 61, 1, 62, 1, 62, 5, 62, 513,
		8, 62, 10, 62, 12, 62, 516, 9, 62, 1, 62, 1, 62, 1, 62, 1, 62, 5, 62, 522,
		8, 62, 10, 62, 12, 62, 525, 9, 62, 1, 62, 1, 62, 3, 62, 529, 8, 62, 1,
		63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1, 66, 3, 66,
		541, 8, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 67, 1, 67, 3, 67, 550,
		8, 67, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1,
		68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 3, 68, 568, 8, 68, 1, 69,
		1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 1, 69, 3, 69, 579, 8,
		69, 1, 70, 1, 70, 1, 70, 3, 70, 584, 8, 70, 1, 71, 1, 71, 1, 72, 1, 72,
		1, 72, 1, 73, 1, 73, 1, 73, 1, 74, 1, 74, 1, 74, 1, 75, 1, 75, 1, 75, 1,
		76, 1, 76, 1, 76, 1, 77, 1, 77, 1, 77, 1, 78, 1, 78, 1, 78, 1, 79, 1, 79,
		1, 79, 1, 80, 1, 80, 1, 80, 1, 81, 1, 81, 1, 81, 1, 82, 1, 82, 1, 82, 1,
		82, 5, 82, 622, 8, 82, 10, 82, 12, 82, 625, 9, 82, 1, 82, 3, 82, 628, 8,
		82, 1, 83, 1, 83, 1, 83, 1, 83, 5, 83, 634, 8, 83, 10, 83, 12, 83, 637,
		9, 83, 1, 83, 3, 83, 640, 8, 83, 1, 84, 1, 84, 3, 84, 644, 8, 84, 1, 85,
		1, 85, 1, 85, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1,
		86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86,
		1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1,
		86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86,
		1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1,
		86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86,
		1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 1, 86, 3, 86, 717, 8, 86, 12,
		410, 419, 430, 439, 450, 459, 470, 479, 492, 503, 514, 523, 0, 87, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41,
		21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59,
		30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77,
		39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95,
		48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56,
		113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 64,
		129, 0, 131, 0, 133, 0, 135, 0, 137, 0, 139, 0, 141, 0, 143, 0, 145, 0,
		147, 0, 149, 0, 151, 0, 153, 0, 155, 0, 157, 0, 159, 0, 161, 0, 163, 0,
		165, 0, 167, 0, 169, 0, 171, 0, 173, 0, 1, 0, 17, 2, 0, 9, 9, 32, 32, 3,
		0, 10, 10, 13, 13, 8232, 8233, 5, 0, 10, 10, 13, 13, 39, 39, 92, 92, 8232,
		8233, 2, 0, 39, 39, 92, 92, 2, 0, 34, 34, 92, 92, 1, 0, 48, 57, 2, 0, 65,
		90, 97, 122, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 9, 0, 34, 34,
		39, 39, 92, 92, 97, 98, 102, 102, 110, 110, 114, 114, 116, 116, 118, 118,
		1, 0, 48, 51, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 6, 0, 10, 10,
		13, 13, 42, 42, 47, 47, 91, 92, 8232, 8233, 5, 0, 10, 10, 13, 13, 47, 47,
		91, 92, 8232, 8233, 4, 0, 10, 10, 13, 13, 92, 93, 8232, 8233, 7, 0, 77,
		77, 100, 100, 104, 104, 109, 109, 115, 115, 119, 119, 121, 121, 764, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0,
		0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1,
		0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93,
		1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0,
		101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0,
		0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115,
		1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0,
		0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 1, 175, 1,
		0, 0, 0, 3, 180, 1, 0, 0, 0, 5, 185, 1, 0, 0, 0, 7, 189, 1, 0, 0, 0, 9,
		192, 1, 0, 0, 0, 11, 194, 1, 0, 0, 0, 13, 197, 1, 0, 0, 0, 15, 199, 1,
		0, 0, 0, 17, 201, 1, 0, 0, 0, 19, 203, 1, 0, 0, 0, 21, 205, 1, 0, 0, 0,
		23, 207, 1, 0, 0, 0, 25, 209, 1, 0, 0, 0, 27, 211, 1, 0, 0, 0, 29, 214,
		1, 0, 0, 0, 31, 217, 1, 0, 0, 0, 33, 220, 1, 0, 0, 0, 35, 222, 1, 0, 0,
		0, 37, 225, 1, 0, 0, 0, 39, 227, 1, 0, 0, 0, 41, 229, 1, 0, 0, 0, 43, 231,
		1, 0, 0, 0, 45, 233, 1, 0, 0, 0, 47, 235, 1, 0, 0, 0, 49, 237, 1, 0, 0,
		0, 51, 240, 1, 0, 0, 0, 53, 242, 1, 0, 0, 0, 55, 244, 1, 0, 0, 0, 57, 246,
		1, 0, 0, 0, 59, 248, 1, 0, 0, 0, 61, 250, 1, 0, 0, 0, 63, 252, 1, 0, 0,
		0, 65, 254, 1, 0, 0, 0, 67, 256, 1, 0, 0, 0, 69, 258, 1, 0, 0, 0, 71, 260,
		1, 0, 0, 0, 73, 262, 1, 0, 0, 0, 75, 264, 1, 0, 0, 0, 77, 268, 1, 0, 0,
		0, 79, 272, 1, 0, 0, 0, 81, 276, 1, 0, 0, 0, 83, 279, 1, 0, 0, 0, 85, 283,
		1, 0, 0, 0, 87, 286, 1, 0, 0, 0, 89, 289, 1, 0, 0, 0, 91, 291, 1, 0, 0,
		0, 93, 293, 1, 0, 0, 0, 95, 298, 1, 0, 0, 0, 97, 302, 1, 0, 0, 0, 99, 320,
		1, 0, 0, 0, 101, 322, 1, 0, 0, 0, 103, 333, 1, 0, 0, 0, 105, 345, 1, 0,
		0, 0, 107, 348, 1, 0, 0, 0, 109, 354, 1, 0, 0, 0, 111, 400, 1, 0, 0, 0,
		113, 404, 1, 0, 0, 0, 115, 424, 1, 0, 0, 0, 117, 444, 1, 0, 0, 0, 119,
		464, 1, 0, 0, 0, 121, 484, 1, 0, 0, 0, 123, 508, 1, 0, 0, 0, 125, 528,
		1, 0, 0, 0, 127, 530, 1, 0, 0, 0, 129, 534, 1, 0, 0, 0, 131, 536, 1, 0,
		0, 0, 133, 538, 1, 0, 0, 0, 135, 549, 1, 0, 0, 0, 137, 567, 1, 0, 0, 0,
		139, 578, 1, 0, 0, 0, 141, 580, 1, 0, 0, 0, 143, 585, 1, 0, 0, 0, 145,
		587, 1, 0, 0, 0, 147, 590, 1, 0, 0, 0, 149, 593, 1, 0, 0, 0, 151, 596,
		1, 0, 0, 0, 153, 599, 1, 0, 0, 0, 155, 602, 1, 0, 0, 0, 157, 605, 1, 0,
		0, 0, 159, 608, 1, 0, 0, 0, 161, 611, 1, 0, 0, 0, 163, 614, 1, 0, 0, 0,
		165, 627, 1, 0, 0, 0, 167, 639, 1, 0, 0, 0, 169, 643, 1, 0, 0, 0, 171,
		645, 1, 0, 0, 0, 173, 716, 1, 0, 0, 0, 175, 176, 5, 102, 0, 0, 176, 177,
		5, 117, 0, 0, 177, 178, 5, 110, 0, 0, 178, 179, 5, 99, 0, 0, 179, 2, 1,
		0, 0, 0, 180, 181, 5, 112, 0, 0, 181, 182, 5, 114, 0, 0, 182, 183, 5, 113,
		0, 0, 183, 184, 5, 108, 0, 0, 184, 4, 1, 0, 0, 0, 185, 186, 5, 114, 0,
		0, 186, 187, 5, 101, 0, 0, 187, 188, 5, 116, 0, 0, 188, 6, 1, 0, 0, 0,
		189, 190, 5, 45, 0, 0, 190, 191, 5, 62, 0, 0, 191, 8, 1, 0, 0, 0, 192,
		193, 5, 61, 0, 0, 193, 10, 1, 0, 0, 0, 194, 195, 5, 58, 0, 0, 195, 196,
		5, 61, 0, 0, 196, 12, 1, 0, 0, 0, 197, 198, 5, 43, 0, 0, 198, 14, 1, 0,
		0, 0, 199, 200, 5, 45, 0, 0, 200, 16, 1, 0, 0, 0, 201, 202, 5, 42, 0, 0,
		202, 18, 1, 0, 0, 0, 203, 204, 5, 94, 0, 0, 204, 20, 1, 0, 0, 0, 205, 206,
		5, 47, 0, 0, 206, 22, 1, 0, 0, 0, 207, 208, 5, 37, 0, 0, 208, 24, 1, 0,
		0, 0, 209, 210, 5, 126, 0, 0, 210, 26, 1, 0, 0, 0, 211, 212, 5, 61, 0,
		0, 212, 213, 5, 61, 0, 0, 213, 28, 1, 0, 0, 0, 214, 215, 5, 33, 0, 0, 215,
		216, 5, 61, 0, 0, 216, 30, 1, 0, 0, 0, 217, 218, 5, 62, 0, 0, 218, 219,
		5, 61, 0, 0, 219, 32, 1, 0, 0, 0, 220, 221, 5, 62, 0, 0, 221, 34, 1, 0,
		0, 0, 222, 223, 5, 60, 0, 0, 223, 224, 5, 61, 0, 0, 224, 36, 1, 0, 0, 0,
		225, 226, 5, 60, 0, 0, 226, 38, 1, 0, 0, 0, 227, 228, 5, 124, 0, 0, 228,
		40, 1, 0, 0, 0, 229, 230, 5, 58, 0, 0, 230, 42, 1, 0, 0, 0, 231, 232, 5,
		44, 0, 0, 232, 44, 1, 0, 0, 0, 233, 234, 5, 46, 0, 0, 234, 46, 1, 0, 0,
		0, 235, 236, 5, 36, 0, 0, 236, 48, 1, 0, 0, 0, 237, 238, 5, 46, 0, 0, 238,
		239, 5, 46, 0, 0, 239, 50, 1, 0, 0, 0, 240, 241, 5, 91, 0, 0, 241, 52,
		1, 0, 0, 0, 242, 243, 5, 93, 0, 0, 243, 54, 1, 0, 0, 0, 244, 245, 5, 40,
		0, 0, 245, 56, 1, 0, 0, 0, 246, 247, 5, 41, 0, 0, 247, 58, 1, 0, 0, 0,
		248, 249, 5, 123, 0, 0, 249, 60, 1, 0, 0, 0, 250, 251, 5, 125, 0, 0, 251,
		62, 1, 0, 0, 0, 252, 253, 3, 37, 18, 0, 253, 64, 1, 0, 0, 0, 254, 255,
		3, 33, 16, 0, 255, 66, 1, 0, 0, 0, 256, 257, 5, 95, 0, 0, 257, 68, 1, 0,
		0, 0, 258, 259, 5, 96, 0, 0, 259, 70, 1, 0, 0, 0, 260, 261, 5, 34, 0, 0,
		261, 72, 1, 0, 0, 0, 262, 263, 5, 39, 0, 0, 263, 74, 1, 0, 0, 0, 264, 265,
		5, 34, 0, 0, 265, 266, 5, 34, 0, 0, 266, 267, 5, 34, 0, 0, 267, 76, 1,
		0, 0, 0, 268, 269, 5, 39, 0, 0, 269, 270, 5, 39, 0, 0, 270, 271, 5, 39,
		0, 0, 271, 78, 1, 0, 0, 0, 272, 273, 5, 97, 0, 0, 273, 274, 5, 110, 0,
		0, 274, 275, 5, 100, 0, 0, 275, 80, 1, 0, 0, 0, 276, 277, 5, 111, 0, 0,
		277, 278, 5, 114, 0, 0, 278, 82, 1, 0, 0, 0, 279, 280, 5, 110, 0, 0, 280,
		281, 5, 111, 0, 0, 281, 282, 5, 116, 0, 0, 282, 84, 1, 0, 0, 0, 283, 284,
		5, 63, 0, 0, 284, 285, 5, 63, 0, 0, 285, 86, 1, 0, 0, 0, 286, 287, 5, 110,
		0, 0, 287, 288, 5, 97, 0, 0, 288, 88, 1, 0, 0, 0, 289, 290, 5, 64, 0, 0,
		290, 90, 1, 0, 0, 0, 291, 292, 5, 33, 0, 0, 292, 92, 1, 0, 0, 0, 293, 294,
		7, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 6, 46, 0, 0, 296, 94, 1, 0,
		0, 0, 297, 299, 5, 13, 0, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0,
		299, 300, 1, 0, 0, 0, 300, 301, 5, 10, 0, 0, 301, 96, 1, 0, 0, 0, 302,
		306, 5, 35, 0, 0, 303, 305, 8, 1, 0, 0, 304, 303, 1, 0, 0, 0, 305, 308,
		1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 309, 1, 0,
		0, 0, 308, 306, 1, 0, 0, 0, 309, 310, 3, 95, 47, 0, 310, 98, 1, 0, 0, 0,
		311, 312, 5, 116, 0, 0, 312, 313, 5, 114, 0, 0, 313, 314, 5, 117, 0, 0,
		314, 321, 5, 101, 0, 0, 315, 316, 5, 102, 0, 0, 316, 317, 5, 97, 0, 0,
		317, 318, 5, 108, 0, 0, 318, 319, 5, 115, 0, 0, 319, 321, 5, 101, 0, 0,
		320, 311, 1, 0, 0, 0, 320, 315, 1, 0, 0, 0, 321, 100, 1, 0, 0, 0, 322,
		328, 3, 103, 51, 0, 323, 324, 3, 45, 22, 0, 324, 325, 3, 105, 52, 0, 325,
		327, 1, 0, 0, 0, 326, 323, 1, 0, 0, 0, 327, 330, 1, 0, 0, 0, 328, 326,
		1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 102, 1, 0, 0, 0, 330, 328, 1, 0,
		0, 0, 331, 334, 3, 131, 65, 0, 332, 334, 3, 67, 33, 0, 333, 331, 1, 0,
		0, 0, 333, 332, 1, 0, 0, 0, 334, 340, 1, 0, 0, 0, 335, 339, 3, 131, 65,
		0, 336, 339, 3, 129, 64, 0, 337, 339, 3, 67, 33, 0, 338, 335, 1, 0, 0,
		0, 338, 336, 1, 0, 0, 0, 338, 337, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340,
		338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 104, 1, 0, 0, 0, 342, 340,
		1, 0, 0, 0, 343, 346, 3, 103, 51, 0, 344, 346, 3, 17, 8, 0, 345, 343, 1,
		0, 0, 0, 345, 344, 1, 0, 0, 0, 346, 106, 1, 0, 0, 0, 347, 349, 3, 129,
		64, 0, 348, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0,
		350, 351, 1, 0, 0, 0, 351, 108, 1, 0, 0, 0, 352, 355, 3, 107, 53, 0, 353,
		355, 3, 101, 50, 0, 354, 352, 1, 0, 0, 0, 354, 353, 1, 0, 0, 0, 355, 356,
		1, 0, 0, 0, 356, 359, 3, 49, 24, 0, 357, 360, 3, 107, 53, 0, 358, 360,
		3, 101, 50, 0, 359, 357, 1, 0, 0, 0, 359, 358, 1, 0, 0, 0, 360, 366, 1,
		0, 0, 0, 361, 364, 3, 41, 20, 0, 362, 365, 3, 107, 53, 0, 363, 365, 3,
		101, 50, 0, 364, 362, 1, 0, 0, 0, 364, 363, 1, 0, 0, 0, 365, 367, 1, 0,
		0, 0, 366, 361, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 110, 1, 0, 0, 0,
		368, 370, 3, 129, 64, 0, 369, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371,
		369, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 377,
		3, 45, 22, 0, 374, 376, 3, 129, 64, 0, 375, 374, 1, 0, 0, 0, 376, 379,
		1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 381, 1, 0,
		0, 0, 379, 377, 1, 0, 0, 0, 380, 382, 3, 133, 66, 0, 381, 380, 1, 0, 0,
		0, 381, 382, 1, 0, 0, 0, 382, 401, 1, 0, 0, 0, 383, 385, 3, 129, 64, 0,
		384, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386,
		387, 1, 0, 0, 0, 387, 389, 1, 0, 0, 0, 388, 390, 3, 133, 66, 0, 389, 388,
		1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 401, 1, 0, 0, 0, 391, 393, 3, 45,
		22, 0, 392, 394, 3, 129, 64, 0, 393, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0,
		0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 398, 1, 0, 0, 0, 397,
		399, 3, 133, 66, 0, 398, 397, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 401,
		1, 0, 0, 0, 400, 369, 1, 0, 0, 0, 400, 384, 1, 0, 0, 0, 400, 391, 1, 0,
		0, 0, 401, 112, 1, 0, 0, 0, 402, 405, 3, 135, 67, 0, 403, 405, 8, 2, 0,
		0, 404, 402, 1, 0, 0, 0, 404, 403, 1, 0, 0, 0, 405, 114, 1, 0, 0, 0, 406,
		410, 3, 73, 36, 0, 407, 409, 3, 113, 56, 0, 408, 407, 1, 0, 0, 0, 409,
		412, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 411, 413,
		1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 414, 3, 73, 36, 0, 414, 425, 1,
		0, 0, 0, 415, 419, 3, 71, 35, 0, 416, 418, 3, 113, 56, 0, 417, 416, 1,
		0, 0, 0, 418, 421, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 419, 417, 1, 0, 0,
		0, 420, 422, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 422, 423, 3, 71, 35, 0,
		423, 425, 1, 0, 0, 0, 424, 406, 1, 0, 0, 0, 424, 415, 1, 0, 0, 0, 425,
		116, 1, 0, 0, 0, 426, 430, 3, 145, 72, 0, 427, 429, 3, 113, 56, 0, 428,
		427, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 430, 428,
		1, 0, 0, 0, 431, 433, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 434, 3, 73,
		36, 0, 434, 445, 1, 0, 0, 0, 435, 439, 3, 147, 73, 0, 436, 438, 3, 113,
		56, 0, 437, 436, 1, 0, 0, 0, 438, 441, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0,
		439, 437, 1, 0, 0, 0, 440, 442, 1, 0, 0, 0, 441, 439, 1, 0, 0, 0, 442,
		443, 3, 71, 35, 0, 443, 445, 1, 0, 0, 0, 444, 426, 1, 0, 0, 0, 444, 435,
		1, 0, 0, 0, 445, 118, 1, 0, 0, 0, 446, 450, 3, 149, 74, 0, 447, 449, 3,
		113, 56, 0, 448, 447, 1, 0, 0, 0, 449, 452, 1, 0, 0, 0, 450, 451, 1, 0,
		0, 0, 450, 448, 1, 0, 0, 0, 451, 453, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0,
		453, 454, 3, 73, 36, 0, 454, 465, 1, 0, 0, 0, 455, 459, 3, 151, 75, 0,
		456, 458, 3, 113, 56, 0, 457, 456, 1, 0, 0, 0, 458, 461, 1, 0, 0, 0, 459,
		460, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 460, 462, 1, 0, 0, 0, 461, 459,
		1, 0, 0, 0, 462, 463, 3, 71, 35, 0, 463, 465, 1, 0, 0, 0, 464, 446, 1,
		0, 0, 0, 464, 455, 1, 0, 0, 0, 465, 120, 1, 0, 0, 0, 466, 470, 3, 153,
		76, 0, 467, 469, 3, 113, 56, 0, 468, 467, 1, 0, 0, 0, 469, 472, 1, 0, 0,
		0, 470, 471, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 471, 473, 1, 0, 0, 0, 472,
		470, 1, 0, 0, 0, 473, 474, 3, 73, 36, 0, 474, 485, 1, 0, 0, 0, 475, 479,
		3, 153, 76, 0, 476, 478, 3, 113, 56, 0, 477, 476, 1, 0, 0, 0, 478, 481,
		1, 0, 0, 0, 479, 480, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 482, 1, 0,
		0, 0, 481, 479, 1, 0, 0, 0, 482, 483, 3, 71, 35, 0, 483, 485, 1, 0, 0,
		0, 484, 466, 1, 0, 0, 0, 484, 475, 1, 0, 0, 0, 485, 122, 1, 0, 0, 0, 486,
		487, 3, 157, 78, 0, 487, 492, 3, 165, 82, 0, 488, 491, 3, 167, 83, 0, 489,
		491, 8, 3, 0, 0, 490, 488, 1, 0, 0, 0, 490, 489, 1, 0, 0, 0, 491, 494,
		1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 492, 490, 1, 0, 0, 0, 493, 495, 1, 0,
		0, 0, 494, 492, 1, 0, 0, 0, 495, 496, 3, 73, 36, 0, 496, 509, 1, 0, 0,
		0, 497, 498, 3, 159, 79, 0, 498, 503, 3, 165, 82, 0, 499, 502, 3, 167,
		83, 0, 500, 502, 8, 4, 0, 0, 501, 499, 1, 0, 0, 0, 501, 500, 1, 0, 0, 0,
		502, 505, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 503, 501, 1, 0, 0, 0, 504,
		506, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 506, 507, 3, 71, 35, 0, 507, 509,
		1, 0, 0, 0, 508, 486, 1, 0, 0, 0, 508, 497, 1, 0, 0, 0, 509, 124, 1, 0,
		0, 0, 510, 514, 3, 161, 80, 0, 511, 513, 3, 113, 56, 0, 512, 511, 1, 0,
		0, 0, 513, 516, 1, 0, 0, 0, 514, 515, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0,
		515, 517, 1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 517, 518, 3, 73, 36, 0, 518,
		529, 1, 0, 0, 0, 519, 523, 3, 163, 81, 0, 520, 522, 3, 113, 56, 0, 521,
		520, 1, 0, 0, 0, 522, 525, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 523, 521,
		1, 0, 0, 0, 524, 526, 1, 0, 0, 0, 525, 523, 1, 0, 0, 0, 526, 527, 3, 71,
		35, 0, 527, 529, 1, 0, 0, 0, 528, 510, 1, 0, 0, 0, 528, 519, 1, 0, 0, 0,
		529, 126, 1, 0, 0, 0, 530, 531, 3, 107, 53, 0, 531, 532, 3, 41, 20, 0,
		532, 533, 3, 173, 86, 0, 533, 128, 1, 0, 0, 0, 534, 535, 7, 5, 0, 0, 535,
		130, 1, 0, 0, 0, 536, 537, 7, 6, 0, 0, 537, 132, 1, 0, 0, 0, 538, 540,
		7, 7, 0, 0, 539, 541, 7, 8, 0, 0, 540, 539, 1, 0, 0, 0, 540, 541, 1, 0,
		0, 0, 541, 542, 1, 0, 0, 0, 542, 543, 3, 107, 53, 0, 543, 134, 1, 0, 0,
		0, 544, 545, 5, 92, 0, 0, 545, 550, 7, 9, 0, 0, 546, 550, 3, 137, 68, 0,
		547, 550, 3, 141, 70, 0, 548, 550, 3, 139, 69, 0, 549, 544, 1, 0, 0, 0,
		549, 546, 1, 0, 0, 0, 549, 547, 1, 0, 0, 0, 549, 548, 1, 0, 0, 0, 550,
		136, 1, 0, 0, 0, 551, 552, 5, 92, 0, 0, 552, 553, 5, 117, 0, 0, 553, 554,
		3, 143, 71, 0, 554, 555, 3, 143, 71, 0, 555, 556, 3, 143, 71, 0, 556, 557,
		3, 143, 71, 0, 557, 568, 1, 0, 0, 0, 558, 559, 5, 92, 0, 0, 559, 560, 5,
		117, 0, 0, 560, 561, 5, 123, 0, 0, 561, 562, 3, 143, 71, 0, 562, 563, 3,
		143, 71, 0, 563, 564, 3, 143, 71, 0, 564, 565, 3, 143, 71, 0, 565, 566,
		5, 125, 0, 0, 566, 568, 1, 0, 0, 0, 567, 551, 1, 0, 0, 0, 567, 558, 1,
		0, 0, 0, 568, 138, 1, 0, 0, 0, 569, 570, 5, 92, 0, 0, 570, 571, 7, 10,
		0, 0, 571, 572, 7, 11, 0, 0, 572, 579, 7, 11, 0, 0, 573, 574, 5, 92, 0,
		0, 574, 575, 7, 11, 0, 0, 575, 579, 7, 11, 0, 0, 576, 577, 5, 92, 0, 0,
		577, 579, 7, 11, 0, 0, 578, 569, 1, 0, 0, 0, 578, 573, 1, 0, 0, 0, 578,
		576, 1, 0, 0, 0, 579, 140, 1, 0, 0, 0, 580, 581, 5, 92, 0, 0, 581, 583,
		3, 143, 71, 0, 582, 584, 3, 143, 71, 0, 583, 582, 1, 0, 0, 0, 583, 584,
		1, 0, 0, 0, 584, 142, 1, 0, 0, 0, 585, 586, 7, 12, 0, 0, 586, 144, 1, 0,
		0, 0, 587, 588, 5, 102, 0, 0, 588, 589, 5, 39, 0, 0, 589, 146, 1, 0, 0,
		0, 590, 591, 5, 102, 0, 0, 591, 592, 5, 34, 0, 0, 592, 148, 1, 0, 0, 0,
		593, 594, 5, 114, 0, 0, 594, 595, 5, 39, 0, 0, 595, 150, 1, 0, 0, 0, 596,
		597, 5, 114, 0, 0, 597, 598, 5, 34, 0, 0, 598, 152, 1, 0, 0, 0, 599, 600,
		5, 112, 0, 0, 600, 601, 5, 39, 0, 0, 601, 154, 1, 0, 0, 0, 602, 603, 5,
		112, 0, 0, 603, 604, 5, 34, 0, 0, 604, 156, 1, 0, 0, 0, 605, 606, 5, 120,
		0, 0, 606, 607, 5, 39, 0, 0, 607, 158, 1, 0, 0, 0, 608, 609, 5, 120, 0,
		0, 609, 610, 5, 34, 0, 0, 610, 160, 1, 0, 0, 0, 611, 612, 5, 100, 0, 0,
		612, 613, 5, 39, 0, 0, 613, 162, 1, 0, 0, 0, 614, 615, 5, 100, 0, 0, 615,
		616, 5, 34, 0, 0, 616, 164, 1, 0, 0, 0, 617, 628, 8, 13, 0, 0, 618, 628,
		3, 171, 85, 0, 619, 623, 5, 91, 0, 0, 620, 622, 3, 169, 84, 0, 621, 620,
		1, 0, 0, 0, 622, 625, 1, 0, 0, 0, 623, 621, 1, 0, 0, 0, 623, 624, 1, 0,
		0, 0, 624, 626, 1, 0, 0, 0, 625, 623, 1, 0, 0, 0, 626, 628, 5, 93, 0, 0,
		627, 617, 1, 0, 0, 0, 627, 618, 1, 0, 0, 0, 627, 619, 1, 0, 0, 0, 628,
		166, 1, 0, 0, 0, 629, 640, 8, 14, 0, 0, 630, 640, 3, 171, 85, 0, 631, 635,
		5, 91, 0, 0, 632, 634, 3, 169, 84, 0, 633, 632, 1, 0, 0, 0, 634, 637, 1,
		0, 0, 0, 635, 633, 1, 0, 0, 0, 635, 636, 1, 0, 0, 0, 636, 638, 1, 0, 0,
		0, 637, 635, 1, 0, 0, 0, 638, 640, 5, 93, 0, 0, 639, 629, 1, 0, 0, 0, 639,
		630, 1, 0, 0, 0, 639, 631, 1, 0, 0, 0, 640, 168, 1, 0, 0, 0, 641, 644,
		8, 15, 0, 0, 642, 644, 3, 171, 85, 0, 643, 641, 1, 0, 0, 0, 643, 642, 1,
		0, 0, 0, 644, 170, 1, 0, 0, 0, 645, 646, 5, 92, 0, 0, 646, 647, 8, 1, 0,
		0, 647, 172, 1, 0, 0, 0, 648, 649, 5, 109, 0, 0, 649, 650, 5, 105, 0, 0,
		650, 651, 5, 99, 0, 0, 651, 652, 5, 114, 0, 0, 652, 653, 5, 111, 0, 0,
		653, 654, 5, 115, 0, 0, 654, 655, 5, 101, 0, 0, 655, 656, 5, 99, 0, 0,
		656, 657, 5, 111, 0, 0, 657, 658, 5, 110, 0, 0, 658, 659, 5, 100, 0, 0,
		659, 717, 5, 115, 0, 0, 660, 661, 5, 109, 0, 0, 661, 662, 5, 105, 0, 0,
		662, 663, 5, 108, 0, 0, 663, 664, 5, 108, 0, 0, 664, 665, 5, 105, 0, 0,
		665, 666, 5, 115, 0, 0, 666, 667, 5, 101, 0, 0, 667, 668, 5, 99, 0, 0,
		668, 669, 5, 111, 0, 0, 669, 670, 5, 110, 0, 0, 670, 671, 5, 100, 0, 0,
		671, 717, 5, 115, 0, 0, 672, 673, 5, 115, 0, 0, 673, 674, 5, 101, 0, 0,
		674, 675, 5, 99, 0, 0, 675, 676, 5, 111, 0, 0, 676, 677, 5, 110, 0, 0,
		677, 678, 5, 100, 0, 0, 678, 717, 5, 115, 0, 0, 679, 680, 5, 109, 0, 0,
		680, 681, 5, 105, 0, 0, 681, 682, 5, 110, 0, 0, 682, 683, 5, 117, 0, 0,
		683, 684, 5, 116, 0, 0, 684, 685, 5, 101, 0, 0, 685, 717, 5, 115, 0, 0,
		686, 687, 5, 104, 0, 0, 687, 688, 5, 111, 0, 0, 688, 689, 5, 117, 0, 0,
		689, 690, 5, 114, 0, 0, 690, 717, 5, 115, 0, 0, 691, 692, 5, 100, 0, 0,
		692, 693, 5, 97, 0, 0, 693, 694, 5, 121, 0, 0, 694, 717, 5, 115, 0, 0,
		695, 696, 5, 119, 0, 0, 696, 697, 5, 101, 0, 0, 697, 698, 5, 101, 0, 0,
		698, 699, 5, 107, 0, 0, 699, 717, 5, 115, 0, 0, 700, 701, 5, 109, 0, 0,
		701, 702, 5, 111, 0, 0, 702, 703, 5, 110, 0, 0, 703, 704, 5, 116, 0, 0,
		704, 705, 5, 104, 0, 0, 705, 717, 5, 115, 0, 0, 706, 707, 5, 121, 0, 0,
		707, 708, 5, 101, 0, 0, 708, 709, 5, 97, 0, 0, 709, 710, 5, 114, 0, 0,
		710, 717, 5, 115, 0, 0, 711, 712, 5, 117, 0, 0, 712, 717, 5, 115, 0, 0,
		713, 714, 5, 109, 0, 0, 714, 717, 5, 115, 0, 0, 715, 717, 7, 16, 0, 0,
		716, 648, 1, 0, 0, 0, 716, 660, 1, 0, 0, 0, 716, 672, 1, 0, 0, 0, 716,
		679, 1, 0, 0, 0, 716, 686, 1, 0, 0, 0, 716, 691, 1, 0, 0, 0, 716, 695,
		1, 0, 0, 0, 716, 700, 1, 0, 0, 0, 716, 706, 1, 0, 0, 0, 716, 711, 1, 0,
		0, 0, 716, 713, 1, 0, 0, 0, 716, 715, 1, 0, 0, 0, 717, 174, 1, 0, 0, 0,
		54, 0, 298, 306, 320, 328, 333, 338, 340, 345, 350, 354, 359, 364, 366,
		371, 377, 381, 386, 389, 395, 398, 400, 404, 410, 419, 424, 430, 439, 444,
		450, 459, 464, 470, 479, 484, 490, 492, 501, 503, 508, 514, 523, 528, 540,
		549, 567, 578, 583, 623, 627, 635, 639, 643, 716, 1, 6, 0, 0,
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

// preludioLexerInit initializes any static state used to implement preludioLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewpreludioLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PreludioLexerInit() {
	staticData := &preludiolexerLexerStaticData
	staticData.once.Do(preludiolexerLexerInit)
}

// NewpreludioLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewpreludioLexer(input antlr.CharStream) *preludioLexer {
	PreludioLexerInit()
	l := new(preludioLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &preludiolexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "preludioLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// preludioLexer tokens.
const (
	preludioLexerFUNC                = 1
	preludioLexerPRQL                = 2
	preludioLexerRET                 = 3
	preludioLexerARROW               = 4
	preludioLexerASSIGN              = 5
	preludioLexerDECLARE             = 6
	preludioLexerPLUS                = 7
	preludioLexerMINUS               = 8
	preludioLexerSTAR                = 9
	preludioLexerEXP                 = 10
	preludioLexerDIV                 = 11
	preludioLexerMOD                 = 12
	preludioLexerMODEL               = 13
	preludioLexerEQ                  = 14
	preludioLexerNE                  = 15
	preludioLexerGE                  = 16
	preludioLexerGT                  = 17
	preludioLexerLE                  = 18
	preludioLexerLT                  = 19
	preludioLexerBAR                 = 20
	preludioLexerCOLON               = 21
	preludioLexerCOMMA               = 22
	preludioLexerDOT                 = 23
	preludioLexerDOLLAR              = 24
	preludioLexerRANGE               = 25
	preludioLexerLBRACKET            = 26
	preludioLexerRBRACKET            = 27
	preludioLexerLPAREN              = 28
	preludioLexerRPAREN              = 29
	preludioLexerLBRACE              = 30
	preludioLexerRBRACE              = 31
	preludioLexerLANG                = 32
	preludioLexerRANG                = 33
	preludioLexerUNDERSCORE          = 34
	preludioLexerBACKTICK            = 35
	preludioLexerDOUBLE_QUOTE        = 36
	preludioLexerSINGLE_QUOTE        = 37
	preludioLexerTRIPLE_DOUBLE_QUOTE = 38
	preludioLexerTRIPLE_SINGLE_QUOTE = 39
	preludioLexerAND                 = 40
	preludioLexerOR                  = 41
	preludioLexerNOT                 = 42
	preludioLexerCOALESCE            = 43
	preludioLexerNA                  = 44
	preludioLexerINDEXING            = 45
	preludioLexerFUNCTION_CALL       = 46
	preludioLexerWHITESPACE          = 47
	preludioLexerNEWLINE             = 48
	preludioLexerSINGLE_LINE_COMMENT = 49
	preludioLexerBOOLEAN_LIT         = 50
	preludioLexerIDENT               = 51
	preludioLexerIDENT_START         = 52
	preludioLexerIDENT_NEXT          = 53
	preludioLexerINTEGER_LIT         = 54
	preludioLexerRANGE_LIT           = 55
	preludioLexerFLOAT_LIT           = 56
	preludioLexerSTRING_CHAR         = 57
	preludioLexerSTRING_LIT          = 58
	preludioLexerSTRING_INTERP_LIT   = 59
	preludioLexerSTRING_RAW_LIT      = 60
	preludioLexerSTRING_PATH_LIT     = 61
	preludioLexerREGEX_LIT           = 62
	preludioLexerDATE_LIT            = 63
	preludioLexerDURATION_LIT        = 64
)
