# RSS Feed Aggregator

## About This Project
Gator is a Go-based RSS Feed Aggregator project that allows you to collect and manage RSS feeds efficiently. This repository contains the source code and documentation for the Gator project.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)

## Installation
Pre-Requisite
- go v1.23 or above
- Postgresql server v15
- Goose Migration ([installation](https://github.com/pressly/goose#install))

\
Clone the repository
```bash
git clone https://github.com/hyperneutr0n/gator.git
```

Navigate to the project directory
```bash
cd gator
```

Download Go dependencies
```bash
go mod download
```

Create .gatorconfig.json config file in your home directory
```bash
conn_string="postgres://username:password@localhost:5432/dbname" && 
echo "{\"db_url\":\"$conn_string\"}" > ~/.gatorconfig.json
```

Install the Project
```bash
go install .
```

Run migrations
```bash
goose "postgres://username:password@localhost:5432/dbname" up
```

## Usage
To use this RSS Feed Aggregator:

First register and login
```bash
gator register <username>
gator login <username>
```

See all registered users
```bash
gator users
```

Delete all users
```bash
gator reset
```

Add a feed (will automatically follow the feed as well)
```bash
gator addfeed <feed url>
# do include the protocol (https://)
```

See all feeds that have been added
```bash
gator feeds
```

Follow a feed
```bash
gator follow <feed url>
```

List all feeds you've followed
```bash
gator following
```

Unfollow a feed
```bash
gator unfollow <feed url>
```

Start Aggregation
```bash
gator agg <time between fetching>
# example: 1s, 1m, 1h
```

Browse posts that have been aggregated
```bash
gator browse <limit, default=2>
```

## Features
- RSS Feed aggregation and management
- Written in Go for high performance
- Clean and maintainable codebase
- Easy to integrate

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License
This project is open source and available under the MIT License.