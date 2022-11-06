## Surface Types

These types are from physics data and show what type of contact each wheel is experiencing.

| ID  | Surface      |
|-----|--------------|
| 0   | Tarmac       | 
| 1   | Rumble strip | 
| 2   | Concrete     | 
| 3   | Rock         | 
| 4   | Gravel       | 
| 5   | Mud          | 
| 6   | Sand         | 
| 7   | Grass        | 
| 8   | Water        | 
| 9   | Cobblestone  | 
| 10  | Metal        | 
| 11  | Ridged       | 

## Button Flags

These flags are used in the telemetry packet to determine if any buttons are being held on the controlling device. If
the value below logical ANDed with the button status is set then the corresponding button is being held.

| Bit Flag   | Button            | 
|------------|-------------------|
| 0x00000001 | Cross or A        | 
| 0x00000002 | Triangle or Y     | 
| 0x00000004 | Circle or B       | 
| 0x00000008 | Square or X       | 
| 0x00000010 | D-pad Left        | 
| 0x00000020 | D-pad Right       | 
| 0x00000040 | D-pad Up          | 
| 0x00000080 | D-pad Down        | 
| 0x00000100 | Options or Menu   | 
| 0x00000200 | L1 or LB          | 
| 0x00000400 | R1 or RB          | 
| 0x00000800 | L2 or LT          | 
| 0x00001000 | R2 or RT          | 
| 0x00002000 | Left Stick Click  | 
| 0x00004000 | Right Stick Click |
| 0x00008000 | Right Stick Left  | 
| 0x00010000 | Right Stick Right | 
| 0x00020000 | Right Stick Up    |
| 0x00040000 | Right Stick Down  |
| 0x00080000 | Special           |
| 0x00100000 | UDP Action 1      |
| 0x00200000 | UDP Action 2      |
| 0x00400000 | UDP Action 3      |
| 0x00800000 | UDP Action 4      |
| 0x01000000 | UDP Action 5      |
| 0x02000000 | UDP Action 6      |
| 0x04000000 | UDP Action 7      |
| 0x08000000 | UDP Action 8      |
| 0x10000000 | UDP Action 9      |
| 0x20000000 | UDP Action 10     |
| 0x40000000 | UDP Action 11     |
| 0x80000000 | UDP Action 12     |

## Penalty Types

| ID  | Penalty meaning                                  |
|-----|--------------------------------------------------|
| 0   | Drive through                                    | 
| 1   | Stop Go                                          | 
| 2   | Grid penalty                                     | 
| 3   | Penalty reminder                                 | 
| 4   | Time penalty                                     | 
| 5   | Warning                                          | 
| 6   | Disqualified                                     | 
| 7   | Removed from formation lap                       | 
| 8   | Parked too long timer                            | 
| 9   | Tyre regulations                                 | 
| 10  | This lap invalidated                             | 
| 11  | This and next lap invalidated                    | 
| 12  | This lap invalidated without reason              | 
| 13  | This and next lap invalidated without reason     | 
| 14  | This and previous lap invalidated                | 
| 15  | This and previous lap invalidated without reason | 
| 16  | Retired                                          | 
| 17  | Black flag timer                                 |

## Infringement types

| ID  | Infringement meaning                            | 
|-----|-------------------------------------------------|
| 0   | Blocking by slow driving                        | 
| 1   | Blocking by wrong way driving                   | 
| 2   | Reversing off the start line                    | 
| 3   | Big Collision                                   | 
| 4   | Small Collision                                 | 
| 5   | Collision failed to hand back position single   | 
| 6   | Collision failed to hand back position multiple | 
| 7   | Corner cutting gained time                      | 
| 8   | Corner cutting overtake single                  | 
| 9   | Corner cutting overtake multiple                | 
| 10  | Crossed pit exit lane                           | 
| 11  | Ignoring blue flags                             | 
| 12  | Ignoring yellow flags                           | 
| 13  | Ignoring drive through                          | 
| 14  | Too many drive throughs                         | 
| 15  | Drive through reminder serve within n laps      | 
| 16  | Drive through reminder serve this lap           | 
| 17  | Pit lane speeding                               | 
| 18  | Parked for too long                             | 
| 19  | Ignoring tyre regulations                       | 
| 20  | Too many penalties                              | 
| 21  | Multiple warnings                               | 
| 22  | Approaching disqualification                    | 
| 23  | Tyre regulations select single                  | 
| 24  | Tyre regulations select multiple                | 
| 25  | Lap invalidated corner cutting                  | 
| 26  | Lap invalidated running wide                    | 
| 27  | Corner cutting ran wide gained time minor       | 
| 28  | Corner cutting ran wide gained time significant | 
| 29  | Corner cutting ran wide gained time extreme     | 
| 30  | Lap invalidated wall riding                     | 
| 31  | Lap invalidated flashback used                  | 
| 32  | Lap invalidated reset to track                  | 
| 33  | Blocking the pitlane                            | 
| 34  | Jump start                                      | 
| 35  | Safety car to car collision                     | 
| 36  | Safety car illegal overtake                     | 
| 37  | Safety car exceeding allowed pace               | 
| 38  | Virtual safety car exceeding allowed pace       | 
| 39  | Formation lap below allowed speed               | 
| 40  | Retired mechanical failure                      | 
| 41  | Retired terminally damaged                      | 
| 42  | Safety car falling too far back                 | 
| 43  | Black flag timer                                | 
| 44  | Unserved stop go penalty                        | 
| 45  | Unserved drive through penalty                  | 
| 46  | Engine component change                         | 
| 47  | Gearbox change                                  | 
| 48  | League grid penalty                             | 
| 49  | Retry penalty                                   | 
| 50  | Illegal time gain                               | 
| 51  | Mandatory pitstop                               | 
| 52  | Illegal time gain                               |
| 53  | Mandatory pitstop                               |
| 54  | Attribute assigned                              |
