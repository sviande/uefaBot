
# UEFA slack bot

A simple tool that check update  on match and send you a slack message.

## Feedback

If you have any feedback, please reach me on [twitter](https://twitter.com/sviande/)

## Dependencies

You need a slack Webhook URL, you can follow [the slack guide to create an app](https://api.slack.com/messaging/webhooks)

  
## Installation

First install Go.

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN=$(pwd) GOPATH=$(mktemp -d) go get github.com/sviande/uefaBot
```
    
## Run

### Monitor all UEFA matches
```bash
  WEBHOOK_URL=https://hooks.slack.com/service/XXXXXXXX/XXXXXX uefaBot
```
### Monitor EURO 2021
```bash
  WEBHOOK_URL=https://hooks.slack.com/service/XXXXXXXX/XXXXXX uefaBot -competionId=3
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

  
