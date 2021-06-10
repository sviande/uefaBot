
# UEFA slack bot
[![Go](https://github.com/sviande/uefaBot/actions/workflows/go.yml/badge.svg)](https://github.com/sviande/uefaBot/actions/workflows/go.yml)

A simple tool that check update  on match and send you a slack message.

## Feedback

If you have any feedback, please reach me on [twitter](https://twitter.com/sviande/)

## Dependencies

You need a slack Webhook URL, you can follow [the slack guide to create an app](https://api.slack.com/messaging/webhooks)

## Installation (standalone)

Fetch you binary from [the latest release](https://github.com/sviande/uefaBot/releases/latest)

## Installation (with go)

First install [Go](https://golang.org/doc/install).

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

## Preview
![Screen Shot 2021-06-10 at 13 24 28](https://user-images.githubusercontent.com/1902504/121517369-94751580-c9ef-11eb-9521-a57102458938.png)
![Screen Shot 2021-06-10 at 13 25 27](https://user-images.githubusercontent.com/1902504/121517437-a8b91280-c9ef-11eb-949f-8d49dcdfdda4.png)

## License

[MIT](https://choosealicense.com/licenses/mit/)
