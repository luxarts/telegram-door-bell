#include "detector.h"

Detector::Detector(int _pin) {
    pin = _pin;
    prevState = 0;
}
void Detector::begin(void) {
    pinMode(pin, INPUT_PULLUP);
}

bool Detector::fallingEdgeDetected(void) {
    // Falling edge
    if(!digitalRead(pin) && prevState) {
        prevState = 0;
        return true;
    }

    // Rising edge
    if(digitalRead(pin) && !prevState) {
        prevState = 1;
    }

    return false;
}