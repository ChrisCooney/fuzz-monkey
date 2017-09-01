# Fuzz Monkey

![Build Status](https://travis-ci.org/ChrisCooney/fuzz-monkey.svg?branch=master "Build Status")
[![Coverage](https://codecov.io/gh/ChrisCooney/fuzz-monkey/branch/master/graph/badge.svg)](https://codecov.io/gh/ChrisCooney/fuzz-monkey)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/c59ca3588d3548868c64f71e1cc8f20e)](https://www.codacy.com/app/chris_cooney/fuzz-monkey?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ChrisCooney/fuzz-monkey&amp;utm_campaign=Badge_Grade)

Fuzz Monkey is a bit like chaos monkey only with more fur and instead of tearing
down infrastructure like some kind of crazed baboon in a shoe shop, it carefully
and surgically flings its poop at specific http endpoints. It ain't pretty,
but it's damn sure fuzzy.

![CLI](/assets/cliv1.1.png?raw=true "CLI")

## Building the Binary

Check out this project and, from the `app` folder run

```
go build -o monkey .
```

## Priming the Monkey for all out war

### The "Chaos Monkey" way

The Monkey is a chaotic but loyal warrior. You tell it where to fling and by the grace
of God, it'll fling. The Monkey's instructions come in the form of a JSON file. When
you send the Monkey off into battle, you can either target it's wrath like this:

```
./monkey path/to/config.json
```

or you can simply run the script and it will automatically root around for a file named
_fuzz-monkey.json_.

The Configuration file has a specific format, otherwise the Monkey gets confused. In the root
of the config file is the `endpoints` field. This specifies the targets for the monkey to
attack.

```
{
  endpoints: []
}
```

In endpoints, you specify details for each of the endpoints you want the monkey to attack.
For example:

```
{
  endpoints: [
    {
      "name": "Chris",
      "host": "localhost",
      "port": "80",
      "path": "/orders/1",
      "protocol": "http",
      "attacks": [ ]
    }
  ]
}
```

Each endpoint must have at least one attack registered against it. An attack requires a type
field and the config parameters for that type of attack. The current attack types are:

| Attack Type | Description  |
| -------------|-----|
| HTTP_SPAM     | Goes to town on an endpoint with randomly selected HTTP requests. |
| CORRUPT_HTTP  | Opens a TCP connection and makes corrupt HTTP requests at the endpoint. |
| URL_QUERY_SPAM  | Takes a provided list of parameters and tries known dangerous values  |

For example, in your config, your attack might look something like:

```
{
  endpoints: [
    {
      "name": "Chris",
      "host": "localhost",
      "port": "80",
      "path": "/orders/1",
      "protocol": "http",
      "attacks": [
        {
          "type": "CORRUPT_HTTP",
          "expectedStatus": "400"
        },
        {
          "type": "HTTP_SPAM",
          "expectedStatus": "200",
          "concurrents": 20,
          "messagesPerConcurrent": 100
        },
        {
          "type": "URL_QUERY_SPAM",
          "expectedStatus": "400",
          "parameters": "a,b,c"
        }
      ]
    }
  ]
}
```

The following will randomly run two attacks at the endpoint. The first will randomly fire corrupted
HTTP requests over TCP at the endpoint. The second will randomly open up 20 concurrent connections
and they will each fire 100 requests at the endpoint.

If you don't specify a method in your config, then the Monkey will randomly select one for you
because it enjoys a wide and varied diet. If, however, you wish to specify a HTTP method to use
then simply include the method field in your attack config:

```
{
  "type": "HTTP_SPAM",
  "expectedStatus": "200",
  "concurrents": 20,
  "messagesPerConcurrent": 100,
  "method": "GET"
}
```

### CI Mode

![CLI](/assets/ci-mode.png?raw=true "CI Mode")

If you just want to run a single test through, you can do this by simply adding the `-c` switch
to your command, for example:

    ./monkey -c

This will cause the application to run in CI mode. This will go through each of the attacks once
and return with an error code if any of the attacks fail.
