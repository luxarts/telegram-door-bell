#include "ota.h"

OTA::OTA(int _port) {
    port = _port;
}

void OTA::begin(void) {
    httpUpdater.setup(&httpServer);
    httpServer.begin(port);
}

void OTA::handle() {
    httpServer.handleClient();
}