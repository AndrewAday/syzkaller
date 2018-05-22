#define _GNU_SOURCE
#include <fcntl.h>
#include <sys/ioctl.h>
#include <sys/mman.h>
#include <stdlib.h>
#include <unistd.h>

#define KCOV_INIT_TRACE _IOR('c', 1, unsigned long)
#define KCOV_ENABLE _IO('c', 100)
#define KCOV_DISABLE _IO('c', 101)
#define COVER_SIZE (16 << 20)


void kcov_setup() {
    unsigned long *cover;
    int fd;

    fd = open("/sys/kernel/debug/kcov", O_RDWR);
    if (fd == -1)
        perror("open"), exit(1);
    if (ioctl(fd, KCOV_INIT_TRACE, COVER_SIZE))
        perror("ioctl"), exit(1);
    cover = (unsigned long*)mmap(NULL,
                                COVER_SIZE * sizeof(unsigned long),
                                PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if ((void*)cover == MAP_FAILED)
        perror("mmap"), exit(1);
    if (ioctl(fd, KCOV_ENABLE, 0))
        perror("ioctl"), exit(1);
    
}

void main() {
    int i;
    for (i = 0; i < 4; i++)
        kcov_setup();
    sleep(10);
}
