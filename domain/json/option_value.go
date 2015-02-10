package json

type OptionValue struct {
	Id                     int64  `json:"id"`
	Name                   string `json:"name"`
	Presentation           string `json:"presentation"`
	OptionTypeName         string `json:"option_type_name"`
	OptionTypeId           int64  `json:"option_type_id"`
	OptionTypePresentation string `json:"option_type_presentation"`
}
