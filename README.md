# iyzipay-go
This is an unofficial Iyzico Payment API client

### Information
This package is still in development. I welcome suggestions for changes that will bring it closer to compliance without overly complicating the code, or useful test cases to add to the test suite.

## 💈&nbsp; Installation

```
go get github.com/yahya077/iyzipay-go
```

## 🚀&nbsp; Usage

### Initialization
```
  //add your own credentials
  options := client.IyzicoClientOptions{
      BaseUrl:   iyzipay.SANDBOX_URL,
      APIKey:    "your-api-key",
      APISecret: "your-secret-key",
  }
  
  //this will be our client
  c := client.NewIyzicoClient(options)
```
### Payment
```
  //add what you need for request
  request := requestSchema.CreatePayment{...}
  
  //now lets get payment from service
  response, _ := iyzipay.New(c).Payment().Create(request)

```
### Card Storage
```
  //add what you need for request
  request := requestSchema.RetrieveCard{...}
  
  //now lets get payment from service
  response, _ := iyzipay.New(c).CardStorage().Get(request)

```



## 📫&nbsp; Have a question? Want to chat? Ran into a problem?

#### *Website [yahyahindioglu.com](https://yahyahindioglu.com)*

#### *[LinkedIn](https://www.linkedin.com/in/yahyahindioglu/) Account*

## License
This project is available under the [MIT](https://opensource.org/licenses/mit-license.php) license.
