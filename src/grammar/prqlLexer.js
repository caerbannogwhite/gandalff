// Generated from .\prql.g4 by ANTLR 4.9.2
// jshint ignore: start
import antlr4 from 'antlr4';



const serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786",
    "\u5964\u00028\u018c\b\u0001\u0004\u0002\t\u0002\u0004\u0003\t\u0003",
    "\u0004\u0004\t\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007",
    "\t\u0007\u0004\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004",
    "\f\t\f\u0004\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0004\u0010",
    "\t\u0010\u0004\u0011\t\u0011\u0004\u0012\t\u0012\u0004\u0013\t\u0013",
    "\u0004\u0014\t\u0014\u0004\u0015\t\u0015\u0004\u0016\t\u0016\u0004\u0017",
    "\t\u0017\u0004\u0018\t\u0018\u0004\u0019\t\u0019\u0004\u001a\t\u001a",
    "\u0004\u001b\t\u001b\u0004\u001c\t\u001c\u0004\u001d\t\u001d\u0004\u001e",
    "\t\u001e\u0004\u001f\t\u001f\u0004 \t \u0004!\t!\u0004\"\t\"\u0004#",
    "\t#\u0004$\t$\u0004%\t%\u0004&\t&\u0004\'\t\'\u0004(\t(\u0004)\t)\u0004",
    "*\t*\u0004+\t+\u0004,\t,\u0004-\t-\u0004.\t.\u0004/\t/\u00040\t0\u0004",
    "1\t1\u00042\t2\u00043\t3\u00044\t4\u00045\t5\u00046\t6\u00047\t7\u0004",
    "8\t8\u00049\t9\u0004:\t:\u0004;\t;\u0003\u0002\u0003\u0002\u0003\u0002",
    "\u0003\u0003\u0003\u0003\u0003\u0004\u0003\u0004\u0003\u0005\u0003\u0005",
    "\u0003\u0006\u0003\u0006\u0003\u0007\u0003\u0007\u0003\u0007\u0003\u0007",
    "\u0003\b\u0003\b\u0003\b\u0003\b\u0003\t\u0003\t\u0003\t\u0003\n\u0003",
    "\n\u0003\u000b\u0003\u000b\u0003\f\u0003\f\u0003\r\u0003\r\u0003\u000e",
    "\u0003\u000e\u0003\u000e\u0003\u000f\u0003\u000f\u0003\u000f\u0003\u0010",
    "\u0003\u0010\u0003\u0010\u0003\u0011\u0003\u0011\u0003\u0011\u0003\u0012",
    "\u0003\u0012\u0003\u0012\u0003\u0012\u0003\u0013\u0003\u0013\u0003\u0013",
    "\u0003\u0014\u0003\u0014\u0003\u0014\u0003\u0015\u0003\u0015\u0003\u0015",
    "\u0003\u0015\u0003\u0015\u0003\u0015\u0003\u0015\u0003\u0015\u0003\u0015",
    "\u0003\u0015\u0003\u0015\u0003\u0015\u0003\u0015\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016",
    "\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0016\u0003\u0017",
    "\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017\u0003\u0017",
    "\u0003\u0017\u0003\u0018\u0003\u0018\u0003\u0018\u0003\u0018\u0003\u0018",
    "\u0003\u0018\u0003\u0018\u0003\u0018\u0003\u0019\u0003\u0019\u0003\u0019",
    "\u0003\u0019\u0003\u0019\u0003\u0019\u0003\u001a\u0003\u001a\u0003\u001a",
    "\u0003\u001a\u0003\u001a\u0003\u001b\u0003\u001b\u0003\u001b\u0003\u001b",
    "\u0003\u001b\u0003\u001b\u0003\u001c\u0003\u001c\u0003\u001c\u0003\u001c",
    "\u0003\u001c\u0003\u001c\u0003\u001c\u0003\u001d\u0003\u001d\u0003\u001d",
    "\u0003\u001d\u0003\u001d\u0003\u001d\u0003\u001e\u0003\u001e\u0003\u001e",
    "\u0003\u001e\u0003\u001e\u0003\u001f\u0003\u001f\u0003\u001f\u0003\u001f",
    "\u0003\u001f\u0003 \u0003 \u0003 \u0003 \u0003 \u0003 \u0003!\u0003",
    "!\u0003\"\u0003\"\u0003#\u0003#\u0003$\u0003$\u0003%\u0003%\u0003&\u0003",
    "&\u0003\'\u0003\'\u0003(\u0003(\u0003)\u0003)\u0003*\u0003*\u0003+\u0003",
    "+\u0003,\u0003,\u0003-\u0003-\u0003.\u0003.\u0003/\u0003/\u0003/\u0003",
    "/\u0003/\u00030\u00030\u00030\u00030\u00030\u00030\u00030\u00030\u0003",
    "0\u00050\u012e\n0\u00031\u00061\u0131\n1\r1\u000e1\u0132\u00032\u0003",
    "2\u00033\u00033\u00034\u00034\u00054\u013b\n4\u00034\u00034\u00035\u0006",
    "5\u0140\n5\r5\u000e5\u0141\u00035\u00035\u00075\u0146\n5\f5\u000e5\u0149",
    "\u000b5\u00035\u00055\u014c\n5\u00035\u00065\u014f\n5\r5\u000e5\u0150",
    "\u00035\u00055\u0154\n5\u00035\u00035\u00065\u0158\n5\r5\u000e5\u0159",
    "\u00035\u00055\u015d\n5\u00055\u015f\n5\u00036\u00036\u00036\u00036",
    "\u00076\u0165\n6\f6\u000e6\u0168\u000b6\u00037\u00037\u00037\u00057",
    "\u016d\n7\u00037\u00037\u00037\u00077\u0172\n7\f7\u000e7\u0175\u000b",
    "7\u00038\u00038\u00058\u0179\n8\u00039\u00039\u00039\u00039\u0003:\u0005",
    ":\u0180\n:\u0003:\u0003:\u0003;\u0003;\u0007;\u0186\n;\f;\u000e;\u0189",
    "\u000b;\u0003;\u0003;\u0002\u0002<\u0003\u0003\u0005\u0004\u0007\u0005",
    "\t\u0006\u000b\u0007\r\b\u000f\t\u0011\n\u0013\u000b\u0015\f\u0017\r",
    "\u0019\u000e\u001b\u000f\u001d\u0010\u001f\u0011!\u0012#\u0013%\u0014",
    "\'\u0015)\u0016+\u0017-\u0018/\u00191\u001a3\u001b5\u001c7\u001d9\u001e",
    ";\u001f= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]0_1a\u0002c\u0002e\u0002g",
    "\u0002i2k3m4o5q6s7u8\u0003\u0002\b\u0003\u00022;\u0004\u0002C\\c|\u0004",
    "\u0002GGgg\u0004\u0002--//\u0004\u0002\u000b\u000b\"\"\u0004\u0002\f",
    "\f\u000f\u000f\u0002\u019c\u0002\u0003\u0003\u0002\u0002\u0002\u0002",
    "\u0005\u0003\u0002\u0002\u0002\u0002\u0007\u0003\u0002\u0002\u0002\u0002",
    "\t\u0003\u0002\u0002\u0002\u0002\u000b\u0003\u0002\u0002\u0002\u0002",
    "\r\u0003\u0002\u0002\u0002\u0002\u000f\u0003\u0002\u0002\u0002\u0002",
    "\u0011\u0003\u0002\u0002\u0002\u0002\u0013\u0003\u0002\u0002\u0002\u0002",
    "\u0015\u0003\u0002\u0002\u0002\u0002\u0017\u0003\u0002\u0002\u0002\u0002",
    "\u0019\u0003\u0002\u0002\u0002\u0002\u001b\u0003\u0002\u0002\u0002\u0002",
    "\u001d\u0003\u0002\u0002\u0002\u0002\u001f\u0003\u0002\u0002\u0002\u0002",
    "!\u0003\u0002\u0002\u0002\u0002#\u0003\u0002\u0002\u0002\u0002%\u0003",
    "\u0002\u0002\u0002\u0002\'\u0003\u0002\u0002\u0002\u0002)\u0003\u0002",
    "\u0002\u0002\u0002+\u0003\u0002\u0002\u0002\u0002-\u0003\u0002\u0002",
    "\u0002\u0002/\u0003\u0002\u0002\u0002\u00021\u0003\u0002\u0002\u0002",
    "\u00023\u0003\u0002\u0002\u0002\u00025\u0003\u0002\u0002\u0002\u0002",
    "7\u0003\u0002\u0002\u0002\u00029\u0003\u0002\u0002\u0002\u0002;\u0003",
    "\u0002\u0002\u0002\u0002=\u0003\u0002\u0002\u0002\u0002?\u0003\u0002",
    "\u0002\u0002\u0002A\u0003\u0002\u0002\u0002\u0002C\u0003\u0002\u0002",
    "\u0002\u0002E\u0003\u0002\u0002\u0002\u0002G\u0003\u0002\u0002\u0002",
    "\u0002I\u0003\u0002\u0002\u0002\u0002K\u0003\u0002\u0002\u0002\u0002",
    "M\u0003\u0002\u0002\u0002\u0002O\u0003\u0002\u0002\u0002\u0002Q\u0003",
    "\u0002\u0002\u0002\u0002S\u0003\u0002\u0002\u0002\u0002U\u0003\u0002",
    "\u0002\u0002\u0002W\u0003\u0002\u0002\u0002\u0002Y\u0003\u0002\u0002",
    "\u0002\u0002[\u0003\u0002\u0002\u0002\u0002]\u0003\u0002\u0002\u0002",
    "\u0002_\u0003\u0002\u0002\u0002\u0002i\u0003\u0002\u0002\u0002\u0002",
    "k\u0003\u0002\u0002\u0002\u0002m\u0003\u0002\u0002\u0002\u0002o\u0003",
    "\u0002\u0002\u0002\u0002q\u0003\u0002\u0002\u0002\u0002s\u0003\u0002",
    "\u0002\u0002\u0002u\u0003\u0002\u0002\u0002\u0003w\u0003\u0002\u0002",
    "\u0002\u0005z\u0003\u0002\u0002\u0002\u0007|\u0003\u0002\u0002\u0002",
    "\t~\u0003\u0002\u0002\u0002\u000b\u0080\u0003\u0002\u0002\u0002\r\u0082",
    "\u0003\u0002\u0002\u0002\u000f\u0086\u0003\u0002\u0002\u0002\u0011\u008a",
    "\u0003\u0002\u0002\u0002\u0013\u008d\u0003\u0002\u0002\u0002\u0015\u008f",
    "\u0003\u0002\u0002\u0002\u0017\u0091\u0003\u0002\u0002\u0002\u0019\u0093",
    "\u0003\u0002\u0002\u0002\u001b\u0095\u0003\u0002\u0002\u0002\u001d\u0098",
    "\u0003\u0002\u0002\u0002\u001f\u009b\u0003\u0002\u0002\u0002!\u009e",
    "\u0003\u0002\u0002\u0002#\u00a1\u0003\u0002\u0002\u0002%\u00a5\u0003",
    "\u0002\u0002\u0002\'\u00a8\u0003\u0002\u0002\u0002)\u00ab\u0003\u0002",
    "\u0002\u0002+\u00b8\u0003\u0002\u0002\u0002-\u00c5\u0003\u0002\u0002",
    "\u0002/\u00cd\u0003\u0002\u0002\u00021\u00d5\u0003\u0002\u0002\u0002",
    "3\u00db\u0003\u0002\u0002\u00025\u00e0\u0003\u0002\u0002\u00027\u00e6",
    "\u0003\u0002\u0002\u00029\u00ed\u0003\u0002\u0002\u0002;\u00f3\u0003",
    "\u0002\u0002\u0002=\u00f8\u0003\u0002\u0002\u0002?\u00fd\u0003\u0002",
    "\u0002\u0002A\u0103\u0003\u0002\u0002\u0002C\u0105\u0003\u0002\u0002",
    "\u0002E\u0107\u0003\u0002\u0002\u0002G\u0109\u0003\u0002\u0002\u0002",
    "I\u010b\u0003\u0002\u0002\u0002K\u010d\u0003\u0002\u0002\u0002M\u010f",
    "\u0003\u0002\u0002\u0002O\u0111\u0003\u0002\u0002\u0002Q\u0113\u0003",
    "\u0002\u0002\u0002S\u0115\u0003\u0002\u0002\u0002U\u0117\u0003\u0002",
    "\u0002\u0002W\u0119\u0003\u0002\u0002\u0002Y\u011b\u0003\u0002\u0002",
    "\u0002[\u011d\u0003\u0002\u0002\u0002]\u011f\u0003\u0002\u0002\u0002",
    "_\u012d\u0003\u0002\u0002\u0002a\u0130\u0003\u0002\u0002\u0002c\u0134",
    "\u0003\u0002\u0002\u0002e\u0136\u0003\u0002\u0002\u0002g\u0138\u0003",
    "\u0002\u0002\u0002i\u015e\u0003\u0002\u0002\u0002k\u0160\u0003\u0002",
    "\u0002\u0002m\u016c\u0003\u0002\u0002\u0002o\u0178\u0003\u0002\u0002",
    "\u0002q\u017a\u0003\u0002\u0002\u0002s\u017f\u0003\u0002\u0002\u0002",
    "u\u0183\u0003\u0002\u0002\u0002wx\u0007/\u0002\u0002xy\u0007@\u0002",
    "\u0002y\u0004\u0003\u0002\u0002\u0002z{\u0007b\u0002\u0002{\u0006\u0003",
    "\u0002\u0002\u0002|}\u0007<\u0002\u0002}\b\u0003\u0002\u0002\u0002~",
    "\u007f\u0007$\u0002\u0002\u007f\n\u0003\u0002\u0002\u0002\u0080\u0081",
    "\u0007)\u0002\u0002\u0081\f\u0003\u0002\u0002\u0002\u0082\u0083\u0007",
    "$\u0002\u0002\u0083\u0084\u0007$\u0002\u0002\u0084\u0085\u0007$\u0002",
    "\u0002\u0085\u000e\u0003\u0002\u0002\u0002\u0086\u0087\u0007)\u0002",
    "\u0002\u0087\u0088\u0007)\u0002\u0002\u0088\u0089\u0007)\u0002\u0002",
    "\u0089\u0010\u0003\u0002\u0002\u0002\u008a\u008b\u00070\u0002\u0002",
    "\u008b\u008c\u00070\u0002\u0002\u008c\u0012\u0003\u0002\u0002\u0002",
    "\u008d\u008e\u0007#\u0002\u0002\u008e\u0014\u0003\u0002\u0002\u0002",
    "\u008f\u0090\u0007,\u0002\u0002\u0090\u0016\u0003\u0002\u0002\u0002",
    "\u0091\u0092\u00071\u0002\u0002\u0092\u0018\u0003\u0002\u0002\u0002",
    "\u0093\u0094\u0007\'\u0002\u0002\u0094\u001a\u0003\u0002\u0002\u0002",
    "\u0095\u0096\u0007?\u0002\u0002\u0096\u0097\u0007?\u0002\u0002\u0097",
    "\u001c\u0003\u0002\u0002\u0002\u0098\u0099\u0007#\u0002\u0002\u0099",
    "\u009a\u0007?\u0002\u0002\u009a\u001e\u0003\u0002\u0002\u0002\u009b",
    "\u009c\u0007@\u0002\u0002\u009c\u009d\u0007?\u0002\u0002\u009d \u0003",
    "\u0002\u0002\u0002\u009e\u009f\u0007>\u0002\u0002\u009f\u00a0\u0007",
    "?\u0002\u0002\u00a0\"\u0003\u0002\u0002\u0002\u00a1\u00a2\u0007c\u0002",
    "\u0002\u00a2\u00a3\u0007p\u0002\u0002\u00a3\u00a4\u0007f\u0002\u0002",
    "\u00a4$\u0003\u0002\u0002\u0002\u00a5\u00a6\u0007q\u0002\u0002\u00a6",
    "\u00a7\u0007t\u0002\u0002\u00a7&\u0003\u0002\u0002\u0002\u00a8\u00a9",
    "\u0007A\u0002\u0002\u00a9\u00aa\u0007A\u0002\u0002\u00aa(\u0003\u0002",
    "\u0002\u0002\u00ab\u00ac\u0007o\u0002\u0002\u00ac\u00ad\u0007k\u0002",
    "\u0002\u00ad\u00ae\u0007e\u0002\u0002\u00ae\u00af\u0007t\u0002\u0002",
    "\u00af\u00b0\u0007q\u0002\u0002\u00b0\u00b1\u0007u\u0002\u0002\u00b1",
    "\u00b2\u0007g\u0002\u0002\u00b2\u00b3\u0007e\u0002\u0002\u00b3\u00b4",
    "\u0007q\u0002\u0002\u00b4\u00b5\u0007p\u0002\u0002\u00b5\u00b6\u0007",
    "f\u0002\u0002\u00b6\u00b7\u0007u\u0002\u0002\u00b7*\u0003\u0002\u0002",
    "\u0002\u00b8\u00b9\u0007o\u0002\u0002\u00b9\u00ba\u0007k\u0002\u0002",
    "\u00ba\u00bb\u0007n\u0002\u0002\u00bb\u00bc\u0007n\u0002\u0002\u00bc",
    "\u00bd\u0007k\u0002\u0002\u00bd\u00be\u0007u\u0002\u0002\u00be\u00bf",
    "\u0007g\u0002\u0002\u00bf\u00c0\u0007e\u0002\u0002\u00c0\u00c1\u0007",
    "q\u0002\u0002\u00c1\u00c2\u0007p\u0002\u0002\u00c2\u00c3\u0007f\u0002",
    "\u0002\u00c3\u00c4\u0007u\u0002\u0002\u00c4,\u0003\u0002\u0002\u0002",
    "\u00c5\u00c6\u0007u\u0002\u0002\u00c6\u00c7\u0007g\u0002\u0002\u00c7",
    "\u00c8\u0007e\u0002\u0002\u00c8\u00c9\u0007q\u0002\u0002\u00c9\u00ca",
    "\u0007p\u0002\u0002\u00ca\u00cb\u0007f\u0002\u0002\u00cb\u00cc\u0007",
    "u\u0002\u0002\u00cc.\u0003\u0002\u0002\u0002\u00cd\u00ce\u0007o\u0002",
    "\u0002\u00ce\u00cf\u0007k\u0002\u0002\u00cf\u00d0\u0007p\u0002\u0002",
    "\u00d0\u00d1\u0007w\u0002\u0002\u00d1\u00d2\u0007v\u0002\u0002\u00d2",
    "\u00d3\u0007g\u0002\u0002\u00d3\u00d4\u0007u\u0002\u0002\u00d40\u0003",
    "\u0002\u0002\u0002\u00d5\u00d6\u0007j\u0002\u0002\u00d6\u00d7\u0007",
    "q\u0002\u0002\u00d7\u00d8\u0007w\u0002\u0002\u00d8\u00d9\u0007t\u0002",
    "\u0002\u00d9\u00da\u0007u\u0002\u0002\u00da2\u0003\u0002\u0002\u0002",
    "\u00db\u00dc\u0007f\u0002\u0002\u00dc\u00dd\u0007c\u0002\u0002\u00dd",
    "\u00de\u0007{\u0002\u0002\u00de\u00df\u0007u\u0002\u0002\u00df4\u0003",
    "\u0002\u0002\u0002\u00e0\u00e1\u0007y\u0002\u0002\u00e1\u00e2\u0007",
    "g\u0002\u0002\u00e2\u00e3\u0007g\u0002\u0002\u00e3\u00e4\u0007m\u0002",
    "\u0002\u00e4\u00e5\u0007u\u0002\u0002\u00e56\u0003\u0002\u0002\u0002",
    "\u00e6\u00e7\u0007o\u0002\u0002\u00e7\u00e8\u0007q\u0002\u0002\u00e8",
    "\u00e9\u0007p\u0002\u0002\u00e9\u00ea\u0007v\u0002\u0002\u00ea\u00eb",
    "\u0007j\u0002\u0002\u00eb\u00ec\u0007u\u0002\u0002\u00ec8\u0003\u0002",
    "\u0002\u0002\u00ed\u00ee\u0007{\u0002\u0002\u00ee\u00ef\u0007g\u0002",
    "\u0002\u00ef\u00f0\u0007c\u0002\u0002\u00f0\u00f1\u0007t\u0002\u0002",
    "\u00f1\u00f2\u0007u\u0002\u0002\u00f2:\u0003\u0002\u0002\u0002\u00f3",
    "\u00f4\u0007h\u0002\u0002\u00f4\u00f5\u0007w\u0002\u0002\u00f5\u00f6",
    "\u0007p\u0002\u0002\u00f6\u00f7\u0007e\u0002\u0002\u00f7<\u0003\u0002",
    "\u0002\u0002\u00f8\u00f9\u0007r\u0002\u0002\u00f9\u00fa\u0007t\u0002",
    "\u0002\u00fa\u00fb\u0007s\u0002\u0002\u00fb\u00fc\u0007n\u0002\u0002",
    "\u00fc>\u0003\u0002\u0002\u0002\u00fd\u00fe\u0007v\u0002\u0002\u00fe",
    "\u00ff\u0007c\u0002\u0002\u00ff\u0100\u0007d\u0002\u0002\u0100\u0101",
    "\u0007n\u0002\u0002\u0101\u0102\u0007g\u0002\u0002\u0102@\u0003\u0002",
    "\u0002\u0002\u0103\u0104\u0007-\u0002\u0002\u0104B\u0003\u0002\u0002",
    "\u0002\u0105\u0106\u0007/\u0002\u0002\u0106D\u0003\u0002\u0002\u0002",
    "\u0107\u0108\u0007?\u0002\u0002\u0108F\u0003\u0002\u0002\u0002\u0109",
    "\u010a\u0007~\u0002\u0002\u010aH\u0003\u0002\u0002\u0002\u010b\u010c",
    "\u0007.\u0002\u0002\u010cJ\u0003\u0002\u0002\u0002\u010d\u010e\u0007",
    "0\u0002\u0002\u010eL\u0003\u0002\u0002\u0002\u010f\u0110\u0007&\u0002",
    "\u0002\u0110N\u0003\u0002\u0002\u0002\u0111\u0112\u0007a\u0002\u0002",
    "\u0112P\u0003\u0002\u0002\u0002\u0113\u0114\u0007>\u0002\u0002\u0114",
    "R\u0003\u0002\u0002\u0002\u0115\u0116\u0007@\u0002\u0002\u0116T\u0003",
    "\u0002\u0002\u0002\u0117\u0118\u0007]\u0002\u0002\u0118V\u0003\u0002",
    "\u0002\u0002\u0119\u011a\u0007_\u0002\u0002\u011aX\u0003\u0002\u0002",
    "\u0002\u011b\u011c\u0007*\u0002\u0002\u011cZ\u0003\u0002\u0002\u0002",
    "\u011d\u011e\u0007+\u0002\u0002\u011e\\\u0003\u0002\u0002\u0002\u011f",
    "\u0120\u0007p\u0002\u0002\u0120\u0121\u0007w\u0002\u0002\u0121\u0122",
    "\u0007n\u0002\u0002\u0122\u0123\u0007n\u0002\u0002\u0123^\u0003\u0002",
    "\u0002\u0002\u0124\u0125\u0007v\u0002\u0002\u0125\u0126\u0007t\u0002",
    "\u0002\u0126\u0127\u0007w\u0002\u0002\u0127\u012e\u0007g\u0002\u0002",
    "\u0128\u0129\u0007h\u0002\u0002\u0129\u012a\u0007c\u0002\u0002\u012a",
    "\u012b\u0007n\u0002\u0002\u012b\u012c\u0007u\u0002\u0002\u012c\u012e",
    "\u0007g\u0002\u0002\u012d\u0124\u0003\u0002\u0002\u0002\u012d\u0128",
    "\u0003\u0002\u0002\u0002\u012e`\u0003\u0002\u0002\u0002\u012f\u0131",
    "\u0005c2\u0002\u0130\u012f\u0003\u0002\u0002\u0002\u0131\u0132\u0003",
    "\u0002\u0002\u0002\u0132\u0130\u0003\u0002\u0002\u0002\u0132\u0133\u0003",
    "\u0002\u0002\u0002\u0133b\u0003\u0002\u0002\u0002\u0134\u0135\t\u0002",
    "\u0002\u0002\u0135d\u0003\u0002\u0002\u0002\u0136\u0137\t\u0003\u0002",
    "\u0002\u0137f\u0003\u0002\u0002\u0002\u0138\u013a\t\u0004\u0002\u0002",
    "\u0139\u013b\t\u0005\u0002\u0002\u013a\u0139\u0003\u0002\u0002\u0002",
    "\u013a\u013b\u0003\u0002\u0002\u0002\u013b\u013c\u0003\u0002\u0002\u0002",
    "\u013c\u013d\u0005a1\u0002\u013dh\u0003\u0002\u0002\u0002\u013e\u0140",
    "\u0005c2\u0002\u013f\u013e\u0003\u0002\u0002\u0002\u0140\u0141\u0003",
    "\u0002\u0002\u0002\u0141\u013f\u0003\u0002\u0002\u0002\u0141\u0142\u0003",
    "\u0002\u0002\u0002\u0142\u0143\u0003\u0002\u0002\u0002\u0143\u0147\u0005",
    "K&\u0002\u0144\u0146\u0005c2\u0002\u0145\u0144\u0003\u0002\u0002\u0002",
    "\u0146\u0149\u0003\u0002\u0002\u0002\u0147\u0145\u0003\u0002\u0002\u0002",
    "\u0147\u0148\u0003\u0002\u0002\u0002\u0148\u014b\u0003\u0002\u0002\u0002",
    "\u0149\u0147\u0003\u0002\u0002\u0002\u014a\u014c\u0005g4\u0002\u014b",
    "\u014a\u0003\u0002\u0002\u0002\u014b\u014c\u0003\u0002\u0002\u0002\u014c",
    "\u015f\u0003\u0002\u0002\u0002\u014d\u014f\u0005c2\u0002\u014e\u014d",
    "\u0003\u0002\u0002\u0002\u014f\u0150\u0003\u0002\u0002\u0002\u0150\u014e",
    "\u0003\u0002\u0002\u0002\u0150\u0151\u0003\u0002\u0002\u0002\u0151\u0153",
    "\u0003\u0002\u0002\u0002\u0152\u0154\u0005g4\u0002\u0153\u0152\u0003",
    "\u0002\u0002\u0002\u0153\u0154\u0003\u0002\u0002\u0002\u0154\u015f\u0003",
    "\u0002\u0002\u0002\u0155\u0157\u0005K&\u0002\u0156\u0158\u0005c2\u0002",
    "\u0157\u0156\u0003\u0002\u0002\u0002\u0158\u0159\u0003\u0002\u0002\u0002",
    "\u0159\u0157\u0003\u0002\u0002\u0002\u0159\u015a\u0003\u0002\u0002\u0002",
    "\u015a\u015c\u0003\u0002\u0002\u0002\u015b\u015d\u0005g4\u0002\u015c",
    "\u015b\u0003\u0002\u0002\u0002\u015c\u015d\u0003\u0002\u0002\u0002\u015d",
    "\u015f\u0003\u0002\u0002\u0002\u015e\u013f\u0003\u0002\u0002\u0002\u015e",
    "\u014e\u0003\u0002\u0002\u0002\u015e\u0155\u0003\u0002\u0002\u0002\u015f",
    "j\u0003\u0002\u0002\u0002\u0160\u0166\u0005m7\u0002\u0161\u0162\u0005",
    "K&\u0002\u0162\u0163\u0005o8\u0002\u0163\u0165\u0003\u0002\u0002\u0002",
    "\u0164\u0161\u0003\u0002\u0002\u0002\u0165\u0168\u0003\u0002\u0002\u0002",
    "\u0166\u0164\u0003\u0002\u0002\u0002\u0166\u0167\u0003\u0002\u0002\u0002",
    "\u0167l\u0003\u0002\u0002\u0002\u0168\u0166\u0003\u0002\u0002\u0002",
    "\u0169\u016d\u0005e3\u0002\u016a\u016d\u0005M\'\u0002\u016b\u016d\u0005",
    "O(\u0002\u016c\u0169\u0003\u0002\u0002\u0002\u016c\u016a\u0003\u0002",
    "\u0002\u0002\u016c\u016b\u0003\u0002\u0002\u0002\u016d\u0173\u0003\u0002",
    "\u0002\u0002\u016e\u0172\u0005e3\u0002\u016f\u0172\u0005c2\u0002\u0170",
    "\u0172\u0005O(\u0002\u0171\u016e\u0003\u0002\u0002\u0002\u0171\u016f",
    "\u0003\u0002\u0002\u0002\u0171\u0170\u0003\u0002\u0002\u0002\u0172\u0175",
    "\u0003\u0002\u0002\u0002\u0173\u0171\u0003\u0002\u0002\u0002\u0173\u0174",
    "\u0003\u0002\u0002\u0002\u0174n\u0003\u0002\u0002\u0002\u0175\u0173",
    "\u0003\u0002\u0002\u0002\u0176\u0179\u0005m7\u0002\u0177\u0179\u0007",
    ",\u0002\u0002\u0178\u0176\u0003\u0002\u0002\u0002\u0178\u0177\u0003",
    "\u0002\u0002\u0002\u0179p\u0003\u0002\u0002\u0002\u017a\u017b\t\u0006",
    "\u0002\u0002\u017b\u017c\u0003\u0002\u0002\u0002\u017c\u017d\b9\u0002",
    "\u0002\u017dr\u0003\u0002\u0002\u0002\u017e\u0180\u0007\u000f\u0002",
    "\u0002\u017f\u017e\u0003\u0002\u0002\u0002\u017f\u0180\u0003\u0002\u0002",
    "\u0002\u0180\u0181\u0003\u0002\u0002\u0002\u0181\u0182\u0007\f\u0002",
    "\u0002\u0182t\u0003\u0002\u0002\u0002\u0183\u0187\u0007%\u0002\u0002",
    "\u0184\u0186\n\u0007\u0002\u0002\u0185\u0184\u0003\u0002\u0002\u0002",
    "\u0186\u0189\u0003\u0002\u0002\u0002\u0187\u0185\u0003\u0002\u0002\u0002",
    "\u0187\u0188\u0003\u0002\u0002\u0002\u0188\u018a\u0003\u0002\u0002\u0002",
    "\u0189\u0187\u0003\u0002\u0002\u0002\u018a\u018b\u0005s:\u0002\u018b",
    "v\u0003\u0002\u0002\u0002\u0015\u0002\u012d\u0132\u013a\u0141\u0147",
    "\u014b\u0150\u0153\u0159\u015c\u015e\u0166\u016c\u0171\u0173\u0178\u017f",
    "\u0187\u0003\b\u0002\u0002"].join("");


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

