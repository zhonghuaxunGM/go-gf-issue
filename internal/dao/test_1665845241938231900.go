// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-gf-issue/internal/dao/internal"
)

// internalTest1665845241938231900Dao is internal type for wrapping internal DAO implements.
type internalTest1665845241938231900Dao = *internal.Test1665845241938231900Dao

// test1665845241938231900Dao is the data access object for table test_1665845241938231900.
// You can define custom methods on it to extend its functionality as you wish.
type test1665845241938231900Dao struct {
	internalTest1665845241938231900Dao
}

var (
	// Test1665845241938231900 is globally public accessible object for table test_1665845241938231900 operations.
	Test1665845241938231900 = test1665845241938231900Dao{
		internal.NewTest1665845241938231900Dao(),
	}
)

// Fill with you ideas below.
