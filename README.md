# Taberneiro
Travis-CI: [![Travis-CI Build Status](https://travis-ci.org/betorvs/taberneiro.svg?branch=master)](https://travis-ci.org/betorvs/taberneiro)


This is a slack bot written in Golang with Helm charts to deploy it.

Taberneiro is a portuguese word that means a tavern's salesman. 

## How it works

First you need to create a app inside Slack. Deploy it inside some workspace and get this credentials. Make as a bot user and configure a interactive messages too.

### Go Installation

Install go

Install dependencies [dep](https://golang.github.io/dep/docs/installation.html)

### Configure

```sh
dep ensure -update -v
```

### Run

```sh
go build
./taberneiro
```

## Environment Variables

* SLACK_TOKEN: Starts with xoxb-.
* SLACK_APP_NAME: One name
* SLACK_APP_ID: ID from Slack App.
* SLACK_DIRECT_MESSAGE: User ID who will receive all orders
* SLACK_MAGIC_WORD: How can we call the slack bot
* SLACK_MENU_TEXT: Menu Text
* SLACK_HELLO_MESSAGE: Hello message when bot connect to slack
* SLACK_IMAGE_TEXT: Text with menu image.
* SLACK_MENU_LINK_IMAGE: menu image it self.


## References

Thanks for all posts and shared code:

https://github.com/nlopes/slack/

https://medium.com/mercari-engineering/writing-an-interactive-message-bot-for-slack-in-golang-6337d04f36b9 and https://github.com/tcnksm/go-slack-interactive

https://github.com/sebito91/nhlslackbot

https://blog.zikes.me/post/how-i-ruined-office-productivity-with-a-slack-bot/
