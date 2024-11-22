package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	// "http_client/cmd"
	// "io"
	// "net/http"
	// "regexp"

	"github.com/spf13/cobra"
)

type PostRequest struct {
	BaseUrl string
	Exit    int
	Body    string
}

var PostRequestVar *PostRequest
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "send POST request",
	Run: func(cmd *cobra.Command, args []string) {

		if PostRequestVar.BaseUrl == "" {
			fmt.Println("pls enter your baseURL")
			os.Exit(PostRequestVar.Exit)

		}

		fmt.Println("URL:", PostRequestVar.BaseUrl)
		val, err := json.Marshal(PostRequestVar.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("URL:", string(val))
		 resp,err:=http.Post(PostRequestVar.BaseUrl,"application/json",bytes.NewBuffer(val)) 
		 if err!=nil{
			fmt.Println("result not found",err)
		}

		result,err:=io.ReadAll(resp.Body)
		if err!=nil{
			fmt.Println("post request error:",err)
		}
		resp.Body.Close()
		fmt.Println("\n",string(result))
	},
}

func init() {
	PostRequestVar = &PostRequest{
		BaseUrl: "",
		Exit:    0,
	}
	postCmd.Flags().StringVarP(&PostRequestVar.BaseUrl, "url", "u", "", "Set POST Request URL")
	postCmd.Flags().StringVarP(&PostRequestVar.Body, "body", "b", "", "Set Body")
	postCmd.Flags().IntVarP(&PostRequestVar.Exit, "exit", "e", 0, "Exit common")
}
