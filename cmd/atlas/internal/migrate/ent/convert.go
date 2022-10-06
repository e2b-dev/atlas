// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import "ariga.io/atlas/sql/migrate"

// SetRevision takes the values for each field from the given migrate.Revision.
func (rc *RevisionCreate) SetRevision(rev *migrate.Revision) *RevisionCreate {
	rc.SetID(rev.Version)
	rc.SetDescription(rev.Description)
	rc.SetType(rev.Type)
	rc.SetApplied(rev.Applied)
	rc.SetTotal(rev.Total)
	rc.SetExecutedAt(rev.ExecutedAt)
	rc.SetExecutionTime(rev.ExecutionTime)
	rc.SetError(rev.Error)
	rc.SetErrorStmt(rev.ErrorStmt)
	rc.SetHash(rev.Hash)
	rc.SetPartialHashes(rev.PartialHashes)
	rc.SetOperatorVersion(rev.OperatorVersion)
	return rc
}

// AtlasRevision returns an migrate.Revision from the current Revision.
func (r *Revision) AtlasRevision() *migrate.Revision {
	return &migrate.Revision{
		Version:         r.ID,
		Description:     r.Description,
		Type:            r.Type,
		Applied:         r.Applied,
		Total:           r.Total,
		ExecutedAt:      r.ExecutedAt,
		ExecutionTime:   r.ExecutionTime,
		Error:           r.Error,
		ErrorStmt:       r.ErrorStmt,
		Hash:            r.Hash,
		PartialHashes:   r.PartialHashes,
		OperatorVersion: r.OperatorVersion,
	}
}