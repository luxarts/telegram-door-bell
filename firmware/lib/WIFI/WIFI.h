#ifndef WIFI_H
#define WIFI_H

#include <ESP8266WiFi.h>
#include <ESP8266mDNS.h>
#include <logger.h>
#include "credentials.h"

#define DEVICE_NAME "TgDoorBell"
#define WIFI_TIMEOUT 100

class WIFI {
    public:
        WIFI(Logger _logger);
        void begin(void);
    private:
        Logger logger;
};

#endif