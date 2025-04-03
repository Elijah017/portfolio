package router

import (
	"errors"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
)

var ErrInvalidPath = errors.New("not a valid path for a route")

var ComponentDir string = "./components"

type Router struct {
	Prefix     string
	Routes     []string
	Parent     *Router
	Children   []*Router
	Components *[]string
}

func loadComponents() (*[]string, error) {
	var components []string

	err := filepath.Walk(ComponentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".gotmpl" {
			components = append(components, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &components, nil
}

func isDir(path string) (bool, error) {
	if matched, err := regexp.Match(`^\/([a-zA-Z0-9\.\-]+\/)*$`, []byte(path)); !matched || err != nil {
		if err == nil {
			err = ErrInvalidPath
		}
		return matched, err
	}
	return true, nil
}

func isFile(path string) (bool, error) {
	if matched, err := regexp.Match(`^\/(([a-zA-Z0-9\.\-]+\/)*([a-zA-Z0-9\.\-]+\.[a-zA-Z0-9\-]+))$`, []byte(path)); !matched || err != nil {
		if err == nil {
			err = ErrInvalidPath
		}
		return matched, err
	}
	return true, nil
}

// Create a new Router object
//
// Parameters:
//
//	prefix string: the prefix on the route
//
// Returns:
//
//	*Router: Initialised Router or nil on error
//	error: The error that occured if one occured
//	  - ErrInvalidPath
func NewRouter(prefix string) (*Router, error) {
	if _, err := isDir(prefix); err != nil {
		return nil, err
	}

	components, err := loadComponents()
	if err != nil {
		return nil, err
	}

	return &Router{
		Prefix:     prefix,
		Components: components,
	}, nil
}

func (rtr *Router) NewSubRouter(prefix string) (*Router, error) {
	if _, err := isDir(prefix); err != nil {
		return nil, err
	}

	subrouter := &Router{
		Prefix:     rtr.Prefix + prefix[1:],
		Parent:     rtr,
		Components: rtr.Components,
	}

	rtr.Children = append(rtr.Children, subrouter)

	return subrouter, nil
}

func (rtr *Router) AddRoute(path string) error {
	if _, err := isFile(path); err != nil {
		return err
	}

	root := rtr
	for root.Parent != nil {
		root = root.Parent
	}

	root.Routes = append(root.Routes, rtr.Prefix+path[1:])

	return nil
}

func openBuildFile(path string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(path), 0744)
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (rtr *Router) GeneratePages(genDir string, tmplDir string) error {
	root := rtr
	for root.Parent != nil {
		root = root.Parent
	}

	for _, route := range root.Routes {
		tmpl, err := template.New("base").
			Funcs(template.FuncMap{}).
			ParseFiles(*root.Components...)
		if err != nil {
			return err
		}

		tmpl.ParseFiles(tmplDir + route)
		file, err := openBuildFile(genDir + route)
		if err != nil {
			return err
		}

		tmpl.Execute(file, map[string]any{
			"Title": "A Generic Title",
		})
	}

	return nil
}
