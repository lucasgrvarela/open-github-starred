# Open Starred GitHub Repositories

This program lists your starred GitHub repositories and opens them in the default web browser.

## Usage

1. Set up your GitHub credentials by exporting environment variables:

   ```bash
   export GITHUB_USERNAME=your_username
   export GITHUB_TOKEN=your_token
   ```

2. go run main.go OR go run main.go -open-per-page=3
    ```bash
    $ go run main.go -open-per-page 3
    jwt
    echo
    vscode-solargraph
    ```

3. Follow the on-screen prompts to continue opening more repositories.
    ```bash
    Continue to open more 3 repositories? (Press Enter or Y to confirm): 
    Confirmation received. Proceeding...
    asdf-kubent
    gimp
    haproxy

    Continue to open more 3 repositories? (Press Enter or Y to confirm): 
    Confirmation received. Proceeding...
    No more repositories found. Exiting...
    ```