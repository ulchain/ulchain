#serial 10

AU_ALIAS([AC_JNI_INCLUDE_DIR], [AX_JNI_INCLUDE_DIR])
AC_DEFUN([AX_JNI_INCLUDE_DIR],[

JNI_INCLUDE_DIRS=""

if test "x$JAVA_HOME" != x; then
	_JTOPDIR="$JAVA_HOME"
else
	if test "x$JAVAC" = x; then
		JAVAC=javac
	fi
	AC_PATH_PROG([_ACJNI_JAVAC], [$JAVAC], [no])
	if test "x$_ACJNI_JAVAC" = xno; then
		AC_MSG_WARN([cannot find JDK; try setting \$JAVAC or \$JAVA_HOME])
	fi
	_ACJNI_FOLLOW_SYMLINKS("$_ACJNI_JAVAC")
	_JTOPDIR=`echo "$_ACJNI_FOLLOWED" | sed -e 's://*:/:g' -e 's:/[[^/]]*$::'`
fi

case "$host_os" in
        darwin*)        _JTOPDIR=`echo "$_JTOPDIR" | sed -e 's:/[[^/]]*$::'`
                        _JINC="$_JTOPDIR/Headers";;
        *)              _JINC="$_JTOPDIR/include";;
esac
_AS_ECHO_LOG([_JTOPDIR=$_JTOPDIR])
_AS_ECHO_LOG([_JINC=$_JINC])

# On Mac OS X 10.6.4, jni.h is a symlink:
# /System/Library/Frameworks/JavaVM.framework/Versions/Current/Headers/jni.h
# -> ../../CurrentJDK/Headers/jni.h.

AC_CACHE_CHECK(jni headers, ac_cv_jni_header_path,
[
if test -f "$_JINC/jni.h"; then
  ac_cv_jni_header_path="$_JINC"
  JNI_INCLUDE_DIRS="$JNI_INCLUDE_DIRS $ac_cv_jni_header_path"
else
  _JTOPDIR=`echo "$_JTOPDIR" | sed -e 's:/[[^/]]*$::'`
  if test -f "$_JTOPDIR/include/jni.h"; then
    ac_cv_jni_header_path="$_JTOPDIR/include"
    JNI_INCLUDE_DIRS="$JNI_INCLUDE_DIRS $ac_cv_jni_header_path"
  else
    ac_cv_jni_header_path=none
  fi
fi
])



# get the likely subdirectories for system specific java includes
case "$host_os" in
bsdi*)          _JNI_INC_SUBDIRS="bsdos";;
darwin*)        _JNI_INC_SUBDIRS="darwin";;
freebsd*)       _JNI_INC_SUBDIRS="freebsd";;
linux*)         _JNI_INC_SUBDIRS="linux genunix";;
osf*)           _JNI_INC_SUBDIRS="alpha";;
solaris*)       _JNI_INC_SUBDIRS="solaris";;
mingw*)		_JNI_INC_SUBDIRS="win32";;
cygwin*)	_JNI_INC_SUBDIRS="win32";;
*)              _JNI_INC_SUBDIRS="genunix";;
esac

if test "x$ac_cv_jni_header_path" != "xnone"; then
  # add any subdirectories that are present
  for JINCSUBDIR in $_JNI_INC_SUBDIRS
  do
      if test -d "$_JTOPDIR/include/$JINCSUBDIR"; then
           JNI_INCLUDE_DIRS="$JNI_INCLUDE_DIRS $_JTOPDIR/include/$JINCSUBDIR"
      fi
  done
fi
])

# _ACJNI_FOLLOW_SYMLINKS <path>
# Follows symbolic links on <path>,
# finally setting variable _ACJNI_FOLLOWED
# ----------------------------------------
AC_DEFUN([_ACJNI_FOLLOW_SYMLINKS],[
# find the include directory relative to the javac executable
_cur="$1"
while ls -ld "$_cur" 2>/dev/null | grep " -> " >/dev/null; do
        AC_MSG_CHECKING([symlink for $_cur])
        _slink=`ls -ld "$_cur" | sed 's/.* -> //'`
        case "$_slink" in
        /*) _cur="$_slink";;
        # 'X' avoids triggering unwanted echo options.
        *) _cur=`echo "X$_cur" | sed -e 's/^X//' -e 's:[[^/]]*$::'`"$_slink";;
        esac
        AC_MSG_RESULT([$_cur])
done
_ACJNI_FOLLOWED="$_cur"
])# _ACJNI
