package repositories

import (
	"testing"
)

func TestQueryParser(t *testing.T) {
	ransak := NewRansakEmulator()

	//cont / or / and
	expected := "first_name LIKE '%cone%' OR last_name LIKE '%cone%'"
	sql := ransak.ToSql("first_name_or_last_name_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "first_name LIKE '%cone%' AND last_name LIKE '%cone%'"
	sql = ransak.ToSql("first_name_and_last_name_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_cont / or / and
	expected = "first_name NOT LIKE '%cone%' OR last_name NOT LIKE '%cone%'"
	sql = ransak.ToSql("first_name_or_last_name_not_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//matches / or
	expected = "first_name LIKE 'cone' OR last_name LIKE 'cone'"
	sql = ransak.ToSql("first_name_or_last_name_matches", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//does_not_match / or
	expected = "first_name NOT LIKE 'cone' OR last_name NOT LIKE 'cone'"
	sql = ransak.ToSql("first_name_or_last_name_does_not_match", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//lt
	expected = "age < 29"
	sql = ransak.ToSql("age_lt", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//gt
	expected = "age > 29"
	sql = ransak.ToSql("age_gt", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//eq / or / and
	expected = "first_name = 'cone' AND last_name = 'cone'"
	sql = ransak.ToSql("first_name_and_last_name_eq", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = 29"
	sql = ransak.ToSql("age_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = 29 OR years = 29"
	sql = ransak.ToSql("age_or_years_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_eq / or
	expected = "age <> 29 OR years <> 29"
	sql = ransak.ToSql("age_or_years_not_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "name <> 'cone' OR last_name <> 'cone'"
	sql = ransak.ToSql("name_or_last_name_not_eq", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//Has word "not" but is not "not_equal" nor "not_in"
	//so it must be part of the field's name
	expected = "field_not_operator = 29"
	sql = ransak.ToSql("field_not_operator_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}
}
