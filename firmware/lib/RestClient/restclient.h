#ifndef RESTCLIENT_H
#define RESTCLIENT_H

#include <Arduino.h>
#include <WiFiClient.h>
#include <ESP8266HTTPClient.h>
#include "logger.h"

class RestClient {
    public:
        RestClient(Logger _logger);
        void POST(const __FlashStringHelper *url, const __FlashStringHelper *token);
    private:
        Logger logger;
        WiFiClient client;
        HTTPClient http;
};

#endif