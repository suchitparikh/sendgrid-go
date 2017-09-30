# Send a Single Email to a Single Recipient

The following code assumes you are storing the API key in an environment variable (recommended). 

This is the minimum code needed to send an message.

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
                To:     []mail.Email{
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

# Send Multiple Emails to Multiple Recipients

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
                From: mail.Email{"from@example.com", "Example Sender"},
        }
        
        numRecipients := 3
        
        tos := [numRecipients]mail.Email{
                mail.Email{"test1@example.com", "Test Recipient"},
                mail.Email{"test2@example.com", "Test Recipient"}.
                mail.Email{"test3@example.com", "Test Recipient"},
        }
        message.AddTos(tos)
        
        substitutions := []map[string]string
        sub1 := [numRecipients]map[string]string{
                "-name", "Alain",
                "-name", "Elmer",
                "-name", "Matt",
        }
        substitutions := append(substitutions, sub1)
        
        sub2 := [numRecipients]map[string]string{
                "-github-", "http://github.com/test1",
                "-github-", "http://github.com/thinkingserious",
                "-github-", "http://github.com/test2",
        }
        substitutions = append(substitutions, sub2)
        message.AddSubstitutions(substitutions)
        
        message.AddGlobalSubstitutions([]map[string]string{"-time-": "<Current Time>"})
        
        message.SetSubject("Hi -name-!")
        message.SetTextContent("Hello -name-, your github is -github-, email sent at -time-")
        message.SetHTMLContent("<strong>Hello -name-, your github is <a href=\"-github-\">here</a></strong> email sent at -time-")
        
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

