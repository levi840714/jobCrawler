# jobCrawler
Golang job crawler from 104、CakeResume website

## Docker directly deploy environment

Please copy the `.env.example` to `.env` file and set the config data 

> `BOT_TOKEN` is your telegram bot token

> `CHANNEL_ID` is the telegram channel to push the job message

Then execute commands

<pre>
# {keyword} Enter the search keywords you want
# If you have multiple keywords, separate them with ","  Example: keyword=golang,python

docker-compose build --build-arg keyword={<b>keyword</b>} service

docker-compose up -d
</pre>
