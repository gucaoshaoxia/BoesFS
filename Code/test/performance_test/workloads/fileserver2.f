#
# CDDL HEADER START
#
# The contents of this file are subject to the terms of the
# Common Development and Distribution License (the "License").
# You may not use this file except in compliance with the License.
#
# You can obtain a copy of the license at usr/src/OPENSOLARIS.LICENSE
# or http://www.opensolaris.org/os/licensing.
# See the License for the specific language governing permissions
# and limitations under the License.
#
# When distributing Covered Code, include this CDDL HEADER in each
# file and include the License file at usr/src/OPENSOLARIS.LICENSE.
# If applicable, add the following below this CDDL HEADER, with the
# fields enclosed by brackets "[]" replaced with your own identifying
# information: Portions Copyright [yyyy] [name of copyright owner]
#
# CDDL HEADER END
#
#
# Copyright 2008 Sun Microsystems, Inc.  All rights reserved.
# Use is subject to license terms.
#
set $dir=/home/boes/Final/boesfs/Code/test/performance_test/tmp2
set $meandirwidth=100
set $ndirs=10000
set $nthreads=50

set mode quit firstdone

define fileset name=bigdirset,path=$dir,size=0,leafdirs=$ndirs,dirwidth=$meandirwidth

define process name=filereader,instances=1
{
  thread name=dirreaderthread,memsize=10m,instances=$nthreads
  {
    flowop makedir name=mkdir1,filesetname=bigdirset
  }
}

echo  "BoesFS Mkdir Test successfully loaded"

run 60