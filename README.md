# Github Activity command line application

this is my solution to the challenge [roadmap.sh](https://roadmap.sh/projects/github-user-activity).


you can start by clonnning the repo first using:
```sh
git clone https://github.com/amruth6002/github_user_activity
```

go to the folder where you installed the project files in command line using cd

if you are using wsl the you can run:
```sh
go build -o github-activity ./src
sudo mv github-activity /usr/local/bin/
```
so that the application runs globally

you can run the command using
```sh
github-activity <username>
```
