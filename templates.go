package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/apodhrad/ubi-init/log"
)

//go:embed templates
var templates embed.FS

func copyTemplate(template string, dir string) error {
	log.Info("Initiliaze UBI in '%v'", dir)

	info, ok := exists(dir)
	if !ok {
		log.Info("Dir '%v' doesn't exist", dir)
		log.Info("Create dir '%v'", dir)
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
		log.Info("Dir '%v' created", dir)
	} else if info.IsDir() {
		log.Info("Dir '%v' already exists", dir)
	} else {
		return errors.New(fmt.Sprintf("Cannot initialize UBI in '%v' as it is file", dir))
	}

	log.Info("Use template '%v'", template)
	templateRoot := "templates/" + template
	var hierarchyErr error
	fs.WalkDir(templates, templateRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		templatePath := templatePath(templateRoot, path, d)
		if templatePath == "" {
			return nil
		}
		log.Info("  %v", templatePath)
		target := filepath.Join(dir, templatePath)
		if _, ok := exists(target); ok {
			msg := fmt.Sprintf("Cannot generate '%v' as it already exists", target)
			hierarchyErr = errors.Join(hierarchyErr, errors.New(msg))
		}
		return nil
	})
	if hierarchyErr != nil {
		return hierarchyErr
	}

	log.Info("Generate files and dirs")
	return fs.WalkDir(templates, "templates/"+template, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		templatePath := templatePath(templateRoot, path, d)
		if templatePath == "" {
			return nil
		}
		log.Info("  %v", templatePath)
		target := filepath.Join(dir, templatePath)
		if d.IsDir() {
			return os.Mkdir(target, 0755)
		}
		data, err := fs.ReadFile(templates, path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	})
}
