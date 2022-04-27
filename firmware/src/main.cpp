#include <Arduino.h>
#include <logger.h>
#include <WiFi.h>
#include <ota.h>
#include <restclient.h>
#include <detector.h>

#define OTA_PORT 80
#define PIN_SIGNAL 12

Logger logger;
WIFI wifi(logger);
OTA ota(OTA_PORT);
RestClient restClient(logger);
Detector detector(PIN_SIGNAL);

void setup() {
  wifi.begin();
  ota.begin();
  detector.begin();
}

void loop() {
  ota.handle();

  if(detector.fallingEdgeDetected()) {
    logger.Writeln(F("Pulse detected!"));
    restClient.POST(F("<server>"), F("Bearer <token>"));
  }
}