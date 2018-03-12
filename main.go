package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"

	"github.com/rodrigo-brito/hub-spoke-go/util/log"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		size int32
		cpu int32
		verbose bool
	)

	var rootCmd = &cobra.Command{
		Use:   "hub-spoke-go [input-file]",
		Short: "Read input file and create a network design",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("args = ", args, size, verbose, cpu)
		},
	}

	rootCmd.Flags().Int32VarP(&size, "size", "n", 0, "(required) network size, number of nodes")
	rootCmd.Flags().Int32VarP(&cpu, "cpu", "c", 1, "number of CPUs to use")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "active verbose mode")
	rootCmd.MarkFlagRequired("size")

	logLevel := logrus.ErrorLevel
	if verbose {
		logLevel = logrus.DebugLevel
	}
	log.Init(logLevel)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
