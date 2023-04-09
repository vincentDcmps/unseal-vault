package vault

import(
    vaultapi "github.com/hashicorp/vault/api"
    "fmt"
    "os"
)

type Vault struct {
    client *vaultapi.Client
}


func CreateVault(url string) (Vault) {
    vault  := Vault{}
    client,err := vaultapi.NewClient(&vaultapi.Config{Address:url})
    if err!= nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }else{
        vault.client=client
    }
    return vault
}

func (v *Vault) Vaultisseal () (bool){
  sys:=v.client.Sys()
  health,err:= sys.Health()
  if err != nil {
    fmt.Println(err.Error())
    return false
  }
  return health.Sealed
}

func (v *Vault) Vaultunseal (key string) (bool){
    sys:=v.client.Sys()
    status,err := sys.Unseal(key)
    if err != nil {
       fmt.Println(err.Error())
    }
    if status.Sealed{
        return false
    }else{
        return true
    }
}
