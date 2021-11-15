package main

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	Decode DecodeCmd `cmd:"" help:"Decodes PTD server response."`
}

func main() {
	cli := CLI{}

	ctx := kong.Parse(&cli,
		kong.Name("ptd-response-decoder"),
		kong.Description("Decodes encrypted response sent by PTD server. Omit when decoding GetNativeToken/Login response."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
