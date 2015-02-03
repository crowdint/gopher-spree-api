package repositories

import (
	"testing"
)

func TestQueryParser(t *testing.T) {
	lex("first_name_or_last_name_cont")
	t.Error(template)
}
