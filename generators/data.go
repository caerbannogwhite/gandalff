package main

import (
	"fmt"
	"go/ast"

	"github.com/caerbannogwhite/gandalff/meta"
)

type MakeOperationType func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr

type OperationApplyTo struct {
	SeriesName    string
	SeriesType    meta.BaseType
	MakeOperation MakeOperationType
}

type Operation struct {
	OpCode  meta.OPCODE
	ApplyTo []OperationApplyTo
}

type SeriesFile struct {
	SeriesName            string
	SeriesType            meta.BaseType
	SeriesTypeStr         string
	SeriesGoTypeStr       string
	SeriesGoOuterTypeStr  string
	SeriesNullableTypeStr string
	DefaultValue          string
	IsGoTypePtr           bool
	IsTimeType            bool
	Operations            map[string]Operation
}

var DATA_BASE_METHODS = map[string]SeriesFile{
	"series_bool_base.go": {
		SeriesName:            "SeriesBool",
		SeriesTypeStr:         "BoolType",
		SeriesGoTypeStr:       "bool",
		SeriesGoOuterTypeStr:  "bool",
		SeriesNullableTypeStr: "NullableBool",
		DefaultValue:          "false",
	},

	"series_int_base.go": {
		SeriesName:            "SeriesInt",
		SeriesTypeStr:         "IntType",
		SeriesGoTypeStr:       "int",
		SeriesGoOuterTypeStr:  "int",
		SeriesNullableTypeStr: "NullableInt",
		DefaultValue:          "0",
	},

	"series_int64_base.go": {
		SeriesName:            "SeriesInt64",
		SeriesTypeStr:         "Int64Type",
		SeriesGoTypeStr:       "int64",
		SeriesGoOuterTypeStr:  "int64",
		SeriesNullableTypeStr: "NullableInt64",
		DefaultValue:          "0",
	},

	"series_float64_base.go": {
		SeriesName:            "SeriesFloat64",
		SeriesTypeStr:         "Float64Type",
		SeriesGoTypeStr:       "float64",
		SeriesGoOuterTypeStr:  "float64",
		SeriesNullableTypeStr: "NullableFloat64",
		DefaultValue:          "0",
	},

	"series_string_base.go": {
		SeriesName:            "SeriesString",
		SeriesTypeStr:         "StringType",
		SeriesGoTypeStr:       "*string",
		SeriesGoOuterTypeStr:  "string",
		SeriesNullableTypeStr: "NullableString",
		DefaultValue:          "s.ctx.stringPool.naTextPtr",
		IsGoTypePtr:           true,
	},

	"series_time_base.go": {
		SeriesName:            "SeriesTime",
		SeriesTypeStr:         "TimeType",
		SeriesGoTypeStr:       "time.Time",
		SeriesGoOuterTypeStr:  "time.Time",
		SeriesNullableTypeStr: "NullableTime",
		DefaultValue:          "time.Time{}",
		IsTimeType:            true,
	},

	"series_duration_base.go": {
		SeriesName:            "SeriesDuration",
		SeriesTypeStr:         "DurationType",
		SeriesGoTypeStr:       "time.Duration",
		SeriesGoOuterTypeStr:  "time.Duration",
		SeriesNullableTypeStr: "NullableDuration",
		DefaultValue:          "time.Duration(0)",
	},
}

