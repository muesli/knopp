int pinA = 12;  // Connected to CLK on KY-040
int pinB = 13;  // Connected to DT on KY-040
int pinSW = 14; // Connected to SW on KY-040

int encoderPos = 0;
int lastA;
boolean bCW;
boolean pressed = false;

void setup() { 
    pinMode(pinA, INPUT);
    pinMode(pinB, INPUT);
    pinMode(pinSW, INPUT_PULLUP);

    lastA = digitalRead(pinA);   

    Serial.begin(115200);
} 

void loop() { 
    int aVal = digitalRead(pinA);
    int bVal = digitalRead(pinB);
    int swVal = digitalRead(pinSW);

    if (swVal == 0 && !pressed) {
        Serial.println("D");
        pressed = true;
    } else if (swVal == 1 && pressed) {
        Serial.println("U");
        pressed = false;
    }

    if (aVal != lastA && aVal == LOW) {
        // Means the knob is rotating
        // if the knob is rotating, we need to determine direction
        // We do that by reading pin B.
        bCW = (bVal == LOW);
        if (bCW) {
            encoderPos++;
        } else {
            encoderPos--;
        }

        // Serial.print("Rotated: ");
        if (bCW) {
            // Serial.println("clockwise");
        } else {
            // Serial.println("counterclockwise");
        }
        // Serial.print("Encoder Position: ");
        Serial.println(encoderPos);
    }

    lastA = aVal;
 }
 
