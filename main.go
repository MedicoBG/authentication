package main

import (
	"authentication/utils"
)

func init() {
	utils.LoadDatabaseConfig()
	utils.LoadHashingCost()
}

func main() {

}
