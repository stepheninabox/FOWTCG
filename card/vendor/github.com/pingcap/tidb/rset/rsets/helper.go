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

package rsets

import (
	"github.com/pingcap/tidb/expression"
	"github.com/pingcap/tidb/field"
)

// GetAggFields gets aggregate fields position map.
func GetAggFields(fields []*field.Field) map[int]struct{} {
	aggFields := map[int]struct{}{}
	for i, v := range fields {
		if expression.ContainAggregateFunc(v.Expr) {
			aggFields[i] = struct{}{}
		}
	}
	return aggFields
}

// HasAggFields checks whether has aggregate field.
func HasAggFields(fields []*field.Field) bool {
	aggFields := GetAggFields(fields)
	return len(aggFields) > 0
}
