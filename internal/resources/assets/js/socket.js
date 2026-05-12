export class Event {
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}

export class GOSocket {
    constructor(URL) {
        this.URL = URL;
        this.OpenCallback = null;
        this.CloseCallback = null;
        this.conn = undefined;
        this.routeEventList = new Map();
        this.readyState = WebSocket.CLOSED;
    }

    connectionState() {
        return this.conn ? this.conn.readyState : WebSocket.CLOSED;
    }

    async start() {
        // Check if the browser supports WebSocket
        if (!window["WebSocket"]) {
            alert("Not supporting websockets");
            return;
        }

        this.conn = new WebSocket(this.URL);
        this.readyState = WebSocket.CONNECTING;

        this.conn.onopen = () => {
            this.readyState = WebSocket.OPEN;
            if (this.OpenCallback) this.OpenCallback();
        };

        this.conn.onclose = () => {
            this.readyState = WebSocket.CLOSED;
            if (this.CloseCallback) this.CloseCallback();
        };

        this.conn.onmessage = (evt) => {
            const eventData = JSON.parse(evt.data);
            let event = new Event(eventData.type, eventData.payload);
            this.routeEvent(event);
        };
    }

    routeEvent(event) {
        let callback = this.routeEventList.get(event.type);
        if (callback !== undefined) {
            // Vi spreder payload ud som argumenter (spread operator)
            callback(...event.payload);
        }
    }

    close() {
        if (this.conn) {
            this.conn.close();
        }
    }

    onOpen(callback) {
        this.OpenCallback = callback;
    }

    onClose(callback) {
        this.CloseCallback = callback;
    }

    on(command, callback) {
        this.routeEventList.set(command, callback);
    }

    // send a message to the server
    invoke(command, ...inputs) {
        if (this.conn) {
            const event = new Event(command, inputs);
            this.conn.send(JSON.stringify(event));
        }
    }

    off(command) {
        this.routeEventList.delete(command);
    }
}