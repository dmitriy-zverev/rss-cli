# RSS CLI Aggregator
This is an RSS feed aggregator in Go. It allows users to add RSS feeds from across the internet to be collected, store the collected posts in a PostgreSQL database, follow and unfollow RSS feeds that other users have added, view summaries of the aggregated posts in the terminal, with a link to the full post

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