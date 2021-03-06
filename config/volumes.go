package config

/*
 Copyright 2019 - 2020 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

// volume configuration settings used by the pgBackRest repo mount
const VOLUME_PGBACKREST_REPO_NAME = "backrestrepo"
const VOLUME_PGBACKREST_REPO_MOUNT_PATH = "/backrestrepo"

// volume configuration settings used by the SSHD secret
const VOLUME_SSHD_NAME = "sshd"
const VOLUME_SSHD_MOUNT_PATH = "/sshd"

// volume configuration settings used by tablespaces

// the pattern for the volume name used on a tablespace, which follows
// "tablespace-<tablespaceName>"
const VOLUME_TABLESPACE_NAME_PREFIX = "tablespace-"

// the pattern for the path used to mount the volume of a tablespace, which
// follows "/tablespace/<pvcName>"
const VOLUME_TABLESPACE_PATH_PREFIX = "/tablespaces/"

// the pattern for the name of a tablespace PVC, which is off the form:
// "<clusterName>-tablespace-<tablespaceName>"
const VOLUME_TABLESPACE_PVC_NAME_FORMAT = "%s-tablespace-%s"
