package main

import (
	"fmt"
	"path/filepath"
	"os"
	"github.com/gophercloud/utils/openstack/clientconfig"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/adrg/xdg"
)

type CustomProject struct {
	Id    string
	Name   string

}
type Message struct {
    Type     string
    Level   string
    Message string
    Array []CustomProject
}

func getServers(cloudName string)  Message{

	var customProjects []CustomProject

	cloudsFile := filepath.Join(xdg.DataHome, "GuiOpenstack.yaml")
	fmt.Println(cloudsFile)

	os.Setenv("OS_CLIENT_CONFIG_FILE", cloudsFile)
	os.Unsetenv("https_proxy")
	os.Unsetenv("http_proxy")
	fmt.Println(os.Getenv("http_proxy"), "oo")
	fmt.Println(os.Getenv("https_proxy"), "oo")

	opts := new(clientconfig.ClientOpts)
	opts.Cloud = cloudName

	identityClient, err := clientconfig.NewServiceClient("identity", opts)
	if err != nil {
		fmt.Println(err)
		return Message{Type: "Message", Level: "Error", Message: fmt.Sprint(err), Array: []CustomProject{}}
	}

	// Récupération des projets (projects)
	projectListOpts := projects.ListOpts{}
	projectPages, err := projects.List(identityClient, projectListOpts).AllPages()
	if err != nil {
		fmt.Println(err)
		return Message{Type: "Message", Level: "Error", Message: fmt.Sprint(err), Array: []CustomProject{}}
	}

	projectList, err := projects.ExtractProjects(projectPages)
	if err != nil {
		fmt.Println(err)
		return Message{Type: "Message", Level: "Error", Message: fmt.Sprint(err), Array: []CustomProject{}}
	}

	for _, project := range projectList {
		customProjects = append(customProjects, CustomProject{
			Id:   project.ID,
			Name: project.Name,
			// Ajoutez d'autres champs du projet que vous souhaitez inclure
		})
	}


		
		return Message{Type: "CustomProjects", Level: "Info", Message: "", Array: customProjects }

}