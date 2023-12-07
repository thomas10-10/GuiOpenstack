package main

import (
	"context"
	"fmt"
	"path/filepath"
	"os"
	"io/ioutil"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	//"github.com/gophercloud/gophercloud/openstack/utils"
	"github.com/gophercloud/utils/openstack/clientconfig"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	//"github.com/gophercloud/gophercloud/pagination"
	"github.com/adrg/xdg"
//    	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
//	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {

	fmt.Println("Home data directory:", xdg.DataHome)

	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://identity.example.com/v3",
		Username:         "votre_nom_utilisateur",
		Password:         "votre_mot_de_passe",
		TenantName:       "votre_tenant",
		DomainName:       "votre_domaine",
	}

	client, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println("Erreur lors de la création du client OpenStack:", err)
	}
	fmt.Println(client)
	


	return fmt.Sprintf("Hello %s, It's show time!", name)

}
func (a *App) Readconf() string {
	filePath := filepath.Join(xdg.DataHome, "GuiOpenstack.yaml")
	content, err := ioutil.ReadFile(filePath)

	if os.IsNotExist(err) {
		ioutil.WriteFile(filePath, []byte{}, 0644)
		fmt.Printf("Fichier cree")
	}

	fmt.Println("fichier present")
	return string(content)
}
func (a *App) Writeconf(content string)  {
	filePath := filepath.Join(xdg.DataHome, "GuiOpenstack.yaml")
	_ , err := ioutil.ReadFile(filePath)

	if !os.IsNotExist(err) {
		ioutil.WriteFile(filePath, []byte(content), 0644)
		fmt.Printf("Fichier modifie")
	}

	fmt.Println("fichier modifie")
	
}

type CustomServer struct {
    ID     string
    Name   string
    Status string
}


func (a *App) GetServers()  string{
		cloudsFile := filepath.Join(xdg.DataHome, "GuiOpenstack.yaml")
		fmt.Println(cloudsFile)
		// Chemin vers le fichier clouds.yaml
		os.Setenv("OS_CLIENT_CONFIG_FILE", cloudsFile)

		opts := new(clientconfig.ClientOpts)
		opts.Cloud = "openstack"

//		provider, err := clientconfig.AuthenticatedClient(opts)
			fmt.Println("fffffffff")
		computeClient, err := clientconfig.NewServiceClient("compute", opts)

		fmt.Println(err)
	





		listOpts := servers.ListOpts{
			AllTenants: false,
		}
		
		allPages, err := servers.List(computeClient, listOpts).AllPages()
		if err != nil {
			fmt.Println(err)
		}
		
		allServers, err := servers.ExtractServers(allPages)
		if err != nil {
			fmt.Println(err)
		}
		
		for _, server := range allServers {
			fmt.Println("%+v\n", server)
		}
	
		
		return ""

	
}
