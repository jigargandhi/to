# to: Command line shortcut 

This tool allows you to configure and visit urls with some different characters as parameter.

## Installation
1. Install golang 
2. Run 
```shell
go install go install github.com/jigargandhi/to@latest
```

## Example

If you frequently need to visit https://jira.yourcompany.com/browse/JIRA-TICKET-ID then you can add the url to the tool with the following command

```shell
to url add jira https://jira.yourcompany.com/browse/{TAG}
```
and then run if you want to open JIRA ticket PROJECT-1234 then you can run

```shell
to url jira PROJECT-1234
```
it will open a browser to that location.

You can also do google search from command line by adding

```shell
to url add google https://google.com/?q={TAG}
```
and to search for `golang` then run 

```shell
to url google golang
```
The urls and their parameters are stored in `~\.to\shortcuts.yml` folder

Its not mandatory to add `{TAG}` in the url.