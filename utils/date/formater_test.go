package date

import (
	"testing"

	"github.com/ehlxr/go-utils/log"
)

func TestDateFormater(t *testing.T) {
	log.Infof("now tims is %s", Formater("yyyy-MM-dd HH:mm:ss"))
}
