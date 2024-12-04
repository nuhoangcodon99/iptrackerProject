/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the location of an IP address",
	Long:  `Trace the location of an IP address using the IP Tracker CLI Application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
				showData(ip)
			}
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type IPData struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	var data IPData
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		fmt.Println(err)
	}

	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("IP Tracker Results")

	fmt.Println("IP Address:", data.IP)
	fmt.Println("Hostname:", data.Hostname)
	fmt.Println("City:", data.City)
	fmt.Println("Region:", data.Region)
	fmt.Println("Country:", data.Country)
	fmt.Println("Location:", data.Loc)
	fmt.Println("Organization:", data.Org)
	fmt.Println("Postal:", data.Postal)
	fmt.Println("Timezone:", data.Timezone)
}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return responseByte
}
