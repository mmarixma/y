!#/bin/bash
export KEY_INTERVIEW=$(cat ./streets/Buckingham_Place | sed -n '179p' | grep -oE '[0-9]+')
echo $KEY_INTERVIEW
cat ./interviews/interview-$KEY_INTERVIEW
echo $MAIN_SUSPECT
