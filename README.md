## **🔥 Introduction**
This documentation provides an overview of the APIs available for a shopping app. These APIs allow users and sellers to interact with the shopping platform by performing actions such as user registration, login, product management, cart management, and order management etc.

## **💥 APIs**

```go
r.POST("/api/user/login", userauth.Login)
r.POST("/api/user/register", userauth.Register)
r.POST("/api/seller/login", sellerauth.Login)
r.POST("/api/seller/register", sellerauth.Register)
r.POST("/api/seller/addproduct", sellerhandler.AddProduct)
r.GET("/api/seller/getproduct", sellerhandler.Get_Product)
r.GET("/api/getallproducts", userhandler.Search_Product)
r.DELETE("/api/seller/delete", sellerhandler.RemoveProduct)
r.PUT("/api/seller/update", sellerhandler.Update_Product)
r.POST("/api/user/addcart", userhandler.Add_Item_In_Cart)
r.DELETE("/api/user/removecart", userhandler.Remove_Item_In_Cart)
r.GET("/api/user/viewcart", userhandler.View_Item_In_Cart)
r.POST("/api/user/placeorder", userhandler.Place_Order)
r.DELETE("/api/user/cancelorder", userhandler.Cancel_Order)
r.GET("/api/user/vieworder", userhandler.View_Purchased_Items)
```
## **🛠️ Local Development** :

1. Open your terminal and then type
    ```shell
    $ git clone https://github.com/Pushkarm029/Amazon-Shopping-Backend
    ```
2. cd into the folder
    ```shell
    $ cd Amazon-Shopping-Backend/
    ```
3. install all go dependencies
   ```shell
   $ go get ./...
   ```
4. Start Database on :5432
   ```shell
   docker compose up
   ```
5. Start the server on :8080
    ```shell
    go run main.go
    ```
6. Testing API's
   ```shell
   Add postman_examples/Amazon Backend.postman_collection.json in Postman Collection
   ```

**Note : You need to restart backend server after every change in any .go file.**

## **❤️ Learnings** :
- Implementation of GORM.
- Basics of PostgreSQL.
- use of Postman Collections.

## **⛑️ Maintenance** :

Feel free to open issue to add a feature request or report any BUG. It will be appreciated from the depth of my heart❤️.
