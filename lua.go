package crypto

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

type crypto struct{}

func Constructor(env assert.Environment) {
	env.Set("crypto", lua.NewAnyData(&crypto{}, lua.Reflect(lua.FUNC)))
}
