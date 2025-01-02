package database

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can'tfind product")
	ErrUserIdIsNotValid   = errors.New("this user id is not valid")
	ErrCantUpdateUser     = errors.New("cannot update user")
	ErrCantRemoveItemCart = errors.New("cannot remove item from cart")
	ErrCantGetItem        = errors.New("was not able to get item from the cart")
	ErrCantBuysCartItem   = errors.New("cannot update the purchase")
)

func AddProductToCart() {}

func RemoveCartItem() {}

func BuyItemFromCart() {}

func InstantBuyer() {}
