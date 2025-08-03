
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
* users -returns list of users and identifies which user you are logged in as
* addfeed <name> <url> -adds a rss feed to your database
* feeds -returns list of feeds in your database
* follow <url> -add a feed from your database to your logged in users feed following list. If url does not exist in database, add feed first using addfeed command
* unfollow <url> -unfollows a feed