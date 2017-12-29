package messages

var Usage = `
       __ _ _   _
      / _  | | | |
     | (_| | |_| |
      \__,_|\__,_|
au shortcuts for the anki_viewer_umbrella project
Usage:
  au root                       cd into umbrella project root
  au ..                         same as: au root
  au web                        cd into project web directory
  au assets                     cd into project assets directory
  au anki                       cd into anki directory
  au .                          same as: au anki
  au nodeapp                    cd into anki/node_app directory
  au install                    runs the necessary installations
  au build                      builds files
    -w                          watches for changes and rebuilds
    -js [-w]                    static js files
    -css [-w]                   css files
    -elm [-w]                   elm files
    -static [-w]                static files
  au test                       runs the phoenix tests
    -w                          watches these tests
    -node [-w]                  runs the node tests
    -js [-w]                    runs the js tests
    -elm [-w]                   runs the elm tests
    -nw [-w]                    runs the nightwatch tests
    -all                        runs all the tests
  au cover                      runs the test coverage
    -html                       generates html coverage report
    -open                       opens coverage report
    -htmlopen                   opens coverage report after generating it
  au start                      starts the server
    -prod                       starts in prod environment
  au versions                   prints the versions of project technologies
  au deploy                     deploys the app
Examples:
  # builds our js static assets
  # and watches for changes
  $ au build -js -w
  # tests
  $ au test apps/anki/test/anki_test.ex -w`

var InputRequired = `Input required!
For usage, see:
$ au help`

var NotInApp = `Need to be inside anki_viewer app to run au`

var NotInRoot = `Need to be in project root to run this command`
