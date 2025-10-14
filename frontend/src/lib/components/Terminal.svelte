<script lang="ts">

    import {onMount} from "svelte";
    import {authToken} from "$lib/store/auth";
    import "@xterm/xterm/css/xterm.css";


    export let server: string = "";
    export let onClose: () => void = () => {
    };
    export let onOpen: () => void = () => {
    };
    export let onMessage: (msg: string) => void = (msg) => {
    };
    export let onError: (err: Event) => void = (err) => {
    };
    let termContainer: HTMLDivElement;

    onMount(async () => {
        let xterm = (await import("@xterm/xterm"));
        const term = new xterm.Terminal({
            cursorBlink: true,
            theme: {
                background: '#000000',
                foreground: '#ffffff',
                cursor: '#ffffff',
            }
        });
        term.open(termContainer);
        term.focus();
        let wsProtocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const socket = new WebSocket(wsProtocol + '//' + window.location.host + '/api/ws/' + server + '/ssh');
// Événement: La connexion WebSocket est ouverte
        socket.onopen = (event) => {
            onOpen()
            //socket.send("docker exec -it fakeubuntu bash\n");
            term.write('\x1b[32m--- Connection established ---\r\n\x1b[0m');
            let unsub;
            unsub = authToken.subscribe((value) => {
                if (!value?.token) {
                    return
                }
                socket.send(JSON.stringify({type: "auth", token: value.token}));
            })
        };

        socket.onmessage = (event) => {
            onMessage(event.data);
            term.write(event.data);
        };
        term.onData(data => {
            socket.send(data);
        });

        socket.onclose = (event) => {
            onClose()
            term.write('\x1b[31m--- Connection closed ---\r\n\x1b[0m');

        };

        socket.onerror = (error) => {
            onError(error)
            term.write(`\x1b[31m[WebSocket Error] ${error.message}\r\n\x1b[0m`);
        };
    })
</script>

<div>
    <div id="terminal" bind:this={termContainer} class="bg-black">

    </div>
</div>