// hello worldをシリアルモニタに出し続ける
// シリアルモニタテスト

void setup()
{
  Serial.begin(9600);
  while (! Serial); // Wait untilSerial is ready
}

void loop()
{
  for (int i=0; i<10; i++) {
      Serial.print("hello world");
      Serial.print(i);
      Serial.println("");
  }
}
