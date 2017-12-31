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

var FileStructure = `
├── ankiviewer
    ├── av_umbrella (github.com/ankiviewer/av_umbrella)
    ├── nodeapp (github.com/ankiviewer/nodeapp)
    └── assets (github.com/ankiviewer/assets)
`
var NoInput = `
       __ _ _   _
      / _  | | | |
     | (_| | |_| |
      \__,_|\___/

# Quick Start:

You will need to set up this file structure: ` + FileStructure + `
This can be set up simply with:

av setup filestructure 

# ensure you have the correct technologies installed
av versions

# set up the file structure, environment variables and aliases
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

var NotInApp = `Need to be inside ankiviewer app to run av`

var NotInRoot = `Need to be in project root to run this command`
