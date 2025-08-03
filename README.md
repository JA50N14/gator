
## Installation
* Will need Postgres and Go installed to run this program.
* To install this gator CLI program, run 'go install github.com/JA50N14/gator'
* This install command will download, build, and install the binary.

## Run Program
1. Create a .gatorconfig.json file in your home directory
2. Inside the .gatorconfig.json file enter:
"db_url": "YourPostgresConnectionString?sslmode=disable"
3. Run this command in your CLI: gator register YourUsername
4. The 'register' command will create your account and log you in. Enter any other available commands


## Command List
* addfeed <name> <url> -adds a rss feed to your database
* feeds -returns list of feeds in your database
* follow <url> -adds a feed from your database to current user. If url does not exist in database, add the feed first using addfeed command
* unfollow <url> -unfollows a feed from current user
* following -return list of all feeds current user is following
* agg <time_between_requests> -run in a separate CLI window. Will pull all posts associated with each feed, and insert them into your database.
* browse <int> -returns a list of posts current user follows, ordered by newest to oldest. Will return the number of posts specified in the <int> argument
* users -returns list of users and identifies which user you are logged in as
* login <username> -Logs in a username
* reset -removes all users from the database. Used for development of this CLI tool for testing purposes

