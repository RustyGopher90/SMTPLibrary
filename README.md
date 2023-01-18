# LibSMTP - An SMTP library for GO
### Who This is For:
If you are using a private smtp server and want to send an anonymous email this library is right for you.

### Note:
You must change the SMTPConfig struct in the test suite to pass all tests.

### Usage:
- This library allows you to send urgent and non urgent emails with GO.
- To use this library, you just need to pass an SMTPConfig struct to the SendEmail function.
- All fields must be filled in for the SMTPConfig struct or all functions will return an error.
- Clone this repo, copy the folder and place it inside C:\Program Files\Go\src or wherever you installed GO. After that just import the package like you would any other standard package.
```go
type SMTPConfig struct {
    From        string
    ReplyTo     string
    SMTPServer  string
    SMTPPort    int
    SMTPSubject string
    Recipients  []string
    Message     string
    Urgent      bool
}
```
 
### Functions:
<details>
  <summary>func SendEmail</summary>
 
  #### Send an email
 
  ##### Example
  ```go
  package main
 
  import (
    "libsmtp"
  )
 
  var server = "someAddress@someServer.com"
 
  func main() {
    var smtpconfig = SMTPConfig{
	From:        "reporting@someServer.com",
	ReplyTo:     "SomePersonalAddr@someServer.com",
	SMTPServer:  server,
	SMTPPort:    25,
	SMTPSubject: "Super Cool Libraries",
	Recipients:  []string{"SomePersonalAddr@someServer.com"},
	Message:     "Man this library is awesome!",
	Urgent:      false,
}
   
    libsmtp.SendEmail(&smtpconfig)
}
  ```
</details>