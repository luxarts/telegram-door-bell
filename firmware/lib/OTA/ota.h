#ifndef OTA_H
#define OTA_H

#include <ESP8266WebServer.h>
#include <ESP8266HTTPUpdateServer.h>

class OTA {
    public:
        OTA(int port);
        void begin(void);
        void handle(void);
    private:
        int port;
        ESP8266WebServer httpServer;
        ESP8266HTTPUpdateServer httpUpdater;
};

#endif