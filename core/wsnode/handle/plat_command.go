package handle

import (
	"coredx/utils"
	"strings"

	"github.com/tidwall/gjson"
)

func plat_login_handle(msg []byte) []byte {
	node_id := gjson.Get(string(msg), "node_id")
	node_unit := gjson.Get(string(msg), "node_unit")
	node_deploy := gjson.Get(string(msg), "node_deploy")
	node_operation := gjson.Get(string(msg), "node_operation")
	login_code := gjson.Get(string(msg), "login_code")
	node_public := gjson.Get(string(msg), "node_public")
	node_pubic_md5 := gjson.Get(string(msg), "node_pubic_md5")
	node_private_md5 := gjson.Get(string(msg), "node_private_md5")
	node_algo := gjson.Get(string(msg), "node_algo")
	node_store_key := gjson.Get(string(msg), "node_store_key")
	if strings.EqualFold(utils.MD5([]byte(node_public.String())), node_pubic_md5.String()) {

	}

}
