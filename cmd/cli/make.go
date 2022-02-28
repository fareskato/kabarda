package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"strings"
	"time"
)

func doMake(arg2, arg3 string) error {
	switch arg2 {
	case "migration":
		// get database type(postgres, mysql)
		dbType := kbr.DB.DataBaseType
		if arg3 == "" {
			exitGracefully(errors.New("please give a migration descriptive name"))
		}
		// build migration name
		fileName := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg3)
		upFile := kbr.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
		downFile := kbr.RootPath + "/migrations/" + fileName + "." + dbType + ".down.sql"

		err := copyFileFromTemplate("templates/migrations/migration."+dbType+".up.sql", upFile)
		if err != nil {
			exitGracefully(err)
		}
		err = copyFileFromTemplate("templates/migrations/migration."+dbType+".down.sql", downFile)
		if err != nil {
			exitGracefully(err)
		}
	case "auth":
		err := doAuth()
		if err != nil {
			exitGracefully(err)
		}
	case "handler":
		if arg3 == "" {
			exitGracefully(errors.New("please type handler name"))
		}
		// handler file name
		handlerFileName := kbr.RootPath + "/handlers/" + strings.ToLower(arg3) + ".go"
		if isFileExists(handlerFileName) {
			exitGracefully(errors.New(fmt.Sprintf("%s aleady exists", handlerFileName)))
		}
		// read the context of file into data
		data, err := templateFS.ReadFile("templates/handlers/handler.go.txt")
		if err != nil {
			exitGracefully(err)
		}
		handler := string(data)
		handler = strings.ReplaceAll(handler, "$HANDLERNAME$", strcase.ToCamel(arg3))

		err = ioutil.WriteFile(handlerFileName, []byte(handler), 0644)
		if err != nil {
			exitGracefully(err)
		}
	case "model":
		if arg3 == "" {
			exitGracefully(errors.New("please type model name"))
		}
		data, err := templateFS.ReadFile("templates/data/model.go.txt")
		if err != nil {
			exitGracefully(err)
		}
		model := string(data)
		// pluralize
		pluralName := pluralize.NewClient()
		var modelName = arg3
		var tableName = arg3
		// check if already plural
		if pluralName.IsPlural(arg3) {
			modelName = pluralName.Singular(arg3)
			tableName = strings.ToLower(tableName)
		} else {
			tableName = strings.ToLower(pluralName.Plural(arg3))
		}
		modelFileName := kbr.RootPath + "/data/" + strings.ToLower(modelName) + ".go"
		if isFileExists(modelFileName) {
			exitGracefully(errors.New(fmt.Sprintf("%s aleady exists", modelFileName)))
		}
		model = strings.ReplaceAll(model, "$MODELNAME$", strcase.ToCamel(modelName))
		model = strings.ReplaceAll(model, "$TABLENAME$", tableName)
		err = copyDataToFile([]byte(model), modelFileName)
		if err != nil {
			exitGracefully(err)
		}
	case "session":
		err := doSessionTable()
		if err != nil {
			exitGracefully(err)
		}

	case "key":
		rnd := kbr.RandomString(32)
		color.Blue("32 character encryption key: %s", rnd)

	case "mail":
		if arg3 == "" {
			exitGracefully(errors.New("please type mail template name"))
		}
		// mail templates(html and plain text)
		htmlMail := kbr.RootPath + "/mail/" + strings.ToLower(arg3) + ".html.gohtml"
		plainTextlMail := kbr.RootPath + "/mail/" + strings.ToLower(arg3) + ".plain.gohtml"
		// copy files
		err := copyFileFromTemplate("templates/mailer/mail.html.gohtml", htmlMail)
		if err != nil {
			exitGracefully(err)
		}
		err = copyFileFromTemplate("templates/mailer/mail.plain.gohtml", plainTextlMail)
		if err != nil {
			exitGracefully(err)
		}
	}
	return nil
}
