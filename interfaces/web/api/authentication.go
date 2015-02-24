package api

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
)

var (
	readRoutesPattern = []string{
		`^` + namespace() + `/api/products(/?)$`,
		`^` + namespace() + `/api/products/\d+$`,
	}
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authRequired := spree.IsAuthenticationRequired()

		// GET + authentication (false) + readAction => next
		// POST + authentication (false) + token (spreeToken) => next
		// authentication (true) + token (spreeToken || orderToken) => next
		if isReadAction(c.Request) && !authRequired {
			nextHandler(c, &models.User{})
			return
		} else {
			if err := verifySpreeTokenAccess(c, authRequired); err != nil {
				return
			}
		}
	}
}

func verifySpreeTokenAccess(c *gin.Context, authRequired bool) error {
	var err error
	user := &models.User{}
	isGuestUser := false
	spreeToken := getSpreeToken(c)
	dbRepo := repositories.NewDatabaseRepository()

	// If spreeToken is empty, check if orderToken is set and action is orders show
	if spreeToken == "" {
		if isGuestUser, err = verifyOrderTokenAndAction(c, dbRepo, authRequired); err != nil {
			return err
		}
	}

	// if user is not guest then find him by spree token
	if !isGuestUser {
		if err = findUserBySpreeApiKey(c, dbRepo, user, spreeToken); err != nil {
			return err
		}
	} else {
		user.Id = -1
	}

	c.Set(SPREE_TOKEN, spreeToken)
	nextHandler(c, user)
	return nil
}

func verifyOrderTokenAndAction(c *gin.Context, dbRepo *repositories.DbRepository, authRequired bool) (bool, error) {
	if isOrdersShowAction(c.Request.URL.Path) {
		if err := verifyOrderTokenAccess(c, dbRepo, authRequired); err != nil {
			return false, err
		}

		return true, nil
	} else {
		unauthorizedAuthRequiredMsg(c, authRequired)
		return false, errors.New("Unathorized to perform that action.")
	}

}

func verifyOrderTokenAccess(c *gin.Context, dbRepo *repositories.DbRepository, authRequired bool) error {
	// Get order token
	orderToken := getOrderToken(c)

	// Return if order token is not provided
	if orderToken == "" {
		unauthorizedAuthRequiredMsg(c, authRequired)
		return errors.New("Order Token is not present")
	}

	// Find the order by guest token (order token)
	order := &models.Order{}
	err := dbRepo.FindBy(order, nil, map[string]interface{}{"guest_token": orderToken})
	if err != nil {
		unauthorized(c, "You are not authorized to perform that action.")
		return err
	}

	// Get order number (from path) and verify if it's equal to the order's number (from guest token)
	orderNumber := getOrderNumber(c.Request.URL.Path)
	if order.Number != orderNumber {
		unauthorized(c, "You are not authorized to perform that action.")
		return errors.New("The order number param is not equal to order's number")
	}

	c.Set("Order", order)
	return nil
}

func unauthorizedAuthRequiredMsg(c *gin.Context, authRequired bool) {
	if authRequired {
		unauthorized(c, "You must specify an API key.")
	} else {
		unauthorized(c, "You are not authorized to perform that action.")
	}
}

func nextHandler(c *gin.Context, user *models.User) {
	c.Set("CurrentUser", user)
	c.Next()
}

func findUserBySpreeApiKey(c *gin.Context, dbRepo *repositories.DbRepository, user *models.User, spreeToken string) error {
	err := dbRepo.FindBy(user, nil, map[string]interface{}{"spree_api_key": spreeToken})

	if err != nil {
		unauthorized(c, "Invalid API key ("+spreeToken+") specified.")
		return err
	}

	dbRepo.UserRoles(user)
	return nil
}

func isReadAction(req *http.Request) bool {
	readAction := false
	path := req.URL.Path

	for _, pattern := range readRoutesPattern {
		if readAction, _ = regexp.MatchString(pattern, path); readAction {
			// readAction is true when => [Country, OptionType, OptionValue, Product, ProductProperty, Property, State, Taxon, Taxonomy, Variant, Zone]
			break
		}
	}

	return req.Method == "GET" && readAction
}

func isOrdersShowAction(path string) bool {
	match, _ := regexp.MatchString(`^`+namespace()+`/api/orders/\w+$`, path)
	return match
}

func getOrderNumber(path string) string {
	// path => /api/orders/R12293
	pathArray := strings.Split(path, "/") // => [api, orders, R12293]
	return pathArray[len(pathArray)-1]    // => R12293
}
