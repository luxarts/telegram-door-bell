#include <WIFI.h>

WIFI::WIFI(Logger _logger) {
	logger = _logger;
}
void WIFI::begin(void) {
	// Setup
	logger.Writeln(F("Setting mode WIFI_STA"));
	WiFi.mode(WIFI_STA);
	logger.Write(F("Setting device name to "));
	logger.Writeln(DEVICE_NAME);
	WiFi.hostname(DEVICE_NAME);

	// Connect
	logger.Write(F("Connecting to "));
	logger.Writeln(WIFI_SSID);
	WiFi.begin(WIFI_SSID, WIFI_PASS);
	while(WiFi.waitForConnectResult() != WL_CONNECTED) {
		logger.Writeln(F("Retrying connection..."));
		WiFi.begin(WIFI_SSID, WIFI_PASS);
		delay(WIFI_TIMEOUT);
	}

	logger.Writeln(F("Connected!"));
	logger.Write(F("IP: "));
	logger.Writeln(WiFi.localIP().toString().c_str());
}