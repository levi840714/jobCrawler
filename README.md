# jobCrawler
Golang job crawler form 104、CakeResume website

## Docker directly deploy environment

Please go to the **.env.example** file to set the config data 

**BOT_TOKEN** is your telegram bot token

**CHANNEL_ID** is the telegram channel to push the job message

Then execute commands

<pre>
# {keyword} Enter the search keywords you want

docker-compose build --build-arg keyword={<b>keyword</b>} service

docker-compose up -d
</pre>
