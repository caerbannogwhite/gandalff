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
LE: '<=';
GE: '>=';

AT: '@';
BAR: '|';
COLON: ':';
COMMA: ',';
DOT: '.';
DOLLAR: '$';
RANGE: '..';
LANG: '<';
RANG: '>';
LBRACKET: '[';
RBRACKET: ']';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
UNDERSCORE: '_';

BACKTICK: '`';
DOUBLE_QUOTE: '"';
SINGLE_QUOTE: '\'';
TRIPLE_DOUBLE_QUOTE: '"""';
TRIPLE_SINGLE_QUOTE: '\'\'\'';

AND: 'and';
OR: 'or';
NOT: 'not';
COALESCE: '??';
NULL_: 'na';

IDENT: IDENT_START (DOT IDENT_NEXT)*;
IDENT_START: (LETTER | UNDERSCORE) (LETTER | DIGIT | UNDERSCORE)*;
IDENT_NEXT: IDENT_START | STAR;

WHITESPACE: (' ' | '\t') -> skip;
NEWLINE: '\r'? '\n';

SINGLE_LINE_COMMENT: '#' ~[\r\n\u2028\u2029]* NEWLINE;

// Literals
BOOLEAN_LIT: 'true' | 'false';

INTEGER_LIT: DIGIT+;

RANGE_LIT: (INTEGER_LIT | IDENT) RANGE (INTEGER_LIT | IDENT) (
		COLON (INTEGER_LIT | IDENT)
	)?;

FLOAT_LIT:
	DIGIT+ DOT DIGIT* EXPONENT?
	| DIGIT+ EXPONENT?
	| DOT DIGIT+ EXPONENT?;

STRING_CHAR: (ESC | ~[\\'\r\n\u2028\u2029]);
STRING_RAW_CHAR: ~[\\'\r\n\u2028\u2029];

STRING_LIT: SINGLE_QUOTE STRING_CHAR*? SINGLE_QUOTE;
STRING_INTERP_LIT:
	STRING_INTERP_START STRING_CHAR*? SINGLE_QUOTE;
STRING_RAW_LIT: STRING_RAW_START STRING_RAW_CHAR*? SINGLE_QUOTE;
STRING_PATH_LIT: STRING_PATH_START STRING_CHAR*? SINGLE_QUOTE;

REGEX_LIT:
	REGEX_START REGEX_FIRST_CHAR (REGEX_CHAR | ~[\\'])*? SINGLE_QUOTE;

DATE_LIT: DATE_START STRING_CHAR*? SINGLE_QUOTE;

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

fragment STRING_INTERP_START: 'f\'';
fragment STRING_RAW_START: 'r\'';
fragment STRING_PATH_START: 'p\'';
fragment REGEX_START: 'x\'';
fragment DATE_START: 'd\'';

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