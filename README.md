# FetchmojiAI

<img
    src="https://avatars.slack-edge.com/2023-12-08/6311721858790_698d80808b06dfb298e8_512.jpg"
    alt="FetchmojiAI Logo"
    width="256"
    height="256"
/>

A bot to help generate emojis through AI to use in Slack chats. The bot uses OpenAI's image generation API to generate images given a description. This is built for the Fetch Frostbyte Hackathon and the Hack and Splash Hackathon.

## Using the app

A channel needs to have the [FetchmojiAI](https://fetchrewards.slack.com/apps/A0699CL6AVB-fetchmojiai) app added.

1. Send a prompt with `@FetchmojiAI`

```text
@FetchmojiAI a cat playing with a dog
```

2. After a few seconds, your message will have a thread reply with your imagined emoji

## Running the service

1. Install dependencies

```bash
make install
```

2. Install go modules

```bash
go mod tidy
```

3. Copy the contents of the `.env.example` file into a new `.env` file

```bash
cp .env.example .env
```

> Follow the links to obtain in the environment variables

4. Run `ngrok` to create a remote request URL from `localhost:8080`

```bash
ngrok http http://localhost:8080
```

> Follow installation instructions [here](https://dashboard.ngrok.com/get-started/setup/macos).

5. Copy the `ngrok` "Forwarding" address and paste it as the "Request URL" on the Slack App [Event Subscriptions]((https://a.slack-edge.com/80588/img/api/event_url_verification.png)) dashbaord

<img
    src="https://a.slack-edge.com/80588/img/api/event_url_verification.png"
    alt="FetchmojiAI Logo"
    width="763"
    height="183"
/>

6. Run the service

```bash
make run
```
