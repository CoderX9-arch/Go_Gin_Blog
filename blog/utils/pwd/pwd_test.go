package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Printf(HsahPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$OGDcdogNsTQvP7wGB6E4gOhIUcXqmAcOoYjHLjgHozYXjEp/63L66", "1234"))
}
