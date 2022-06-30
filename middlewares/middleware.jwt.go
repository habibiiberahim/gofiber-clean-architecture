package middlewares

// //JWTProtected func for specify route group with JWT authentication.
// func JWTProtected() func(*fiber.Ctx) error {
// 	// Create config for JWT authentication middleware.
// 	jwtwareConfig := jwtware.Config{
// 		SigningKey:     []byte(configs.AppCfg().JWTSecretKey),
// 		ContextKey:     "user", // used in private route
// 		ErrorHandler:   jwtError,
// 		SuccessHandler: verifyTokenExpiration,
// 	}

// 	return jwtware.New(jwtwareConfig)
// }

// func verifyTokenExpiration(c *fiber.Ctx) error {
// 	user := c.Locals("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	expires := int64(claims["exp"].(float64))
// 	if time.Now().Unix() > expires {
// 		return jwtError(c, errors.New("token expired"))
// 	}

// 	return c.Next()
// }

// func jwtError(c *fiber.Ctx, err error) error {
// 	// Return status 401 and failed authentication error.
// 	if err.Error() == "Missing or malformed JWT" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"msg": err.Error(),
// 		})
// 	}

// 	// Return status 401 and failed authentication error.
// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"msg": err.Error(),
// 	})
// }
