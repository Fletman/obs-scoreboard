import json
import obspython as obs
import _thread
from modules import websocket

ws = None

# ---------- WebSocket Functions ----------------

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
    # text params: props obj, prop name, prop desc, format
    # number params: props obj, prop name, prop desc, min val, max val, increment size

    obs.obs_properties_add_text(props, "Test", "Test Input String", obs.OBS_TEXT_DEFAULT)

    return props

def script_update(settings):
    """Called when settings are updated by user"""
    global test_string
    test_string = obs.obs_data_get_string(settings, "Test")
    obs.script_log(obs.LOG_INFO, test_string)

def script_defaults(settings):
    """Set default values for script settings/parameters"""
    obs.obs_data_set_default_string(settings, "Test", "bip")

def script_load(settings):
    """Called on script startup"""
    global ws
    websocket.enableTrace(False)
    port = 8080
    host = "ws://localhost:{}/live".format(port)
    obs.script_log(obs.LOG_INFO, "Attempting connection to host: {}".format(host))
    ws = websocket.WebSocketApp(
        host,
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close
    )
    ws.run_forever()

def script_unload():
    """Called when script is unloaded"""
    global ws
    ws.close()