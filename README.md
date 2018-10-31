# Bike Checker

This tool is setup to check regularly to see if a bike is in stock on a website. 

You need to setup some environment variables on your system in order for it to run correctly.
Please set the following:
`SCRAPER_EMAIL_ADDRESS`,
`SCRAPER_PASSWORD`,
`DESTINATION_ADDRESS`,

By default it will check every 5 seconds, however by passing the flag `--checkTime (number)` when running the program, you can decide what is reasonable to you.


## Future Plans
Extend the checker to be able to take in small snippets of html so that it can look up anything.