# Send a Single Email to a Single Recipient

The following code assumes you are storing the API key in an environment variable (recommended). 

This is the minimum code needed to send an email.

```go
package main

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	message := &mail.Message{
		To:          []mail.Email{"test@example.com", "Example Recipient"},
        	From:        mail.Email{"from@example.com", "Example Sender"},
        	Subject:     "Test Email Subject",
		TextContent: "Text Email Content",
		HTMLContent: "<strong>HTML Email Content<strong>",
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```

# Send a Single Email to Multiple Recipients

The following code assumes you are storing the API key in an environment variable (recommended). 

```go
package main

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	message := &mail.Message{
		To:          []mail.Email{
			         mail.Email{"test1@example.com", "Example Recipient1"},
				 mail.Email{"test2@example.com", "Example Recipient3"},
				 mail.Email{"test3@example.com", "Example Recipient3"},
			     },
		From:        mail.Email{"from@example.com", "Example Sender"},
		Subject:     "Test Email Subject",
		TextContent: "Text Email Content",
		HTMLContent: "<strong>HTML Email Content<strong>",
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```
