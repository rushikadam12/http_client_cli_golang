package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type DeleteStruct struct {
	BaseUrl string
	Body    string
}

var DeleteStructVar *DeleteStruct
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "send DELETE request",
	Run: func(cmd *cobra.Command, args []string) {
		if DeleteStructVar.BaseUrl == "" {
			fmt.Println("Pls enter URL for Delete Request")
			return
		}

		request, err := http.NewRequest("DELETE", DeleteStructVar.BaseUrl, bytes.NewBuffer([]byte(DeleteStructVar.Body)))
		if err != nil {
			fmt.Println("error found in Delete Request:", err)
		}
		
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println("Error found:", err)
		}
		result,err:=io.ReadAll(resp.Body)
		if err!=nil{
			fmt.Println("error while parsing response")
		}
		defer resp.Body.Close()

		fmt.Println("\n")
		fmt.Println("STATUS:", resp.Status)
		fmt.Println("\n")
		fmt.Println("request:")
		fmt.Println("\n")
		fmt.Println(string(result)	)
	},
}

func init() {
	DeleteStructVar = &DeleteStruct{
		BaseUrl: "",
		Body:    "",
	}
	deleteCmd.Flags().StringVarP(&DeleteStructVar.BaseUrl, "url", "u", "", "Set Base URL")
	deleteCmd.Flags().StringVarP(&DeleteStructVar.Body, "body", "b", "", "Set Body for Delete Request")
}
