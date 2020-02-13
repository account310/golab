#!/bin/sh

# POST-COMMIT HOOK
#   [1] REPOS-PATH   (the path to this repository)
#   [2] REV          (the number of the revision just committed)
export LANG=zh_CN.UTF-8

# *************************************************************
# * Edit the following lines to suit your environment         *
# *************************************************************
# Set full path to svnlook.exe
SVNLOOK=/usr/bin/svnlook 
# Set full path to svnplus.exe
POSTEXE=/opt/mjc/svn/golab/src/mattermost-plugin-svn/bin/svnplus
POSTCONFIG=/opt/mjc/svn/golab/src/mattermost-plugin-svn/conf/config.toml
# *************************************************************
# * This sets the arguments supplied by Subversion            *
# *************************************************************
REPOS=%1
REV=%2


# *************************************************************
# * Get Author and comment                                    *
# *************************************************************
setlocal EnableDelayedExpansion

# get comments
for COMMNET in `$SVNLOOK log -r $REV $REPOS`;
do
  echo "COMMNET1" $COMMNET >> /opt/aaa.log
  $COMMNET=${COMMNET}"[___]"$COMMNET
  echo "COMMNET2" $COMMNET >> /opt/aaa.log
done
# get author
for AUTHOR in `$SVNLOOK author $REPOS -r $REV`;
do
  echo "AUTHOR1" $AUTHOR >> /opt/aaa.log
  $AUTHOR=$AUTHOR
  echo "AUTHOR2" $COMMNET >> /opt/aaa.log
done
# get file list
for FILELIST in `$SVNLOOK changed $REPOS -r $REV `;
do
  echo "FILELIST1" $FILELIST >> /opt/aaa.log
  $FILELIST=${FILELIST}"[___]"$FILELIST
  echo "FILELIST2" $FILELIST >> /opt/aaa.log
done

# *************************************************************
# * Hand it to commit                                       *
# *************************************************************
"$POSTEXE" -conf="$POSTCONFIG" -author="$AUTHOR" -comments="$COMMENT" -filelist="$FILELIST" -projectpath="$REPOS" -sendtype="" -rev="$REV"
exit 0
