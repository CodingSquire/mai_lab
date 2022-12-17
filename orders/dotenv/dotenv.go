package dotenv

import (
	"bufio"
	"os"
	"strings"
)

type Dotenv struct {
	_params map[string]string
}

func Config() *Dotenv {
	ret := make(map[string]string)
	dotenv, err := os.Open(".env");

	if !os.IsNotExist(err) {
		scanner := bufio.NewScanner(dotenv);

		for scanner.Scan() {
			keyValue := strings.Split(scanner.Text(), "=")
			ret[strings.ToLower(keyValue[0])] = keyValue[1]
			os.Setenv(keyValue[0], keyValue[1])
		}
	}

	return &Dotenv{_params: ret}
}

func (d *Dotenv) Get(key string) string {
	return d._params[strings.ToLower(key)]
}
