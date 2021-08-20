# install
use scoop:
```
scoop bucket add scoop-rehtt "https://github.com/Rehtt/scoop-rehtt.git"
scoop install tcping
```

use go:
```
go get -u github.com/Rehtt/tcping
```

or download [file](https://github.com/Rehtt/tcping/releases/)


# use
TCP Ping v0.2.0
https://github.com/rehtt/tcping
Use: tcping [-w] [-l] [-t] <IP address / Host> [Port (default: 80)]
Must fill in IP address or Host.
You can choose to fill in the port, you can add multiple ports or use "-" to specify the range, port default 80.
-w 5    : ping every 5 seconds, default 1
-l 5    : send 5 pings, default 3
-t 5    : timeout 5 seconds, default 2
eg: tcping google.com
eg: tcping google.com 443
eg: tcping google.com 80 443
eg: tcping google.com 80-85 443-448
eg: tcping -w 10 -l 6 -t 3 google.com 443
