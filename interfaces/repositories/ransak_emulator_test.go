package repositories

import (
	"testing"
)

func TestQueryParser(t *testing.T) {
	ransak := NewRansakEmulator()

	expected := "first_name LIKE '%cone%' OR last_name LIKE '%cone%' "
	sql := ransak.ToSql("first_name_or_last_name_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "first_name LIKE '%cone%' AND last_name LIKE '%cone%' "
	sql = ransak.ToSql("first_name_and_last_name_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "first_name = 'cone' AND last_name = 'cone' "
	sql = ransak.ToSql("first_name_and_last_name_eq", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = 29 "
	sql = ransak.ToSql("age_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = 29 OR years = 29 "
	sql = ransak.ToSql("age_or_years_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}
}
