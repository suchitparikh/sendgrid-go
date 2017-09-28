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
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	message := &mail.SGMessage{
		To:          []mail.Email{"Example Recipient", "test@example.com"},
        From:        mail.Email{"Example Sender", "from@example.com"},
        Subject:     "Test Email Subject",
		TextContent: "Text Email Content",
		HTMLContent: "<strong>HTML Email Content<strong>",
	}

	response, err := client.SendMail(message)
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
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	message := &mail.SGMessage{
		To:          []mail.Email{"Example Recipient", "test@example.com"},
		From:        mail.Email{"Example Sender", "from@example.com"},
		Subject:     "Test Email Subject",
		TextContent: "Text Email Content",
		HTMLContent: "<strong>HTML Email Content<strong>",
	}

	// Add additional To addresses
	message.AddTos([]mail.Email{
		mail.Email{"Example Recipient2", "test@example.com"},
		mail.Email{"Example Recipient3", "test@example.com"},
		mail.Email{"Example Recipient4", "test@example.com"},
	})

	response, err := client.SendMail(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```
