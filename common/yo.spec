################################################################################

# rpmbuilder:relative-pack true

################################################################################

%define  debug_package %{nil}

################################################################################

%define _posixroot        /
%define _root             /root
%define _bin              /bin
%define _sbin             /sbin
%define _srv              /srv
%define _home             /home
%define _opt              /opt
%define _lib32            %{_posixroot}lib
%define _lib64            %{_posixroot}lib64
%define _libdir32         %{_prefix}%{_lib32}
%define _libdir64         %{_prefix}%{_lib64}
%define _logdir           %{_localstatedir}/log
%define _rundir           %{_localstatedir}/run
%define _lockdir          %{_localstatedir}/lock/subsys
%define _cachedir         %{_localstatedir}/cache
%define _spooldir         %{_localstatedir}/spool
%define _crondir          %{_sysconfdir}/cron.d
%define _loc_prefix       %{_prefix}/local
%define _loc_exec_prefix  %{_loc_prefix}
%define _loc_bindir       %{_loc_exec_prefix}/bin
%define _loc_libdir       %{_loc_exec_prefix}/%{_lib}
%define _loc_libdir32     %{_loc_exec_prefix}/%{_lib32}
%define _loc_libdir64     %{_loc_exec_prefix}/%{_lib64}
%define _loc_libexecdir   %{_loc_exec_prefix}/libexec
%define _loc_sbindir      %{_loc_exec_prefix}/sbin
%define _loc_bindir       %{_loc_exec_prefix}/bin
%define _loc_datarootdir  %{_loc_prefix}/share
%define _loc_includedir   %{_loc_prefix}/include
%define _loc_mandir       %{_loc_datarootdir}/man
%define _rpmstatedir      %{_sharedstatedir}/rpm-state
%define _pkgconfigdir     %{_libdir}/pkgconfig

################################################################################

Summary:         Command-line YAML processor
Name:            yo
Version:         0.5.0
Release:         0%{?dist}
Group:           Applications/System
License:         EKOL
URL:             https://github.com/essentialkaos/yo

Source0:         https://source.kaos.st/%{name}/%{name}-%{version}.tar.bz2

BuildRoot:       %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

BuildRequires:   golang >= 1.13

Provides:        %{name} = %{version}-%{release}

################################################################################

%description
Command-line YAML processor.

################################################################################

%prep
%setup -q

%build
export GOPATH=$(pwd)
go build src/github.com/essentialkaos/%{name}/%{name}.go

%install
rm -rf %{buildroot}

install -dm 755 %{buildroot}%{_bindir}
install -pm 755 %{name} %{buildroot}%{_bindir}/

%clean
rm -rf %{buildroot}

################################################################################

%files
%defattr(-,root,root,-)
%doc LICENSE.EN LICENSE.RU
%{_bindir}/%{name}

################################################################################

%changelog
* Fri Jan 10 2020 Anton Novojilov <andy@essentialkaos.com> - 0.5.0-0
- ek package updated to the latest stable version

* Sat Jun 15 2019 Anton Novojilov <andy@essentialkaos.com> - 0.4.0-0
- ek package updated to the latest stable version

* Wed Oct 17 2018 Anton Novojilov <andy@essentialkaos.com> - 0.3.2-0
- go-simpleyaml updated to v2

* Tue Mar 27 2018 Anton Novojilov <andy@essentialkaos.com> - 0.3.1-0
- ek package updated to latest stable release

* Thu May 25 2017 Anton Novojilov <andy@essentialkaos.com> - 0.3.0-0
- ek package updated to v9

* Sat Apr 15 2017 Anton Novojilov <andy@essentialkaos.com> - 0.2.0-0
- ek package updated to v8

* Wed Mar 08 2017 Anton Novojilov <andy@essentialkaos.com> - 0.1.0-0
- ek package updated to v7

* Tue Feb 14 2017 Anton Novojilov <andy@essentialkaos.com> - 0.0.2-0
- Fixed output for arrays with maps and sub arrays

* Tue Feb 14 2017 Anton Novojilov <andy@essentialkaos.com> - 0.0.1-0
- Initial release
