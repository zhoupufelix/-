<?php

/**
 * 单例模式，是一种常用的软件设计模式。
 * 在它的核心结构中只包含一个被称为单例的特殊类。
 * 通过单例模式可以保证系统中，应用该模式的类一个类只有一个实例。
 * 即一个类只有一个对象实例。
 */


class DB{
  //私有的静态变量储存实例
  private static $instance = null;

  //防止被克隆 违背了只有一个实例的概念
  private function __clone(){}

  //防止自己实例化
  private function __construct(){}

  /**
   * 公有的静态方法获取实例
   * @return [type] [description]
   */
  public static function getInstance(){
    if ( self::$instance == null ) {
      self::$instance = new self();
    }
    return self::$instance;
  }

}

//调用
$db = DB::getInstance();
