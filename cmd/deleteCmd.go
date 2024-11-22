package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type DeleteStruct struct {
	BaseUrl string
	Body    string
}

var DeleteStructVar *DeleteStruct
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete request",
	Run: func(cmd *cobra.Command, args []string) {
		if DeleteStructVar.BaseUrl==""{
			fmt.Println("Pls enter URL for Delete Request")
			return
		}
		
	},
}

func init() {

	deleteCmd.Flags().StringVarP(&DeleteStructVar.BaseUrl, "url", "u", "", "Set Base URL")
	deleteCmd.Flags().StringVarP(&DeleteStructVar.Body, "body", "b", "", "Set Body for Delete Request")
}
