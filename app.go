package main


import (
	"context"
	"fmt"
	"path/filepath"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"io/ioutil"
	"log"
	"time"
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
	Cloud  string
    Id     string
    Name   string
    Status string
}
type Message struct {
    Type     string
    Level   string
    Message string
    Array []CustomServer
}
type MessageCloudNames struct {
    Type    string
    Level   string
    Message string
    Array []string
}
func (a *App) GetCloudNames()  MessageCloudNames{

		var cloudNames []string
		cloudsFile := filepath.Join(xdg.DataHome, "GuiOpenstack.yaml")
		os.Setenv("OS_CLIENT_CONFIG_FILE", cloudsFile)
    	clouds, err := clientconfig.LoadCloudsYAML()
    	if err != nil {
     	   fmt.Println(err)
       	 	return MessageCloudNames{Type: "Message", Level: "ERROR", Message: fmt.Sprint(err), Array: []string{}}
    	}

    for cloudName := range clouds {
        cloudNames = append(cloudNames, cloudName)
    }

    return MessageCloudNames{Type: "Message", Level: "INFO", Message: fmt.Sprint(err), Array: cloudNames}



}

func createConfigDir(){
	var configDirName := "guiOpenstack"
	dirName := configDirName
	err := os.Mkdir(filepath.Join(xdg.DataHome, dirName), 0755)
	if err!=null {
		fmt.Println("err")
	}
}


func (a *App) GetContentFile(path) {
	createConfigDir()
	cloudsFile := filepath.Join(xdg.DataHome, dirName, "GuiOpenstack.yaml")
	fmt.Println(cloudsFile)

} 




func getServers(cloudName string)  Message{

		var customServers []CustomServer
		cloudsFile := filepath.Join(xdg.DataHome, "GuiOpenstack.yaml")
		fmt.Println(cloudsFile)
		// Chemin vers le fichier clouds.yaml
		os.Setenv("OS_CLIENT_CONFIG_FILE", cloudsFile)
		os.Unsetenv("https_proxy")
		os.Unsetenv("http_proxy")
		fmt.Println(os.Getenv("http_proxy"), "oo" )
		fmt.Println(os.Getenv("https_proxy"), "oo")
		//os.Setenv("no_proxy", "stargate-int.enedis.fr:13000")
		//os.Setenv("NO_PROXY", "stargate-int.enedis.fr:13000")



		opts := new(clientconfig.ClientOpts)
		opts.Cloud = cloudName
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
		
		fmt.Println(allServers[0])
		for _, server := range allServers {

			customServers = append(customServers, CustomServer{
				Id:   server.ID,
				Name: server.Name,
    			Status: server.Status,
				Cloud: cloudName,
				// Ajoutez d'autres champs du serveur que vous souhaitez inclure
			})
		}
	
		
		return Message{Type: "CustomServer", Level: "Info", Message: "", Array: customServers }

}



func (a *App) GetServers(cloudName string)  Message{
	return getServers(cloudName)
}



func (a *App) AsyncGetServersViaEvent(cloudName string) {
	go func() {
			runtime.EventsEmit(a.ctx, "rcv:updateStateRequestGetServers",[]interface{}{cloudName, "begin"} )
			log.Println("debut ",cloudName)
			var result=getServers(cloudName)
			if (result.Level == "Error"){
			log.Println("ERROR",cloudName, result.Message)
			runtime.EventsEmit(a.ctx, "rcv:updateStateRequestGetServers",[]interface{}{cloudName, "error"} )
			runtime.EventsEmit(a.ctx, "rcv:displayError",[]interface{}{cloudName, result.Message}  )
			}
			if (result.Level == "Info"){
			runtime.EventsEmit(a.ctx, "rcv:getServers",result.Array)
			runtime.EventsEmit(a.ctx, "rcv:updateStateRequestGetServers",[]interface{}{cloudName, "okey"} )
			}
			log.Println("fin ",cloudName)
			
	}()
}



func (a *App) GreetAsyncViaEvent() {
	go func() {
		for _, m := range []string{"Hello A (from server)", "Hello B (from server)", "Hello C (from server)", "DONE"} {
			runtime.EventsEmit(a.ctx, "rcv:greet", m)
			log.Println(m)
			time.Sleep(1 * time.Second)
		}
	}()
}


