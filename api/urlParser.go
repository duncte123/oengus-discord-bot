package api

import (
	"net/url"
)

func GetBotParamsFromUrl(botUrl string) (*BotHookParams, error) {
	parsed, e := url.Parse(botUrl)

	if e != nil {
		return nil, e
	}

	params := parsed.Query()
	obj := &BotHookParams{}

	if params.Has("editsub") {
		obj.EditSub = params.Get("editsub")
	}

	if params.Has("newsub") {
		obj.NewSub = params.Get("newsub")
	}

	if params.Has("donation") {
		obj.Donation = params.Get("donation")
	}

	if params.Has("marathon") {
		obj.MarathonId = params.Get("marathon")
	}

	return obj, nil
}
