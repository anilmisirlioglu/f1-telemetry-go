## Packet ID Table

The packets IDs are as follows:

| Packet Name | Value | Description |
| --- | --- | --- |
| Motion | 0 | Contains all motion data for player’s car – only sent while player is in control |
| Session | 1 | Data about the session – track, time left |
| Lap Data | 2 | Data about all the lap times of cars in the session |
| Event | 3 | Various notable events that happen during a session |
| Participants | 4 | List of participants in the session, mostly relevant for multiplayer |
| Car Setups | 5 | Packet detailing car setups for cars in the race |
| Car Telemetry | 6 | Telemetry data for all cars |
| Car Status | 7 | Status data for all cars such as damage |
| Final Classification | 8 | Final classification confirmation at the end of a race |
| Lobby Info | 9 | Information about players in a multiplayer lobby |