# go-deployment-samples - Revel

To deploy this application to Azure's Golang Blessed Images:

- Install Revel so it's accessible via the CLI - [docs](https://revel.github.io/tutorial/gettingstarted.html)
- Run `revel build . ./somepathyoucreated`. Somepath will contain the build output
- On the App Service, add a custom startup command set to **./somepathyoucreated/run.sh**
- In this case, its `./build/run.sh`

