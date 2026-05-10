

export class Event {
    type : string 
    payload : any[]

    constructor(type : string , payload : any) {
        this.type = type;
        this.payload = payload;
    }
}

export class GOSocket {
    URL : string
    OpenCallback?  : (...inputs: any) => void 
    CloseCallback? : (...inputs: any) => void
    conn : WebSocket | undefined 
    routeEventList = new Map<string, (...inputs: any) => void>()
    readyState : number = WebSocket.CLOSED


    constructor(URL : string) {
        this.URL = URL 
    }

    connectionState() : number {
        return this.conn!.readyState
    }

    async start() {

        // Check if the browser supports WebSocket
        if (!window["WebSocket"]) {
            alert("Not supporting websockets");
            return
        }
        this.conn = new WebSocket(this.URL);
        //console.log("readyState:" + this.conn.readyState)
        this.readyState = WebSocket.CONNECTING

        this.conn.onopen = () => {
            this.readyState = WebSocket.OPEN
            if(this.OpenCallback) this.OpenCallback()
        }

        this.conn.onclose = () => {
            if(this.CloseCallback) this.CloseCallback()
        }

       
        this.conn.onmessage = (evt : MessageEvent<any>) => {
            const eventData = JSON.parse(evt.data)
            let event = new Event(eventData.type, eventData.payload) 
            this.routeEvent(event)
        }  
    }

    routeEvent(event : Event) {
        let callback = this.routeEventList.get(event.type)
        if(callback != undefined) {
            callback(...event.payload)
        }
    }

    close() {
        this.conn?.close()
    }

    onOpen(callback: (...inputs: any) => void) {
       this.OpenCallback = callback
    }

    onClose(callback: (...inputs: any) => void) {
        this.CloseCallback = callback
    }

    on(command: string, callback: (...args: any) => any): void {
        this.routeEventList.set(command, callback)
    }
    
    // Handler to send a message to the sever
    invoke(command : string, ...inputs: any) {
        
        if(this.conn) {
            const event = new Event(command, inputs)
            this.conn.send(JSON.stringify(event))
        }
    }
    
    off(command: string) {
        this.routeEventList.delete(command)
    }
}


