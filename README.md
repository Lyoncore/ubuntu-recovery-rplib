# ubuntu-recovery-rplib

## logger in rplib
The rplib logger function is calling snapd/logger.
[snapd/logger](https://github.com/snapcore/snapd/)
Before you start to log, please call logger.SimpleSetup() first.

There are three levels for this logger.
    Notice: logger.Noticef()
        The log will be in /var/log/syslog and Stdout
    Debug:  logger.Debugf()
        The log will be in syslog only.
        Stdout will show if env variable SNAPD_DEBUG=1
    Panic:  logger.Panicf()
        The log will be in syslog and Stdout. Also with panic.


