// Copyright © 2018 Patrick Dube <pdube.devbox@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var domain string
var recordID string

func init() {
	rootCmd.AddCommand(recordsCmd)

	recordsCmd.Flags().StringVarP(&domain, "domainID", "d", "", "Domain ID")
	recordsCmd.AddCommand(createRecordsCmd)
	recordsCmd.AddCommand(updateRecordsCmd)
	recordsCmd.AddCommand(deleteRecordsCmd)

	createRecordsCmd.Flags().StringVarP(&domain, "domainID", "d", "", "Domain ID")

	updateRecordsCmd.Flags().StringVarP(&domain, "domainID", "d", "", "Domain ID")
	updateRecordsCmd.Flags().StringVarP(&recordID, "recordID", "r", "", "Record ID")

	deleteRecordsCmd.Flags().StringVarP(&domain, "domainID", "d", "", "Domain ID")
	deleteRecordsCmd.Flags().StringVarP(&recordID, "recordID", "r", "", "Record ID")
}

//{"name":"www","type":"A","value":"1.1.1.1","id":"57181329","gtdLocation":"DEFAULT","ttl":86400}
type record struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Value       string `json:"value,omitempty"`
	GTDLocation string `json:"gtdLocation,omitempty"`
	TTL         int    `json:"ttl,omitempty"`
}

// recordsCmd represents the records command
var recordsCmd = &cobra.Command{
	Use:   "records",
	Short: "List the records for a domain",
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			fmt.Println("Domain (-d) is required")
			os.Exit(1)
		}
		err := doRequest("GET", "/dns/managed/"+domain+"/records", nil)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var createRecordsCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an A record",
	Run: func(cmd *cobra.Command, args []string) {

		if domain == "" {
			fmt.Println("Domain (-d) is required")
			os.Exit(1)
		}

		if len(args) < 2 {
			fmt.Println("Missing name and value")
			os.Exit(1)
		}

		r := record{
			Name:        args[0],
			Value:       args[1],
			Type:        "A",
			GTDLocation: "DEFAULT",
			TTL:         3600,
		}

		err := doRequest("POST", "/dns/managed/"+domain+"/records/", r)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var updateRecordsCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an A record",
	Run: func(cmd *cobra.Command, args []string) {

		if domain == "" {
			fmt.Println("Domain (-d) is required")
			os.Exit(1)
		}
		if recordID == "" {
			fmt.Println("Record (-r) is required")
			os.Exit(1)
		}

		if len(args) < 2 {
			fmt.Println("Missing name and value")
			os.Exit(1)
		}

		r := record{
			ID:          recordID,
			Name:        args[0],
			Value:       args[1],
			Type:        "A",
			GTDLocation: "DEFAULT",
			TTL:         1800,
		}

		err := doRequest("PUT", "/dns/managed/"+domain+"/records/"+recordID, r)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var deleteRecordsCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a record",
	Run: func(cmd *cobra.Command, args []string) {

		if domain == "" {
			fmt.Println("Domain (-d) is required")
			os.Exit(1)
		}
		if recordID == "" {
			fmt.Println("Record (-r) is required")
			os.Exit(1)
		}

		err := doRequest("DELETE", "/dns/managed/"+domain+"/records/"+recordID, nil)
		if err != nil {
			fmt.Println(err)
		}
	},
}
