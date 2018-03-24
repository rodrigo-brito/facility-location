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
		cpu     int
		verbose bool
	)

	var rootCmd = &cobra.Command{
		Use:   "hub-spoke-go [input-file]",
		Short: "Read input file and create a network design",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			run(args[0], cpu, verbose)
		},
	}

	rootCmd.Flags().IntVarP(&cpu, "cpu", "c", 1, "number of CPUs to use")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "active verbose mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
