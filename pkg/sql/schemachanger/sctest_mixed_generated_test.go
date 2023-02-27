// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by sctestgen, DO NOT EDIT.

package schemachanger

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/sctest"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

func TestValidateMixedVersionElements_add_column(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_add_column_default_seq(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column_default_seq", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_add_column_default_unique(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column_default_unique", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_add_column_no_default(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/add_column_no_default", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_add_check_udf(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_check_udf", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_add_check_unvalidated(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_check_unvalidated", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_add_check_vanilla(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_check_vanilla", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_add_check_with_seq_and_udt(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_check_with_seq_and_udt", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_add_foreign_key(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_foreign_key", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_add_primary_key_drop_rowid(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_primary_key_drop_rowid", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_add_unique_without_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_add_unique_without_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_alter_column_set_not_null(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_alter_column_set_not_null", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_alter_primary_key_drop_rowid(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_alter_primary_key_drop_rowid", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_alter_primary_key_vanilla(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_alter_primary_key_vanilla", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_drop_constraint_check(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_drop_constraint_check", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_drop_constraint_fk(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_drop_constraint_fk", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_drop_constraint_uwi(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_drop_constraint_uwi", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_alter_table_validate_constraint(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/alter_table_validate_constraint", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_create_function(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/create_function", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_create_function_in_txn(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/create_function_in_txn", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/create_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_column_basic(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_basic", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_column_computed_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_computed_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_column_create_index_separate_statements(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_create_index_separate_statements", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_column_unique_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_unique_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_column_with_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_with_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_column_with_partial_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_with_partial_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_column_with_udf_default(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_column_with_udf_default", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_function(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_function", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_index_hash_sharded_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_hash_sharded_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_index_partial_expression_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_partial_expression_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_index_vanilla_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_vanilla_index", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_index_with_materialized_view_dep(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_index_with_materialized_view_dep", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_multiple_columns_separate_statements(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_multiple_columns_separate_statements", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_schema(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_schema", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_table(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_table", sctest.SingleNodeMixedCluster)
}
func TestValidateMixedVersionElements_drop_table_udf_default(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	sctest.ValidateMixedVersionElements(t, "pkg/sql/schemachanger/testdata/end_to_end/drop_table_udf_default", sctest.SingleNodeMixedCluster)
}