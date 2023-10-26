//go:build rocksdbBackend
// +build rocksdbBackend

package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/sei-protocol/sei-db/benchmark/dbbackend"
	"github.com/spf13/cobra"
)

func BenchmarkReadCmd() *cobra.Command {
	benchmarkReadCmd := &cobra.Command{
		Use:   "benchmark-read",
		Short: "Benchmark read is designed to measure read performance of different db backends",
		Run:   benchmarkRead,
	}

	benchmarkReadCmd.PersistentFlags().StringP("db-backend", "d", "", "DB Backend")
	benchmarkReadCmd.PersistentFlags().StringP("raw-kv-input-dir", "r", "", "Input Directory for benchmark which contains the raw kv data")
	benchmarkReadCmd.PersistentFlags().StringP("output-dir", "o", "", "Output Directory")
	benchmarkReadCmd.PersistentFlags().IntP("concurrency", "c", 1, "Concurrency while writing to db")
	benchmarkReadCmd.PersistentFlags().Int64P("max-operations", "p", 1000, "Max operations to run")
	benchmarkReadCmd.PersistentFlags().IntP("num-versions", "v", 1, "number of versions in db")

	return benchmarkReadCmd
}

func benchmarkRead(cmd *cobra.Command, args []string) {
	dbBackend, _ := cmd.Flags().GetString("db-backend")
	rawKVInputDir, _ := cmd.Flags().GetString("raw-kv-input-dir")
	outputDir, _ := cmd.Flags().GetString("output-dir")
	numVersions, _ := cmd.Flags().GetInt("num-versions")
	concurrency, _ := cmd.Flags().GetInt("concurrency")
	maxOps, _ := cmd.Flags().GetInt64("max-operations")

	if dbBackend == "" {
		panic("Must provide db backend when benchmarking")
	}

	if rawKVInputDir == "" {
		panic("Must provide raw kv input dir when benchmarking")
	}

	if outputDir == "" {
		panic("Must provide output dir")
	}

	_, isAcceptedBackend := ValidDBBackends[dbBackend]
	if !isAcceptedBackend {
		panic(fmt.Sprintf("Unsupported db backend: %s\n", dbBackend))
	}

	BenchmarkRead(rawKVInputDir, numVersions, outputDir, dbBackend, concurrency, maxOps)
}

// Benchmark read latencies and throughput of db backend
func BenchmarkRead(inputKVDir string, numVersions int, outputDir string, dbBackend string, concurrency int, maxOps int64) {
	// Create output directory
	err := os.MkdirAll(outputDir, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	// Iterate over files in directory
	fmt.Printf("Reading Raw Keys and Values from %s\n", inputKVDir)

	if dbBackend == RocksDBBackendName {
		backend := dbbackend.RocksDBBackend{}
		backend.BenchmarkDBRead(inputKVDir, numVersions, outputDir, concurrency, maxOps)
	}

	return
}
