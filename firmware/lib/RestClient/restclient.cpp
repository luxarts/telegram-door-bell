#include "restclient.h"

RestClient::RestClient(Logger _logger) {
    logger = _logger;
}

void RestClient::POST(const __FlashStringHelper *url, const __FlashStringHelper *token) {
    http.begin(client, url);
    http.addHeader("Authorization", token);
    
    logger.Write(F("Sending POST to "));
    logger.Write(url);
    logger.Write(F(" with token "));
    logger.Writeln(token);
    
    int response = http.POST("");
    
    logger.Write(F("API response: "));
    logger.Writeln(response);

    http.end();
}