# Kitchen Sink - an example with all settings used

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
                From: mail.Email{"from@example.com", "Example Sender"},
        }
        
        message.AddTo(mail.Email{"test1@example.com", "Test Recipient"})
        tos := []mail.Email{
                mail.Email{"test2@example.com", "Test Recipient"}.
                mail.Email{"test3@example.com", "Test Recipient"},
        }
        message.AddTos(tos)
        
        message.AddCC(mail.Email{"cc1@example.com", "Test Recipient"})
        ccs := []mail.Email{
                mail.Email{"cc2@example.com", "Test Recipient"}.
                mail.Email{"cc3@example.com", "Test Recipient"},
        }
        message.AddCCs(ccs)
        
        message.AddBCC(mail.Email{"bcc1@example.com", "Test Recipient"})
        bccs := []mail.Email{
                mail.Email{"bcc2@example.com", "Test Recipient"}.
                mail.Email{"bcc3@example.com", "Test Recipient"},
        }
        message.AddBCCs(bccs)
        
        message.AddHeader("X-Hdr1", "Test1")
        hdrs := []map[string]string{
                "X-Hdr2", "Test2",
                "X-Hdr3", "Test3",
        }
        message.AddHeaders(hdrs)
        
        message.AddSubstitution("%City1%", "Denver")
        message.AddSubstitution("%City2%", "Orange")
        substitutions := []map[string]string{
                "%name1%": "Name 1",
                "%name2%": "Name 2",
        }
        message.AddSubstitutions(substitutions)
        
        message.AddCustomArg("Message-Category", "Marketing")
        customArgs := []map[string]string{
                "Campaign-ID": "Mkt-123",
                "Location": "US-West",
        }
        message.AddCustomArgs(customArgs)
        
        message.SetSendAt(1461775051)
        
        // If you need to add more [Personalizations](https://sendgrid.com/docs/Classroom/Send/v3_Mail_Send/personalizations.html), here is an example of adding another Personalization by passing in a personalization index
        
        message.SetReplyTo(mail.Email{"replyto@example.com", "Reply To"})
        
        message.SetSubject("Sending Email is fun!")
        
        message.AddContent(mail.Content{
                "Type": "text/plain",
                "Value": "Text content",
        })
        message.AddContent(mail.Content{
                "Type": "text/html",
                "Value": "<strong> HTML content</strong>",
        })
        
        message.AddAttachment(mail.Attachment{
                "filename": "balance_001.pdf",
                "content": "base64 encoded string",
                "type": "application/pdf",
                "disposition": "attachment",
        })
        
        message.SetTemplateId("13b8f94f-bcae-4ec6-b752-70d6cb59f932")
        
        message.AddSection("%section1", "Substitution for Section 1 Tag")
        sections := []map[string]string{
                "%section2%", "Substitution for Section 2 Tag",
                "%section3%", "Substitution for Section 3 Tag",
        }
        message.AddSections(sections)
        
        message.AddCategory("customer")
        categories := []string{"new_account", "aws"}
        message.AddCategories(categories")
        
        message.AddGlobalCustomArg("campaign", "welcome")
        globalCustomArgs := []map[string]string{
                "sequence2", "2",
                "sequence3", "3",
        }
        message.AddGlobalCustomArgs(globalCustomArgs)
        
        asmGroupIds := []int{1, 4, 5}
        message.SetAsm(3, asmGroupIds)
        
        message.SetGlobalSendAt(1461775051)
        
        message.SetIpPoolName("23")
        
        // mail settings
        message.SetBccSetting(true, "test@example.com")
        message.SetBypassListManagement(true)
        message.SetFooterSetting(true, "Some Footer HTML", "Some Footer Text")
        message.SetSandBoxMode(true)
        message.SetSpamCheck(true, 1, "https://gotchya.example.com")
        
        // tracking settings
        message.SetClickTracking(true, false)
        message.SetOpenTracking(true, "Optional tag to replace with the open image in the body of the message")
        message.SetSubscriptionTracking(true,
                                        "HTML to insert into the text / html portion of the message",
                                        "text to insert into the text/plain portion of the message",
                                        "substitution tag")

        message.SetGoogleAnalytics(true, "some campaign", "some content", "some medium", "some source", "some term")
        
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

# Attachments

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
                To:          []mail.Email{"test@example.com", "Example Recipient"},
                From:        mail.Email{"from@example.com", "Example Sender"},
                Subject:     "Test Email Subject",
                TextContent: "Text Email Content",
                HTMLContent: "<strong>HTML Email Content<strong>",
        }
        
        message.AddAttachment(mail.Attachment{
                "filename": "yosemite.jpg",
                "content": "base64 encoded string",
                "type": "image/jpg",
                "disposition": "inline",
                "content_id": "image_1",
        })
        
        attachments := []mail.Attachment{
                mail.Attachment{
                        "filename": "el_capitan.jpg",
                        "content": "base64 encoded string 2",
                        "type": "image/jpg",
                        "disposition": "inline",
                        "content_id": "image_2",
                },
                mail.Attachment{
                        "filename": "sierra.jpg",
                        "content": "base64 encoded string 3",
                        "type": "image/jpg",
                        "disposition": "inline",
                        "content_id": "image_3",
                },
        }
        message.AddAttachments(attachments)

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

# Transactional Templates

The following code assumes you are storing the API key in an environment variable (recommended). 

For this example, we assume you have created a [transactional template](https://sendgrid.com/docs/User_Guide/Transactional_Templates/index.html). Following is the template content we used for testing.

Template ID (replace with your own):

```text
13b8f94f-bcae-4ec6-b752-70d6cb59f932
```

Email Subject:

```text
<%subject%>
```

Template Body:

```html
<html>
<head>
	<title></title>
</head>
<body>
Hello -name-,
<br /><br/>
I'm glad you are trying out the template feature!
<br /><br/>
<%body%>
<br /><br/>
I hope you are having a great day in -city- :)
<br /><br/>
</body>
</html>
```

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
        }
        
        message.SetTemplateId("13b8f94f-bcae-4ec6-b752-70d6cb59f932")
        message.AddSubstitution("-name-", "Example User")
        message.AddSubstitution("-city-", "Orange")
        message.AddSubstitution("%subject%", "A message with Template")

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
