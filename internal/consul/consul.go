package consul

import(
    consulapi "github.com/hashicorp/consul/api"
    "fmt"
    "os"
)

type Consul struct {
    client *consulapi.Client
}

func CreateConsul (url string) (Consul){
    consul := Consul{}
    client,err := consulapi.NewClient(&consulapi.Config{Address:url})
    if err!=nil{
        fmt.Println(err.Error())
        os.Exit(1)
    }else{
        consul.client=client
    }
    return consul
}

func  (c *Consul)SearchService(name string) []*consulapi.CatalogService{
  catalog:= c.client.Catalog()
  q := &consulapi.QueryOptions{}
  services, _, _ := catalog.Service(name,"",q)
  return services
}
