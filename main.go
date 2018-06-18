package main

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solver"
	"github.com/spf13/cobra"

	"github.com/rodrigo-brito/facility-location/util/log"
)

func run(inputFile string, asyncLimit int, verbose bool, targetValue float64) {
	log.Init(verbose)

	network, err := network.FromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	network.MaxAsyncTask = asyncLimit

	solver := solver.New(
		solver.WithNetworkData(network),
		solver.WithMaxAsyncTasks(asyncLimit),
		solver.WithTarget(targetValue),
	)

	solver.Solve()
}

func main() {
	var (
		async       int
		verbose     bool
		targetValue float64
	)

	var rootCmd = &cobra.Command{
		Use:   "hub-spoke-go [input-file]",
		Short: "Read input file and create a network design",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			run(args[0], async, verbose, targetValue)
		},
	}

	rootCmd.Flags().IntVarP(&async, "async", "a", 4, "number of async tasks")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "active verbose mode")
	rootCmd.Flags().Float64VarP(&targetValue, "best", "b", 0, "value of the best solution")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
