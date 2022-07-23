package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	command := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viper.GetString("Flag"))
		},
	}

	viper.SetDefault("Flag", "Flag Value")

	command.Execute()
}
