package animal

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
)


func Introduce()  {
	mysql.NewConfig();
	fmt.Println("I am a cat");
}