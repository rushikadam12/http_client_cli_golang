package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

type FlagStruct struct {
	BaseUrl string
	Exit    int
}

func ValidateBaseURL(url *FlagStruct) bool {
	pattern := `^(http[s]?:\/\/)?(([0-9]{1,3}\.){3}[0-9]{1,3}|([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,})(:\d+)?(\/[^\s]*)?$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(url.BaseUrl)
}

func GetResponse(flag *FlagStruct){
	resp,err:=http.Get(flag.BaseUrl);
	if err!=nil{
		fmt.Println("Error while making GET request:",err)
		return
	}

	defer resp.Body.Close()

	respData,err:=io.ReadAll(resp.Body)

	if err!=nil{
		fmt.Println("error while parsing response",err)
		return
	}

	
	fmt.Print("GET REQUEST RESPONSE")
	

	fmt.Println(" \n")
	fmt.Print("STATUS:",resp.StatusCode)
	fmt.Println("\n")
	
	fmt.Println("\n")
	fmt.Println(string(respData))
	fmt.Println("\n")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send the GET request",
	Run: func(cmd *cobra.Command, args []string) {

		Flag := &FlagStruct{"", 0}

		if Flag.BaseUrl == "" {
			fmt.Println("Enter your Base Url:")
			fmt.Scanln(&Flag.BaseUrl)
		}

		for {

			println("")
			fmt.Print("URL:", Flag.BaseUrl)
			fmt.Println("")

			validation_result := ValidateBaseURL(Flag)

			if !validation_result {

				fmt.Println("!!Pls enter valid url")
				fmt.Println("Enter your Base Url again:")
				fmt.Scanln(&Flag.BaseUrl)
				continue
			}
			GetResponse(Flag)
			fmt.Println("Enter 1 for exit or 0 for continue")
			fmt.Scanln(&Flag.Exit)
			fmt.Println("\n")
			if Flag.Exit == 1 {
				os.Exit(1)
			}

		}

	},
}

func init() {
	flag := &FlagStruct{}

	getCmd.Flags().StringVarP(&flag.BaseUrl, "base url", "b", "", "Set the base URlL")
}
