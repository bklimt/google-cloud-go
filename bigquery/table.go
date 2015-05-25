// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bigquery

import bq "google.golang.org/api/bigquery/v2"

// A Table is a reference to a BigQuery table.
type Table struct {
	// ProjectID, DatasetID and TableID must be set.
	// All other fields are optional.
	ProjectID string
	DatasetID string
	TableID   string

	CreateDisposition CreateDisposition // default is CreateIfNeeded.

	WriteDisposition WriteDisposition // default is WriteAppend.
}

// CreateDisposition specifies the circumstances under which destination table will be created.
type CreateDisposition string

const (
	// The table will be created if it does not already exist.  Tables are created atomically on successful completion of a job.
	CreateIfNeeded CreateDisposition = "CREATE_IF_NEEDED"

	// The table must already exist and will not be automatically created.
	CreateNever CreateDisposition = "CREATE_NEVER"
)

// WriteDisposition specifies how existing data in a destination table is treated.
type WriteDisposition string

const (
	// Data will be appended to any existing data in the destination table.
	// Data is appended atomically on successful completion of a job.
	WriteAppend WriteDisposition = "WRITE_APPEND"

	// Existing data in the destination table will be overwritten.
	// Data is overwritten atomically on successful completion of a job.
	WriteTruncate WriteDisposition = "WRITE_TRUNCATE"

	// Writes will fail if the destination table already contains data.
	WriteEmpty WriteDisposition = "WRITE_EMPTY"
)

func (t *Table) implementsSource() {
}

func (t *Table) implementsDestination() {
}

func (t *Table) customizeLoadDst(conf *bq.JobConfigurationLoad) {
	conf.DestinationTable = &bq.TableReference{
		ProjectId: t.ProjectID,
		DatasetId: t.DatasetID,
		TableId:   t.TableID,
	}
	conf.CreateDisposition = string(t.CreateDisposition)
	conf.WriteDisposition = string(t.WriteDisposition)
}

func (t *Table) customizeExtractSrc(conf *bq.JobConfigurationExtract) {
	conf.SourceTable = &bq.TableReference{
		ProjectId: t.ProjectID,
		DatasetId: t.DatasetID,
		TableId:   t.TableID,
	}
}

func (t *Table) customizeCopySrc(conf *bq.JobConfigurationTableCopy) {
	// TODO: support copying multiple tables.
	conf.SourceTable = &bq.TableReference{
		ProjectId: t.ProjectID,
		DatasetId: t.DatasetID,
		TableId:   t.TableID,
	}
}

func (t *Table) customizeCopyDst(conf *bq.JobConfigurationTableCopy) {
	conf.DestinationTable = &bq.TableReference{
		ProjectId: t.ProjectID,
		DatasetId: t.DatasetID,
		TableId:   t.TableID,
	}
	conf.CreateDisposition = string(t.CreateDisposition)
	conf.WriteDisposition = string(t.WriteDisposition)
}