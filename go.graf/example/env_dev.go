/*
Development (non-production) settings go here.

Obviously, best practices would be to avoid hardcoding database passwords,
etc, into this file, and instead to pull from environment variables (or the like).
*/
package main

import (
	"fmt"
	"os"

	"github.com/carbocation/go.graf"
)

type environment string

func Environment() *graf.ConfigFile {
	var logdir = `/tmp`

	res := &graf.ConfigFile{
		//These are passed to templates
		Public: &graf.ConfigPublic{
			Site:         "Форум Антонинкиного класса родителей",
			Url:          "http://www.nezabu.tk",
			ContactEmail: "uvelichitel@gmail.com",
			GACode:       "UA-00000000-0",
			GAUrl:        "http://www.nezabu.tk",
		},

		DB: &graf.ConfigDB{
			User:     "gtfo",
			Password: "",
			DBName:   "forum",
			Port:     "5432",
			PoolSize: 95,
		},

		App: &graf.ConfigApp{
			LogAccess:   os.Stdout,
			LogError:    os.Stderr,
			Identifier:  "graf",
			Environment: "dev",

			//Port that nginx (for reverse proxy) or the browser has to be pointed at
			Host: "localhost",
			WsHost: "localhost",
			Port: "9996",
			WsPort: "9996",
			//64 bit random string generated with `openssl rand -base64 64`
			Secret: `/TsvkZlJD/ZLtx+ffq4ldgupCneonDNUmCp8jpXx4ECqRX9LF5JoI9BWH5ysBtjjUcAsLyEwHNZ8X360jBP+tw==`,

			//The ID
			RootForumID: "1",
		},
	}

	al, err := os.OpenFile(fmt.Sprintf(logdir+`/%s_access.log`, res.App.Identifier), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		panic(err)
	}
	res.App.LogAccess = al

	el, err := os.OpenFile(fmt.Sprintf(logdir+`/%s_error.log`, res.App.Identifier), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0640)
	if err != nil {
		panic(err)
	}
	res.App.LogError = el

	return res
}
