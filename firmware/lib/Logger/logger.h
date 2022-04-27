#ifndef LOGGER_H
#define LOGGER_H

#include <Arduino.h>

#define SERIAL_BAUDRATE 115200

class Logger {
    public:
        Logger();
        void Write(char *msg);
        void Write(const char *msg);
        void Write(const __FlashStringHelper *msg);
        void Write(int msg);
        void Writeln(char *msg);
        void Writeln(const char *msg);
        void Writeln(const __FlashStringHelper *msg);
        void Writeln(int msg);
};

#endif