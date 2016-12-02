#!/bin/bash

# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# start rpcbind if it is not started yet
s=`/etc/init.d/rpcbind status`
if [[ $s == *"not running"* ]]; then
   echo "Starting rpcbind: /sbin/rpcbind -w"
   /sbin/rpcbind -w
fi
echo `/etc/init.d/rpcbind status`

# mount with passing paramaters
/bin/mount "${@}"

# kill rpcbindn process after mount finishes
if [[ $s == *"not running"* ]]; then
      echo "Kill rpcbind pid $( pidof rpcbind )"
      kill $( pidof rpcbind )
fi
