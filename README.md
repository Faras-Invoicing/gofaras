# gofaras
go package to integrate with Faras API using go language

### Example ### 
```
 customer := gofaras.Customer{
    Name: "Ahmad,
    PhoneNumber : 96655555555,
    ID:    "122",
    Email: "ahmad@example.com",
 }
 product := gofaras.Product{
    ProductName: "Coffee Mug",
    Quantity: 1,
    Price: 20.5,
 }
 products := make([]gofaras.Product, 0)
 products = append(products, product)
 inv := gofaras.Invoice{
    Key : "your api key",
    Products: products,
    Customer: customer,
    TestMode: 0, // 1 is test, 0 is production
  }
 isErr, Errmsg, url, pdfurl := gofaras.NewInvoice(inv)
```
