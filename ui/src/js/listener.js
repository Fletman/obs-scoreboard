/**
 * Class for handling live score updates
 */
 class ScoreListener {
    /**
     * @param {string} host WebSocket endpoint to connect to
     * @todo Server/client auth?
     */
    constructor(host = null) {
        if(host) {
            this.host = host;
        } else if(process.env.VUE_APP_SOCKET_HOST) {
            this.host = process.env.VUE_APP_SOCKET_HOST;
        } else {
            this.host = "ws://localhost:8080/live";
        }
        this.event_handlers = {};
    }

    /**
     * Register an event handler
     * @param {string} event_name Event name
     * @param {Function} f Function handling event, takes a single event object
     * @public
     */
     on(event_name, f) {
        if(event_name in this.event_handlers) {
            this.event_handlers[event_name].push(f);
        } else {
            this.event_handlers[event_name] = [f];
        }
    }

    /**
     * Fire an event and invoke registered handlers
     * @param {string} event_name Event name
     * @param {Event} event Object containing event details
     * @private
     */
    emit(event_name, event) {
        const handlers = this.event_handlers[event_name];
        if(handlers && handlers.length > 0) {
            for(const handler of handlers) {
                handler(event);
            }
        }
    }

    /**
     * Connect listener to backend service to begin receiving events
     * @todo Server jwt auth
     */
    connect() {
        this.ws = new WebSocket(this.host);

        this.ws.addEventListener('open', (event) => {
            this.emit('open', event)
        });
        this.ws.addEventListener('message', (event) => {
            this.emit('message', event)
        });
        this.ws.addEventListener('error', (event) => {
            this.emit('error', event)
        });
        this.ws.addEventListener('close', (event) => {
            this.emit('close', event)
        });
    }
}

let handler;
module.exports = {
    /**
     *
     * @returns {ScoreListener} Socket handler
     */
    get_socket_handler() {
        if(!handler) {
            handler = new ScoreListener();
            handler.connect();
        }
        return handler;
    }
}