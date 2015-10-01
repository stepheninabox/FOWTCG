// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package plans_test

import (
	. "github.com/pingcap/check"
	"github.com/pingcap/tidb/column"
	"github.com/pingcap/tidb/expression"
	"github.com/pingcap/tidb/field"
	"github.com/pingcap/tidb/model"
	mysql "github.com/pingcap/tidb/mysqldef"
	"github.com/pingcap/tidb/parser/opcode"
	"github.com/pingcap/tidb/plan/plans"
	"github.com/pingcap/tidb/rset/rsets"
	"github.com/pingcap/tidb/util/mock"
	"github.com/pingcap/tidb/util/types"
)

type testJoinSuit struct {
	cols []*column.Col
}

var _ = Suite(&testJoinSuit{})

func (s *testJoinSuit) TestJoin(c *C) {
	s.cols = []*column.Col{
		{
			ColumnInfo: model.ColumnInfo{
				ID:           0,
				Name:         model.NewCIStr("id"),
				Offset:       0,
				DefaultValue: 0,
				FieldType:    *types.NewFieldType(mysql.TypeLonglong),
			},
		},
		{
			ColumnInfo: model.ColumnInfo{
				ID:           1,
				Name:         model.NewCIStr("name"),
				Offset:       1,
				DefaultValue: nil,
				FieldType:    *types.NewFieldType(mysql.TypeVarchar),
			},
		},
	}

	var testData1 = []*testRowData{
		{1, []interface{}{10, "10"}},
		{2, []interface{}{10, "20"}},
		{3, []interface{}{10, "30"}},
		{4, []interface{}{40, "40"}},
		{6, []interface{}{60, "60"}},
	}

	var testData2 = []*testRowData{
		{1, []interface{}{10, "10"}},
		{2, []interface{}{10, "20"}},
		{3, []interface{}{10, "30"}},
		{4, []interface{}{40, "40"}},
		{6, []interface{}{60, "60"}},
	}

	tblPlan1 := &testTablePlan{testData1, []string{"id", "name"}, 0}
	tblPlan2 := &testTablePlan{testData2, []string{"id", "name"}, 0}

	joinPlan := &plans.JoinPlan{
		Left:   tblPlan1,
		Right:  tblPlan2,
		Type:   "CROSS",
		Fields: []*field.ResultField{},
		On:     expression.Value{Val: true},
	}
	rset := rsets.Recordset{
		Plan: joinPlan,
		Ctx:  mock.NewContext()}

	rset.Do(func(data []interface{}) (bool, error) {
		return true, nil
	})
	tblPlan1.Close()
	tblPlan2.Close()
	rset.Plan = &plans.JoinPlan{
		Left:   tblPlan1,
		Right:  tblPlan2,
		Type:   "LEFT",
		Fields: []*field.ResultField{},
		On:     expression.Value{Val: true},
	}
	rset.Do(func(data []interface{}) (bool, error) {
		return true, nil
	})
	tblPlan1.Close()
	tblPlan2.Close()
	joinPlan = &plans.JoinPlan{
		Left:   tblPlan1,
		Right:  tblPlan2,
		Type:   "RIGHT",
		Fields: []*field.ResultField{},
		On:     expression.Value{Val: true},
	}

	expr := &expression.BinaryOperation{
		Op: opcode.LT,
		L: &expression.Ident{
			CIStr: model.NewCIStr("id"),
		},
		R: expression.Value{
			Val: 100,
		},
	}
	np, _, err := joinPlan.Filter(nil, expr)
	c.Assert(np, NotNil)
	c.Assert(err, IsNil)
	rset.Plan = joinPlan
	rset.Do(func(data []interface{}) (bool, error) {
		return true, nil
	})
	tblPlan1.Close()
	tblPlan2.Close()
	rset.Plan = &plans.JoinPlan{
		Left:   tblPlan1,
		Right:  tblPlan2,
		Type:   "FULL",
		Fields: []*field.ResultField{},
		On:     expression.Value{Val: true},
	}

	rset.Do(func(data []interface{}) (bool, error) {
		return true, nil
	})

}
