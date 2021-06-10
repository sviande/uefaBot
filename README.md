
# UEFA slack bot

A simple tool that check update  on match and send you a slack message.

## Feedback

If you have any feedback, please reach out to us at https://twitter.com/sviande/

  
## Installation (Development)

First install Go.

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN=$(pwd) GOPATH=$(mktemp -d) go get github.com/sviande/uefaBot
```
    
## Run

--competitionId=3 Euro 2021
```bash
  WEBHOOK_URL=https://hooks.slack.com/service/XXXXXXXX/XXXXXX uefaBot --competitionId=3
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

  
