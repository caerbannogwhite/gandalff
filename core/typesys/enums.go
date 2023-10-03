package typesys

type OPCODE uint8

const (
	OP_START_STMT OPCODE = iota
	OP_END_STMT
	OP_START_PIPELINE
	OP_END_PIPELINE
	OP_START_FUNC_CALL
	OP_MAKE_FUNC_CALL
	OP_START_LIST
	OP_END_LIST
	OP_ADD_FUNC_PARAM
	OP_ADD_EXPR_TERM
	OP_PUSH_NAMED_PARAM
	OP_PUSH_ASSIGN_IDENT
	OP_PUSH_TERM
	OP_END_CHUNCK
	OP_VAR_DECL
	OP_VAR_ASSIGN
	OP_INDEXING
	OP_GOTO

	OP_BINARY_MUL
	OP_BINARY_DIV
	OP_BINARY_MOD
	OP_BINARY_ADD
	OP_BINARY_SUB
	OP_BINARY_EXP

	OP_BINARY_EQ
	OP_BINARY_NE
	OP_BINARY_GE
	OP_BINARY_LE
	OP_BINARY_GT
	OP_BINARY_LT

	OP_BINARY_AND
	OP_BINARY_OR
	OP_BINARY_XOR
	OP_BINARY_COALESCE
	OP_BINARY_MODEL

	OP_BINARY_LSHIFT
	OP_BINARY_RSHIFT

	OP_UNARY_ADD
	OP_UNARY_SUB
	OP_UNARY_NOT

	NO_OP = 255
)

type PARAM1 uint8

const (
	TERM_NULL PARAM1 = iota
	TERM_BOOLEAN
	TERM_INTEGER
	TERM_RANGE
	TERM_FLOAT
	TERM_STRING
	TERM_STRING_RAW
	TERM_STRING_PATH
	TERM_REGEX
	TERM_DATE
	TERM_DURATION_MICROSECOND
	TERM_DURATION_MILLISECOND
	TERM_DURATION_SECOND
	TERM_DURATION_MINUTE
	TERM_DURATION_HOUR
	TERM_DURATION_DAY
	TERM_DURATION_MONTH
	TERM_DURATION_YEAR
	TERM_LIST
	TERM_PIPELINE
	TERM_SYMBOL
)

const (
	SYMBOL_NULL  = "na"
	SYMBOL_TRUE  = "true"
	SYMBOL_FALSE = "false"

	SYMBOL_INDEXING  = "@"
	SYMBOL_FUNC_CALL = "$"
	SYMBOL_RANGE     = ".."
	SYMBOL_COLON     = ":"

	SYMBOL_BINARY_MUL      = "*"
	SYMBOL_BINARY_DIV      = "/"
	SYMBOL_BINARY_MOD      = "%"
	SYMBOL_BINARY_EXP      = "^"
	SYMBOL_BINARY_ADD      = "+"
	SYMBOL_BINARY_SUB      = "-"
	SYMBOL_BINARY_EQ       = "=="
	SYMBOL_BINARY_NE       = "!="
	SYMBOL_BINARY_GE       = ">="
	SYMBOL_BINARY_LE       = "<="
	SYMBOL_BINARY_GT       = ">"
	SYMBOL_BINARY_LT       = "<"
	SYMBOL_BINARY_AND      = "not"
	SYMBOL_BINARY_OR       = "or"
	SYMBOL_BINARY_MODEL    = "~"
	SYMBOL_BINARY_COALESCE = "??"

	SYMBOL_UNARY_ADD = "+"
	SYMBOL_UNARY_SUB = "-"
	SYMBOL_UNARY_NOT = "not"

	SYMBOL_DURATION_MICROSECOND_SHORT = "us"
	SYMBOL_DURATION_MILLISECOND_SHORT = "ms"
	SYMBOL_DURATION_SECOND_SHORT      = "s"
	SYMBOL_DURATION_MINUTE_SHORT      = "m"
	SYMBOL_DURATION_HOUR_SHORT        = "h"
	SYMBOL_DURATION_DAY_SHORT         = "d"
	SYMBOL_DURATION_MONTH_SHORT       = "M"
	SYMBOL_DURATION_YEAR_SHORT        = "y"
	SYMBOL_DURATION_MICROSECOND       = "microseconds"
	SYMBOL_DURATION_MILLISECOND       = "milliseconds"
	SYMBOL_DURATION_SECOND            = "seconds"
	SYMBOL_DURATION_MINUTE            = "minutes"
	SYMBOL_DURATION_HOUR              = "hours"
	SYMBOL_DURATION_DAY               = "days"
	SYMBOL_DURATION_MONTH             = "months"
	SYMBOL_DURATION_YEAR              = "years"
)

type LOG_TYPE uint8

const (
	LOG_INFO LOG_TYPE = iota
	LOG_WARNING
	LOG_ERROR
	LOG_DEBUG
)

type LogEnty struct {
	LogType LOG_TYPE `json:"logType"`
	Level   uint8    `json:"level"`
	Message string   `json:"message"`
}

type Columnar struct {
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	ActualLength int      `json:"actualLength"` // actual length of the column
	Data         []string `json:"data"`
	Nulls        []bool   `json:"nulls"`
}

type PreludioOutput struct {
	Log  []LogEnty    `json:"log"`
	Data [][]Columnar `json:"data"`
}
