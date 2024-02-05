// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package distri_ai

import (
	"bytes"
	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestEncodeDecode_PlaceOrder(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("PlaceOrder"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(PlaceOrder)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(PlaceOrder)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}
