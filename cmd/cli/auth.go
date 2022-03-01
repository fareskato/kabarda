package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

func doAuth() error {
	// create migrations
	dbType := kbr.DB.DataBaseType
	fileName := fmt.Sprintf("%d_create_auth_tables", time.Now().UnixMicro())
	upFile := kbr.RootPath + "/migrations/" + fileName + ".up.sql"
	downFile := kbr.RootPath + "/migrations/" + fileName + ".down.sql"
	err := copyFileFromTemplate("templates/migrations/auth_tables."+dbType+".sql", upFile)
	if err != nil {
		exitGracefully(err)
	}
	err = copyDataToFile([]byte("drop table if exists users cascade; drop table if exists tokens cascade;"+
		" drop table if exists remember_tokens;"), downFile)
	if err != nil {
		exitGracefully(err)
	}
	// run migrations
	err = doMigrate("up", "")
	if err != nil {
		exitGracefully(err)
	}
	// copy all auth functionality related files(users and token models)
	err = copyFileFromTemplate("templates/data/user.go.txt", kbr.RootPath+"/data/user.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/data/token.go.txt", kbr.RootPath+"/data/token.go")
	if err != nil {
		exitGracefully(err)
	}
	// copy remember token
	err = copyFileFromTemplate("templates/data/remember_token.go.txt", kbr.RootPath+"/data/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}

	// copy auth middlewares
	err = copyFileFromTemplate("templates/middlewares/auth.go.txt", kbr.RootPath+"/middlewares/auth.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/middlewares/auth-token.go.txt", kbr.RootPath+"/middlewares/auth-token.go")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/middlewares/remember_token.go.txt", kbr.RootPath+"/middlewares/remember_token.go")
	if err != nil {
		exitGracefully(err)
	}
	// copy handlers
	err = copyFileFromTemplate("templates/handlers/auth-handlers.go.txt", kbr.RootPath+"/handlers/auth-handlers.go")
	if err != nil {
		exitGracefully(err)
	}
	// copy views and mailer views
	err = copyFileFromTemplate("templates/mailer/password-reset.html.gohtml", kbr.RootPath+"/mail/password-reset.html.gohtml")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/mailer/password-reset.plain.gohtml", kbr.RootPath+"/mail/password-reset.plain.gohtml")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/views/forgot.jet", kbr.RootPath+"/views/forgot.jet")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/views/reset-password.jet", kbr.RootPath+"/views/reset-password.jet")
	if err != nil {
		exitGracefully(err)
	}
	err = copyFileFromTemplate("templates/views/login.jet", kbr.RootPath+"/views/login.jet")
	if err != nil {
		exitGracefully(err)
	}
	color.Yellow("\tUpdating source files... ")
	fmt.Println(os.Getwd())
	updateSource()
	// inform developer
	color.Yellow(" - users, tokens and remember_tokens migrations created in migrations directory and ran")
	color.Yellow(" - user and token models created in data directory")
	color.Yellow(" - auth middlewares created in middlewares directory")
	color.Yellow("")
	color.Green(" - Please add user and token models in data/models.go, and add appropriate middlewares to your routes")
	return nil
}
