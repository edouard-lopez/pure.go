package go_version

import "runtime"

var getGoVersion = runtime.Version // so we can mock it

func Get() string {
	return getGoVersion()
}
