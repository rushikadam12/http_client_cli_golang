package cmd

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"http_client/utility"

	"github.com/spf13/cobra"
)

type PutStruct struct {
	BaseUrl string
	Body    string
}

var PutStructVar *PutStruct

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "send PUT request",
	Run: func(cmd *cobra.Command, args []string) {
		if PutStructVar.BaseUrl == "" {
			log.Fatal("Pls provide the valid url")
		}

		req, err := http.NewRequest("PUT", PutStructVar.BaseUrl, bytes.NewBuffer([]byte(PutStructVar.Body)))
		if err != nil {
			log.Fatal("error while making PUT request:", err)
		}
		req.Header.Set("Content-Type","application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("error while making request:", err)
		}
		defer resp.Body.Close()
		result,err:=io.ReadAll(resp.Body)

		if err!=nil{
			log.Fatal("error while parsing response",err)
		}
	
		utility.ResultFormat(resp.StatusCode,string(result))

	},
}

func init() {
	PutStructVar = &PutStruct{
		BaseUrl: "",
		Body:    "",
	}
	putCmd.Flags().StringVarP(&PutStructVar.BaseUrl, "url", "u", "", "set the base url")
	putCmd.Flags().StringVarP(&PutStructVar.Body, "body", "b", "", "set the body")
}
