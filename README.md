# gh-teamlabel

This GitHub CLI extension can be used to label a pull request with team names
that the author of the pull request is part of.

[![Total alerts](https://img.shields.io/lgtm/alerts/g/koozz/gh-teamlabel.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/koozz/gh-teamlabel/alerts/)
[![Language grade: Go](https://img.shields.io/lgtm/grade/go/g/koozz/gh-teamlabel.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/koozz/gh-teamlabel/context:go)

## Usage

From the commandline:

```bash
gh extension install koozz/gh-teamlabel # Just once
gh teamlabel -org my-org team1_slug:Team1Label team2_slug:Team2Label
```

In GitHub Actions, add a step like this:

```yaml
  - name: Team labeling
    env:
      GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    run: |
      gh extension install koozz/gh-teamlabel --pin 2.0.3
      gh teamlabel -org my-org team1_slug:Team1Label team2_slug:Team2Label
```

### Caveat

The used GitHub token must have enough privileges, either use a GitHub App or a Personal Access Token.

## Adding labels

Adding labels on a pull request only works if these labels have been added to
the repository.

You could prepend the above scripts with calls to [gh label create](https://cli.github.com/manual/gh_label_create)

```bash
gh label create Team1Label --color 336699 --description "Team one label description" --force
gh label create Team2Label --color 336699 --description "Team two label description" --force

gh teamlabel -org my-org team1:Team1Label team2:Team2Label
```

## License

MIT
