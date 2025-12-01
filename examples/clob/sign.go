package main

import (
	"fmt"

	"github.com/lixvyang/polymarket-sdk-go/auth"
)

func main() {

	hmacSig := auth.BuildPolyHmacSignature("g2GweqLAsNoTXJ_jUeRjG63V_esnXTpV50YEbt6_DIQ=", 1764562634, "GET", "/auth/api-keys", nil)
	fmt.Printf("HMAC signature: %s\n", hmacSig)

}
