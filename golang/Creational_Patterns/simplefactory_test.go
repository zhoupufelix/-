package simplefactory

import "testing"

func TestNewOmsApi(t *testing.T){
    omsApi := NewApi("oms")
    s := omsApi.Say("felix")
    expectString := "oms,felix"

    if s != expectString {
      t.Fatal("TestOms Failed")
    }
}

func TestNewWmsApi(t *testing.T){
    wmsApi := NewApi("wms")
    s := wmsApi.Say("felix")
    expectString := "wms,felix"

    if s != expectString {
      t.Fatal("TestWms Failed")
    }
}
