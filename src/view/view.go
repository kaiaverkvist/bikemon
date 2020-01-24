package view

import (
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/kaiaverkvist/stick"
	"github.com/kaiaverkvist/stick/twig"
	"github.com/yarf-framework/yarf"
	"os"
	"path/filepath"
)

type View struct {
	Name string

	Variables map[string]stick.Value
	Context *yarf.Context
}

func New(ctx *yarf.Context) *View {
	view := &View{}

	view.Variables = make(map[string]stick.Value)
	view.Context = ctx

	return view
}

func (v *View) Render() error {

	// Set up proper pathing for the template.
	workingDirectory, _ := os.Getwd()
	templatePath := filepath.Join(workingDirectory, config.AppConfig.ViewConfig.Folder)
	documentName := v.Name + "." + config.AppConfig.ViewConfig.Extension

	// Create the template from the path
	stickTemplate := twig.New(stick.NewFilesystemLoader(templatePath))

	// Render the template, or alternatively the error returned:
	err := stickTemplate.Execute(documentName, v.Context.Response, v.Variables)
	if err != nil {
		return err
	}
	// No error, return nil!
	return nil
}