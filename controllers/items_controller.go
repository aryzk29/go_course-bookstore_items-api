package controllers

import (
	"fmt"
	"github.com/aryzk29/bookstore_oauth-go/oauth"
	"github.com/aryzk29/go_course-bookstore_items-api/domain/items"
	"github.com/aryzk29/go_course-bookstore_items-api/services"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
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
