package crypto

import (
	"encoding/base64"
	"github.com/vela-security/vela-public/lua"
)

func (c *crypto) Base64Encode(L *lua.LState) lua.LValue {
	data := L.CheckString(1)
	r := base64.StdEncoding.EncodeToString(lua.S2B(data))
	return lua.LString(r)
}

func (c *crypto) Base64Decode(L *lua.LState) lua.LValue {
	data := L.CheckString(1)
	r, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return lua.LNil
	}
	return lua.LString(r)
}
