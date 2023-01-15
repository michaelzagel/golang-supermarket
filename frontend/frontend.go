package frontend

type RegisterInput struct {
	Email           string `json:"email"`
	DeliveryAddress string `json:"delivery_address"`
	Password        string `json:"password"`
}

type RegisterOutput struct {
	CustomerId int64 `json:"customer_id"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInOutput struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
}

type StartOrderInput struct {
	AccessToken string `json:"access_token"`
}

type StartOrderOutput struct {
	OrderId int64 `json:"order_id"`
}

type AddItemInput struct {
	AccessToken string  `json:"access_token"`
	OrderId     int64   `json:"order_id"`
	Name        string  `json:"name"`
	Quantity    float32 `json:"quantity"`
}

type AddItemOutput struct {
	ItemId int64 `json:"item_id"`
}

type FinishOrderInput struct {
	AccessToken string `json:"access_token"`
	OrderId     int64  `json:"order_id"`
	CreditCard  string `json:"credit_card"`
}

type FinishOrderOutput struct {
	WindowTo   string `json:"window_to"`
	WindowFrom string `json:"window_from"`
}
