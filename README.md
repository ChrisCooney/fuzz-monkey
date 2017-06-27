# Fuzz Monkey

Fuzz Monkey is a bit like chaos monkey only with more fur and instead of tearing
down infrastructure like some kind of crazed baboon in a shoe shop, it carefully
and surgically flings its poop at specific http endpoints. It ain't pretty,
but it's damn sure fuzzy.

## Use

Check out this project and, from the `app` folder run

```
go build -o monkey .
```

When running the binary, you will need to create a config file to tell the monkey
in which direction to fling.
