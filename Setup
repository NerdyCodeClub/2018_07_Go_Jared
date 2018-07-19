# Quick Setup

These are the steps I followed to get started with GO.

# Install VS Code & GO

  - Install the latest version of [VS Code](https://code.visualstudio.com/download) 
  - Install the latest version of [GO](https://golang.org/dl/)
  - Install the VS Code extention for [GO Lang Support](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

# Configure Proxy Settings
From what I've seen most of the packages go uses are hosted on github. So when you try to install a package using the **go get** command, behind the scenes a **git clone** will be run. If you're behind a proxy this probably won't work the first time. 

To get around this you'll need to add your proxy settings to git. Run the following commands to configure git to use a proxy. **Be sure to replace the values in square brackets with your details**.

```sh
git config --global http.proxy http://[domain]\[username]:[password]@[proxyServer]:[port]
git config http.sslVerify "false"
```

After running the above commands you should be able to clone git repos & run the **go get** command within VSCode.
