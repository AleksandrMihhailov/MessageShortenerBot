Telegram message shortener bot
==============================

If message length grater than 140 characters, bot will post it to [telegra.ph](https://telegra.ph) and replay the message to the chat which containing link to [telegra.ph](https://telegra.ph)


Don't know how to run? Ha

    docker build -t message-shortener-bot .

and

    docker run -e TOKEN="<your bot token here>" message-shortener-bot
