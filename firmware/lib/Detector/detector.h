#ifndef DETECTOR_H
#define DETECTOR_H

#include <Arduino.h>

class Detector {
    public:
        Detector(int _pin);
        void begin(void);
        bool fallingEdgeDetected(void);
    private:
        int pin;
        bool prevState;
};

#endif