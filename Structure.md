A typical architecture of (fiber) web framework

**What is the usage of 'Group()' function in fiber web framwork**

By using Group(), you can logically group routes together and apply middleware to multiple routes at once, reducing code duplication and providing a cleaner structure to your application.


```
/app
    /internal
        /aws
        /database
        /server
            server.go
        /shop
            /controller
                product.go (3) - NewProductAdminHandlers
            /message_broker
            /mocks
            /repository
                product.go (1) - CRUD functions
            /routes
                routes.go (4) - MapProductAdminRoutes
            /service
                product.go (2) - NewProductServiceImpl
    /message
    /migrations
    /pkg
        /models
            product.go (0)
    /scripts
```

```
# Step 1 -- UPDATE models
pkg/models/product.go

# Step 2 -- UPDATE CRUD function ORM
internal/shop/repository/product.go

# Step 3 -- UPDATE up migration file
migrations/000016_GS-6761.up.sql

```