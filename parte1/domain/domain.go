package domain

type ShoppingDate struct {
	Date string `param:"dayId"`
	Day  int    `query:"dia"`
}

type Register struct {
	ClientID int     `json:"clientId"`
	Nombre   string  `json:"nombre"`
	Compro   bool    `json:"compro"`
	TDC      string  `json:"tdc,omitempty"`
	Monto    float64 `json:"monto,omitempty"`
}

type Summary struct {
	Total         float64            `json:"total"`
	ComprasPorTDC map[string]float64 `json:"comprasPorTDC"`
	NoCompraron   int                `json:"nocompraron"`
	CompraMasAlta float64            `json:"compraMasAlta"`
}
