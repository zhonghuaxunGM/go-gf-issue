// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-gf-issue/internal/dao/internal"
)

// internalTestDao is internal type for wrapping internal DAO implements.
type internalTestDao = *internal.TestDao

// testDao is the data access object for table test.
// You can define custom methods on it to extend its functionality as you wish.
type testDao struct {
	internalTestDao
}

var (
	// Test is globally public accessible object for table test operations.
	Test = testDao{
		internal.NewTestDao(),
	}
)

// Fill with you ideas below.
