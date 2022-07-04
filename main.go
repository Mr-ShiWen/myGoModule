package main

import (
	"fmt"
	"github.com/Mr-ShiWen/myGoModule/animal"
	"github.com/Mr-ShiWen/myGoModule/plant"
	"github.com/go-sql-driver/mysql"
)


func main() {
	fmt.Println("Hello world, I am go module");
	mysql.NewConfig();

	animal.Introduce();
	plant.Introduce();
}
