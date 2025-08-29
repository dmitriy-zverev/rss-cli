# RSS CLI - The Gator Aggregator
This is an RSS feed aggregator in Go. It allows users to add RSS feeds from across the internet to be collected, store the collected posts in a PostgreSQL database, follow and unfollow RSS feeds that other users have added, view summaries of the aggregated posts in the terminal, with a link to the full post

## Requirments
* Postgres 15.0 (or newer)
* Golang 12.4.5 (or newer)

## How to install
1. Run the command:
```
go install github.com/dmitriy-zverev/rss-cli@latest
mv $(which rss-cli) ~/go/bin/gator
```
2. Start the Postgres service with: `brew services start postgresql@15` if on macOS.
3. Change the config in your HOME directory named `.gatorconfig.json` and instead of `username` use your computer username.
4. Use the tool with `gator <command>`

## How it works
`login` - logins you with provided login if registerd earlier

```
gator login dmitriy
```
```
Hello dmitriy! Welcome back
```

`register` - registers a new user with provided login

```
gator register dmitriy
```
```
Nice choice! The username dmitriy was reserved for you
```

`addfeed` - adds new feed and auto follows it

```
gator addfeed "HN Newest" https://news.ycombinator.com/newest
```
```
dmitriy - HN Newest
Feed 'HN Newest' (https://news.ycombinator.com/newest) was successfully added!
```

`feeds` - lists all of the feeds

```
gator feeds
```
```
Hacker News RSS (https://hnrss.org/newest) - dmitiry
TechCrunch (https://techcrunch.com/feed/) - another-user
HN Newest (https://news.ycombinator.com/newest) - dmitriy
```

`follow` - follows the given feed for the current user

```
gator follow https://hnrss.org/newest
```
```
dmitriy is following Hacker News RSS
```

`unfollow` - unfollows the given feed for the current user

```
gator unfollow https://hnrss.org/newest
```
```
dmitriy has unfollowed https://hnrss.org/newest
```

`following` - lists all of the feed names you have been following

```
gator following
```
```
dmitriy's feeds:

    - TechCrunch
    - Hacker News
```

`users` - lists all of the registered users

```
gator users
```
```
dmitriy (current)
another-user 1
another-user 2
```

`reset` - deletes all users (be cautious using this command)

```
gator reset
```
```
The database has been reset
```

## Found a Bug
Text me: https://t.me/dmitry_zverev
