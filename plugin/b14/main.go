// Package b14coder base16384 与 tea 加解密
package b14coder

import (
	rei "github.com/fumiama/ReiBot"

	base14 "github.com/fumiama/go-base16384"

	"github.com/FloatTech/floatbox/crypto"
	ctrl "github.com/FloatTech/zbpctrl"

	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

func init() {
	en := rei.Register("base16384", &ctrl.Options[*rei.Ctx]{
		DisableOnDefault: false,
		Help: "base16384加解密\n" +
			"- 加密xxx\n- 解密xxx\n- 用yyy加密xxx\n- 用yyy解密xxx",
	})
	en.OnMessageRegex(`^加密\s*(.+)$`).SetBlock(true).
		Handle(func(ctx *rei.Ctx) {
			str := ctx.State["regex_matched"].([]string)[1]
			es := base14.EncodeString(str)
			if es != "" {
				_, _ = ctx.SendPlainMessage(false, es)
			} else {
				_, _ = ctx.SendPlainMessage(false, "加密失败!")
			}
		})
	en.OnMessageRegex(`^解密\s*([一-踀]+[㴁-㴆]?)$`).SetBlock(true).
		Handle(func(ctx *rei.Ctx) {
			str := ctx.State["regex_matched"].([]string)[1]
			es := base14.DecodeString(str)
			if es != "" {
				_, err := ctx.SendPlainMessage(false, es)
				if err != nil {
					_, _ = ctx.SendPlainMessage(false, "ERROR: ", err)
				}
			} else {
				_, _ = ctx.SendPlainMessage(false, "解密失败!")
			}
		})
	en.OnMessageRegex(`^用(.+)加密\s*(.+)$`).SetBlock(true).
		Handle(func(ctx *rei.Ctx) {
			key, str := ctx.State["regex_matched"].([]string)[1], ctx.State["regex_matched"].([]string)[2]
			t := crypto.GetTEA(key)
			es, err := base14.UTF16BE2UTF8(base14.Encode(t.Encrypt(helper.StringToBytes(str))))
			if err == nil {
				_, _ = ctx.SendPlainMessage(false, helper.BytesToString(es))
			} else {
				_, _ = ctx.SendPlainMessage(false, "加密失败!")
			}
		})
	en.OnMessageRegex(`^用(.+)解密\s*([一-踀]+[㴁-㴆]?)$`).SetBlock(true).
		Handle(func(ctx *rei.Ctx) {
			key, str := ctx.State["regex_matched"].([]string)[1], ctx.State["regex_matched"].([]string)[2]
			t := crypto.GetTEA(key)
			es, err := base14.UTF82UTF16BE(helper.StringToBytes(str))
			if err == nil {
				_, err := ctx.SendPlainMessage(false, helper.BytesToString(t.Decrypt(base14.Decode(es))))
				if err != nil {
					_, _ = ctx.SendPlainMessage(false, "ERROR: ", err)
				}
			} else {
				_, _ = ctx.SendPlainMessage(false, "解密失败!")
			}
		})
}
