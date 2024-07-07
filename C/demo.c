#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <termios.h>
#include <errno.h>

#define MAX_BUF_SIZE 100

typedef struct {
    int fd;
    char comName[20];
} SerialPort;

SerialPort serialPorts[10];
int numPorts = 0;

void init() {
    numPorts = 0;
}

int openSerialPort(const char *comName) {
    int fd = open(comName, O_RDWR | O_NOCTTY | O_NDELAY);
    if (fd == -1) {
        perror("openSerialPort - Unable to open port");
        return -1;
    }
    struct termios options;
    tcgetattr(fd, &options);
    cfsetispeed(&options, B9600);
    cfsetospeed(&options, B9600);
    options.c_cflag |= (CLOCAL | CREAD);
    options.c_cflag &= ~PARENB;
    options.c_cflag &= ~CSTOPB;
    options.c_cflag &= ~CSIZE;
    options.c_cflag |= CS8;
    options.c_lflag &= ~(ICANON | ECHO | ECHOE | ISIG);
    options.c_oflag &= ~OPOST;
    options.c_cc[VMIN] = 0;
    options.c_cc[VTIME] = 10;
    tcsetattr(fd, TCSANOW, &options);
    return fd;
}

void serialListen(const char *comName) {
    int fd = openSerialPort(comName);
    if (fd == -1) {
        return;
    }
    serialPorts[numPorts].fd = fd;
    strcpy(serialPorts[numPorts].comName, comName);
    numPorts++;
    printf("Open serial port: %s\n", comName);
    char buf[MAX_BUF_SIZE];
    while (1) {
        ssize_t n = read(fd, buf, sizeof(buf));
        if (n > 0) {
            printf("Received data: ");
            for (int i = 0; i < n; ++i) {
                printf("%02X ", buf[i]);
            }
            printf("\n");
            // Process the received data and send response if needed
            // Example: write(fd, response, responseSize);
        } else if (n == -1) {
            perror("serialListen - Read error");
            break;
        }
        usleep(100000); // Sleep for 100ms
    }
}

void closePort(const char *comName) {
    for (int i = 0; i < numPorts; ++i) {
        if (strcmp(serialPorts[i].comName, comName) == 0) {
            close(serialPorts[i].fd);
            printf("Closed serial port: %s\n", comName);
            // Remove closed port from array by shifting elements
            for (int j = i; j < numPorts - 1; ++j) {
                serialPorts[j] = serialPorts[j + 1];
            }
            numPorts--;
            break;
        }
    }
}

//go generate arm-openwrt-linux-muslgnueabi-gcc -o demo demo.c
int main() {
    init();
    serialListen("/dev/ttyS1"); // Change the COM port name as needed
    return 1;
}