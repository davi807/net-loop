package browser

import "io/ioutil"

func bindLoader() {
	ui.Bind("load", func(name string) string {
		res, err := ioutil.ReadFile("./assets" + name)
		if err != nil {
			return err.Error()
		}

		return string(res)
	})
}

func loadStartContent() {
	ui.Eval("load('/css/uikit.min.css').then(res => insert('uikit.min.css', res))")
	ui.Eval("load('/css/main.css').then(res => insert('main.css', res))")

	ui.Eval("load('/js/vue.min.js').then(res => insert('vue.min.js', res))")
	ui.Eval("load('/js/uikit.min.js').then(res => insert('uikit.min.js', res))")
	ui.Eval("load('/js/main.js').then(res => insert('main.js', res))")
}
