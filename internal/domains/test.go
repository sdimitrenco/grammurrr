package domains

import "github.com/sdimitrenco/grammurrr/internal/infrastructure/logging"


func Test(log *logging.Logger) {
	log.Error("bad something")
}