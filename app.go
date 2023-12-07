package main

import (
	"context"
	"fmt"
	"path/filepath"
	"os"
	"io/ioutil"
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
type Message struct {
    Type     string
    Level   string
    Message string
    Array []CustomServer
}

func (a *App) GetServers()  Message{
		var customServers []CustomServer
		cloudsFile := filepath.Join(xdg.DataHome, "GuiOpenstack.yaml")
		fmt.Println(cloudsFile)
		// Chemin vers le fichier clouds.yaml
		os.Setenv("OS_CLIENT_CONFIG_FILE", cloudsFile)
		os.Unsetenv("https_proxy")
		os.Unsetenv("http_proxy")
		fmt.Println(os.Getenv("http_proxy"), "oo" )
		fmt.Println(os.Getenv("https_proxy"), "oo")
		//os.Setenv("https_proxy", "")
		//os.Setenv("http_proxy", "")
		//os.Setenv("HTTPS_PROXY", "")
		//os.Setenv("HTTP_PROXY", "")
		//os.Setenv("no_proxy", "stargate-int.enedis.fr:13000")
		//os.Setenv("NO_PROXY", "stargate-int.enedis.fr:13000")


		
		opts := new(clientconfig.ClientOpts)
		opts.Cloud = "stargate-int"
		//TO-DO  FAIRE PROXY INDEPENDAMENT DU CLUSTER
//		provider, err := clientconfig.AuthenticatedClient(opts)
		fmt.Println("fffffffff")

		computeClient, err := clientconfig.NewServiceClient("compute", opts)
		if err != nil{
			fmt.Println(err)
			return Message{Type: "Message", Level: "Error", Message: fmt.Sprint(err), Array: []CustomServer{} }
		}
	

		listOpts := servers.ListOpts{
			AllTenants: true,
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

			customServers = append(customServers, CustomServer{
				ID:   server.ID,
				Name: server.Name,
    			Status: server.Status,
				// Ajoutez d'autres champs du serveur que vous souhaitez inclure
			})
			fmt.Println("%+v\n", server)
		}
	
		
		return Message{Type: "CustomServer", Level: "Info", Message: "", Array: customServers }

	
}
