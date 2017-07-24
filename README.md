[![Build Status](https://travis-ci.org/rnitame/daily.svg?branch=master)](https://travis-ci.org/rnitame/daily)

# daily
Get daily events from GitHub

## Usage

```
# Show all GitHub events
$ daily

# Show organization events only
$ daily -org <org_name>

# Show all GHE events
$ daily -ghe <ghe_domain_name>
```

## Set GitHub personal token

```
$ git config --global "github.token" xxxxx
```

## Set GHE personal token

```
$ git config --global "ghe.token" xxxxx
```
