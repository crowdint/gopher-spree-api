package slugged

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/nu7hatch/gouuid"
)

type Sluggeable interface {
	SlugCandidates() []interface{}
	SetSlug(string)
}

func GenerateSlugs(sluggeable Sluggeable, separator string) []string {
	slugs := []string{}
	uuidV4, err := uuid.NewV4()
	if err != nil {
		return slugs
	}

	for _, candidate := range sluggeable.SlugCandidates() {
		reflectValue := reflect.ValueOf(candidate)
		switch reflectValue.Kind() {
		case reflect.String:
			slugs = append(slugs, transformCandidate(candidate.(string), separator))
			break
		case reflect.Slice:
			slug := appendCandidateFromSlice(candidate.([]interface{}))
			slugs = append(slugs, transformCandidate(slug, separator))
			break
		}
	}

	if len(slugs) > 0 {
		slugs = append(slugs, slugs[0]+separator+uuidV4.String())
	}

	return slugs
}

func appendCandidateFromSlice(sliceCandidate []interface{}) string {
	slug := ""
	for i := 0; i < len(sliceCandidate); i++ {
		candidateItem := sliceCandidate[i]
		if candidateItem == "" {
			continue
		}
		slug += candidateItem.(string) + " "
	}
	return slug
}

func transformCandidate(candidate, separator string) string {
	candidate = strings.ToLower(candidate)
	candidate = strings.Trim(candidate, " ")
	reg, err := regexp.Compile("[.,]")
	if err != nil {
		return ""
	}

	candidate = reg.ReplaceAllString(candidate, "")
	candidate = strings.Replace(candidate, " ", separator, -1)

	return candidate
}
