package main

import(
  "fmt"
  "unseal-vault/internal/vault"
  "unseal-vault/internal/consul"
)

func main() {
  var key string
  consulClient := consul.CreateConsul("consul.service.consul:8500")
  services:=consulClient.SearchService("vault")
  for _, element :=range services {
    url:=fmt.Sprint("http://",element.ServiceAddress, ":", element.ServicePort)
    vaultClient:=vault.CreateVault(url)
    if vaultClient.Vaultisseal(){
      fmt.Println(url + "is sealed")
      if key == "" {
        fmt.Print("enter your key: ")
        fmt.Scan(&key)
      }
      if vaultClient.Vaultunseal(key){
        fmt.Println("unseal success")
      }
     }else{
      fmt.Println(url + " not sealed")
    }
   }
}


