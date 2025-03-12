# bimager
I try to conveniently rename bonus image files using data from the Rallymaster.

ScoreMaster will need access to an image file associated with each bonus and the filename used could be anything but, given that some environments are fussier than others, the most convenient format is *bonus.ext* where 'bonus' is the (case-sensitive on some systems) unique code of the bonus and 'ext' is the filetype, typically jpg or png.

Either -smf or -fld must be used to specify the folder affected, either the ScoreMaster installation root folder (-smf) or the image folder itself (-fld)

The pattern matching the bonusid is specified with -bre
