It's like Imgur, but for sixel images over the command line.

Currently only configured around and targeting unsecured local HTTP connections. That will change once I learn more about making secure and reliable web services with Go.

Quick notes:
- use either flags or cobra/viper libraries for the cli
- optional crude encryption through password-locked archives
- instead of having the endpoint be the filename, randomly generate an all-lowercase alphanumerical passcode
- be able to select a directory or a grouping of files to zip and send
- ability to save, load, and then list endpoints for repeated/scheduled serving of the same files
