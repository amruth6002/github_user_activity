package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type GitHubEvent struct{
	Type string `json:"type"`
	Repo struct{
		Name string `json:"name"`
	} `json:"repo"`
}

var rootcmd=&cobra.Command{
	Use: "github-activity <username>",
	Short: "Fetch recent Gitub user activity",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command ,args []string)  {
		username :=args[0]
		url:=fmt.Sprintf("https://api.github.com/users/%s/events", username)
		resp,err:=http.Get(url)
		if err!=nil{
			fmt.Println("Error fetching data:",err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode!=http.StatusOK{
			fmt.Println("failed to fetch activity for\"%s\".HTTP status:%s\n",username,resp.Status)
			os.Exit(1)
		}
		body,err:=io.ReadAll(resp.Body)
		if err!=nil{
			fmt.Println("error reading response:",err)
			os.Exit(1)
		}
		var events []GitHubEvent
		err=json.Unmarshal(body,&events)
		if err!=nil{
			fmt.Println("error parsing json:",err)
			os.Exit(1)
		}
		if len(events)==0{
			fmt.Println("No recent activity found for user:",username)
			return
		}
		for _,event:=range events{
			switch event.Type {
            case "PushEvent":
                fmt.Printf("Pushed commits to %s\n", event.Repo.Name)
            case "IssuesEvent":
                fmt.Printf("Opened an issue in %s\n", event.Repo.Name)
            case "WatchEvent":
                fmt.Printf("Starred %s\n", event.Repo.Name)
            default:
                fmt.Printf("Activity %s in %s\n", event.Type, event.Repo.Name)
            }
		}

	},
}


func main(){
if err:=rootcmd.Execute();err!=nil{
	fmt.Println(err)
	os.Exit(1)
}
}