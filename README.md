# WIP repo puller

this to simplify server side projects that need to just update the workdir on server and doing something when there is push to the repo. so its just support `push event`

## Support
- github push master

## Usage
```
make install
repohook -branch=master -path=/path/repo/wanna/pull -script=/script/after/event
```