export default class prqlLexer extends antlr4.Lexer {

    static grammarFileName = "prql.g4";
    static channelNames = [ "DEFAULT_TOKEN_CHANNEL", "HIDDEN" ];
	static modeNames = [ "DEFAULT_MODE" ];
	static literalNames = [ null, "'->'", "'`'", "':'", "'\"'", "'''", "'\"\"\"'", 
                         "'''''", "'..'", "'!'", "'*'", "'/'", "'%'", "'=='", 
                         "'!='", "'>='", "'<='", "'and'", "'or'", "'??'", 
                         "'microseconds'", "'milliseconds'", "'seconds'", 
                         "'minutes'", "'hours'", "'days'", "'weeks'", "'months'", 
                         "'years'", "'func'", "'prql'", "'table'", "'+'", 
                         "'-'", "'='", "'|'", "','", "'.'", "'$'", "'_'", 
                         "'<'", "'>'", "'['", "']'", "'('", "')'", "'null'" ];
	static symbolicNames = [ null, null, null, null, null, null, null, null, 
                          null, null, null, null, null, null, null, null, 
                          null, null, null, null, null, null, null, null, 
                          null, null, null, null, null, "FUNC", "PRQL", 
                          "TABLE", "PLUS", "MINUS", "EQUAL", "BAR", "COMMA", 
                          "DOT", "DOLLAR", "UNDERSCORE", "LANG", "RANG", 
                          "LBRACKET", "RBRACKET", "LPAREN", "RPAREN", "NULL_", 
                          "BOOLEAN", "NUMBER", "IDENT", "IDENT_START", "IDENT_NEXT", 
                          "WHITESPACE", "NEWLINE", "COMMENT" ];
	static ruleNames = [ "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", 
                      "T__7", "T__8", "T__9", "T__10", "T__11", "T__12", 
                      "T__13", "T__14", "T__15", "T__16", "T__17", "T__18", 
                      "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
                      "T__25", "T__26", "T__27", "FUNC", "PRQL", "TABLE", 
                      "PLUS", "MINUS", "EQUAL", "BAR", "COMMA", "DOT", "DOLLAR", 
                      "UNDERSCORE", "LANG", "RANG", "LBRACKET", "RBRACKET", 
                      "LPAREN", "RPAREN", "NULL_", "BOOLEAN", "INT", "DIGIT", 
                      "LETTER", "EXP", "NUMBER", "IDENT", "IDENT_START", 
                      "IDENT_NEXT", "WHITESPACE", "NEWLINE", "COMMENT" ];

