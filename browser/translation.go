package browser

import (
	"io/ioutil"
)

const defaultLang = "en"

var translation string

// SetupTranslation send translation string to browser
func SetupTranslation() {
	var lang string
	var notConf = false

	if lang = config["lang"]; len(lang) == 0 {
		lang = defaultLang
		notConf = true
	}

	data, _ := ioutil.ReadFile("./assets/translations/" + lang + ".json")
	ui.Eval("var i18n =" + "`" + string(data) + "`;")

	if notConf {
		go setConfig("lang", lang)
	}
}
