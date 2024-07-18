# yaro
An open source launcher for EA games

## Current Status

- [X] Authorization
- [ ] Download games
- [ ] Update games
- [ ] DRM/Online play
- [ ] Download DLC

## How to build
1. Create a new directory to store the Go workspace. We'll name it "yaro_workspace"
2. Run "git clone https://github.com/ahmouse15/yaro"
3. Run "git clone https://github.com/ahmouse15/webview_go"
4. Create a new "go.work" file.
5. Inside of it, write "use (./webview_go ./yaro)"
