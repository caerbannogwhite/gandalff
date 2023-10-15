lexer grammar preludioLexer;

FUNC: 'func';
PRQL: 'prql';
RET: 'ret';
ARROW: '->';
ASSIGN: '=';
DECLARE: ':=';

PLUS: '+';
MINUS: '-';
STAR: '*';
EXP: '^';
DIV: '/';
MOD: '%';
MODEL: '~';

EQ: '==';
NE: '!=';
GE: '>=';
GT: '>';
LE: '<=';
LT: '<';

BAR: '|';
COLON: ':';
COMMA: ',';
DOT: '.';
DOLLAR: '$';
RANGE: '..';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LANG: LT;
RANG: GT;
UNDERSCORE: '_';

BACKTICK: '`';
DOUBLE_QUOTE: '"';
SINGLE_QUOTE: '\'';
TRIPLE_DOUBLE_QUOTE: '"""';
TRIPLE_SINGLE_QUOTE: '\'\'\'';

AND: 'and';
OR: 'or';
NOT: 'not';
REVERSE: 'rev';
IF: 'if';
DO: 'do';
ELSE: 'else';
FOR: 'for';
IN: 'in';
END: 'end';
HELP: '?';
COALESCE: '??';
NA: 'na';
INDEXING: '@';
FUNCTION_CALL: '!';

WHITESPACE: (' ' | '\t') -> skip;
NEWLINE: '\r'? '\n';

SINGLE_LINE_COMMENT: '#' ~[\r\n\u2028\u2029]* NEWLINE;

// Literals
BOOLEAN_LIT: 'true' | 'false';

IDENT: IDENT_START (DOT IDENT_NEXT)*;
IDENT_START: (LETTER | UNDERSCORE) (LETTER | DIGIT | UNDERSCORE)*;
IDENT_NEXT: IDENT_START | STAR;

INTEGER_LIT: DIGIT+;

RANGE_LIT: (INTEGER_LIT | IDENT) RANGE (INTEGER_LIT | IDENT) (
		COLON (INTEGER_LIT | IDENT)
	)?;

FLOAT_LIT:
	DIGIT+ DOT DIGIT* EXPONENT?
	| DIGIT+ EXPONENT?
	| DOT DIGIT+ EXPONENT?;

STRING_CHAR: (ESC | ~[\\'\r\n\u2028\u2029]);

STRING_LIT:
	SINGLE_QUOTE STRING_CHAR*? SINGLE_QUOTE
	| DOUBLE_QUOTE STRING_CHAR*? DOUBLE_QUOTE;
STRING_INTERP_LIT:
	STRING_INTERP_START_SINGLE STRING_CHAR*? SINGLE_QUOTE
	| STRING_INTERP_START_DOUBLE STRING_CHAR*? DOUBLE_QUOTE;
STRING_RAW_LIT:
	STRING_RAW_START_SINGLE STRING_CHAR*? SINGLE_QUOTE
	| STRING_RAW_START_DOUBLE STRING_CHAR*? DOUBLE_QUOTE;
STRING_PATH_LIT:
	STRING_PATH_START_SINGLE STRING_CHAR*? SINGLE_QUOTE
	| STRING_PATH_START_SINGLE STRING_CHAR*? DOUBLE_QUOTE;

REGEX_LIT:
	REGEX_START_SINGLE REGEX_FIRST_CHAR (REGEX_CHAR | ~[\\'])*? SINGLE_QUOTE
	| REGEX_START_DOUBLE REGEX_FIRST_CHAR (REGEX_CHAR | ~[\\"])*? DOUBLE_QUOTE;

DATE_LIT:
	DATE_START_SINGLE STRING_CHAR*? SINGLE_QUOTE
	| DATE_START_DOUBLE STRING_CHAR*? DOUBLE_QUOTE;

DURATION_LIT: INTEGER_LIT COLON DURATION_KIND;

fragment DIGIT: [0-9];
fragment LETTER: [a-zA-Z];
fragment EXPONENT: ('E' | 'e') ('+' | '-')? INTEGER_LIT;

fragment ESC:
	'\\' [abtnfrv"'\\]
	| UNICODE_ESCAPE
	| HEX_ESCAPE
	| OCTAL_ESCAPE;

fragment UNICODE_ESCAPE:
	'\\' 'u' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
	| '\\' 'u' '{' HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT '}';

fragment OCTAL_ESCAPE:
	'\\' [0-3] [0-7] [0-7]
	| '\\' [0-7] [0-7]
	| '\\' [0-7];

fragment HEX_ESCAPE: '\\' HEXDIGIT HEXDIGIT?;

fragment HEXDIGIT: ('0' ..'9' | 'a' ..'f' | 'A' ..'F');

fragment STRING_INTERP_START_SINGLE: 'f\'';
fragment STRING_INTERP_START_DOUBLE: 'f"';
fragment STRING_RAW_START_SINGLE: 'r\'';
fragment STRING_RAW_START_DOUBLE: 'r"';
fragment STRING_PATH_START_SINGLE: 'p\'';
fragment STRING_PATH_START_DOUBLE: 'p"';
fragment REGEX_START_SINGLE: 'x\'';
fragment REGEX_START_DOUBLE: 'x"';
fragment DATE_START_SINGLE: 'd\'';
fragment DATE_START_DOUBLE: 'd"';

fragment REGEX_FIRST_CHAR:
	~[*\r\n\u2028\u2029\\/[]
	| REGEX_BACK_SEQ
	| '[' REGEX_CLASS_CHAR* ']';

fragment REGEX_CHAR:
	~[\r\n\u2028\u2029\\/[]
	| REGEX_BACK_SEQ
	| '[' REGEX_CLASS_CHAR* ']';

fragment REGEX_CLASS_CHAR:
	~[\r\n\u2028\u2029\]\\]
	| REGEX_BACK_SEQ;

fragment REGEX_BACK_SEQ: '\\' ~[\r\n\u2028\u2029];

fragment DURATION_KIND:
	'microseconds'
	| 'milliseconds'
	| 'seconds'
	| 'minutes'
	| 'hours'
	| 'days'
	| 'weeks'
	| 'months'
	| 'years'
	| 'us'
	| 'ms'
	| 's'
	| 'm'
	| 'h'
	| 'd'
	| 'w'
	| 'M'
	| 'y';