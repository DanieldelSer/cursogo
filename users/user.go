package users

import (
	"fmt"
	"time"

	"github.com/cursogo/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Daniel", time.Now(), true)
	fmt.Println(u)
}
