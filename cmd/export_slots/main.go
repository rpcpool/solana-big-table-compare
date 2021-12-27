package main

import (
	"context"
	"flag"
	"fmt"

	"cloud.google.com/go/bigtable"
	"github.com/rpcpool/solana-big-table-compare/cmd"
	"github.com/rpcpool/solana-big-table-compare/ledger"
	"go.uber.org/zap"
)

func main() {
	project := flag.String("project", "", "Google Cloud Project ID")
	instance := flag.String("instance", "solana-ledger", "Solana Bigtable instance name")
	start := flag.Int64("start", 0, "Start slot number")
	maxRows := flag.Int64("rows", -1, "Max number of slots to return")
	noPad := flag.Bool("no-pad", false, "Print numbers without padding")
	flag.Parse()

	log := cmd.MustNewLogger()

	if *project == "" {
		flag.Usage()
		log.Fatal("Missing -project flag")
	}
	if *instance == "" {
		flag.Usage()
		log.Fatal("Missing -instance flag")
	}

	ctx := context.Background()

	// Create BigTable instance client.
	btClient, err := bigtable.NewClient(ctx, *project, *instance)
	if err != nil {
		log.Fatal("Failed to connect to BigTable", zap.Error(err))
	}
	// Create Solana-Ledger client.
	client := ledger.NewClient(btClient)
	defer func() {
		if err := client.Close(); err != nil {
			log.Error("Error while closing BigTable instance", zap.Error(err))
		}
	}()

	// Test connection.
	firstSlot, err := client.GetFirstAvailableBlock(ctx)
	if err != nil {
		log.Fatal("Failed to get first available block", zap.Error(err))
	}
	log.Info("Connected to BigTable", zap.Int64("first_slot", firstSlot))

	// Iterate block numbers.
	lastPrint := int64(0)
	rows := int64(0)
	pad := !*noPad
	err = client.IterateSlots(ctx, *start, func(slot int64) bool {
		if pad {
			fmt.Printf("%12d\n", slot)
		} else {
			fmt.Println(slot)
		}
		rows++
		if rows-lastPrint > 10000 {
			log.Info("Scanning blocks", zap.Int64("current_slot", slot))
			lastPrint = rows
		}
		return *maxRows <= 0 || *maxRows >= rows
	})
	if err != nil {
		log.Error("Error iterating slots", zap.Error(err))
	}
}
