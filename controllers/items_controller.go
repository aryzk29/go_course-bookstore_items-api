package controllers

import (
	"fmt"
	"github.com/aryzk29/go_course-bookstore_items-api/domain/items"
	"github.com/aryzk29/go_course-bookstore_items-api/services"
	"github.com/federicoleon/bookstore_oauth-go/oauth"
	"net/http"
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		return
	}

	fmt.Println(result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
