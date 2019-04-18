package main

import (
	"fmt"
	"moqikaka.com/mongodb_go/src/bll/herobll"
	"moqikaka.com/mongodb_go/src/dal"
)

func main() {
	// connect mongodb
	dal.StartConn()
	heroList := herobll.GetData("04183f22-d990-4d5c-9d1a-5ef4e6800088_769")
	fmt.Println(heroList[0].MoldeID)
	//herobll.UpdateInfo1()

	// test get data
}
