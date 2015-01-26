package domain

type OptionValue struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	Presentation           string `json:"presentation"`
	OptionTypeName         string `json:"option_type_name"`
	OptionTypeID           int64  `json:"option_type_id"`
	OptionTypePresentation string `json:"option_type_presentation"`
}
