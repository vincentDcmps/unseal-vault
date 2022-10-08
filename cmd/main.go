package main

import(
  consulapi "github.com/hashicorp/consul/api"
  "fmt"
  "os"
  vaultapi "github.com/hashicorp/vault/api"

)

func main() {
  var key string
  consulClient,err := consulapi.NewClient(&consulapi.Config{Address: "consul.service.consul:8500"})
  if err!= nil {
     fmt.Println(err.Error())
     os.Exit(1)
  }
  services:=searchService("vault",consulClient)
  for _, element :=range services {
    url:=fmt.Sprint("http://",element.ServiceAddress, ":", element.ServicePort)
    vaultClient, err:=vaultapi.NewClient(&vaultapi.Config{Address: url})
    if err!= nil {
     fmt.Println(err.Error())
     os.Exit(1)
    }
    if vaultisseal(vaultClient){
      fmt.Println(url + "is sealed")
      if key == "" {
        fmt.Print("enter your key: ")
        fmt.Scan(&key)
      }
      status,err:=vaultunseal(vaultClient,key)
      if err != nil {
       fmt.Println(err.Error())
      }
      if status.Sealed == false {
        fmt.Println("unseal success")
      }
     }else{
      fmt.Println(url + " not sealed")
    }
   }
}


func  searchService(name string,consul *consulapi.Client) []*consulapi.CatalogService{
  catalog:= consul.Catalog()
  q := &consulapi.QueryOptions{}
  services, _, _ := catalog.Service(name,"",q)
  return services
}

func vaultisseal (vaultClient *vaultapi.Client) (bool){
  sys:=vaultClient.Sys()
  health,err:= sys.Health()
  if err != nil {
    fmt.Println(err.Error())
    return false
  }
  return health.Sealed

}

func vaultunseal (vaultClient *vaultapi.Client,key string) (*vaultapi.SealStatusResponse,error){
  sys:=vaultClient.Sys()
  status,err := sys.Unseal(key)
  return status,err


}
