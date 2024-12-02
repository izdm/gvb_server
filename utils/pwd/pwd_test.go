package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}
func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$1DIiaTy28aswHqhGQ6MLneanZhY0QeUeRhI0lazWzGHK0cKBhJYEu", "123456"))
}
