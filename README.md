# Fuzz Monkey

Fuzz Monkey is a bit like chaos monkey only with more fur and instead of tearing
down infrastructure like some kind of crazed baboon in a shoe shop, it carefully
and surgically flings its poop at specific http endpoints. It ain't pretty,
but it's damn sure fuzzy.

![CLI](/assets/cli.png?raw=true "CLI")

## Building the Binary

Check out this project and, from the `app` folder run

```
go build -o monkey .
```

## Priming the Monkey for all out war

The Monkey is a chaotic but loyal warrior. You tell it where to fling and by the grace
of God, it'll fling. The Monkey's instructions come in the form of a JSON file. When
you send the Monkey off into battle, you can target it's wrath like this:

```
./monkey path/to/config.json
```

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
          "maxResponseTime": 5000,
          "expectedStatus": "400"
        },
        {
          "type": "HTTP_SPAM",
          "maxResponseTime": 5000,
          "expectedStatus": "200",
          "concurrents": 20,
          "messagesPerConcurrent": 100
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
  "maxResponseTime": 5000,
  "expectedStatus": "200",
  "concurrents": 20,
  "messagesPerConcurrent": 100,
  "method": "GET"
}
```
