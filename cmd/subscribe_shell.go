package cmd

import (
	"Txray/core/node"
	"Txray/core/setting"
	"Txray/log"
	"Txray/tools"
	"Txray/tools/format"
	"github.com/abiosoft/ishell"
	"os"
	"strings"
)

func InitSubscribeShell(shell *ishell.Shell) {
	sub := &ishell.Cmd{
		Name: "sub",
		Help: "订阅管理, 使用sub查看帮助信息",
		Func: func(c *ishell.Context) {
			var key string
			if len(c.Args) == 1 {
				key = c.Args[0]
			} else {
				key = "all"
			}
			format.ShowSub(os.Stdout, node.GetSubscribe(key)...)
		},
	}
	sub.AddCmd(&ishell.Cmd{
		Name: "help",
		Help: "",
		Func: func(c *ishell.Context) {
			c.Println(HelpSub())
		},
	})
	// 添加订阅
	sub.AddCmd(&ishell.Cmd{
		Name: "add",
		Help: "添加订阅",
		Func: func(c *ishell.Context) {
			argMap := FlagsParse(c.Args, map[string]string{
				"r": "remarks",
			})
			if len(c.Args) >= 1 {
				if sublink, ok := argMap["data"]; ok {
					if remarksArg, ok := argMap["remarks"]; ok {
						node.AddSubscribe(sublink, remarksArg)
					} else {
						node.AddSubscribe(sublink, "remarks")
					}
				} else {
					log.Warn("需要输入一个订阅链接")
				}
			} else if len(c.Args) == 0 {
				log.Warn("还需要输入一个订阅链接")
			}
		},
	})
	// 删除订阅
	sub.AddCmd(&ishell.Cmd{
		Name: "rm",
		Help: "删除订阅",
		Func: func(c *ishell.Context) {
			if len(c.Args) == 1 {
				key := c.Args[0]
				node.DelSubs(key)
			} else if len(c.Args) == 0 {
				log.Warn("还需要输入一个索引")
			} else {
				log.Warn(strings.Join(c.Args, " "), ": 参数过多")
			}
		},
	})
	// 修改订阅
	sub.AddCmd(&ishell.Cmd{
		Name: "mv",
		Help: "修改订阅",
		Func: func(c *ishell.Context) {
			argMap := FlagsParse(c.Args, map[string]string{
				"r": "remarks",
				"u": "url",
			})
			if key, ok := argMap["data"]; ok {

				url := argMap["url"]
				remarks := argMap["remarks"]

				using := ""
				if value, ok := argMap["using"]; ok {
					if value == "y" {
						using = "true"
					} else if value == "n" {
						using = "false"
					}
				}
				node.SetSubs(key, using, url, remarks)
			}
		},
	})
	// 更新节点
	sub.AddCmd(&ishell.Cmd{
		Name: "update-node",
		Help: "更新节点",
		Func: func(c *ishell.Context) {
			argMap := FlagsParse(c.Args, map[string]string{
				"s": "socks5",
				"h": "http",
				"a": "addr",
			})
			key := argMap["data"]
			var port uint
			mode := "none"
			addr := "127.0.0.1"
			if socks5, ok := argMap["socks5"]; ok {
				mode = "socks5"
				if tools.IsNetPort(socks5) {
					port = tools.StrToUint(socks5)
				} else {
					port = setting.SocksPort()
				}
			} else if http, ok := argMap["http"]; ok {
				mode = "http"
				if tools.IsNetPort(http) {
					port = tools.StrToUint(http)
				} else {
					port = setting.HttpPort()
				}
			}
			if address, ok := argMap["addr"]; ok {
				addr = address
			}
			node.AddNodeBySub(key, mode, addr, port)
		},
	})
	shell.AddCmd(sub)
}
