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
