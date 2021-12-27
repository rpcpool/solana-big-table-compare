// Package ledger interfaces with the Solana BigTable database.
package ledger

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/bigtable"
	"go.uber.org/zap"
)

// Client accesses data stored on the "solana-ledger" BigTable instances.
type Client struct {
	Client *bigtable.Client
	Log    *zap.Logger

	blocks   *bigtable.Table
	txByAddr *bigtable.Table
	txs      *bigtable.Table
}

// NewClient takes ownership of the provided BigTable client.
func NewClient(client *bigtable.Client) *Client {
	return &Client{
		Client: client,
		Log:    zap.NewNop(),

		blocks:   client.Open("blocks"),
		txByAddr: client.Open("tx-by-addr"),
		txs:      client.Open("txs"),
	}
}

// Close permanently shuts down the client and terminates any existing connections or background resources.
func (c *Client) Close() (err error) {
	c.blocks = nil
	c.txByAddr = nil
	c.txs = nil
	err = c.Client.Close()
	c.Client = nil
	return
}

// GetFirstAvailableBlock returns the first available slot that contains a block.
func (c *Client) GetFirstAvailableBlock(ctx context.Context) (int64, error) {
	var firstKey string
	err := c.blocks.ReadRows(ctx,
		bigtable.InfiniteRange(""),
		func(row bigtable.Row) bool {
			firstKey = row.Key()
			return true
		},
		bigtable.LimitRows(1),
		bigtable.RowFilter(bigtable.CellsPerRowLimitFilter(1)),
		bigtable.RowFilter(bigtable.StripValueFilter()))
	if err != nil {
		return -1, err
	}
	return KeyToSlot(firstKey)
}

// IterateSlots reads the blocks table starting at slot number "start".
//
// onSlot is called for each block. If onSlot returns false, the iterator is stopped.
func (c *Client) IterateSlots(ctx context.Context, start int64, onSlot func(int64) bool) error {
	return c.blocks.ReadRows(ctx,
		bigtable.InfiniteRange(SlotToKey(start)),
		func(row bigtable.Row) bool {
			slot, err := KeyToSlot(row.Key())
			if err != nil {
				c.Log.Warn("Invalid block table entry in DB", zap.String("row_key", row.Key()))
				return true
			}
			return onSlot(slot)
		},
		bigtable.RowFilter(bigtable.CellsPerRowLimitFilter(1)),
		bigtable.RowFilter(bigtable.StripValueFilter()))
}

// KeyToSlot decodes a blocks table key.
func KeyToSlot(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 64)
}

// SlotToKey returns the blocks table row key for a given slot number.
func SlotToKey(slot int64) string {
	if slot < 0 {
		panic(fmt.Sprintf("negative slot number: %d", slot))
	}
	return fmt.Sprintf("%016x", slot)
}
