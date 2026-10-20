package handlers

import (
	"github.com/libost/gotgbot/v2"
	"github.com/libost/gotgbot/v2/ext"
)

type Response func(b *gotgbot.Bot, ctx *ext.Context) error
