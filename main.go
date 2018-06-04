package main

import (
	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/model/solver"
	"github.com/spf13/cobra"

	"github.com/rodrigo-brito/hub-spoke-go/util/log"
)

func run(inputFile string, asyncLimit int, verbose bool) {
	log.Init(verbose)

	network, err := network.FromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	solver := solver.New(
		solver.WithNetworkData(network),
		solver.WithMaxAsyncTasks(asyncLimit),
	)

	solver.Solve()
}

func main() {
	var (
		async   int
		verbose bool
	)

	var rootCmd = &cobra.Command{
		Use:   "hub-spoke-go [input-file]",
		Short: "Read input file and create a network design",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			run(args[0], async, verbose)
		},
	}

	rootCmd.Flags().IntVarP(&async, "async", "a", 4, "number of async tasks")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "active verbose mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
