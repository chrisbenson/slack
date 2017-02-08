# Slack
##### This is a tiny convenience library, in the form of a Go package, that makes it trivial to send a message to a channel in your Slack account.  To use it, an incoming webhook must be established in the Slack account, which provides the URL that an HTTPS POST request is made to.

##### Import the package:
```import "github.com/chrisbenson/slack"```

##### Code example:
```
m := slack.Message{Channel: "#mychannel",
                   Username: "My Bot Name",
                   Text: "Testing, testing, 1 2 3.",
                   IconEmoji: ":myicon:"}
err := slack.Send(`https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX`, m)
if err != nil {
   return err
}
return nil
```
