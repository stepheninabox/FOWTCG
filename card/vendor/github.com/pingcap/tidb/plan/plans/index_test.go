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
	"fmt"

	. "github.com/pingcap/check"
	"github.com/pingcap/tidb"
	"github.com/pingcap/tidb/column"
	"github.com/pingcap/tidb/expression"
	"github.com/pingcap/tidb/field"
	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/model"
	mysql "github.com/pingcap/tidb/mysqldef"
	"github.com/pingcap/tidb/parser/opcode"
	"github.com/pingcap/tidb/plan/plans"
	"github.com/pingcap/tidb/rset/rsets"
	"github.com/pingcap/tidb/sessionctx/variable"
	"github.com/pingcap/tidb/table"
	"github.com/pingcap/tidb/table/tables"
	"github.com/pingcap/tidb/util/types"
)

type testIndexSuit struct {
	txn   kv.Transaction
	cols  []*column.Col
	tbl   table.Table
	vars  map[string]interface{}
	store kv.Storage
}

// implement Context interface
func (p *testIndexSuit) GetTxn(forceNew bool) (kv.Transaction, error) { return p.txn, nil }

func (p *testIndexSuit) FinishTxn(rollback bool) error { return nil }

// SetValue saves a value associated with this context for key
func (p *testIndexSuit) SetValue(key fmt.Stringer, value interface{}) {
	p.vars[key.String()] = value
}

// Value returns the value associated with this context for key
func (p *testIndexSuit) Value(key fmt.Stringer) interface{} {
	return p.vars[key.String()]
}

// ClearValue clears the value associated with this context for key
func (p *testIndexSuit) ClearValue(key fmt.Stringer) {}

var _ = Suite(&testIndexSuit{})

func (p *testIndexSuit) SetUpSuite(c *C) {
	store, err := tidb.NewStore(tidb.EngineGoLevelDBMemory)
	c.Assert(err, IsNil)
	p.store = store
	p.vars = map[string]interface{}{}
	p.txn, _ = p.store.Begin()
	p.cols = []*column.Col{
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
	p.tbl = tables.NewTable(2, "t2", "test", p.cols, &simpleAllocator{})

	idxCol := &column.IndexedCol{
		IndexInfo: model.IndexInfo{
			Name:  model.NewCIStr("id"),
			Table: model.NewCIStr("t2"),
			Columns: []*model.IndexColumn{
				{
					Name:   model.NewCIStr("id"),
					Offset: 0,
					Length: 0,
				},
			},
			Unique:  false,
			Primary: false,
		},
		X: kv.NewKVIndex("i", "id", false),
	}
	p.tbl.AddIndex(idxCol)

	variable.BindSessionVars(p)

	var i int64
	for i = 0; i < 10; i++ {
		p.tbl.AddRecord(p, []interface{}{i * 10, "hello"})
	}
}

func (p *testIndexSuit) TestTableNilPlan(c *C) {
	nilPlan := &plans.TableNilPlan{
		T: p.tbl,
	}
	var ids []int
	id := 0
	rset := rsets.Recordset{Plan: nilPlan, Ctx: p}
	err := rset.Do(func(data []interface{}) (bool, error) {
		id++
		ids = append(ids, id)
		return true, nil
	})
	c.Assert(err, IsNil)
	c.Assert(ids, DeepEquals, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func (p *testIndexSuit) TestIndexPlan(c *C) {
	pln := &plans.TableDefaultPlan{
		T: p.tbl,
		Fields: []*field.ResultField{
			field.ColToResultField(p.cols[0], "t"),
			field.ColToResultField(p.cols[1], "t"),
		},
	}

	// expr: id > 0
	expr := &expression.BinaryOperation{
		Op: opcode.GE,
		L: &expression.Ident{
			CIStr: model.NewCIStr("id"),
		},
		R: expression.Value{
			Val: 50,
		},
	}

	expr2 := &expression.BinaryOperation{
		Op: opcode.LT,
		L: &expression.Ident{
			CIStr: model.NewCIStr("id"),
		},
		R: expression.Value{
			Val: 100,
		},
	}

	expr3 := &expression.BinaryOperation{
		Op: opcode.LE,
		L: &expression.Ident{
			CIStr: model.NewCIStr("id"),
		},
		R: expression.Value{
			Val: 100,
		},
	}

	expr4 := &expression.BinaryOperation{
		Op: opcode.GE,
		L: &expression.Ident{
			CIStr: model.NewCIStr("id"),
		},
		R: expression.Value{
			Val: 60,
		},
	}

	expr5 := &expression.Ident{
		CIStr: model.NewCIStr("id"),
	}

	np, _, err := pln.Filter(p, expr)
	c.Assert(err, IsNil)
	c.Assert(np, NotNil)
	np, _, err = np.Filter(p, expr2)
	c.Assert(err, IsNil)
	c.Assert(np, NotNil)
	np, _, err = np.Filter(p, expr3)
	c.Assert(err, IsNil)
	c.Assert(np, NotNil)
	np, _, err = np.Filter(p, expr4)
	c.Assert(err, IsNil)
	c.Assert(np, NotNil)
	np, _, err = np.Filter(p, expr5)
	c.Assert(err, IsNil)
	c.Assert(np, NotNil)

	ret := map[int64]string{}
	rset := rsets.Recordset{
		Plan: np,
		Ctx:  p,
	}
	rset.Do(func(data []interface{}) (bool, error) {
		ret[data[0].(int64)] = data[1].(string)
		return true, nil
	})
	excepted := map[int64]string{}
	for i := 6; i < 10; i++ {
		excepted[int64(i*10)] = "hello"
	}

	expr6 := &expression.UnaryOperation{
		Op: '!',
		V: &expression.Ident{
			CIStr: model.NewCIStr("id"),
		},
	}
	np, _, err = np.Filter(p, expr6)
	c.Assert(err, IsNil)
	c.Assert(np, NotNil)
	c.Assert(ret, DeepEquals, excepted)
}

func (p *testIndexSuit) TearDownSuite(c *C) {
	p.txn.Commit()
}
