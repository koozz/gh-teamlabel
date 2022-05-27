# gh-teamlabel

This GitHub CLI extension can be used to label a pull request with team names
that the author of the pull request is part of.

## Usage

From the commandline:

```bash
gh extension install koozz/gh-teamlabel # Just once
gh teamlabel -org my-org -labels=team1:Team1Label,team2:Team2Label
```

In GitHub Actions, add a step like this:

```yaml
  - name: Team labeling
    run: |
      gh extension install koozz/gh-teamlabel # Just once
      gh teamlabel -org my-org -labels=team1:Team1Label,team2:Team2Label
```

## Adding labels

Adding labels on a pull request only works if these labels have been added to
the repository.

You could prepend the above scripts with calls to [gh label create](https://cli.github.com/manual/gh_label_create)

```bash
gh label create Team1Label --color 336699 --description "Team one label description" --force
gh label create Team2Label --color 336699 --description "Team two label description" --force

gh teamlabel -org my-org -labels=team1:Team1Label,team2:Team2Label
```

## License

MIT
