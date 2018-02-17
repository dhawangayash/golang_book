find . -type f  | grep class$ | xargs -I {} rm {}

git status

echo "*********************"
git add . 

if [ -n "$1" ]; then
    git commit -m "$1"
else
    git commit -m `date +%Y%m%d`
fi

git pull --rebase origin master
git push origin master

