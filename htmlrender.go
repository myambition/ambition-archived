package ambition

import (
	"html/template"
	"io/ioutil"
	"os"
)

func ReadHtml() {
	dat, err := ioutil.ReadFile("./html/index.html")
	check(err)
	t, err := template.New("temp").Parse(string(dat))
	actions, err := database.GetActions()
	t.Execute(os.Stdout, actions)
}
