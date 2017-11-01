rm ./*.json ./*.txt
./agenda register -uUserName --password pass --email=a@xxx.com --contact=123456
cat userList.json
printf "\n"
./agenda login -uUserName --password pass1
cat curUser.json
printf "\n"
./agenda logout
./agenda login -uUserName --password pass
cat curUser.json
printf "\n"
./agenda logout
cat curUser.json
printf "\n"
./agenda query
printf "\n"
./agenda register -upart1 --password pass --email=a@xxx.com --contact=123456
./agenda login -upart1 --password pass
cat userList.json
./agenda cm --title=meeting --part=part1 --start=2001-11-11/12:00 --end=2012-12-13/12:00
cat meetingList.json
