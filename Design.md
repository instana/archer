archer hook before-remove
archer hook after-remove
archer hook before-install
archer hook after-install

archer realize all
archer realize install
archer realize build
archer realize configure

archer collection list
archer collection show
archer collection status
archer collection stop
archer collection start
archer collection restart
archer collection purge

archer state list
archer state purge

archer build

/etc/archer/config.conf
/etc/archer/policy.conf
/ect/archer/collection.d/collectionname.conf
/usr/lib/archer/package.d/packagename.conf
/var/lib/archer/state.db

package resources:

service
user
group
template
execute
  always
  once
dir


Archerfile:

pkg {
    name = "package-name"
    description = "package description"
    vendor = "vendor"
    maintainer = "maintainer"
    url = "url"
    license = "license"
    arch = "arch"
    version = "1.0.0"
    iteration = "1"
    branch = "master"
    vcs_revision = ""
}

requirement {
    name = "whatever"
    method = "depends"
    operation = ""
    version = ""
}

build {
    target_path = "path"
    work_path = "path"
    out_path = "path"
    file_user = "root"
    file_group = "root"
    rpm = true
    deb = true
}