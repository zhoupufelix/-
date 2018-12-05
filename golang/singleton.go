/**
 * 单例模式，是一种常用的软件设计模式。
 * 在它的核心结构中只包含一个被称为单例的特殊类。
 * 通过单例模式可以保证系统中，应用该模式的类一个类只有一个实例。
 * 即一个类只有一个对象实例。
 */

 package singleton

 import (
 	"sync"
 	"fmt"
 )

 type db struct{}

 var (
 	once sync.Once
 	m *db
 )

 func GetInstance() *db{
 	once.Do(
 		func() {
 			m = &db{}
 		}
 	)
 	return m
 }

 func (this *db)Test(a int){
 	fmt.Println("this is test",a)
 }
