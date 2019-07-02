<?php
class Factory{

  public static function NewApi($name=""){
      switch ( $name ) {
          case 'oms':
            return (new Oms);
            break;
          case 'wms':
            return (new Wms);
            break;
          default:
            return (new Oms)
            break;
      }
  }
}


interface IApi{
  function Say($name="");
}

class OmsApi implements IApi{
    function Say($name=""){
        echo "oms,".$name;
    }
}

class WmsApi implements IApi{
    function Say($name=""){
        echo "wms,".$name;
    }
}

//测试
$oms = Factory::NewApi('oms');
$oms->Say('felix');

$oms = Factory::NewApi('wms');
$oms->Say('felix');
