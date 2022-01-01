## OBS Scoreboard
Various web components built with the intention of displaying live scoreboards either from a Web UI or as an OBS plugin.

#### Server
Backend scoreboard server, supports an HTTP interface for scoreboard CRUD operations and a WebSocket interface for streaming realtime score updates. See `server` folder for details.

#### UI
Web UI that supports reading/writing scoreboards as well as viewing live score updates. See `ui` folder for details.

#### OBS Plugin
Plugin script for [Open Broadcast Softare (OBS)](https://obsproject.com/) to retrieve scores in realtime and update a text source. See `obs-client` folder for details.

###### Brought to you by Flet Inc.™