/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	"vidar-scan/basework"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug command",
	Long:  `我用来测试代码的!`,
	Run: func(cmd *cobra.Command, args []string) {
		//data, err := os.ReadFile("fastText/source/tomcat404.html")
		data, err := os.ReadFile("fastText/source/zhihu.html")
		if err != nil {
			panic(err)
		}

		start := time.Now()
		output, err := basework.HTMLPreprocess(string(data))
		if err != nil {
			panic(err)
		}

		Label, err := basework.Predict404(output)
		if err != nil {
			panic(err)
		}

		fmt.Println(Label)
		elapsed := time.Since(start)
		fmt.Printf("elapsed: %s\n", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