    constructor(input) {
        super(input)
        this._interp = new antlr4.atn.LexerATNSimulator(this, atn, decisionsToDFA, new antlr4.PredictionContextCache());
    }

    get atn() {
        return atn;
    }
}

prqlLexer.EOF = antlr4.Token.EOF;
prqlLexer.T__0 = 1;
prqlLexer.T__1 = 2;
prqlLexer.T__2 = 3;
prqlLexer.T__3 = 4;
prqlLexer.T__4 = 5;
prqlLexer.T__5 = 6;
prqlLexer.T__6 = 7;
prqlLexer.T__7 = 8;
prqlLexer.T__8 = 9;
prqlLexer.T__9 = 10;
prqlLexer.T__10 = 11;
prqlLexer.T__11 = 12;
prqlLexer.T__12 = 13;
prqlLexer.T__13 = 14;
prqlLexer.T__14 = 15;
prqlLexer.T__15 = 16;
prqlLexer.T__16 = 17;
prqlLexer.T__17 = 18;
prqlLexer.T__18 = 19;
prqlLexer.T__19 = 20;
prqlLexer.T__20 = 21;
prqlLexer.T__21 = 22;
prqlLexer.T__22 = 23;
prqlLexer.T__23 = 24;
prqlLexer.T__24 = 25;
prqlLexer.T__25 = 26;
prqlLexer.T__26 = 27;
prqlLexer.T__27 = 28;
prqlLexer.FUNC = 29;
prqlLexer.PRQL = 30;
prqlLexer.TABLE = 31;
prqlLexer.PLUS = 32;
prqlLexer.MINUS = 33;
prqlLexer.EQUAL = 34;
prqlLexer.BAR = 35;
prqlLexer.COMMA = 36;
prqlLexer.DOT = 37;
prqlLexer.DOLLAR = 38;
prqlLexer.UNDERSCORE = 39;
prqlLexer.LANG = 40;
prqlLexer.RANG = 41;
prqlLexer.LBRACKET = 42;
prqlLexer.RBRACKET = 43;
prqlLexer.LPAREN = 44;
prqlLexer.RPAREN = 45;
prqlLexer.NULL_ = 46;
prqlLexer.BOOLEAN = 47;
prqlLexer.NUMBER = 48;
prqlLexer.IDENT = 49;
prqlLexer.IDENT_START = 50;
prqlLexer.IDENT_NEXT = 51;
prqlLexer.WHITESPACE = 52;
prqlLexer.NEWLINE = 53;
prqlLexer.COMMENT = 54;



