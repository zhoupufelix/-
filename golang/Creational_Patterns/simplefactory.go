/**
 * 简单工厂
 */
package simplefactory

import "fmt"


//API is an interface
type API interface{
  Say(name string) string
}

//omsApi implements the Api interface
type  omsApi struct{}

func(*omsApi) Say(name string) string{
  return fmt.Sprintf("oms,%s",name)
}

//wmsApi  implemnets the Api interface
type wmsApi struct{}

func (*wmsApi)Say(name string) string{
  return fmt.Sprintf("wms,%s",name)
}

func NewApi(t string) API{
  switch t{
    case "oms":
      return &omsApi{}
    case "wms":
      return &wmsApi{}
    default:
      return nil
  }
}
