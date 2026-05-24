class Event {
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}
  
class GOSocket {
    constructor(URL, options = {}) {
        this.URL = URL;
        this.OpenCallback = null;
        this.CloseCallback = null;
        this.ErrorCallback = null;
        this.conn = undefined;
        this.routeEventList = new Map();
        this.readyState = WebSocket.CLOSED;
        
        // ✨ Reconnection options med nye defaults
        this.reconnectAttempts = 0;
        this.maxReconnectAttempts = options.maxReconnectAttempts ?? 10;
        this.reconnectDelay = options.reconnectDelay ?? 2000;
        this.maxReconnectDelay = options.maxReconnectDelay ?? 60000;
        this.reconnectTimeout = null;
        
        // ✨ Message queueing - bufferer beskeder mens disconnected
        this.messageQueue = [];
        this.enableQueueing = options.enableQueueing ?? true;
    }
  
    connectionState() {
        return this.conn?.readyState ?? WebSocket.CLOSED;
    }
  
    isConnected() {
        return this.connectionState() === WebSocket.OPEN;
    }
  
    async start() {
        if (!window["WebSocket"]) {
            throw new Error("WebSocket not supported");
        }
  
        this.conn = new WebSocket(this.URL);
        this.readyState = WebSocket.CONNECTING;
  
        this.conn.onopen = () => {
            this.readyState = WebSocket.OPEN;
            this.reconnectAttempts = 0;
            console.log("WebSocket connected");
            
            this.flushMessageQueue();
            this.OpenCallback?.();
        };
  
        this.conn.onclose = () => {
            this.readyState = WebSocket.CLOSED;
            this.CloseCallback?.();
            this.attemptReconnect();
        };

        this.conn.onerror = (error) => {
            console.error("WebSocket error:", error);
            this.ErrorCallback?.(error);
        };
  
        this.conn.onmessage = (evt) => {
            try {
                const eventData = JSON.parse(evt.data);
                const event = new Event(eventData.type, eventData.payload);
                this.routeEvent(event);
            } catch (error) {
                console.error("Failed to parse message:", error);
            }
        };
    }

    attemptReconnect() {
        if (this.reconnectAttempts >= this.maxReconnectAttempts) {
            console.error(`Max reconnection attempts (${this.maxReconnectAttempts}) reached`);
            return;
        }

        this.reconnectAttempts++;
        
        const delay = Math.min(
            this.reconnectDelay * Math.pow(2, this.reconnectAttempts - 1),
            this.maxReconnectDelay
        );

        console.log(`Reconnecting in ${delay}ms (attempt ${this.reconnectAttempts}/${this.maxReconnectAttempts})`);

        this.reconnectTimeout = setTimeout(() => {
            this.start();
        }, delay);
    }

    stopReconnect() {
        if (this.reconnectTimeout) {
            clearTimeout(this.reconnectTimeout);
            this.reconnectTimeout = null;
        }
        this.reconnectAttempts = this.maxReconnectAttempts;
    }

    flushMessageQueue() {
        while (this.messageQueue.length > 0) {
            const { command, inputs } = this.messageQueue.shift();
            this.invoke(command, ...inputs);
        }
    }
  
    routeEvent(event) {
        const callback = this.routeEventList.get(event.type);
        if (callback) {
            try {
                callback(...event.payload);
            } catch (error) {
                console.error(`Error in callback for event "${event.type}":`, error);
            }
        }
    }
  
    close() {
        this.stopReconnect();
        this.conn?.close();
    }
  
    onOpen(callback) {
        this.OpenCallback = callback;
    }
  
    onClose(callback) {
        this.CloseCallback = callback;
    }

    onError(callback) {
        this.ErrorCallback = callback;
    }
  
    on(command, callback) {
        this.routeEventList.set(command, callback);
    }
  
    invoke(command, ...inputs) {
        if (this.isConnected()) {
            const event = new Event(command, inputs);
            this.conn.send(JSON.stringify(event));
        } else if (this.enableQueueing) {
            console.warn(`Connection not open, queueing "${command}"`);
            this.messageQueue.push({ command, inputs });
        } else {
            console.warn(`Cannot send "${command}" - connection not open`);
        }
    }
  
    off(command) {
        this.routeEventList.delete(command);
    }
}

// Standard - with queueing
//const socket = new GOSocket("ws://localhost:8080");

// without queueing
/*
const socket = new GOSocket("ws://localhost:8080", {
    enableQueueing: false
});

// Customize reconnection
const socket = new GOSocket("ws://localhost:8080", {
    maxReconnectAttempts: 20,
    reconnectDelay: 5000,
    maxReconnectDelay: 120000
});
*/
