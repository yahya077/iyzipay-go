# iyzipay-go
This is an unofficial Iyzico Payment API client

### Information
This package is still in development. I welcome suggestions for changes that will bring it closer to compliance without overly complicating the code, or useful test cases to add to the test suite.

## 🚀&nbsp; Usage

```
  options := client.IyzicoClientOptions{
      BaseUrl:   iyzipay.SANDBOX_URL,
      APIKey:    "your-api-key",
      APISecret: "your-secret-key",
  }
  c := client.NewIyzicoClient(options)
  
  // now lets get payment from service; CreatePayment
  
  response, _ := iyzipay.New(c).CreatePayment(paymentRequest)

```



## 📫&nbsp; Have a question? Want to chat? Ran into a problem?

#### *Website [yahyahindioglu.com](https://yahyahindioglu.com)*

#### *[LinkedIn](https://www.linkedin.com/in/yahyahindioglu/) Account*

## License
This project is available under the [MIT](https://opensource.org/licenses/mit-license.php) license.
