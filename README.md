# Packages:

- Echo framework: https://github.com/labstack/echo/v4
- Validate struct: https://github.com/go-playground/validator
- Connect to database: https://github.com/lib/pq
- Migrate: https://github.com/rubenv/sql-migrate
- Generate uuid: https://github.com/google/uuid
- JWT: https://github.com/golang-jwt/jwt

# Generate JWT key with nodejs:

```javascript
require('crypto').randomBytes(64).toString('hex');
```

# Custom validator

- To use custom validator, We need to add struct_validator.go and edit tag of struct model like SignIn model in req_signin.go (validate:"pwd"). And update main.go as well as use c.Validate(model) instead of validate.Struct
