// Package main ReiBot-Plugin
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	_ "github.com/FloatTech/ReiBot-Plugin/plugin/b14"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/base64gua"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/baseamasiro"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/bilibili_parse"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/chrev"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/emojimix"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/fortune"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/genshin"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/heisi"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/hyaku"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/lolicon"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/manager"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/moegoe"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/runcode"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/saucenao"
	_ "github.com/FloatTech/ReiBot-Plugin/plugin/tracemoe"

	// -----------------------以下为内置依赖，勿动------------------------ //
	rei "github.com/fumiama/ReiBot"
	tgba "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"

	"github.com/FloatTech/ReiBot-Plugin/kanban"
	// -----------------------以上为内置依赖，勿动------------------------ //
)

func main() {
	token := flag.String("t", "", "telegram api token")
	buffer := flag.Int("b", 256, "message sequence length")
	debug := flag.Bool("d", false, "enable debug-level log output")
	offset := flag.Int("o", 0, "the last Update ID to include")
	timeout := flag.Int("T", 60, "timeout")
	help := flag.Bool("h", false, "print this help")
	flag.Parse()
	if *help {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	sus := make([]int64, 0, 16)
	for _, s := range flag.Args() {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		sus = append(sus, i)
	}

	rei.OnMessageCommandGroup([]string{"help", "帮助", "menu", "菜单"}, rei.OnlyToMe).SetBlock(true).
		Handle(func(ctx *rei.Ctx) {
			_, _ = ctx.SendPlainMessage(false, kanban.Banner)
		})
	rei.Run(rei.Bot{
		Token:  *token,
		Buffer: *buffer,
		UpdateConfig: tgba.UpdateConfig{
			Offset:  *offset,
			Limit:   0,
			Timeout: *timeout,
		},
		SuperUsers: sus,
		Debug:      *debug,
	})
}
