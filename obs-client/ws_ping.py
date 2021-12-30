import json
import _thread
import time
import websocket

def on_message(ws, msg):
    print(json.loads(msg))

def on_error(ws, error):
    print(error)

def on_close(ws, status, close_msg):
    print("Websocket closed with status {}. {}".format(status, close_msg))

def on_open(ws):
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


if __name__ == '__main__':
    websocket.enableTrace(False)
    port = 8080
    ws = websocket.WebSocketApp(
        "ws://localhost:{}/live".format(port),
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close
    )
    ws.run_forever()