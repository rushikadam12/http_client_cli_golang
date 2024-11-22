package cmd

import (
	"fmt"
	"io"
	"net/http"
	
	"regexp"

	"github.com/spf13/cobra"
)

type GetStruct struct {
	BaseUrl string
}

var GetStructVar *GetStruct

func ValidateBaseURL(url *GetStruct) bool {
	pattern := `^(http[s]?:\/\/)?(([0-9]{1,3}\.){3}[0-9]{1,3}|([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,})(:\d+)?(\/[^\s]*)?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(url.BaseUrl)
}

func GetResponse(flag *GetStruct) {
	resp, err := http.Get(flag.BaseUrl)
	if err != nil {
		fmt.Println("Error while making GET request:", err)
		return
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error while parsing response", err)
		return
	}

	fmt.Print("GET REQUEST RESPONSE")

	fmt.Println(" \n")
	fmt.Print("STATUS:", resp.StatusCode)
	fmt.Println("\n")

	fmt.Println("\n")
	fmt.Println(string(respData))
	fmt.Println("\n")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send the GET request",
	Run: func(cmd *cobra.Command, args []string) {

		if GetStructVar.BaseUrl == "" {
			fmt.Println("Pls provide the URL to make GET Request")
			return

		}

		println("")
		fmt.Print("URL:", GetStructVar.BaseUrl)
		fmt.Println("")

		validation_result := ValidateBaseURL(GetStructVar)

		if !validation_result {			fmt.Println("!!Pls enter valid url")
			return
		}
		GetResponse(GetStructVar)

	},
}

func init() {
	GetStructVar=&GetStruct{
		BaseUrl: "",
	}
	getCmd.Flags().StringVarP(&GetStructVar.BaseUrl,"url", "u", "", "Set the base URL")

}
