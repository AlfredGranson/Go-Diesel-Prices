
# doe_fetcher

Go tool to retrieve diesel prices from the US Energy Information Administration API


## Installation

First create an API key by signing up at: https://www.eia.gov/opendata/register.php

Clone the repo

```bash
  git clone git@github.com:AlfredGranson/doe_fetcher.git
```
Create your env and update the EIA_KEY to match your key
```bash
  cd doe_fetcher
  cp .env.example .env
```
Build the project
```bash
  go build
```
## Usage/Examples

```bash
  ./doe_fetcher
``` 
