package messages

// TODO: generate this from Cmds
var Usage = `
av shortcuts for the av_umbrella project

Usage:
  av setup
  av install                    runs the necessary installations
  av build                      builds files
  av test                       runs the phoenix tests
  av cover                      runs the test coverage
  av start                      starts the server
  av versions                   prints the versions of project technologies
  av deploy                     deploys the app

Examples:
  # builds our js static assets
  # and watches for changes
  $ av build -js -w
  # tests
  $ av test apps/anki/test/anki_test.ex -w`

  // TODO: add these for alias output
// find "$(find $HOME -type d | grep Anki2 | head -1)" -type f | grep collection.anki2 for anki sqlite db path
// pwd for proj path
var NoInput = `
       __ _ _   _
      / _  | | | |
     | (_| | |_| |
      \__,_|\___/

# Quick Start:

git clone https://github.com/ankiviewer/av_umbrella.git && cd av_umbrella

# ensure you have the correct technologies installed
av versions

# install the dependencies and run the build
av install

# set up the database, aliases and other required setups
# place the following lines in your .rc file for constant use

export AV_ANKI_SQLITE_PATH=%s
export AV_PROJ_PATH=%s
av setup

# start the app
av start

# For more information on av usage
av help
`

var NotInApp = `Need to be inside ankiviewer app to run av`

var NotInRoot = `Need to be in project root to run this command`
