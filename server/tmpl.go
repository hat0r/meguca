/*
 Compiles HTML templates to be used for rendering pages during dynamic content
 insertion
*/

package server

import (
	"bytes"
	"fmt"
	"github.com/dchest/htmlmin"
	"html/template"
	"io/ioutil"
)

// templateStore stores the static part of HTML templates and the corresponding
// truncated MD5 hash of said template
type templateStore struct {
	Parts [][]byte
	Hash  string
}

// templateMap stores all available templates
type templateMap map[string]templateStore

// resources exports temolates and their hashes by language
var resources templateMap

// compileTemplates reads template HTML from disk, injects dynamic variables,
// hashes and stores them
func compileTemplates() {
	// Only one for now, but there will be more later
	raw := map[string]string{}
	for _, name := range []string{"index"} {
		file, err := ioutil.ReadFile("./tmpl/" + name + ".html")
		throw(err)
		raw[name] = string(file)
	}
	newResources := templateMap{}
	index, mobile := indexTemplate(raw["index"])
	newResources["index"] = index
	newResources["mobile"] = mobile
	resources = newResources
}

// clientFileHash is the combined, shortened MD5 hash of all client files
var clientFileHash string

type templateVars struct {
	Config               template.JS
	Navigation           template.HTML
	ConfigHash, MediaURL string
	IsMobile             bool
}

// indexTemplate compiles the HTML template for thread and board pages of the
// imageboard
func indexTemplate(raw string) (templateStore, templateStore) {
	vars := templateVars{ConfigHash: configHash}
	vars.Config = template.JS(marshalJSON(clientConfig))
	vars.Navigation = boardNavigation()
	tmpl, err := template.New("index").Parse(raw)
	throw(err)

	// Rigt now the desktop and mobile templates are almost identical. This will
	// change, when we get a dedicated mobile GUI.
	return buildIndexTemplate(tmpl, vars, false),
		buildIndexTemplate(tmpl, vars, true)
}

// boardNavigation renders interboard navigation we put in the top banner
func boardNavigation() template.HTML {
	html := `<b id="navTop">[`

	// Actual boards and "/all/" metaboard
	for i, board := range append(config.Boards.Enabled, "all") {
		if board == config.Boards.Staff {
			continue
		}
		html += boardLink(i > 0, board, "../"+board+"/")
	}

	// Add custom URLs to board navigation
	for _, link := range config.Boards.Psuedo {
		html += boardLink(true, link[0], link[1])
	}
	html += `]</b>`
	return template.HTML(html)
}

// Builds a a board link, for the interboard navigation bar
func boardLink(notFirst bool, name, url string) string {
	link := fmt.Sprintf(`<a href="%v">%v</a>`, url, name)
	if notFirst {
		link = " / " + link
	}
	return link
}

// buildIndexTemplate constructs the HTML template array, minifies and hashes it
func buildIndexTemplate(
	tmpl *template.Template,
	vars templateVars,
	isMobile bool,
) templateStore {
	vars.IsMobile = isMobile
	buffer := new(bytes.Buffer)
	throw(tmpl.Execute(buffer, vars))
	minified, err := htmlmin.Minify(buffer.Bytes(), &htmlmin.Options{
		MinifyScripts: true,
	})
	throw(err)
	return templateStore{
		bytes.Split(minified, []byte("<$$$>")),
		hashBuffer(minified),
	}
}
