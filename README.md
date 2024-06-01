# Go Orchestra Agent

> This thing can help you with computing your expressions with multi threading. 
> Just deploy GoOrchestra(https://github.com/J3olchara/GoOrchestra) and then return here

# Installation guide
- Clone this repository
```bash
git clone https://github.com/J3olchara/GoOrchestraAgent && cd GoOrchestraAgent
```


# Linux
- Firstly you need installed docker-compose with minimum version 2.24
> to check version use ```docker-compose --version```
```bash
sudo apt-get install docker-compose
```
- Then create environment directory
```bash
cp ./env_example/dev.agent.env-example ./env/dev.agent.env
```
- Then you need to fill your application environment in ./env directory as written in env-example directory or make defaults by
```bash
mkdir ./env && cp ./env_example/dev.agent.env-example ./env/dev.agent.env
```
- Start Application
```bash
docker-compose -f docker-compose.dev.yml up
```


# MacOS
- Firstly you need installed docker-compose with minimum version 2.24
> to check version use ```docker-compose --version```
```bash
brew install docker-compose
```
- Then create environment directory
```bash
cp ./env_example/dev.agent.env-example ./env/dev.agent.env
```
- Then you need to fill your application environment in ./env directory as written in env-example directory or make defaults by
```bash
mkdir ./env && cp ./env_example/dev.agent.env-example ./env/dev.agent.env
```
- Start Application
```bash
docker-compose -f docker-compose.dev.yml up
```

# Windows
- Firstly you need installed docker-compose with minimum version 2.24
> to check version use ```docker-compose --version```
```bash
winget install docker-compose
```
- Then create environment directory
```bash
cp ./env_example/dev.agent.env-example ./env/dev.agent.env
```
- Then you need to fill your application environment in ./env directory as written in env-example directory or make defaults by
```bash
mkdir ./env && cp ./env_example/dev.agent.env-example ./env/dev.agent.env
```
- Start Application
```bash
docker-compose -f docker-compose.dev.yml up
```