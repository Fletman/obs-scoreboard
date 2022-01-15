import json
import obspython as obs
import _thread
from modules import websocket

ws = None
host = ""

# ---------- WebSocket Functions ----------------
def ws_connect(props, prop):
    """Connect to WebSocket"""
    global ws
    global host
    websocket.enableTrace(False)
    obs.script_log(obs.LOG_INFO, "Attempting connection to host: {}".format(host))
    ws = websocket.WebSocketApp(
        host,
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close
    )
    ws.run_forever()

def ws_close(props, prop):
    """Close WebSocket connection"""
    global ws
    ws.close()

def on_message(ws, msg):
    """Handler for messages received from WebSocket"""
    obs.script_log(obs.LOG_INFO, json.loads(msg))

def on_error(ws, error):
    """Handler for WebSocket errors"""
    obs.script_log(obs.LOG_ERROR, repr(error))

def on_close(ws, status, close_msg):
    """Handler for Websocket closing"""
    obs.script_log(obs.LOG_INFO, "Websocket closed with status {}. {}".format(status, close_msg))

def on_open(ws):
    """Handler for WebSocket connection"""
    obs.script_log(obs.LOG_INFO, "Successfully connected")
    def ping(*args):
        ping_sec = args[0]
        i = 0
        while i < 60 :
            i += 1
            msg = "Ping {}".format(i)
            ws.send(msg)
            time.sleep(ping_sec)
        ws.close()
    _thread.start_new_thread(ping, (5,))

# ---------- OBS Functions ----------------------

def script_description():
    """Return string of script description"""
    return "Connects to scoreboard server and listens for updates to featured scoreboard, then outputs to specified text sources"

def script_properties():
    """Declare properties that can be set/altered in OBS script UI"""
    props = obs.obs_properties_create()
    obs.obs_properties_add_text(props, "host", "URL for scoreboard server", obs.OBS_TEXT_DEFAULT)
    obs.obs_properties_add_button(props, "start_button", "Start", ws_connect)
    obs.obs_properties_add_button(props, "stop_button", "Stop", ws_close)
    return props

def script_update(settings):
    """Called when settings are updated by user"""
    global host
    host = obs.obs_data_get_string(settings, "host")
    obs.script_log(obs.LOG_INFO, host)

def script_defaults(settings):
    """Set default values for script settings/parameters"""
    obs.obs_data_set_default_string(settings, "host", "ws://localhost:8080")

def script_load(settings):
    """Called on script startup"""
    pass

def script_unload():
    """Called when script is unloaded"""
    global ws
    ws.close()