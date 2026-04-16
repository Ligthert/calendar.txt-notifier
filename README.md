# Calendar.txt Notifier (caltxtnot)
A simple and highly configurable tool to send notifications of my calendar.txt file.

This tool is designed to read a simplified `calendar.txt`, grab the events (and birthdays) of the day its executed and print the events and birthdays. 

The `calendar.txt` uses a simple formatting using `MM-DD` as the date format. An example `calendar.txt` is:
```
05-15 vr 14:00 Go to the dentist
05-16 za Friends Bday
```
Same format can be applied for birthdays (Support to calculate age is on the wishlist)

Currently `stdout` and `discord` (via webhooks) are supported as an output. Config, calender, and birthday files can be placed on the local filesystem or some webserver (using the full path, or URI's starting with `file://`, `http://`, and `https://`)


# Usage:
* `--config`: Using the config-file [required]
* `--debug`: Print debug messages
* `--test`: Test if the config-file is valid

`--config` can point to either a file path, file://, http://, or https:// URI resource.

# Config file
For an example config-file please see [config.toml-dist](config.toml-dist). For normal usage copy this file to `config.toml`.

# Roadmap/Wishlist
* Detach and run in daemon mode to frequently poll and continuesly notify during the day
* Read config from env vars to follow the 12-factor app methodology
* Do something with docker
* XMPP support among other messaging platforms

# Background
My server in the datacenter that ran an XMPP server and a simple script that did some onliner magic using `curl`, `date`, and `grep` before sending it off to the XMPP server died. So I wanted some extra flexibility and options and decided to rewrite it in Go (for funzies and XP).

I've been using a variant of `calendar.txt` which was limited to the `MM-DD` format and I tailored this tool. While the extra +600 lines seems exessive, it should allow for future features like:
* Running it as a daemon
* Send reminders or notices of events.