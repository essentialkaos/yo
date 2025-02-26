################################################################################

%define debug_package  %{nil}

################################################################################

Summary:        Command-line YAML processor
Name:           yo
Version:        1.0.1
Release:        0%{?dist}
Group:          Applications/System
License:        Apache License, Version 2.0
URL:            https://kaos.sh/yo

Source0:        https://source.kaos.st/%{name}/%{name}-%{version}.tar.bz2

BuildRoot:      %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

BuildRequires:  golang >= 1.23

Provides:       %{name} = %{version}-%{release}

################################################################################

%description
Command-line YAML processor.

################################################################################

%prep

%setup -q
if [[ ! -d "%{name}/vendor" ]] ; then
  echo -e "----\nThis package requires vendored dependencies\n----"
  exit 1
elif [[ -f "%{name}/%{name}" ]] ; then
  echo -e "----\nSources must not contain precompiled binaries\n----"
  exit 1
fi

%build
pushd %{name}
  go build %{name}.go
  cp LICENSE ..
popd

%install
rm -rf %{buildroot}

install -dm 755 %{buildroot}%{_bindir}
install -dm 755 %{buildroot}%{_mandir}/man1

install -pm 755 %{name}/%{name} %{buildroot}%{_bindir}/

./%{name}/%{name} --generate-man > %{buildroot}%{_mandir}/man1/%{name}.1

%clean
rm -rf %{buildroot}

%post
if [[ -d %{_sysconfdir}/bash_completion.d ]] ; then
  %{name} --completion=bash 1> %{_sysconfdir}/bash_completion.d/%{name} 2>/dev/null
fi

if [[ -d %{_datarootdir}/fish/vendor_completions.d ]] ; then
  %{name} --completion=fish 1> %{_datarootdir}/fish/vendor_completions.d/%{name}.fish 2>/dev/null
fi

if [[ -d %{_datadir}/zsh/site-functions ]] ; then
  %{name} --completion=zsh 1> %{_datadir}/zsh/site-functions/_%{name} 2>/dev/null
fi

%postun
if [[ $1 == 0 ]] ; then
  if [[ -f %{_sysconfdir}/bash_completion.d/%{name} ]] ; then
    rm -f %{_sysconfdir}/bash_completion.d/%{name} &>/dev/null || :
  fi

  if [[ -f %{_datarootdir}/fish/vendor_completions.d/%{name}.fish ]] ; then
    rm -f %{_datarootdir}/fish/vendor_completions.d/%{name}.fish &>/dev/null || :
  fi

  if [[ -f %{_datadir}/zsh/site-functions/_%{name} ]] ; then
    rm -f %{_datadir}/zsh/site-functions/_%{name} &>/dev/null || :
  fi
fi

################################################################################

%files
%defattr(-,root,root,-)
%doc LICENSE
%{_mandir}/man1/%{name}.1.*
%{_bindir}/%{name}

################################################################################

%changelog
* Mon Jun 24 2024 Anton Novojilov <andy@essentialkaos.com> - 1.0.1-0
- Code refactoring
- Dependencies update

* Thu Mar 28 2024 Anton Novojilov <andy@essentialkaos.com> - 1.0.0-0
- Improved support information gathering
- Code refactoring
- Dependencies update

* Tue Dec 19 2023 Anton Novojilov <andy@essentialkaos.com> - 0.5.7-0
- Code refactoring
- Dependencies update

* Mon Mar 06 2023 Anton Novojilov <andy@essentialkaos.com> - 0.5.6-0
- Added verbose info output
- Dependencies update
- Code refactoring

* Wed Nov 23 2022 Anton Novojilov <andy@essentialkaos.com> - 0.5.5-0
- Fixed build using sources from source.kaos.st
- Dependencies update

* Fri May 27 2022 Anton Novojilov <andy@essentialkaos.com> - 0.5.4-0
- Updated for compatibility with the latest version of ek package

* Mon May 09 2022 Anton Novojilov <andy@essentialkaos.com> - 0.5.3-0
- Updated for compatibility with the latest version of ek package

* Tue Mar 29 2022 Anton Novojilov <andy@essentialkaos.com> - 0.5.2-0
- Removed pkg.re usage
- Added module info
- Added Dependabot configuration

* Fri Dec 04 2020 Anton Novojilov <andy@essentialkaos.com> - 0.5.1-0
- ek package updated to the latest stable version
- Added completion generation
- Added man page generation

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