func GenerateOperationsData() map[string]SeriesFile {
	var data = map[string]SeriesFile{
		"series_na_ops.go": {
			SeriesName: "SeriesNA",
			SeriesType: meta.NullType,
			Operations: map[string]Operation{},
		},

		"series_bool_ops.go": {
			SeriesName: "SeriesBool",
			SeriesType: meta.BoolType,
			Operations: map[string]Operation{
				"Mul": {
					OpCode: meta.OP_BINARY_MUL,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] && %s.data[%s] { %s[%s] = 1 }", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Div": {
					OpCode: meta.OP_BINARY_DIV,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nb2 := float64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 / b2", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 / float64(%s.data[%s])", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 / float64(%s.data[%s])", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 / %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Mod": {
					OpCode: meta.OP_BINARY_MOD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nb2 := float64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(b1, b2)", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = math.Mod(b1, float64(%s.data[%s]))", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = math.Mod(b1, float64(%s.data[%s]))", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = math.Mod(b1, float64(%s.data[%s]))", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Exp": {
					OpCode: meta.OP_BINARY_EXP,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nb2 := float64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = int64(math.Pow(b1, b2))", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = int64(math.Pow(b1, float64(%s.data[%s])))", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = int64(math.Pow(b1, float64(%s.data[%s])))", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = float64(math.Pow(b1, float64(%s.data[%s])))", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Add": {
					OpCode: meta.OP_BINARY_ADD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 + b2", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 + %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 + %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 + %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(boolToString(%s.data[%s]) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Sub": {
					OpCode: meta.OP_BINARY_SUB,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 - b2", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 - %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 - %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 - %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Eq": {
					OpCode: meta.OP_BINARY_EQ,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ne": {
					OpCode: meta.OP_BINARY_NE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Lt": {
					OpCode: meta.OP_BINARY_LT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 < b2", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 < %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 < %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 < %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Le": {
					OpCode: meta.OP_BINARY_LE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 <= b2", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 <= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 <= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 <= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Gt": {
					OpCode: meta.OP_BINARY_GT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 > b2", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 > %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 > %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 > %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"Ge": {
					OpCode: meta.OP_BINARY_GE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nb2 := int64(0)\nif %s.data[%s] { b1 = 1 }\nif %s.data[%s] { b2 = 1 }\n%s[%s] = b1 >= b2", op1, op1Index, op2, op2Index, res, resIndex)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 >= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := int64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 >= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b1 := float64(0)\nif %s.data[%s] { b1 = 1 }\n%s[%s] = b1 >= %s.data[%s]", op1, op1Index, res, resIndex, op2, op2Index)}
							},
						},
					},
				},

				"And": {
					OpCode: meta.OP_BINARY_AND,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] && %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Or": {
					OpCode: meta.OP_BINARY_OR,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] || %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},
			},
		},

		"series_int_ops.go": {
			SeriesName: "SeriesInt",
			SeriesType: meta.IntType,
			Operations: map[string]Operation{
				"Mul": {
					OpCode: meta.OP_BINARY_MUL,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Div": {
					OpCode: meta.OP_BINARY_DIV,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = float64(%s.data[%s]) / b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Mod": {
					OpCode: meta.OP_BINARY_MOD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(float64(%s.data[%s]), b2)", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Exp": {
					OpCode: meta.OP_BINARY_EXP,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = int64(math.Pow(float64(%s.data[%s]), b2))", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Add": {
					OpCode: meta.OP_BINARY_ADD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] + b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(intToString(int64(%s.data[%s])) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Sub": {
					OpCode: meta.OP_BINARY_SUB,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] - b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Eq": {
					OpCode: meta.OP_BINARY_EQ,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ne": {
					OpCode: meta.OP_BINARY_NE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Lt": {
					OpCode: meta.OP_BINARY_LT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] < b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Le": {
					OpCode: meta.OP_BINARY_LE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] <= b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Gt": {
					OpCode: meta.OP_BINARY_GT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] > b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ge": {
					OpCode: meta.OP_BINARY_GE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] >= b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(%s.data[%s]) >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},
			},
		},

		"series_int64_ops.go": {
			SeriesName: "SeriesInt64",
			SeriesType: meta.Int64Type,
			Operations: map[string]Operation{
				"Mul": {
					OpCode: meta.OP_BINARY_MUL,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Div": {
					OpCode: meta.OP_BINARY_DIV,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = float64(%s.data[%s]) / b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Mod": {
					OpCode: meta.OP_BINARY_MOD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(float64(%s.data[%s]), b2)", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Exp": {
					OpCode: meta.OP_BINARY_EXP,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = int64(math.Pow(float64(%s.data[%s]), b2))", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = int64(math.Pow(float64(%s.data[%s]), float64(%s.data[%s])))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Add": {
					OpCode: meta.OP_BINARY_ADD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] + b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(intToString(%s.data[%s]) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Sub": {
					OpCode: meta.OP_BINARY_SUB,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] - b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Eq": {
					OpCode: meta.OP_BINARY_EQ,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ne": {
					OpCode: meta.OP_BINARY_NE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != int64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = float64(%s.data[%s]) != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Lt": {
					OpCode: meta.OP_BINARY_LT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] < b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) < %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Le": {
					OpCode: meta.OP_BINARY_LE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] <= b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) <= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Gt": {
					OpCode: meta.OP_BINARY_GT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] > b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) > %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ge": {
					OpCode: meta.OP_BINARY_GE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := int64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] >= b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= int64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = float64(%s.data[%s]) >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},
			},
		},

		"series_float64_ops.go": {
			SeriesName: "SeriesFloat64",
			SeriesType: meta.Float64Type,
			Operations: map[string]Operation{
				"Mul": {
					OpCode: meta.OP_BINARY_MUL,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("if %s.data[%s] { %s[%s] = %s.data[%s] }", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] * %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Div": {
					OpCode: meta.OP_BINARY_DIV,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] / b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] / %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Mod": {
					OpCode: meta.OP_BINARY_MOD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Mod(%s.data[%s], b2)", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Mod(float64(%s.data[%s]), float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Exp": {
					OpCode: meta.OP_BINARY_EXP,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = math.Pow(%s.data[%s], b2)", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(%s.data[%s], float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(%s.data[%s], float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = math.Pow(%s.data[%s], float64(%s.data[%s]))", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Add": {
					OpCode: meta.OP_BINARY_ADD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] + b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(floatToString(%s.data[%s]) + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Sub": {
					OpCode: meta.OP_BINARY_SUB,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] - b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - float64(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Eq": {
					OpCode: meta.OP_BINARY_EQ,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] == %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ne": {
					OpCode: meta.OP_BINARY_NE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] != %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Lt": {
					OpCode: meta.OP_BINARY_LT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] < b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] < %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Le": {
					OpCode: meta.OP_BINARY_LE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] <= b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] <= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Gt": {
					OpCode: meta.OP_BINARY_GT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] > b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] > %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ge": {
					OpCode: meta.OP_BINARY_GE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("b2 := float64(0)\nif %s.data[%s] { b2 = 1 }\n%s[%s] = %s.data[%s] >= b2", op2, op2Index, res, resIndex, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= float64(%s.data[%s])`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = %s.data[%s] >= %s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},
			},
		},

		"series_string_ops.go": {
			SeriesName: "SeriesString",
			SeriesType: meta.StringType,
			Operations: map[string]Operation{
				"Add": {
					OpCode: meta.OP_BINARY_ADD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesBool",
							SeriesType: meta.BoolType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + boolToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt",
							SeriesType: meta.IntType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + intToString(int64(%s.data[%s])))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesInt64",
							SeriesType: meta.Int64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + intToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesFloat64",
							SeriesType: meta.Float64Type,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + floatToString(%s.data[%s]))", res, resIndex, op1, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + *%s.data[%s])", res, resIndex, op1, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + %s.data[%s].String())", res, resIndex, op1, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + %s.data[%s].String())", res, resIndex, op1, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Eq": {
					OpCode: meta.OP_BINARY_EQ,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] == *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ne": {
					OpCode: meta.OP_BINARY_NE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] != *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Lt": {
					OpCode: meta.OP_BINARY_LT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] < *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Le": {
					OpCode: meta.OP_BINARY_LE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] <= *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Gt": {
					OpCode: meta.OP_BINARY_GT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] > *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ge": {
					OpCode: meta.OP_BINARY_GE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf(`%s[%s] = *%s.data[%s] >= *%s.data[%s]`, res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},
			},
		},

		"series_time_ops.go": {
			SeriesName: "SeriesTime",
			SeriesType: meta.TimeType,
			Operations: map[string]Operation{
				"Add": {
					OpCode: meta.OP_BINARY_ADD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(%s.data[%s].String() + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].AddDate(%s.data[%s].Year(), int(%s.data[%s].Month()), %s.data[%s].Day())", res, resIndex, op1, op1Index, op2, op2Index, op2, op2Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Add(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Sub": {
					OpCode: meta.OP_BINARY_SUB,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Sub(%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Add(-%s.data[%s])", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Eq": {
					OpCode: meta.OP_BINARY_EQ,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Compare(%s.data[%s]) == 0", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ne": {
					OpCode: meta.OP_BINARY_NE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Compare(%s.data[%s]) != 0", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Lt": {
					OpCode: meta.OP_BINARY_LT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Compare(%s.data[%s]) == -1", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Le": {
					OpCode: meta.OP_BINARY_LE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Compare(%s.data[%s]) <= 0", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Gt": {
					OpCode: meta.OP_BINARY_GT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Compare(%s.data[%s]) == 1", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ge": {
					OpCode: meta.OP_BINARY_GE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Compare(%s.data[%s]) >= 1", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},
			},
		},

		"series_duration_ops.go": {
			SeriesName: "SeriesDuration",
			SeriesType: meta.DurationType,
			Operations: map[string]Operation{
				"Add": {
					OpCode: meta.OP_BINARY_ADD,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesString",
							SeriesType: meta.StringType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(%s.data[%s].String() + *%s.data[%s])", res, resIndex, op2, op1, op1Index, op2, op2Index)}
							},
						},
						{
							SeriesName: "SeriesTime",
							SeriesType: meta.TimeType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s].Add(%s.data[%s])", res, resIndex, op2, op2Index, op1, op1Index)}
							},
						},
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] + %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Sub": {
					OpCode: meta.OP_BINARY_SUB,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] - %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Eq": {
					OpCode: meta.OP_BINARY_EQ,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] == %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ne": {
					OpCode: meta.OP_BINARY_NE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesDuration",

							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] != %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Lt": {
					OpCode: meta.OP_BINARY_LT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] < %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Le": {
					OpCode: meta.OP_BINARY_LE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] <= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Gt": {
					OpCode: meta.OP_BINARY_GT,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] > %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},

				"Ge": {
					OpCode: meta.OP_BINARY_GE,
					ApplyTo: []OperationApplyTo{
						{
							SeriesName: "SeriesDuration",
							SeriesType: meta.DurationType,
							MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
								return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s] >= %s.data[%s]", res, resIndex, op1, op1Index, op2, op2Index)}
							},
						},
					},
				},
			},
		},
	}

	// Generate entries for SeriesNA
	opNames := []string{"Mul", "Div", "Mod", "Exp", "Add", "Sub", "Eq", "Ne", "Lt", "Le", "Gt", "Ge", "And", "Or"}
	opCodes := []meta.OPCODE{
		meta.OP_BINARY_MUL, meta.OP_BINARY_DIV, meta.OP_BINARY_MOD, meta.OP_BINARY_EXP, meta.OP_BINARY_ADD, meta.OP_BINARY_SUB,
		meta.OP_BINARY_EQ, meta.OP_BINARY_NE, meta.OP_BINARY_LT, meta.OP_BINARY_LE, meta.OP_BINARY_GT, meta.OP_BINARY_GE,
		meta.OP_BINARY_AND, meta.OP_BINARY_OR,
	}

	seriesNames := []string{"SeriesNA", "SeriesBool", "SeriesInt", "SeriesInt64", "SeriesFloat64", "SeriesString", "SeriesTime", "SeriesDuration"}
	seriesTypes := []meta.BaseType{meta.NullType, meta.BoolType, meta.IntType, meta.Int64Type, meta.Float64Type, meta.StringType, meta.TimeType, meta.DurationType}

	for i, opName := range opNames {
		applyTo := []OperationApplyTo{}
		for j, seriesName := range seriesNames {
			resType := ComputeResInnerType(opCodes[i], seriesTypes[j], seriesTypes[j])

			// Special case for string concatenation
			if opCodes[i] == meta.OP_BINARY_ADD && seriesTypes[j] == meta.StringType {
				applyTo = append(applyTo, OperationApplyTo{
					SeriesName: seriesName,
					SeriesType: seriesTypes[j],
					MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
						return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(NA_TEXT + *%s.data[%s])", res, resIndex, op1, op2, op2Index)}
					},
				})
			} else

			// Special case for logical OR
			if opCodes[i] == meta.OP_BINARY_OR && seriesTypes[j] == meta.BoolType {
				applyTo = append(applyTo, OperationApplyTo{
					SeriesName: seriesName,
					SeriesType: seriesTypes[j],
					MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
						return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s]", res, resIndex, op2, op2Index)}
					},
				})
			} else if resType != meta.ErrorType {
				applyTo = append(applyTo, OperationApplyTo{
					SeriesName: seriesName,
					SeriesType: seriesTypes[j],
					MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
						return &ast.Ident{Name: ""}
					},
				})
			}
		}

		data["series_na_ops.go"].Operations[opName] = Operation{OpCode: opCodes[i], ApplyTo: applyTo}
	}

	// Append SeriesNA to all other series
	fileNames := []string{"series_bool_ops.go", "series_int_ops.go", "series_int64_ops.go", "series_float64_ops.go", "series_string_ops.go", "series_time_ops.go", "series_duration_ops.go"}
	seriesTypes = []meta.BaseType{meta.BoolType, meta.IntType, meta.Int64Type, meta.Float64Type, meta.StringType, meta.TimeType, meta.DurationType}

	for i, fileName := range fileNames {
		for j, opName := range opNames {
			resType := ComputeResInnerType(opCodes[j], seriesTypes[i], meta.NullType)

			// Special case for string concatenation
			if opCodes[j] == meta.OP_BINARY_ADD && seriesTypes[i] == meta.StringType {
				data[fileName].Operations[opName] = Operation{
					OpCode: data[fileName].Operations[opName].OpCode,
					ApplyTo: append(data[fileName].Operations[opName].ApplyTo, OperationApplyTo{
						SeriesName: "SeriesNA",
						SeriesType: meta.NullType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.ctx.stringPool.Put(*%s.data[%s] + NA_TEXT)", res, resIndex, op1, op1, op1Index)}
						},
					}),
				}
			} else

			// Special case for logical OR
			if opCodes[j] == meta.OP_BINARY_OR && seriesTypes[i] == meta.BoolType {
				data[fileName].Operations[opName] = Operation{
					OpCode: data[fileName].Operations[opName].OpCode,
					ApplyTo: append(data[fileName].Operations[opName].ApplyTo, OperationApplyTo{
						SeriesName: "SeriesNA",
						SeriesType: meta.NullType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: fmt.Sprintf("%s[%s] = %s.data[%s]", res, resIndex, op1, op1Index)}
						},
					}),
				}
			} else if resType != meta.ErrorType {
				data[fileName].Operations[opName] = Operation{
					OpCode: data[fileName].Operations[opName].OpCode,
					ApplyTo: append(data[fileName].Operations[opName].ApplyTo, OperationApplyTo{
						SeriesName: "SeriesNA",
						SeriesType: meta.NullType,
						MakeOperation: func(res, resIndex, op1, op1Index, op2, op2Index string) ast.Expr {
							return &ast.Ident{Name: ""}
						},
					}),
				}
			}
		}
	}

	return data
}
