import React, { useEffect, useRef, useState } from 'react';

type Message = {
    action: string;
    user: string;
    body: any;
    room: string;
};

function WebChat() {
    const [username, setUsername] = useState('');
    const [room, setRoom] = useState('');
    const [connected, setConnected] = useState(false);

    const [messages, setMessages] = useState<Message[]>([]);
    const [input, setInput] = useState('');

    const socketRef = useRef<WebSocket | null>(null);
    const messageEndRef = useRef<HTMLDivElement | null>(null);

    /* Tells the backend that a user wants to join a specific room, and then sends the body
    or the text content into the server as a string. If a message from the server (Ex. 'User joined')
    or a message from another person within the server is sent, it will be displayed in the chatbox. */

    useEffect(() => {

        if (connected) {

            const socket = new WebSocket(`ws://localhost:8080/ws?username=${username}`);
            socketRef.current = socket;

            socket.onopen = () => {
                const joinMsg = {
                    action: 'join',
                    user: username,
                    room: room,
                    body: null,
                };
                socket.send(JSON.stringify(joinMsg));
            };

            socket.onmessage = (event) => {
                const msg: Message = JSON.parse(event.data);
                setMessages((prev) => [...prev, msg]);
            };

            socket.onerror = (err) => {
                console.error("WebSocket error:", err);
            };

            socket.onclose = () => {
                console.warn("WebSocket closed");
            };

            return () => {
                socket.close();
            };
        }
    }, [connected, username, room]);

    // Simple automatic scroll to the most recent message sent in the web chat

    useEffect(() => {
        if (messageEndRef.current) {
            messageEndRef.current.scrollIntoView({ behavior: 'smooth' });
        }
    }, [messages]);

    /* This gets called when the 'Send' button is pressed; Checks if there's a WebSocket connection
    to the server, checks if the user input isn't just spaces, then creates a message object using
    the same structure constructed in the backend. It gets sent to the WebSocket server, then clears 
    the input box after it is sent. */

    const sendMessage = () => {

        if (socketRef.current && input.trim()) {

            const msg: Message = {
                action: 'message',
                user: username,
                room: room,
                body: input,
            };

            socketRef.current.send(JSON.stringify(msg));
            setInput('');
            
        }

    };

    /* Simply checks if username and room input variables are non-empty. If they are non-empty, it then 
    activates a WebSocket connection */

    const handleJoin = () => {
        if (username && room) {
            setConnected(true);
        }
    };

    /* The screen that allows the user to input a username and a room number to join, only using
    one of the three hard-coded room numbers (0, 1, or 2). This screen is only shown when the
    user is NOT yet connected to any chat room number. */

    if (!connected) {

        return (

            <div style={{ padding: 20 }}>
                <h2>Join Chat Room</h2>

                <input
                    placeholder="Username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />

                <input
                    placeholder="Room ID (e.g. 0, 1, 2)"
                    value={room}
                    onChange={(e) => setRoom(e.target.value)}
                />

                <button onClick={handleJoin}>Join</button>

            </div>

        );

    }

    // Screen for the general chat

    return (

        <div style={{ padding: 20 }}>
            <h2>Chat Room {room}</h2>

            <div style={{ border: '1px solid black', height: '300px', overflowY: 'scroll', marginBottom: 10 }}>
               
                {messages.map((msg, idx) => (

                    <div key={idx}>
                        <strong>{msg.user}:</strong> {msg.body}
                    </div>

                ))}

                <div ref={messageEndRef} />

            </div>

            <input
                type="text"
                value={input}
                onChange={(e) => setInput(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && sendMessage()}
            />

            <button onClick={sendMessage}>Send</button>

        </div>

    );
}

export default WebChat;
