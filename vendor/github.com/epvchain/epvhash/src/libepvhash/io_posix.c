/*
  This file is part of epvhash.

  epvhash is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  epvhash is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with epvhash.  If not, see <http://www.gnu.org/licenses/>.
*/
/** @file io_posix.c
 */

#include "io.h"
#include <sys/types.h>
#include <sys/stat.h>
#include <errno.h>
#include <libgen.h>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <pwd.h>

FILE* epvhash_fopen(char const* file_name, char const* mode)
{
	return fopen(file_name, mode);
}

char* epvhash_strncat(char* dest, size_t dest_size, char const* src, size_t count)
{
	return strlen(dest) + count + 1 <= dest_size ? strncat(dest, src, count) : NULL;
}

bool epvhash_mkdir(char const* dirname)
{
	int rc = mkdir(dirname, S_IRWXU | S_IRWXG | S_IROTH | S_IXOTH);
	return rc != -1 || errno == EEXIST;
}

int epvhash_fileno(FILE *f)
{
	return fileno(f);
}

char* epvhash_io_create_filename(
	char const* dirname,
	char const* filename,
	size_t filename_length
)
{
	size_t dirlen = strlen(dirname);
	size_t dest_size = dirlen + filename_length + 1;
	if (dirname[dirlen] != '/') {
		dest_size += 1;
	}
	char* name = malloc(dest_size);
	if (!name) {
		return NULL;
	}

	name[0] = '\0';
	epvhash_strncat(name, dest_size, dirname, dirlen);
	if (dirname[dirlen] != '/') {
		epvhash_strncat(name, dest_size, "/", 1);
	}
	epvhash_strncat(name, dest_size, filename, filename_length);
	return name;
}

bool epvhash_file_size(FILE* f, size_t* ret_size)
{
	struct stat st;
	int fd;
	if ((fd = fileno(f)) == -1 || fstat(fd, &st) != 0) {
		return false;
	}
	*ret_size = st.st_size;
	return true;
}

bool epvhash_get_default_dirname(char* strbuf, size_t buffsize)
{
	static const char dir_suffix[] = ".epvhash/";
	strbuf[0] = '\0';
	char* home_dir = getenv("HOME");
	if (!home_dir || strlen(home_dir) == 0)
	{
		struct passwd* pwd = getpwuid(getuid());
		if (pwd)
			home_dir = pwd->pw_dir;
	}

	size_t len = strlen(home_dir);
	if (!epvhash_strncat(strbuf, buffsize, home_dir, len)) {
		return false;
	}
	if (home_dir[len] != '/') {
		if (!epvhash_strncat(strbuf, buffsize, "/", 1)) {
			return false;
		}
	}
	return epvhash_strncat(strbuf, buffsize, dir_suffix, sizeof(dir_suffix));
}
