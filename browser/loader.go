package browser

import "io/ioutil"

func loader(name string) string {
	res, err := ioutil.ReadFile("./assets" + name)
	if err != nil {
		return err.Error()
	}

	return string(res)
}
