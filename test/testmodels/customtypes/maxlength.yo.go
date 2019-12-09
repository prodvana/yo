// Code generated by yo. DO NOT EDIT.
// Package customtypes contains the types.
package customtypes

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// MaxLength represents a row from 'MaxLengths'.
type MaxLength struct {
	MaxString string `spanner:"MaxString" json:"MaxString"` // MaxString
	MaxBytes  []byte `spanner:"MaxBytes" json:"MaxBytes"`   // MaxBytes
}

func MaxLengthPrimaryKeys() []string {
	return []string{
		"MaxString",
	}
}

func MaxLengthColumns() []string {
	return []string{
		"MaxString",
		"MaxBytes",
	}
}

func (ml *MaxLength) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "MaxString":
			ret = append(ret, &ml.MaxString)
		case "MaxBytes":
			ret = append(ret, &ml.MaxBytes)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (ml *MaxLength) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "MaxString":
			ret = append(ret, ml.MaxString)
		case "MaxBytes":
			ret = append(ret, ml.MaxBytes)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newMaxLength_Decoder returns a decoder which reads a row from *spanner.Row
// into MaxLength. The decoder is not goroutine-safe. Don't use it concurrently.
func newMaxLength_Decoder(cols []string) func(*spanner.Row) (*MaxLength, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*MaxLength, error) {
		var ml MaxLength
		ptrs, err := ml.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &ml, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (ml *MaxLength) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("MaxLengths", MaxLengthColumns(), []interface{}{
		ml.MaxString, ml.MaxBytes,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (ml *MaxLength) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("MaxLengths", MaxLengthColumns(), []interface{}{
		ml.MaxString, ml.MaxBytes,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (ml *MaxLength) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("MaxLengths", MaxLengthColumns(), []interface{}{
		ml.MaxString, ml.MaxBytes,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (ml *MaxLength) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, MaxLengthPrimaryKeys()...)

	values, err := ml.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "MaxLength.UpdateColumns", "MaxLengths", err)
	}

	return spanner.Update("MaxLengths", colsWithPKeys, values), nil
}

// FindMaxLength gets a MaxLength by primary key
func FindMaxLength(ctx context.Context, db YORODB, maxString string) (*MaxLength, error) {
	key := spanner.Key{maxString}
	row, err := db.ReadRow(ctx, "MaxLengths", key, MaxLengthColumns())
	if err != nil {
		return nil, newError("FindMaxLength", "MaxLengths", err)
	}

	decoder := newMaxLength_Decoder(MaxLengthColumns())
	ml, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindMaxLength", "MaxLengths", err)
	}

	return ml, nil
}

// ReadMaxLength retrieves multiples rows from MaxLength by KeySet as a slice.
func ReadMaxLength(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*MaxLength, error) {
	var res []*MaxLength

	decoder := newMaxLength_Decoder(MaxLengthColumns())

	rows := db.Read(ctx, "MaxLengths", keys, MaxLengthColumns())
	err := rows.Do(func(row *spanner.Row) error {
		ml, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, ml)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadMaxLength", "MaxLengths", err)
	}

	return res, nil
}

// Delete deletes the MaxLength from the database.
func (ml *MaxLength) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := ml.columnsToValues(MaxLengthPrimaryKeys())
	return spanner.Delete("MaxLengths", spanner.Key(values))
}
