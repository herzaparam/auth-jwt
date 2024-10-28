package main

import (
	"github.com/herzaparam/auth-jwt/routes"
)

func main() {
	routes.Run(":5000")
}
