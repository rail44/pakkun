%w(
  openjdk-8-jdk
  schedtool
  git-core
  gnupg
  flex
  bison
  gperf
  build-essential
  zip
  curl
  zlib1g-dev
  gcc-multilib
  g++-multilib
  libc6-dev-i386
  lib32ncurses5-dev
  x11proto-core-dev
  libx11-dev
  lib32z-dev
  ccache
  libgl1-mesa-dev
  libxml2-utils
  xsltproc
  unzip
).each do |pkg|
  package pkg do
    action :install
  end
end
