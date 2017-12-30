package messages

var Usage = `
av shortcuts for the av_umbrella project

Usage:

%s
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

You will need to set up this file structure:

├── ankiviewer
    ├── av_umbrella (github.com/ankiviewer/av_umbrella)
    ├── nodeapp (github.com/ankiviewer/nodeapp)
    └── assets (github.com/ankiviewer/assets)

This can be set up simply with:

av setup filestructure 

# ensure you have the correct technologies installed
av versions

# set up the database, export the correct environment variables and setup aliases
av setup

# install the dependencies
av install

# run the build
av build

# start the app
av start

# For more information on av usage
av help
`
// # set up the database, aliases and other required setups
// # place the following lines in your .rc file for constant use
// 
// export AV_ANKI_SQLITE_PATH=%s
// export AV_PROJ_PATH=%s
// av setup

var NotInApp = `Need to be inside ankiviewer app to run av`

var NotInRoot = `Need to be in project root to run this command`
